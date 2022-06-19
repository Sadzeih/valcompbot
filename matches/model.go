package matches

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

type Upcoming struct {
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
	Teams []struct {
		ID       string `json:"team_id"`
		Name     string `json:"name"`
		Tag      string `json:"tag"`
		MapsWon  string `json:"maps_won"`
		IsWinner bool   `json:"is_winner"`
	} `json:"teams"`
	Maps []struct {
		Name  string       `json:"name"`
		Link  string       `json:"link"`
		Teams MapTeamStats `json:"teams"`
	} `json:"maps"`
	Totals struct {
		Teams TotalTeamStats `json:"teams"`
	} `json:"totals"`
	Streams []struct {
		Provider    string `json:"provider"`
		Name        string `json:"name"`
		Country     string `json:"country"`
		ViewerCount string `json:"viewer_count"`
		Live        string `json:"live"`
		Link        string `json:"link"`
	} `json:"streams"`
}

type EventMatch struct {
	ID        string    `json:"match_id"`
	Timestamp Timestamp `json:"scheduled_ts"`
	EventName string    `json:"event"`
	Series    string    `json:"series"`
	Subseries string    `json:"subseries"`
	Teams     []struct {
		ID       string `json:"team_id"`
		Name     string `json:"name"`
		Tag      string `json:"tag"`
		MapsWon  string `json:"maps_won"`
		IsWinner bool   `json:"is_winner"`
	} `json:"teams"`
}

type Team struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Country string `json:"country"`
	LogoURL string `json:"logo_url"`
	MapsWon string `json:"maps_won"`
}

type MapTeamStats []struct {
	TeamID           string `json:"team_id"`
	Name             string `json:"name"`
	RoundsWon        string `json:"rounds_won"`
	RoundsWonAttack  int    `json:"rounds_won_atk"`
	RoundsWonDefense int    `json:"rounds_won_def"`
	IsWinner         bool   `json:"is_winner"`
	Players          []struct {
		PlayerID  string `json:"player_id"`
		Alias     string `json:"alias"`
		Country   string `json:"country"`
		TwitterID string `json:"twitter_id"`
		Stats     struct {
			Agent      string `json:"agent"`
			Rounds     string `json:"rounds"`
			Kills      string `json:"kills"`
			Deaths     string `json:"deaths"`
			Assists    string `json:"assists"`
			KDA        string `json:"kda_string"`
			Damage     string `json:"damage"`
			ADR        string `json:"adr"`
			FBPR       string `json:"fbpr"`
			FDPR       string `json:"fdpr"`
			ACS        string `json:"combat_score"`
			EconRating string `json:"econ_rating"`
			FB         string `json:"first_bloods"`
			FD         string `json:"first_deaths"`
			Plants     string `json:"plants"`
			Defuses    string `json:"defuses"`
		} `json:"stats"`
	} `json:"players"`
}

type TotalTeamStats []struct {
	TeamID           string `json:"team_id"`
	Name             string `json:"name"`
	RoundsWon        string `json:"rounds_won"`
	RoundsWonAttack  int    `json:"rounds_won_atk"`
	RoundsWonDefense int    `json:"rounds_won_def"`
	IsWinner         bool   `json:"is_winner"`
	Players          []struct {
		PlayerID string `json:"player_id"`
		Alias    string `json:"alias"`
		Stats    struct {
			Agents     []string `json:"agents"`
			Rounds     string   `json:"rounds"`
			Kills      string   `json:"kills"`
			Deaths     string   `json:"deaths"`
			Assists    string   `json:"assists"`
			KDA        string   `json:"kda_string"`
			Damage     string   `json:"damage"`
			ADR        string   `json:"adr"`
			FBPR       string   `json:"fbpr"`
			FDPR       string   `json:"fdpr"`
			ACS        string   `json:"combat_score"`
			EconRating string   `json:"econ_rating"`
			FB         string   `json:"first_bloods"`
			FD         string   `json:"first_deaths"`
			Plants     string   `json:"plants"`
			Defuses    string   `json:"defuses"`
		} `json:"stats"`
	} `json:"players"`
}
