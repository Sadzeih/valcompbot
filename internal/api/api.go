package api

import (
	"context"
	"net/http"

	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/events"
	"github.com/Sadzeih/valcompbot/matches"
	"github.com/gorilla/mux"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func Start(redditClient *reddit.Client, entClient *ent.Client) error {
	r := mux.NewRouter()

	ctx := context.Background()

	eventsHandler := events.NewHandler(ctx, entClient)
	r.HandleFunc("/event", eventsHandler.HandleTrackEvent).
		Methods(http.MethodPost)
	r.HandleFunc("/event/{eventID}", eventsHandler.HandleDeleteTrackedEvent).
		Methods(http.MethodDelete)
	r.HandleFunc("/events", eventsHandler.HandleGetTrackedEvents)

	matchesHandler := matches.NewHandler(ctx, redditClient, entClient)
	r.HandleFunc("/matches/{eventID}", matchesHandler.HandleGetByEventID)
	r.HandleFunc("/match/{ID}", matchesHandler.HandleGetMatch).
		Methods(http.MethodGet)
	r.HandleFunc("/match/{ID}", matchesHandler.HandlePostMatch).
		Methods(http.MethodPost)

	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}

	return nil
}
