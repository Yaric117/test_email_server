package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

const loggerTitle = "[config]"

var App app
var DB db

type app struct {
	Name     string `env:"APP_NAME" envDefault:"Go app"`
	Token    string `env:"APP_TOKEN,required"`
	Url      string `env:"APP_URL" envDefault:"http://localhost"`
	Protocol string `env:"APP_PROTOCOL" envDefault:"http"`
	Mode     string `env:"APP_MODE" envDefault:"development" envDescription:"production|development|local"`
	Cors     struct {
		AllowedOrigins string `env:"APP_CORS_ALLOWED_ORIGINS" envDefault:"*"`
	}
}

type db struct {
	Host             string `env:"DB_HOST,required"`
	Port             string `env:"DB_PORT,required"`
	Database         string `env:"DB_DATABASE,required"`
	Username         string `env:"DB_USERNAME,required"`
	Password         string `env:"DB_PASSWORD,required"`
	MigrationsPath   string `env:"DB_MIGRATIONS_PATH" envDefault:"/app/src/db/migrations"`
	MigrateOnStartup bool   `env:"DB_MIGRATE_ON_STARTUP" envDefault:"false"`
}

func init() {
	if err := env.Parse(&App); err != nil {
		panic(fmt.Sprintf("%s[app]: %s", loggerTitle, err))
	}

	if err := env.Parse(&DB); err != nil {
		panic(fmt.Sprintf("%s[DB]: %s", loggerTitle, err))
	}
}
