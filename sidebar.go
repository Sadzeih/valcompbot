package main

import (
	"context"
	"github.com/Sadzeih/valcompbot/vlr"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func BuildSidebar(client *reddit.Client, subSettings *reddit.SubredditSettings) error {
	matches, err := vlr.GetUpcomingMatches()
	if err != nil {
		return err
	}

	sidebarMd := Format(matches)
	subSettings.Sidebar = &sidebarMd

	_, err = client.Subreddit.Edit(context.Background(), subSettings.ID, subSettings)
	if err != nil {
		return err
	}

	return nil
}
