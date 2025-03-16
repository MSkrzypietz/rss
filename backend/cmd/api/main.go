package main

import (
	"database/sql"
	"fmt"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type application struct {
	db         *database.Queries
	httpClient *http.Client
}

func main() {
	if !isProductionEnv() {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatalln("HTTP_PORT is undefined")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalln("DB_URL is undefined")
	}

	db, err := sql.Open("libsql", dbURL)
	if err != nil {
		log.Fatalf("Cannot open database connection: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot ping database: %v", err)
	}

	app := &application{
		db:         database.New(db),
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
	go app.ContinuousFeedScraping()

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", app.routes()))
	corsMux := middlewareCors(mux)

	server := http.Server{Addr: ":" + httpPort, Handler: corsMux}
	fmt.Printf("Serving on http port: %s\n", httpPort)
	log.Fatal(server.ListenAndServe())
}

func isProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}
