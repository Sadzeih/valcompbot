package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Sadzeih/valcompbot/comments"
	"github.com/Sadzeih/valcompbot/ent/pickemsevent"
	"github.com/Sadzeih/valcompbot/matches"
	"github.com/Sadzeih/valcompbot/pickems"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/migrate"
	"github.com/Sadzeih/valcompbot/highlighter"
	"github.com/Sadzeih/valcompbot/internal/api"
	_ "github.com/lib/pq"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
	if err := config.Parse(); err != nil {
		log.Print(err)
		return
	}
	credentials := reddit.Credentials{
		ID:       config.Get().RedditClientID,
		Secret:   config.Get().RedditClientSecret,
		Username: config.Get().RedditUsername,
		Password: config.Get().RedditPassword,
	}

	redditClient, err := reddit.NewClient(credentials)
	if err != nil {
		log.Print(err)
		return
	}

	entClient, err := ent.Open("postgres", config.Get().PostgresString)
	if err != nil {
		log.Print(err)
		return
	}
	defer entClient.Close()
	ctx := context.Background()
	if err := entClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Printf("failed creating schema resources: %v", err)
		return
	}

	// A wait group for synchronizing routines
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Sidebar ticker routine
	go func() {
		defer wg.Done()
		startSidebar(redditClient)
	}()

	// Comments topic routine
	commentsTopic := comments.New(redditClient)
	go func() {
		defer wg.Done()

		defer commentsTopic.Close()

		commentsTopic.Run()
	}()

	// PMT scheduler
	scheduler := matches.NewScheduler(entClient, redditClient)
	go func() {
		defer wg.Done()
		scheduler.Start(ctx)
	}()

	// API routine
	go func() {
		defer wg.Done()
		err := api.Start(redditClient, entClient)
		if err != nil {
			log.Printf("error while running API: %v", err)
			return
		}
	}()

	// Pickems routine
	if config.Get().EnablePickems {
		go func() {
			defer wg.Done()

			pickemsEvent, err := entClient.PickemsEvent.Query().
				Order(ent.Desc(pickemsevent.FieldTimestamp)).
				Limit(1).
				Only(ctx)
			if ent.MaskNotFound(err) != nil {
				log.Print(fmt.Errorf("could not get latest pickems event"))
				return
			}
			if pickemsEvent != nil {
				pickems.Event = pickemsEvent.EventID
			}

			pickemsService := pickems.New(redditClient, commentsTopic.Subscribe())
			if err := pickemsService.Run(); err != nil {
				return
			}
		}()
	}

	// Comment highlighter routine
	if config.Get().EnableStickies {
		go func() {
			defer wg.Done()

			h, err := highlighter.New(context.Background(), redditClient, entClient, commentsTopic.Subscribe())
			if err != nil {
				log.Printf("failed creating highlighter: %v", err)
				return
			}
			if err := h.Run(); err != nil {
				log.Printf("failed to run highlighter: %v", err)
				return
			}
		}()
	}

	// Sentinels routine
	if config.Get().EnableSentinels {
		go func() {
			defer wg.Done()
			if err := DaysSinceLastSentinelsPost(redditClient); err != nil {
				log.Printf("Error in DaysSinceLastSentinelsPost: %v", err)
				return
			}
		}()
	}

	wg.Wait()
}
