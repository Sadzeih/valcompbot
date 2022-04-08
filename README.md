# [/r/ValorantCompetitive](https://reddit.com/r/ValorantCompetitive) Bot

This bot manages the sidebar and widgets of /r/ValorantCompetitive.

Thanks to [VLR](https://vlr.gg) for the access to their API, and facilitating by adding features to some endpoints.

## Running the bot

You must provide these environment variables
```env
REDDIT_SUBREDDIT="SUBREDDIT"
REDDIT_CLIENT_ID="CLIENT_ID"
REDDIT_CLIENT_SECRET="CLIENT_SECRET"
REDDIT_USERNAME="BOT_USERNAME"
REDDIT_PASSWORD="PASSWORD"
VLR_TOKEN="VLR_TOKEN"
```

and start the bot by running:

```bash
$ go run .
```

## Development

Development was made using [go-reddit](https://github.com/vartanbeno/go-reddit), 
which needed missing features, so a [PR was opened](https://github.com/vartanbeno/go-reddit/pull/32).
