package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
<<<<<<< Updated upstream
	Port               string `env:"PORT"`
	RedditSubreddit    string `env:"REDDIT_SUBREDDIT"`
	RedditClientID     string `env:"REDDIT_CLIENT_ID"`
	RedditClientSecret string `env:"REDDIT_CLIENT_SECRET"`
	RedditUsername     string `env:"REDDIT_USERNAME"`
	RedditPassword     string `env:"REDDIT_PASSWORD"`
	VLRToken           string `env:"VLR_TOKEN"`
	EnableSentinels    bool   `env:"ENABLE_SENTINELS" envDefault:"false"`
	EnableStickies     bool   `env:"ENABLE_STICKIES" envDefault:"false"`
	PostgresString     string `env:"POSTGRES_STRING"`
	AllowOrigin        string `env:"ALLOW_ORIGIN"`
=======
	Port                    string `env:"PORT"`
	SigningKey              string `env:"SIGNING_KEY"`
	RedditSubreddit         string `env:"REDDIT_SUBREDDIT"`
	RedditClientID          string `env:"REDDIT_CLIENT_ID"`
	RedditClientSecret      string `env:"REDDIT_CLIENT_SECRET"`
	RedditClientScopes      string `env:"REDDIT_CLIENT_SCOPES"`
	RedditClientRedirectURL string `env:"REDDIT_CLIENT_REDIRECT_URL"`
	RedditUsername          string `env:"REDDIT_USERNAME"`
	RedditPassword          string `env:"REDDIT_PASSWORD"`
	VLRToken                string `env:"VLR_TOKEN"`
	EnableSentinels         string `env:"ENABLE_SENTINELS"`
	PostgresString          string `env:"POSTGRES_STRING"`
	AllowOrigin             string `env:"ALLOW_ORIGIN"`
>>>>>>> Stashed changes
}

var (
	cfg *Config
)

func Parse() error {
	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		return err
	}
	return nil
}

func Get() *Config {
	return cfg
}
