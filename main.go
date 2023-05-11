package main

import (
	"context"
	"log"
	"sync"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/migrate"
	"github.com/Sadzeih/valcompbot/internal/api"
	_ "github.com/lib/pq"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
	if err := config.Parse(); err != nil {
		log.Fatal(err)
	}
	credentials := reddit.Credentials{
		ID:       config.Get().RedditClientID,
		Secret:   config.Get().RedditClientSecret,
		Username: config.Get().RedditUsername,
		Password: config.Get().RedditPassword,
	}

	redditClient, err := reddit.NewClient(credentials)
	if err != nil {
		log.Fatal(err)
	}

	entClient, err := ent.Open("postgres", config.Get().PostgresString)
	if err != nil {
		log.Fatal(err)
	}
	defer entClient.Close()
	ctx := context.Background()
	if err := entClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// A wait group for synchronizing routines
	wg := sync.WaitGroup{}
	wg.Add(1)
	//
	//// Sidebar ticker routine
	//go func() {
	//	defer wg.Done()
	//	startSidebar(redditClient)
	//}()

	// API routine
	go func() {
		//defer wg.Done()
		api.Start(redditClient, entClient)
	}()

<<<<<<< Updated upstream
	if config.Get().EnableStickies {
		// Comment highlighter routine
		go func() {
			h, err := highlighter.New(context.Background(), redditClient, entClient)
			if err != nil {
				log.Fatalf("failed creating highlighter: %v", err)
			}
			if err := h.Run(); err != nil {
				log.Fatalf("failed to run highlighter: %v", err)
			}
		}()
	}

	if config.Get().EnableSentinels {
		go func() {
			defer wg.Done()
			if err := DaysSinceLastSentinelsPost(redditClient); err != nil {
				log.Fatalf("Error in DaysSinceLastSentinelsPost: %v", err)
			}
		}()
	}
=======
	//// Comment highlighter routine
	//go func() {
	//	h, err := highlighter.New(context.Background(), redditClient, entClient)
	//	if err != nil {
	//		log.Fatalf("failed creating highlighter: %v", err)
	//	}
	//	if err := h.Run(); err != nil {
	//		log.Fatalf("failed to run highlighter: %v", err)
	//	}
	//}()
	//
	//e, err := strconv.ParseBool(config.Get().EnableSentinels)
	//if err != nil {
	//	log.Fatalf("failed to parse bool: %v", err)
	//}
	//if e {
	//	go func() {
	//		defer wg.Done()
	//		if err := DaysSinceLastSentinelsPost(redditClient); err != nil {
	//			log.Fatalf("Error in DaysSinceLastSentinelsPost: %v", err)
	//		}
	//	}()
	//}
>>>>>>> Stashed changes

	wg.Wait()
}
