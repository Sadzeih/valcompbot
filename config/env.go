package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port               string `env:"PORT"`
	RedditSubreddit    string `env:"REDDIT_SUBREDDIT"`
	RedditClientID     string `env:"REDDIT_CLIENT_ID"`
	RedditClientSecret string `env:"REDDIT_CLIENT_SECRET"`
	RedditUsername     string `env:"REDDIT_USERNAME"`
	RedditPassword     string `env:"REDDIT_PASSWORD"`
	VLRToken           string `env:"VLR_TOKEN"`
	PostgresString     string `env:"POSTGRES_STRING"`
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
