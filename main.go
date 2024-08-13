package main

import (
	"database/sql"
	"fmt"
	"github.com/MSkrzypietz/rss/api"
	"github.com/MSkrzypietz/rss/render"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"log"
	"net/http"
	"os"
)

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

	renderCfg := render.NewConfig(db)
	apiCfg := api.NewConfig(db)

	go apiCfg.ContinuousFeedScraping()

	mux := http.NewServeMux()

	mux.Handle("/", renderCfg.Handlers())
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiCfg.Handlers()))

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
