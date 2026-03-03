package main

import (
	"net/http"

	"github.com/Jidetireni/tiny/config"
	"github.com/Jidetireni/tiny/internals/shorten"
	"github.com/go-chi/chi/v5"
)

func NewServer(
	config *config.Config,
	shortenService *shorten.Service,
) http.Handler {
	// Initialize the router
	r := chi.NewRouter()
	// Wire the routes and inject dependencies
	router(
		r,
		shortenService,
	)
	// 3. Wrap with Middlewares (Logger, Auth, etc.)
	return r
}
