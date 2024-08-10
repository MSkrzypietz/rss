package main

import (
	"database/sql"
	"fmt"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/views"
	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"log"
	"net/http"
	"os"
	"time"
)

type apiConfig struct {
	DB         *database.Queries
	httpClient *http.Client
}

func main() {
	if !isProductionEnv() {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is undefined")
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

	apiCfg := apiConfig{
		DB:         database.New(db),
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}

	go apiCfg.continuousFeedFetcher()

	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(views.Index("Michael")))

	mux.HandleFunc("GET /v1/users", apiCfg.authenticate(apiCfg.getUsers))
	mux.HandleFunc("POST /v1/users", apiCfg.createUser)

	mux.HandleFunc("GET /v1/feeds", apiCfg.getFeeds)
	mux.HandleFunc("POST /v1/feeds", apiCfg.authenticate(apiCfg.createFeed))

	mux.HandleFunc("GET /v1/feed_follows", apiCfg.authenticate(apiCfg.getFeedFollows))
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.authenticate(apiCfg.createFeedFollow))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.authenticate(apiCfg.deleteFeedFollow))

	mux.HandleFunc("GET /v1/posts", apiCfg.authenticate(apiCfg.getPosts))

	mux.HandleFunc("GET /v1/readiness", getReadiness)
	mux.HandleFunc("GET /v1/err", getError)

	corsMux := middlewareCors(mux)

	server := http.Server{Addr: ":" + port, Handler: corsMux}
	fmt.Printf("Serving on port: %s\n", port)
	if isProductionEnv() {
		certFile := os.Getenv("CERT_FILE_PATH")
		if certFile == "" {
			log.Fatalln("CERT_FILE_PATH is undefined")
		}

		keyFile := os.Getenv("KEY_FILE_PATH")
		if keyFile == "" {
			log.Fatalln("KEY_FILE_PATH is undefined")
		}

		log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}

func isProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}
