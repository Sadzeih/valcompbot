package vlr

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sadzeih/valcompbot/config"
)

const (
	apiURL              = "https://api.vlr.gg"
	upcomingMatchesPath = "/matches/upcoming"
	tokenFmt            = "?token="
)

func GetUpcomingMatches() ([]UpcomingMatch, error) {
	url := fmt.Sprintf("%s%s%s%s", apiURL, upcomingMatchesPath, tokenFmt, config.Get().VLRToken)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating upcoming matches request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned non-2XX code: %d", resp.StatusCode)
	}

	upcomingMatches := struct {
		Matches []UpcomingMatch `json:"matches"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&upcomingMatches); err != nil {
		return nil, fmt.Errorf("error unmarshalling matches: %w", err)
	}

	return upcomingMatches.Matches, nil
}
