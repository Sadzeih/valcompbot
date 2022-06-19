package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/matches"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func startSidebar(rClient *reddit.Client) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	subSettings, _, err := rClient.Subreddit.GetSettings(context.Background(), config.Get().RedditSubreddit)
	if err != nil {
		fmt.Println(err)
		return
	}

	m, err := matches.GetUpcoming()
	if err != nil {
		fmt.Println(err)
	}
	if err := BuildSidebar(rClient, subSettings, m); err != nil {
		fmt.Println(err)
	}
	if err := BuildWidget(rClient, m); err != nil {
		fmt.Println(err)
	}

	for range ticker.C {
		m, err = matches.GetUpcoming()
		if err != nil {
			fmt.Println(err)
		}

		subSettings, _, err = rClient.Subreddit.GetSettings(context.Background(), config.Get().RedditSubreddit)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := BuildSidebar(rClient, subSettings, m); err != nil {
			fmt.Println(err)
		}
		if err := BuildWidget(rClient, m); err != nil {
			fmt.Println(err)
		}
	}
}

func BuildSidebar(client *reddit.Client, subSettings *reddit.SubredditSettings, m []matches.Upcoming) error {
	sidebarMd, err := format(true, m)
	if err != nil {
		return err
	}
	subSettings.Sidebar = &sidebarMd

	_, err = client.Subreddit.Edit(context.Background(), subSettings.ID, subSettings)
	if err != nil {
		return err
	}

	return nil
}

func BuildWidget(client *reddit.Client, m []matches.Upcoming) error {
	sidebarMd, err := format(false, m)
	if err != nil {
		return err
	}

	widgets, _, err := client.Widget.Get(context.Background(), config.Get().RedditSubreddit)
	if err != nil {
		return err
	}
	var tickerWidget *reddit.TextAreaWidget
	for _, widget := range widgets {
		w, ok := widget.(*reddit.TextAreaWidget)
		if ok && w.Name == "Match Ticker" {
			tickerWidget = w
			break
		}
	}

	if tickerWidget == nil {
		return fmt.Errorf("BuildWidget: could not find ticker widget")
	}

	widgetReq := reddit.TextAreaWidgetRequest{
		Style: tickerWidget.Style,
		Text:  sidebarMd,
		Name:  tickerWidget.Name,
	}

	_, _, err = client.Widget.Update(context.Background(), config.Get().RedditSubreddit, tickerWidget.GetID(), &widgetReq)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
