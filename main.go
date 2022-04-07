package main

import (
	"context"
	"fmt"
	"github.com/Sadzeih/valcompbot/vlr"
	"log"
	"sync"
	"time"

	"github.com/Sadzeih/valcompbot/config"
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
	client, err := reddit.NewClient(credentials)
	if err != nil {
		log.Fatal(err)
	}

	// A wait group for synchronizing routines
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Sidebar ticker routine
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		subSettings, _, err := client.Subreddit.GetSettings(context.Background(), config.Get().RedditSubreddit)
		if err != nil {
			fmt.Println(err)
			return
		}

		matches, err := vlr.GetUpcomingMatches()
		if err != nil {
			fmt.Println(err)
		}
		if err := BuildSidebar(client, subSettings, matches); err != nil {
			fmt.Println(err)
		}
		if err := BuildWidget(client, matches); err != nil {
			fmt.Println(err)
		}

		for range ticker.C {
			matches, err = vlr.GetUpcomingMatches()
			if err != nil {
				fmt.Println(err)
			}
			if err := BuildSidebar(client, subSettings, matches); err != nil {
				fmt.Println(err)
			}
			if err := BuildWidget(client, matches); err != nil {
				fmt.Println(err)
			}
		}
	}()

	wg.Wait()
}
