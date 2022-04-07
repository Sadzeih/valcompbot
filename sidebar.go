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
