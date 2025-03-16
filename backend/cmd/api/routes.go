package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth/login", app.login)

	mux.HandleFunc("GET /v1/users", app.authenticate(app.getAuthenticatedUser))
	mux.HandleFunc("POST /v1/users", app.createUser)

	mux.HandleFunc("GET /v1/feeds", app.authenticate(app.getFeeds))
	mux.HandleFunc("POST /v1/feeds", app.authenticate(app.createFeed))

	mux.HandleFunc("GET /v1/feed_follows", app.authenticate(app.getFeedFollows))
	mux.HandleFunc("POST /v1/feed_follows", app.authenticate(app.createFeedFollow))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", app.authenticate(app.deleteFeedFollow))

	mux.HandleFunc("GET /v1/feed_filters", app.authenticate(app.getFeedFilters))
	mux.HandleFunc("POST /v1/feed_filters", app.authenticate(app.createFeedFilter))
	mux.HandleFunc("DELETE /v1/feed_filters/{feedFilterID}", app.authenticate(app.deleteFeedFilter))

	mux.HandleFunc("GET /v1/posts", app.authenticate(app.getUnreadPosts))
	mux.HandleFunc("POST /v1/posts/{postID}/read", app.authenticate(app.markPostAsRead))

	mux.HandleFunc("GET /v1/healthcheck", healthcheckHandler)
	mux.HandleFunc("GET /v1/readiness", getReadiness)
	mux.HandleFunc("GET /v1/err", getError)

	return enableCORS(mux)
}
