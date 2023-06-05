package api

import (
	"context"
	"github.com/Sadzeih/valcompbot/pickems"
	"net/http"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/events"
	"github.com/Sadzeih/valcompbot/matches"
	"github.com/gorilla/mux"
	"github.com/purini-to/zapmw"
	"github.com/rs/cors"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Start(redditClient *reddit.Client, entClient *ent.Client) error {
	r := mux.NewRouter()
	logger, _ := zap.NewProduction()
	r.Use(
		zapmw.WithZap(logger),
		zapmw.Request(zapcore.InfoLevel, "request"),
		zapmw.Recoverer(zapcore.ErrorLevel, "recover", zapmw.RecovererDefault),
	)

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

	pickemsHandler := pickems.NewHandler(ctx, redditClient, entClient)
	r.HandleFunc("/pickems/event", pickemsHandler.SetEvent).
		Methods(http.MethodPost)
	r.HandleFunc("/pickems/leaderboards", pickemsHandler.GetLeaderboards).
		Methods(http.MethodGet)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{config.Get().AllowOrigin},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}).Handler(r)

	if err := http.ListenAndServe(":8080", c); err != nil {
		return err
	}

	return nil
}
