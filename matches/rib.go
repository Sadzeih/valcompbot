package matches

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ribAPIURL    = "https://rib.gg/api"
	ribMatchPath = "/vlr/%s"
)

type RibMatch struct {
	Found       bool   `json:"found"`
	VlrMatchID  string `json:"vlrMatchId"`
	RibSeriesID string `json:"ribSeriesId"`
	Match       string `json:"match"`
}

// GetRibMatch returns the RIB.gg match page for a VLR.gg match ID.
// It returns nil (and no error) when RIB.gg has no match for the given VLR.gg match ID.
func GetRibMatch(vlrMatchID string) (*RibMatch, error) {
	url := fmt.Sprintf(ribAPIURL+ribMatchPath, vlrMatchID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("GetRibMatch: error creating rib match request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetRibMatch: error doing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetRibMatch: request returned non-2XX code: %d", resp.StatusCode)
	}

	ribMatch := &RibMatch{}
	if err := json.NewDecoder(resp.Body).Decode(ribMatch); err != nil {
		return nil, fmt.Errorf("GetRibMatch: error unmarshalling RibMatch: %w", err)
	}

	if !ribMatch.Found {
		return nil, nil
	}

	return ribMatch, nil
}
