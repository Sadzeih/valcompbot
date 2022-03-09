package main

import (
	"context"
	"fmt"
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

	errChan := make(chan error)

	// Sidebar ticker routine
	wg.Add(1)
	go func(client *reddit.Client) {
		defer wg.Done()
		ticker := time.Tick(1 * time.Minute)
		for next := range ticker {
			subSettings, _, err := client.Subreddit.GetSettings(context.Background(), "valcompbottest")
			if err != nil {
				errChan <- err
				return
			}

			sidebarFmt := `## hello this is cool
next update: %+v`

			sidebar := fmt.Sprintf(sidebarFmt, next)

			subSettings.Sidebar = &sidebar

			resp, err := client.Subreddit.Edit(context.Background(), subSettings.ID, subSettings)
			if err != nil {
				errChan <- err
				return
			}

			fmt.Printf("%+v\n", resp)
		}
	}(client)

	wg.Wait()
	for err := range errChan {
		fmt.Println(err)
	}
}
