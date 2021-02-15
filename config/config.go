package config

import (
	"log"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config represents configuration variables
type Config struct {
	Server struct {
		ENV   string `envconfig:"APP_ENV" required:"true"`
		NAME  string `envconfig:"APP_NAME" required:"true"`
		PORT  string `envconfig:"APP_PORT" required:"true"`
		DEBUG bool   `envconfig:"APP_DEBUG" default:"false"`
	}

	JWT struct {
		SigningKey      string `envconfig:"JWT_SIGNING_KEY" required:"true"`
		TokenExpiration int    `envconfig:"JWT_TOKEN_EXPIRATION" required:"true"`
	}

	Database struct {
		Host     string `envconfig:"DB_HOST" required:"true"`
		Port     string `envconfig:"DB_PORT" required:"true"`
		Username string `envconfig:"DB_USERNAME" required:"true"`
		Password string `envconfig:"DB_PASSWORD" required:"true"`
		DBName   string `envconfig:"DB_NAME" required:"true"`
	}

	Google struct {
		ClientID         string `envconfig:"OAUTH_CLIENT_ID" required:"true"`
		ClientSecret     string `envconfig:"OAUTH_CLIENT_SECRET" required:"true"`
		OauthRedirectURL string `envconfig:"OAUTH_REDIRECT_URL" required:"true"`
	}
}

// LoadTest loads test config
func LoadTest() Config {
	return load("test", ".env.test")
}

// LoadDefault loads default config (default.yml) and override config with env if supplied
func LoadDefault() Config {
	return load("default", ".env")
}

// load config and populate to config struct
func load(file string, env string) Config {
	var config Config

	readEnv(&config, env)
	err := envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}
	return config
}

func readEnv(cfg *Config, env string) {
	err := godotenv.Overload(getSourcePath() + "/../" + env)
	if err != nil {
		log.Print(err)
	}
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
