package main

import (
	"context"
	"fmt"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/vlr"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func BuildSidebar(client *reddit.Client, subSettings *reddit.SubredditSettings, matches []vlr.UpcomingMatch) error {
	sidebarMd, err := format(true, matches)
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

func BuildWidget(client *reddit.Client, matches []vlr.UpcomingMatch) error {
	sidebarMd, err := format(false, matches)
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
