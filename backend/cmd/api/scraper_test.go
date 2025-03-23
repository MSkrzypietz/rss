package main

import (
	"log/slog"
	"net/http"
	"os"
	"testing"
)

func TestFetchFeed(t *testing.T) {
	app := &application{
		httpClient: http.DefaultClient,
		logger:     slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	feedUrls := []string{
		"https://fabiensanglard.net/rss.xml",
		"https://blog.cloudflare.com/rss",
		"https://blog.golang.org/feed.atom",
		"https://hackattic.com/challenges.rss",
		"https://blog.boot.dev/index.xml",
		"https://blog.jetbrains.com/feed",
	}

	for _, feedUrl := range feedUrls {
		_, err := app.getFeedByUrl(feedUrl)
		if err != nil {
			t.Fatal(err)
		}
	}
}
