package vlr

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(int64(ts), 0))
	return nil
}

type UpcomingMatch struct {
	ID           string    `json:"id"`
	Timestamp    Timestamp `json:"timestamp"`
	Tentative    string    `json:"tentative"`
	EventName    string    `json:"event_name"`
	EventIconURL string    `json:"event_icon_url"`
	EventTier    string    `json:"event_tier"`
	MatchLink    string    `json:"match_link"`
	MatchType    string    `json:"match_type"`
	WinnerID     *string   `json:"winner_id"`
	Forfeit      string    `json:"forfeit"`
	Teams        []Team    `json:"teams"`
	Match        *Match
}

type Match struct {
	Info struct {
		ID        string `json:"id"`
		EventID   string `json:"event_id"`
		Completed string `json:"completed"`
		Patch     string `json:"patch"`
		Series    string `json:"series"`
		Subseries string `json:"subseries"`
		Link      string `json:"link"`
	} `json:"info"`
	Streams []struct {
		Provider    string `json:"provider"`
		Name        string `json:"name"`
		Country     string `json:"country"`
		ViewerCount string `json:"viewer_count"`
		Live        string `json:"live"`
		Link        string `json:"link"`
	} `json:"streams"`
}

type Team struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Country string `json:"country"`
	LogoURL string `json:"logo_url"`
	MapsWon string `json:"maps_won"`
}
