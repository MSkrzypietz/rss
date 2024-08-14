package main

import (
	"database/sql"
	"fmt"
	"github.com/MSkrzypietz/rss/api"
	"github.com/MSkrzypietz/rss/render"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"log"
	"net"
	"net/http"
	"os"
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

	httpsPort := os.Getenv("HTTPS_PORT")
	if httpsPort == "" {
		log.Fatalln("HTTPS_PORT is undefined")
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

	if isProductionEnv() {
		certFile := os.Getenv("CERT_FILE_PATH")
		if certFile == "" {
			log.Fatalln("CERT_FILE_PATH is undefined")
		}

		keyFile := os.Getenv("KEY_FILE_PATH")
		if keyFile == "" {
			log.Fatalln("KEY_FILE_PATH is undefined")
		}

		go redirectHttpToHttps(httpPort, httpsPort)
		server := http.Server{Addr: ":" + httpsPort, Handler: corsMux}
		fmt.Printf("Serving on https port: %s\n", httpsPort)
		log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
	} else {
		server := http.Server{Addr: ":" + httpPort, Handler: corsMux}
		fmt.Printf("Serving on http port: %s\n", httpPort)
		log.Fatal(server.ListenAndServe())
	}
}

func isProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}

func redirectHttpToHttps(httpPort, httpsPort string) {
	server := http.Server{
		Addr: ":" + httpPort,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host, _, _ := net.SplitHostPort(r.Host)
			redirectUrl := r.URL
			redirectUrl.Host = net.JoinHostPort(host, httpsPort)
			redirectUrl.Scheme = "https"
			http.Redirect(w, r, redirectUrl.String(), http.StatusMovedPermanently)
		}),
	}
	fmt.Printf("Redirecting http traffic from port %s to %s\n", httpPort, httpsPort)
	log.Println(server.ListenAndServe())
}
