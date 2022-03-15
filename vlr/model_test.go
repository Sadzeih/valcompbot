package vlr

import (
	"encoding/json"
	"testing"
)

const (
	matchesJson = `
[
    {
      "id": "33933",
      "timestamp": "1647547200",
      "tentative": "0",
      "event_name": "Champions Tour North America Stage 1: Challengers ",
      "event_icon_url": "https://owcdn.net/img/6009f963577f4.png",
      "event_tier": "ignition",
      "match_link": "https://www.vlr.gg/77785/cloud9-vs-knights-champions-tour-north-america-stage-1-challengers-ubqf",
      "match_type": "bo3",
      "winner_id": null,
      "forfeit": "0",
      "teams": [
        {
          "id": "188",
          "name": "Cloud9",
          "tag": "C9",
          "country": "US",
          "logo_url": "https://owcdn.net/img/60cedb25c2016.png",
          "maps_won": "0"
        },
        {
          "id": "1615",
          "name": "Knights",
          "tag": "PK",
          "country": "US",
          "logo_url": "https://owcdn.net/img/604022f368433.png",
          "maps_won": "0"
        }
      ]
    }
]
`
)

func TestUnmarshal(t *testing.T) {
	matches := make([]UpcomingMatch, 0)
	if err := json.Unmarshal([]byte(matchesJson), &matches); err != nil {
		t.Errorf("could not unmarshal: %s", err)
	}
	t.Logf("%+v\n", matches)
}
