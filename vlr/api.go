package vlr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sadzeih/valcompbot/config"
)

const (
	apiURL              = "https://api.vlr.gg"
	upcomingMatchesPath = "/matches/upcoming"
	matchPath           = "/match/%s"
	tokenFmt            = "?token=%s&tier=riot"
)

func GetUpcomingMatches() ([]UpcomingMatch, error) {
	url := fmt.Sprintf(apiURL+upcomingMatchesPath+tokenFmt, config.Get().VLRToken)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("GetUpcomingMatches: error creating upcoming upcomingMatches.Matches request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetUpcomingMatches: error doing request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetUpcomingMatches: request returned non-2XX code: %d", resp.StatusCode)
	}

	upcomingMatches := struct {
		Matches []UpcomingMatch `json:"matches"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&upcomingMatches); err != nil {
		return nil, fmt.Errorf("GetUpcomingMatches: error unmarshalling upcomingMatches.Matches: %w", err)
	}

	for idx, match := range upcomingMatches.Matches {
		matchTime := time.Time(match.Timestamp)
		if matchTime.Before(time.Now()) {
			matchRes, err := GetMatch(match.ID)
			if err != nil {
				return nil, err
			}
			upcomingMatches.Matches[idx].Match = matchRes
		}
	}

	return upcomingMatches.Matches, nil
}

func GetMatch(id string) (*Match, error) {
	url := fmt.Sprintf(apiURL+matchPath+tokenFmt, id, config.Get().VLRToken)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("GetMatch: error creating upcoming upcomingMatches.Matches request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetMatch: error doing request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetMatch: request returned non-2XX code: %d", resp.StatusCode)
	}

	match := &Match{}
	if err := json.NewDecoder(resp.Body).Decode(&match); err != nil {
		return nil, fmt.Errorf("GetMatch: error unmarshalling upcomingMatches.Matches: %w", err)
	}

	return match, nil
}
