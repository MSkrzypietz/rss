package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/login", app.login)

	mux.HandleFunc("GET /users", app.authenticate(app.getAuthenticatedUser))
	mux.HandleFunc("POST /users", app.createUser)

	mux.HandleFunc("GET /feeds", app.authenticate(app.getFeeds))
	mux.HandleFunc("POST /feeds", app.authenticate(app.createFeed))

	mux.HandleFunc("GET /feed_follows", app.authenticate(app.getFeedFollows))
	mux.HandleFunc("POST /feed_follows", app.authenticate(app.createFeedFollow))
	mux.HandleFunc("DELETE /feed_follows/{feedFollowID}", app.authenticate(app.deleteFeedFollow))

	mux.HandleFunc("GET /feed_filters", app.authenticate(app.getFeedFilters))
	mux.HandleFunc("POST /feed_filters", app.authenticate(app.createFeedFilter))
	mux.HandleFunc("DELETE /feed_filters/{feedFilterID}", app.authenticate(app.deleteFeedFilter))

	mux.HandleFunc("GET /posts", app.authenticate(app.getUnreadPosts))
	mux.HandleFunc("POST /posts/{postID}/read", app.authenticate(app.markPostAsRead))

	mux.HandleFunc("GET /healthcheck", healthcheckHandler)
	mux.HandleFunc("GET /readiness", getReadiness)
	mux.HandleFunc("GET /err", getError)

	return enableCORS(mux)
}
