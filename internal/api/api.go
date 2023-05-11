package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/events"
	"github.com/Sadzeih/valcompbot/matches"
	"github.com/Sadzeih/valcompbot/oauth2"

	"github.com/gorilla/mux"
	"github.com/purini-to/zapmw"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
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

	oauthRepo := oauth2.New(
		redditClient,
		config.Get().SigningKey,
		config.Get().RedditClientID,
		config.Get().RedditClientSecret,
		strings.Split(config.Get().RedditClientScopes, ","),
		config.Get().RedditClientRedirectURL,
	)

	r.HandleFunc("/oauth/authorize/url", oauthRepo.AuthCodeURLHandler).
		Methods(http.MethodGet)
	r.HandleFunc("/oauth/callback", oauthRepo.Callback).
		Methods(http.MethodGet)

	ar := mux.NewRouter()

	eventsHandler := events.NewHandler(ctx, entClient)
	ar.HandleFunc("/event", eventsHandler.HandleTrackEvent).
		Methods(http.MethodPost)
	ar.HandleFunc("/event/{eventID}", eventsHandler.HandleDeleteTrackedEvent).
		Methods(http.MethodDelete)
	ar.HandleFunc("/events", eventsHandler.HandleGetTrackedEvents)

	matchesHandler := matches.NewHandler(ctx, redditClient, entClient)
	ar.HandleFunc("/matches/{eventID}", matchesHandler.HandleGetByEventID)
	ar.HandleFunc("/match/{ID}", matchesHandler.HandleGetMatch).
		Methods(http.MethodGet)
	ar.HandleFunc("/match/{ID}", matchesHandler.HandlePostMatch).
		Methods(http.MethodPost)

	an := negroni.New(negroni.HandlerFunc(oauthRepo.Middleware), negroni.Wrap(ar))
	r.PathPrefix("/").Handler(an)

	n := negroni.New(negroni.NewRecovery())

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
	n.UseHandler(c)

	if err := http.ListenAndServe(":8080", n); err != nil {
		return err
	}

	return nil
}
