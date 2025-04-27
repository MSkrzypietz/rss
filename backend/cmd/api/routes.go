package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodPost, "/v1/auth/login", app.loginHandler)

	router.HandlerFunc(http.MethodGet, "/v1/users", app.authenticate(app.getAuthenticatedUserHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users", app.createUserHandler)

	router.HandlerFunc(http.MethodGet, "/v1/feeds", app.authenticate(app.listFeedsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/feeds", app.authenticate(app.createFeedHandler))
	router.HandlerFunc(http.MethodPost, "/v1/feeds/:id/fetch", app.authenticate(app.fetchFeedHandler))

	router.HandlerFunc(http.MethodGet, "/v1/feed_follows", app.authenticate(app.listFeedFollowsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/feed_follows", app.authenticate(app.createFeedFollowHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/feed_follows/:id", app.authenticate(app.deleteFeedFollowHandler))

	router.HandlerFunc(http.MethodGet, "/v1/feed_filters", app.authenticate(app.listFeedFiltersHandler))
	router.HandlerFunc(http.MethodPost, "/v1/feed_filters", app.authenticate(app.createFeedFilterHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/feed_filters/:id", app.authenticate(app.deleteFeedFilterHandler))

	router.HandlerFunc(http.MethodGet, "/v1/posts", app.authenticate(app.getUnreadPostsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/posts/:id/read", app.authenticate(app.markPostAsReadHandler))

	router.HandlerFunc(http.MethodPost, "/v1/telegram/echo", app.authenticate(app.telegramEchoHandler))

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/readiness", app.getReadinessHandler)

	return enableCORS(router)
}
