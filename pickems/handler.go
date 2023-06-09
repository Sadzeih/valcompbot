package pickems

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/utils"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"net/http"
	"time"
)

const (
	leaderboardEndpointFmt = `/pickem/leaderboard?event_id=%d&group_id=valcomp`
)

type Handler struct {
	reddit *reddit.Client
	ent    *ent.Client
	ctx    context.Context
}

func NewHandler(ctx context.Context, r *reddit.Client, e *ent.Client) *Handler {
	return &Handler{
		reddit: r,
		ent:    e,
		ctx:    ctx,
	}
}

func (h *Handler) SetEvent(w http.ResponseWriter, r *http.Request) {
	jReq := struct {
		EventID *int `json:"event_id,omitempty"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&jReq); err != nil {
		br := utils.BadRequestError
		br.Context = err.Error()
		utils.WriteError(w, br)
		return
	}

	builder := h.ent.PickemsEvent.Create()

	if jReq.EventID != nil {
		builder = builder.SetEventID(*jReq.EventID)
	}

	e, err := builder.
		SetTimestamp(time.Now()).
		Save(r.Context())
	if err != nil {
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	Event = e.EventID

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetLeaderboards(w http.ResponseWriter, r *http.Request) {
	if Event == nil {
		utils.WriteError(w, utils.APIError{
			Code:    http.StatusNotFound,
			Message: "Event not found",
		})
		return
	}
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(apiURL+leaderboardEndpointFmt+tokenFmt, *Event, config.Get().VLRToken),
		nil,
	)
	if err != nil {
		ie := utils.InternalServerError
		ie.Context = fmt.Errorf("could not create request: %w", err).Error()
		utils.WriteError(w, ie)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ie := utils.InternalServerError
		ie.Context = fmt.Errorf("error while executing request: %w", err).Error()
		utils.WriteError(w, ie)
		return
	}

	if resp.StatusCode != 200 {
		ie := utils.InternalServerError
		ie.Context = fmt.Errorf("request returned non-200: %d", resp.StatusCode).Error()
		utils.WriteError(w, ie)
		return
	}

	leaderboardsResp := &LeaderboardsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&leaderboardsResp); err != nil {
		ie := utils.InternalServerError
		ie.Context = fmt.Errorf("could not decode json: %w", err).Error()
		utils.WriteError(w, ie)
		return
	}

	result := bytes.Buffer{}
	err = leaderboardsTmpl.Execute(&result, leaderboardsResp)
	if err != nil {
		ie := utils.InternalServerError
		ie.Context = fmt.Errorf("could not parse template: %w", err).Error()
		utils.WriteError(w, ie)
		return
	}

	w.Write(result.Bytes())
}
