package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os/exec"
	"testproject/config"
)

var Conn *pgxpool.Pool

func init() {
	var err error

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	Conn, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(fmt.Sprintf("Conn: connect error: %v", err))
	}

	if config.DB.MigrateOnStartup {
		migrate()
	}
}

func migrate() {
	bin := "/usr/local/bin/migrate"
	source := fmt.Sprintf("file://%s", config.DB.MigrationsPath)
	database := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	command := "up"

	cmd := exec.Command(bin, "-source", source, "-database", database, command)

	_, err := cmd.Output()
	if err != nil {
		log.Fatalf("Conn: migrate error: %v", err)
	}
}
