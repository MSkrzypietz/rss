package main

import (
	"context"
	"database/sql"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/joho/godotenv"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type application struct {
	logger     *slog.Logger
	db         *database.Queries
	httpClient *http.Client
}

func main() {
	logFile, err := os.OpenFile("/app/rss-api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		return
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := slog.New(slog.NewTextHandler(multiWriter, nil))

	if !isProductionEnv() {
		if err := godotenv.Load(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		logger.Error("HTTP_PORT is undefined")
		os.Exit(1)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		logger.Error("DB_URL is undefined")
		os.Exit(1)
	}

	db, err := openDB(dbURL)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:     logger,
		db:         database.New(db),
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
	go app.ContinuousFeedScraping()

	srv := http.Server{
		Addr:    ":" + httpPort,
		Handler: app.routes(),
	}

	logger.Info("starting srv", "addr", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}

func isProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}

func openDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("libsql", dbURL)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
