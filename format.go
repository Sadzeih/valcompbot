package main

import (
	"fmt"
	"time"

	"github.com/Sadzeih/valcompbot/matches"
)

const (
	tickerMd = `
%s
---
`

	eventMd = `
**%s**

| Match | Starting in |
|:-----:|------------:|
%s

`

	matchMdFmt = "| [%s](%s) | %s |\n"
)

func format(sidebar bool, m []matches.Upcoming) (string, error) {
	var matchesMd string

	eventMatches := make(map[string][]matches.Upcoming)
	for _, match := range m {
		eventName := match.EventName
		if _, ok := eventMatches[eventName]; !ok {
			eventMatches[eventName] = make([]matches.Upcoming, 0)
		}
		eventMatches[eventName] = append(eventMatches[eventName], match)
	}

	eventsMd := ""
	for eventName, matchesInEvent := range eventMatches {
		for _, match := range matchesInEvent {
			startingIn := ""
			matchTime := time.Time(match.Timestamp)
			if matchTime.Before(time.Now()) {
				if len(match.Match.Streams) != 0 {
					startingIn = fmt.Sprintf("[LIVE](%s)", match.Match.Streams[0].Link)
				} else {
					startingIn = "LIVE"
				}
			} else {
				startingIn = formatDuration(time.Until(matchTime))
			}

			matchStr := fmt.Sprintf("%s vs %s", match.Teams[0].Name, match.Teams[1].Name)
			if match.Teams[0].Name == "" || match.Teams[1].Name == "" {
				matchStr = fmt.Sprintf("*%s: %s*", match.Match.Info.Series, match.Match.Info.Subseries)
			}

			matchesMd += fmt.Sprintf(matchMdFmt,
				matchStr,
				match.MatchLink,
				startingIn,
			)
		}
		eventsMd += fmt.Sprintf(eventMd, eventName, matchesMd)
		matchesMd = ""
	}

	md := fmt.Sprintf(tickerMd, eventsMd)
	if sidebar {
		return fmt.Sprintf(sidebarMd, md), nil
	}
	return md, nil
}

func formatDuration(d time.Duration) string {
	od := d
	dstr := ""
	if d.Hours() >= 24 {
		dstr += fmt.Sprintf("%dd", int(d.Hours()/24))
		d -= time.Duration(int(d.Hours()/24)*24) * time.Hour
	}
	if d.Hours() >= 1 {
		dstr += fmt.Sprintf("%dh", int(d.Minutes()/60))
		d -= time.Duration(int(d.Minutes()/60)) * time.Hour
	}
	if od.Hours() <= 2 {
		dstr += fmt.Sprintf("%dm", int(d.Minutes()))
	}

	return dstr
}
