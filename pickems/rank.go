package pickems

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sadzeih/valcompbot/config"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"net/http"
)

const (
	apiURL          = "https://api.vlr.gg"
	rankEndpointFmt = `/pickem/rank?reddit_id=%s&event_id=%d`
	tokenFmt        = "&token=%s"
)

func (s *Service) RankComment(comment *reddit.Comment) error {
	if Event == nil {
		return fmt.Errorf("event is nil")
	}
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(apiURL+rankEndpointFmt+tokenFmt, comment.Author, *Event, config.Get().VLRToken),
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

	if rankResp.Rank.Global.Absolute == nil {
		_, _, err = s.redditClient.Comment.Submit(context.Background(), comment.FullID, fmt.Sprintf(unrankedCommentMd, *Event))
		if err != nil {
			return fmt.Errorf("failed creating response to !rank command: %w", err)
		}
		return nil
	}

	rankComment, err := format(rankResp)
	if err != nil {
		return fmt.Errorf("could not format !rank comment: %w", err)
	}

	_, _, err = s.redditClient.Comment.Submit(context.Background(), comment.FullID, fmt.Sprintf(rankComment, *Event))
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
