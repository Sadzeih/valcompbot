package main

import (
	"fmt"
	"time"

	"github.com/Sadzeih/valcompbot/vlr"
	"github.com/hako/durafmt"
)

const (
	tickerMd = `
| Match | Starting in |
|-------|-------------|
%s

---
`
	sidebarMd  = "## Upcoming matches\n\n" + tickerMd
	matchMdFmt = "| [%s vs %s](%s) | %s |\n"
)

func format(sidebar bool, matches []vlr.UpcomingMatch) (string, error) {
	var matchesMd string
	for _, match := range matches {
		startingIn := ""
		if match.Match != nil {
			if len(match.Match.Streams) != 0 {
				startingIn = fmt.Sprintf("[LIVE](%s)", match.Match.Streams[0].Link)
			} else {
				startingIn = "LIVE"
			}
		} else {
			dura, err := durafmt.ParseStringShort(time.Until(time.Time(match.Timestamp)).String())
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

		team1 := match.Teams[0].Name
		if team1 == "" {
			team1 = "TBD"
		}

		team2 := match.Teams[1].Name
		if team2 == "" {
			team2 = "TBD"
		}
		matchesMd += fmt.Sprintf(matchMdFmt,
			team1,
			team2,
			match.MatchLink,
			startingIn,
		)
	}

	md := sidebarMd
	if !sidebar {
		md = tickerMd
	}
	return fmt.Sprintf(md,
		matchesMd,
	), nil
}
