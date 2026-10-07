# AGENT.md

## Project Overview

This is **valcompbot**, a Go Reddit bot that manages the sidebar, widgets, and automated posts for [/r/ValorantCompetitive](https://reddit.com/r/ValorantCompetitive). It pulls match data from [VLR.gg](https://vlr.gg), persists state in PostgreSQL, and exposes a REST API for event tracking and pick'ems leaderboards.

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.23 |
| Reddit API | `github.com/vartanbeno/go-reddit/v2` (forked: `github.com/Sadzeih/go-reddit/v2`) |
| ORM | [ent](https://entgo.io/ent) (code-generated, schema-first) |
| Database | PostgreSQL (`github.com/lib/pq`) |
| HTTP Router | `github.com/gorilla/mux` |
| Logging | `go.uber.org/zap` |
| CORS | `github.com/rs/cors` |
| Config | `github.com/caarlos0/env/v6` + `github.com/joho/godotenv` |
| External API | VLR.gg, Liquipedia |
| Deployment | Docker (distroless base image) |

## Project Structure

```
├── main.go              # Entry point — wires up all goroutines
├── sidebar.go           # Sidebar/widget updater (runs every 1 min)
├── sidebar_md.go        # Markdown formatter for sidebar
├── format.go            # Shared formatting helpers
├── sentinels.go         # "Days since last sentinel post" tracker
├── config/
│   └── env.go           # Environment-based config (env.Parse)
├── ent/
│   └── schema/          # ent schema definitions (source of truth for DB)
│       ├── trackedevent.go
│       ├── scheduledmatch.go
│       ├── pickemsevent.go
│       ├── highlightedcomment.go
│       └── pinnedcomment.go
├── internal/
│   └── api/
│       └── api.go       # HTTP server, route registration, CORS
├── matches/
│   ├── model.go         # VLR.gg match data types
│   ├── vlr.go           # VLR.gg API client
│   ├── handler.go       # HTTP handlers for match endpoints
│   ├── scheduler.go     # PMT (Post-Match Thread) scheduler
│   ├── format.go        # Match-to-markdown formatting
│   ├── post.go          # Reddit post creation
│   └── model_test.go    # Tests (uses testdata/)
├── pickems/
│   ├── pickems.go       # Pick'ems service (subscribes to comments)
│   ├── handler.go       # HTTP handlers for pick'ems
│   ├── format.go        # Pick'ems markdown formatting
│   ├── rank.go          # Leaderboard/ranking logic
│   └── models.go        # Pick'ems data types
├── events/
│   └── handler.go       # HTTP handlers for tracked events CRUD
├── comments/
│   └── comments.go      # Reddit comment stream consumer (pub/sub)
├── highlighter/
│   ├── highlighter.go   # Comment highlighter service
│   └── const.go         # Highlight keywords/rules
├── liquipedia/
│   └── liquipedia.go    # Liquipedia scraper/client
├── utils/
│   ├── error.go         # APIError type + WriteError helper
│   └── json.go          # WriteJSON helper
└── testdata/            # Example JSON fixtures for tests
```

## Key Patterns

### Configuration
All config is environment-driven via `config/env.go`. Required env vars:
- `REDDIT_SUBREDDIT`, `REDDIT_CLIENT_ID`, `REDDIT_CLIENT_SECRET`, `REDDIT_USERNAME`, `REDDIT_PASSWORD`
- `VLR_TOKEN`, `POSTGRES_STRING`, `ALLOW_ORIGIN`, `PORT`
- Optional: `ENABLE_SENTINELS`, `ENABLE_STICKIES`, `ENABLE_PICKEMS` (default: `true`)

A `.env` file is auto-loaded via `godotenv/autoload`. `.env` is gitignored.

### Database (ent)
- Schema definitions live in `ent/schema/*.go` — **edit these, not the generated code**
- After changing schemas, regenerate with: `go generate ./ent`
- The app runs `entClient.Debug().Schema.Create()` on startup with `WithDropIndex(true)` and `WithDropColumn(true)` — schema is auto-migrated (dev-friendly, not for production)
- All entities use UUID primary keys with `StorageKey("oid")`

### Service Architecture
`main.go` starts multiple goroutines coordinated by a `sync.WaitGroup`:
1. **Sidebar ticker** — updates subreddit sidebar + widget every minute
2. **Comments topic** — pub/sub stream of new Reddit comments
3. **PMT scheduler** — schedules Post-Match Threads based on VLR.gg data
4. **HTTP API** — REST endpoints (see below)
5. **Pick'ems** — optional, subscribes to comment stream for predictions
6. **Comment highlighter** — optional, highlights notable comments
7. **Sentinels** — optional, tracks days since last sentinel post

### HTTP API
Routes registered in `internal/api/api.go`:

| Method | Path | Handler |
|---|---|---|
| POST | `/event` | Track a new event |
| DELETE | `/event/{eventID}` | Delete a tracked event |
| GET | `/events` | List tracked events |
| GET | `/matches/{eventID}` | Get matches for an event |
| GET | `/match/{ID}` | Get a specific match |
| POST | `/match/{ID}` | Create/update a match |
| POST | `/match/{ID}/schedule` | Schedule a PMT for a match |
| POST | `/pickems/event` | Set the current pick'ems event |
| GET | `/pickems/leaderboards` | Get pick'ems leaderboards |

Error responses use `utils.APIError` JSON shape: `{ "code": int, "message": string, "context": string? }`.

### External APIs
- **VLR.gg** — match schedules, results, player stats. Auth via `VLR_TOKEN` env var. Client in `matches/vlr.go`.
- **Liquipedia** — team/player data scraping. Client in `liquipedia/liquipedia.go`.

## Development

### Prerequisites
- Go 1.23+
- PostgreSQL (connection string via `POSTGRES_STRING`)
- Reddit API credentials
- VLR.gg API token

### Common Commands

```bash
# Run the bot
go run .

# Run tests
go test -v ./...

# Run a specific test
go test -v ./matches/ -run TestToMarkdown

# Regenerate ent code after schema changes
go generate ./ent

# Build binary
go build -o bot .

# Build Docker image
docker build -t valcompbot .

# Format code
gofmt -w .

# Vet
 go vet ./...
```

### Testing
- Tests live alongside code (`*_test.go`)
- Test fixtures in `testdata/` (e.g., `testdata/example_match.json`)
- CI runs `go test -v ./...` on every push (see `.github/workflows/valcompbot.yml`)

### Code Style
- Standard Go formatting (`gofmt`)
- Package names are lowercase, singular
- Error handling: return errors up the stack, log at the boundary
- HTTP handlers follow the pattern: `NewHandler(ctx, deps...)` → `handler.HandleXxx(w, r)`
- Use `utils.WriteJSON` and `utils.WriteError` for HTTP responses
- Context is passed explicitly (stored in handler structs), not via `context.Background()` in handlers

## Important Notes

- **Forked dependency**: `go.mod` replaces `github.com/vartanbeno/go-reddit/v2` with `github.com/Sadzeih/go-reddit/v2 v2.2.0`. Do not remove the `replace` directive.
- **Schema migrations are destructive in dev**: `WithDropIndex(true)` and `WithDropColumn(true)` will drop indexes and columns on startup. Do not run against production data without adjusting.
- **ent code is generated**: Never edit files under `ent/` directly (except `ent/schema/`). Run `go generate ./ent` after schema changes.
- **Goroutine lifecycle**: All services are fire-and-forget goroutines. There is no graceful shutdown — the process runs until killed.
