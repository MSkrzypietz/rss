package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

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

	apiCfg := NewConfig(db)
	go apiCfg.ContinuousFeedScraping()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiCfg.Handlers()))
	corsMux := middlewareCors(mux)

	server := http.Server{Addr: ":" + httpPort, Handler: corsMux}
	fmt.Printf("Serving on http port: %s\n", httpPort)
	log.Fatal(server.ListenAndServe())
}

func isProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}
