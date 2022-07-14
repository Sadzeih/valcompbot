package matches

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math"

	"github.com/Sadzeih/valcompbot/config"
)

const (
	apiURL              = "https://api.vlr.gg"
	upcomingMatchesPath = "/matches/upcoming"
	matchPath           = "/match/%s"
	eventMatchlistPath  = "/matchlist/%s"
	tokenFmt            = "?token=%s&tier=riot"
)

func GetUpcoming() ([]Upcoming, error) {
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
		Matches []Upcoming `json:"matches"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&upcomingMatches); err != nil {
		return nil, fmt.Errorf("GetUpcomingMatches: error unmarshalling upcomingMatches.Matches: %w", err)
	}

	for idx, match := range upcomingMatches.Matches {
		matchRes, err := GetMatch(match.ID)
		if err != nil {
			return nil, err
		}
		upcomingMatches.Matches[idx].Match = matchRes
	}

	return upcomingMatches.Matches[:int(math.Min(float64(len(upcomingMatches.Matches)), float64(15)))], nil
}

func GetMatch(id string) (*Match, error) {
	url := fmt.Sprintf(apiURL+matchPath+tokenFmt, id, config.Get().VLRToken)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("GetMatch: error creating match request: %w", err)
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
		return nil, fmt.Errorf("GetMatch: error unmarshalling Match: %w", err)
	}

	return match, nil
}

func GetByEventID(id string) ([]EventMatch, error) {
	url := fmt.Sprintf(apiURL+eventMatchlistPath+tokenFmt, id, config.Get().VLRToken)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("GetByEventID: error creating event matchlist request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetByEventID: error doing request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetByEventID: request returned non-2XX code: %d", resp.StatusCode)
	}

	m := make([]EventMatch, 0)
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, fmt.Errorf("GetByEventID: error unmarshalling []Match: %w", err)
	}

	return m, nil
}
