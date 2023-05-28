package pickems

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sadzeih/valcompbot/config"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"net/http"
	"text/template"
)

const (
	apiURL          = "https://api.vlr.gg"
	rankEndpointFmt = `/pickem/rank?reddit_id=%s&event_id=%s`
	tokenFmt        = "&token=%s"
	rankCommentMd   = `You are ranked [#{{ .Rank.Global }}]({{ .Rank.Link }}) globally on [VLR.gg](https://vlr.gg) with a score of {{ .Rank.Score }}. Which brings you to the top {{ .Rank.Percentile }}% of users!

You are ranked [#{{ .Rank.Local }}]({{ .Rank.Link }}) on the subreddit leaderboards.

[^(Join the subreddit pickems here.)](https://vlr.gg/event/pickem/%s?code=valcomp)`

	unrankedCommentMd = `You either do not have a rank yet or have not joined the subreddit group.

[Join the subreddit pickems here!](https://vlr.gg/event/pickem/%s?code=valcomp)
`
)

var (
	rankTmpl = template.Must(template.New("rankTmpl").Parse(rankCommentMd))
)

type RankResponse struct {
	Rank *struct {
		Global     string `json:"global,omitempty"`
		Local      string `json:"local,omitempty"`
		Absolute   string `json:"absolute,omitempty"`
		Percentile string `json:"percentile,omitempty"`
		Score      string `json:"score,omitempty"`
		Link       string `json:"link,omitempty"`
	} `json:"rank"`
}

func (s *Service) RankComment(comment *reddit.Comment) error {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(apiURL+rankEndpointFmt+tokenFmt, comment.Author, s.eventID, config.Get().VLRToken),
		nil,
	)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error while executing request: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("request returned non-200: %d", resp.StatusCode)
	}

	rankResp := &RankResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&rankResp); err != nil {
		return fmt.Errorf("could not decode json: %w", err)
	}
	if rankResp.Rank == nil {
		_, _, err = s.redditClient.Comment.Submit(context.Background(), comment.FullID, fmt.Sprintf(unrankedCommentMd, s.eventID))
		if err != nil {
			return fmt.Errorf("failed creating response to !rank command: %w", err)
		}
		return nil
	}

	rankComment, err := format(rankResp)
	if err != nil {
		return fmt.Errorf("could not format !rank comment: %w", err)
	}

	_, _, err = s.redditClient.Comment.Submit(context.Background(), comment.FullID, fmt.Sprintf(rankComment, s.eventID))
	if err != nil {
		return fmt.Errorf("failed creating response to !rank command: %w", err)
	}

	return nil
}

func format(rank *RankResponse) (string, error) {
	result := bytes.Buffer{}
	err := rankTmpl.Execute(&result, rank)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
