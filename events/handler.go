package events

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/scheduledmatch"
	"github.com/Sadzeih/valcompbot/ent/trackedevent"
	"github.com/Sadzeih/valcompbot/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type EventsHandler struct {
	client *ent.Client
	ctx    context.Context
}

func NewHandler(ctx context.Context, client *ent.Client) *EventsHandler {
	return &EventsHandler{
		client: client,
		ctx:    ctx,
	}
}

func (h *EventsHandler) HandleTrackEvent(w http.ResponseWriter, r *http.Request) {
	jReq := struct {
		EventID int    `json:"event_id"`
		Name    string `json:"name"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&jReq); err != nil {
		br := utils.BadRequestError
		br.Context = err.Error()
		utils.WriteError(w, br)
		return
	}

	e, err := h.client.TrackedEvent.
		Create().
		SetEventID(jReq.EventID).
		SetName(jReq.Name).
		Save(h.ctx)
	if err != nil {
		log.Panic(err)
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	utils.WriteJSON(w, e)
}

func (h *EventsHandler) HandleGetTrackedEvents(w http.ResponseWriter, r *http.Request) {
	e, err := h.client.TrackedEvent.Query().WithScheduledmatches(
		func(smq *ent.ScheduledMatchQuery) {
			smq.
				Where(
					scheduledmatch.PostedAtIsNil(),
				)
		},
	).
		All(h.ctx)
	if err != nil {
		log.Panic(err)
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	utils.WriteJSON(w, e)
}

func (h *EventsHandler) HandleDeleteTrackedEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teID, ok := vars["eventID"]
	if !ok || teID == "" {
		br := utils.BadRequestError
		br.Context = "eventID must be provided"
		utils.WriteError(w, br)
		return
	}

	id, err := uuid.Parse(teID)
	if err != nil {
		br := utils.BadRequestError
		br.Context = err.Error()
		utils.WriteError(w, br)
		return
	}

	_, err = h.client.TrackedEvent.Query().Where(trackedevent.ID(id)).Only(h.ctx)
	switch {
	case err != nil && ent.IsNotFound(err):
		utils.WriteError(w, utils.NotFoundError)
		return
	case err != nil:
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	err = h.client.TrackedEvent.DeleteOneID(id).Exec(h.ctx)
	if err != nil {
		utils.WriteError(w, utils.InternalServerError)
		return
	}
}
