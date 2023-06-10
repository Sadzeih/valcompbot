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
	EnableSentinels    bool   `env:"ENABLE_SENTINELS" envDefault:"false"`
	EnableStickies     bool   `env:"ENABLE_STICKIES" envDefault:"false"`
	EnablePickems      bool   `env:"ENABLE_PICKEMS" envDefault:"true"`
	PostgresString     string `env:"POSTGRES_STRING"`
	AllowOrigin        string `env:"ALLOW_ORIGIN"`
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
