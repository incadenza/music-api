package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/incadenza/music-api/interal/data"
	"github.com/jackc/pgx/v4/pgxpool"
)

func setupDB(dsn string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return dbpool, err
}

type config struct {
	dsn            string
	port           string
	trustedOrigins []string
}

type application struct {
	config config
	models *data.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.dsn, "dsn", "", "PostgreSQL connection string")
	flag.StringVar(&cfg.port, "port", "8080", "port")

	flag.Func("trusted-origins", "Space seperated list of trusted origins", func(val string) error {
		cfg.trustedOrigins = strings.Fields(val)
		return nil
	})
	flag.Parse()

	dbpool, err := setupDB(cfg.dsn)
	models := data.NewModels(dbpool)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	app := application{
		config: cfg,
		models: models,
	}

	port := fmt.Sprintf(":%s", cfg.port)
	http.ListenAndServe(port, app.routes())
}
