package matches

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/trackedevent"
	"github.com/Sadzeih/valcompbot/utils"
	"github.com/gorilla/mux"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/exp/slices"
)

const (
	pmtFlairID = "76154df2-c8c8-11ec-8e35-c21565c22f2e"
)

type MatchesHandler struct {
	reddit *reddit.Client
	ent    *ent.Client
	ctx    context.Context
}

func NewHandler(ctx context.Context, r *reddit.Client, e *ent.Client) *MatchesHandler {
	return &MatchesHandler{
		reddit: r,
		ent:    e,
		ctx:    ctx,
	}
}

func (h *MatchesHandler) HandleGetByEventID(w http.ResponseWriter, r *http.Request) {
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

func (h *MatchesHandler) HandleGetMatch(w http.ResponseWriter, r *http.Request) {
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

func (h *MatchesHandler) HandlePostMatch(w http.ResponseWriter, r *http.Request) {
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

	eID, err := strconv.Atoi(m.Info.EventID)
	if err != nil {
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	e, err := h.ent.TrackedEvent.
		Query().
		Where(trackedevent.EventID(eID)).
		Only(h.ctx)
	if err != nil {
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	title := fmt.Sprintf(titleFmt, m.Teams[0].Name, m.Teams[1].Name, e.Name, m.Info.Series)
	if bodyJson.Title != "" {
		title = bodyJson.Title
	}

	md, err := m.ToMarkdown()
	if err != nil {
		log.Println(err)
		utils.WriteError(w, utils.InternalServerError)
		return
	}

	_, _, err = h.reddit.Post.SubmitText(h.ctx, reddit.SubmitTextRequest{
		Subreddit: config.Get().RedditSubreddit,
		Title:     title,
		Text:      md,
		FlairID:   pmtFlairID,
		Spoiler:   true,
	})
	if err != nil {
		log.Println(err)
		utils.WriteError(w, utils.InternalServerError)
		return
	}
}
