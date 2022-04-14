package main

import (
	"fmt"
	"time"

	"github.com/Sadzeih/valcompbot/vlr"
	"github.com/hako/durafmt"
)

const (
	tickerMd = `
%s
---
`

	eventMd = `
| **%s** ||
|:-------:|-----:|
%s

`

	matchMdFmt = "| [%s](%s) | %s |\n"
)

func format(sidebar bool, matches []vlr.UpcomingMatch) (string, error) {
	var matchesMd string

	eventMatches := make(map[string][]vlr.UpcomingMatch)
	for _, match := range matches {
		eventName := match.EventName
		if _, ok := eventMatches[eventName]; !ok {
			eventMatches[eventName] = make([]vlr.UpcomingMatch, 1)
			eventMatches[eventName][0] = match
			continue
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
				dura, err := durafmt.ParseStringShort(time.Until(matchTime).String())
				if err != nil {
					return "", err
				}

				if dura.Duration().Hours() <= 2 {
					dura = dura.LimitFirstN(2)
				}
				if dura.Duration().Hours() < 1 {
					dura = dura.LimitFirstN(1)
				}
				startingIn = dura.String()
			}

			matchStr := fmt.Sprintf("%s vs %s", match.Teams[0].Name, match.Teams[1].Name)
			if match.Teams[0].Name == "" || match.Teams[1].Name == "" {
				// TODO: replace with eventual match name
				matchStr = "TBD"
			}

			matchesMd += fmt.Sprintf(matchMdFmt,
				matchStr,
				match.MatchLink,
				startingIn,
			)
		}
		eventsMd += fmt.Sprintf(eventMd, eventName, matchesMd)
	}

	md := fmt.Sprintf(tickerMd, eventsMd)
	if sidebar {
		return fmt.Sprintf(sidebarMd, md), nil
	}
	return md, nil
}
