package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth/login", app.loginHandler)

	mux.HandleFunc("GET /v1/users", app.authenticate(app.getAuthenticatedUserHandler))
	mux.HandleFunc("POST /v1/users", app.createUserHandler)

	mux.HandleFunc("GET /v1/feeds", app.authenticate(app.listFeedsHandler))
	mux.HandleFunc("POST /v1/feeds", app.authenticate(app.createFeedHandler))

	mux.HandleFunc("GET /v1/feed_follows", app.authenticate(app.listFeedFollowsHandler))
	mux.HandleFunc("POST /v1/feed_follows", app.authenticate(app.createFeedFollowHandler))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", app.authenticate(app.deleteFeedFollowHandler))

	mux.HandleFunc("GET /v1/feed_filters", app.authenticate(app.listFeedFiltersHandler))
	mux.HandleFunc("POST /v1/feed_filters", app.authenticate(app.createFeedFilterHandler))
	mux.HandleFunc("DELETE /v1/feed_filters/{feedFilterID}", app.authenticate(app.deleteFeedFilterHandler))

	mux.HandleFunc("GET /v1/posts", app.authenticate(app.getUnreadPostsHandler))
	mux.HandleFunc("POST /v1/posts/{postID}/read", app.authenticate(app.markPostAsReadHandler))

	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("GET /v1/readiness", app.getReadinessHandler)

	return enableCORS(mux)
}
