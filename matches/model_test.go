package matches

import (
	"encoding/json"
	"os"
	"strings"
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

func TestToMarkdown(t *testing.T) {
	data, err := os.ReadFile("../testdata/example_match.json")
	if err != nil {
		t.Fatal(err)
	}

	m := &Match{}
	if err := json.Unmarshal(data, m); err != nil {
		t.Fatal(err)
	}

	md, err := m.ToMarkdown()
	if err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile("../testdata/test_match_output.md", []byte(md), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestToMarkdownWithRibLink(t *testing.T) {
	data, err := os.ReadFile("../testdata/example_match.json")
	if err != nil {
		t.Fatal(err)
	}

	m := &Match{}
	if err := json.Unmarshal(data, m); err != nil {
		t.Fatal(err)
	}

	m.RibLink = "https://rib.gg/matches/990/paper-rex-vs-team-liquid"

	md, err := m.ToMarkdown()
	if err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile("../testdata/test_match_output_rib.md", []byte(md), 0644); err != nil {
		t.Fatal(err)
	}

	want := "[vlr.gg](" + m.Info.Link + ") | [RIB.gg](" + m.RibLink + ")"
	if !strings.Contains(md, want) {
		t.Errorf("expected markdown to contain %q", want)
	}
}

func TestUnmarshal(t *testing.T) {
	matches := make([]Upcoming, 0)
	if err := json.Unmarshal([]byte(matchesJson), &matches); err != nil {
		t.Errorf("could not unmarshal: %s", err)
	}
	t.Logf("%+v\n", matches)
}
