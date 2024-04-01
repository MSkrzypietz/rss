package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is undefined")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/readiness", getReadiness)
	mux.HandleFunc("GET /v1/err", getError)
	corsMux := middlewareCors(mux)

	server := http.Server{Addr: ":" + port, Handler: corsMux}
	fmt.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
