package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is undefined")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalln("DB_URL is undefined")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Cannot open database connection: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot ping database: %v", err)
	}

	apiCfg := apiConfig{DB: database.New(db)}

	//id, _ := uuid.Parse("4c4984c8-e5b4-48d6-a80b-4448e5e7f43f")
	//apiCfg.DB.MarkFeedFetched(context.Background(), id)

	lastFeeds, err := apiCfg.DB.GetNextFeedsToFetch(context.Background(), 2)
	fmt.Println(lastFeeds[0].LastFetchedAt)
	fmt.Println(lastFeeds[1].LastFetchedAt)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/users", apiCfg.authenticate(apiCfg.getUsers))
	mux.HandleFunc("POST /v1/users", apiCfg.createUser)

	mux.HandleFunc("GET /v1/feeds", apiCfg.getFeeds)
	mux.HandleFunc("POST /v1/feeds", apiCfg.authenticate(apiCfg.createFeed))

	mux.HandleFunc("GET /v1/feed_follows", apiCfg.authenticate(apiCfg.getFeedFollows))
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.authenticate(apiCfg.createFeedFollow))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.authenticate(apiCfg.deleteFeedFollow))

	mux.HandleFunc("GET /v1/readiness", getReadiness)
	mux.HandleFunc("GET /v1/err", getError)

	corsMux := middlewareCors(mux)

	server := http.Server{Addr: ":" + port, Handler: corsMux}
	fmt.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
