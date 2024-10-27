package matches

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/scheduledmatch"
	"github.com/Sadzeih/valcompbot/ent/trackedevent"
	"github.com/Sadzeih/valcompbot/utils"
	"github.com/gorilla/mux"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/exp/slices"
)

const (
	pmtFlairID = "76154df2-c8c8-11ec-8e35-c21565c22f2e"
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

func (h *Handler) HandleGetByEventID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	eventID, ok := vars["eventID"]
	if !ok || eventID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m, err := GetByEventID(eventID)
	if err != nil {
		br := utils.BadRequestError
		br.Context = err.Error()
		utils.WriteError(w, br)
		return
	}

	filtered := make([]EventMatch, 0)

	series := r.FormValue("series")
	if series != "" {
		for _, match := range m {
			if match.Series == series {
				filtered = append(filtered, match)
			}
		}
		m = filtered
	}

	filtered = make([]EventMatch, 0)
	subseries := r.FormValue("subseries")
	if subseries != "" {
		for _, match := range m {
			if match.Subseries == subseries {
				filtered = append(filtered, match)
			}
		}
		m = filtered
	}

	filtered = make([]EventMatch, 0)
	upcoming := r.FormValue("upcoming")
	if upcoming != "" {
		for _, match := range m {
			if upcoming == "true" && time.Time(match.Timestamp).After(time.Now()) {
				filtered = append(filtered, match)
			} else if upcoming == "false" && time.Time(match.Timestamp).Before(time.Now()) {
				filtered = append(filtered, match)
			}
		}
		m = filtered
	}

	slices.SortFunc(m, func(a, b EventMatch) bool {
		return time.Time(a.Timestamp).After(time.Time(b.Timestamp))
	})

	utils.WriteJSON(w, m)
}

func (h *Handler) HandleGetMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["ID"]
	if !ok || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m, err := GetMatch(id)
	if err != nil {
		br := utils.BadRequestError
		br.Context = err.Error()
		utils.WriteError(w, br)
		return
	}

	utils.WriteJSON(w, m)
}

func (h *Handler) HandlePostMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["ID"]
	if !ok || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bodyJson := struct {
		Title string `json:"title"`
	}{}

	if r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(&bodyJson); err != nil {
			br := utils.BadRequestError
			br.Context = err.Error()
			utils.WriteError(w, br)
			return
		}
	}

	m, err := GetMatch(id)
	if err != nil {
		br := utils.BadRequestError
		br.Context = err.Error()
		utils.WriteError(w, br)
		return
	}

	if err := PostMatch(r.Context(), m, h.ent, h.reddit); err != nil {
		log.Println(err)
		utils.WriteError(w, utils.InternalServerError)
		return
	}
}

func (h *Handler) HandleSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["ID"]
	if !ok || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m, err := h.ent.ScheduledMatch.Query().
		Where(
			scheduledmatch.MatchID(id),
		).Only(r.Context())
	if err != nil {
		if !ent.IsNotFound(err) {
			utils.WriteError(w, utils.InternalServerError)
			return
		}
	}
	if m != nil {
		br := utils.BadRequestError
		br.Context = "Match was already scheduled"
		utils.WriteError(w, br)
		return
	}

	vm, err := GetMatch(id)
	if err != nil {
		br := utils.BadRequestError
		br.Context = err.Error()
		utils.WriteError(w, br)
		return
	}

	eID, err := strconv.Atoi(vm.Info.EventID)
	if err != nil {
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	event, err := h.ent.TrackedEvent.Query().
		Where(
			trackedevent.EventID(eID),
		).Only(r.Context())
	if err != nil {
		log.Println(err)
		utils.WriteError(w, utils.InternalServerError)
		return
	}
	if event == nil {
		br := utils.BadRequestError
		br.Context = "Cannot schedule Match from an event that is not tracked"
		utils.WriteError(w, br)
		return
	}

	_, err = h.ent.ScheduledMatch.Create().
		SetMatchID(id).
		SetEventID(event.ID).
		Save(r.Context())
	if err != nil {
		log.Println(err)
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
