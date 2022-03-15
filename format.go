package main

import (
	"fmt"
	"github.com/Sadzeih/valcompbot/vlr"
	"time"
)

const (
	sidebarMd = `
## Upcoming Matches

| Teams | Starting in |
|-------|-------------|
%s

---
`
	matchMdFmt = "| %s %s - %s %s | %s |\n"
)

func Format(matches []vlr.UpcomingMatch) string {
	var matchesMd string
	for _, match := range matches {
		matchesMd += fmt.Sprintf(matchMdFmt,
			match.Teams[0].Name,
			match.Teams[0].MapsWon,
			match.Teams[1].MapsWon,
			match.Teams[1].Name,
			time.Until(time.Time(match.Timestamp)),
		)
	}

	return fmt.Sprintf(sidebarMd,
		matchesMd,
	)
}
