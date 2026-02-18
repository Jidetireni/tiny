package main

import (
	"net/http"

	"github.com/Jidetireni/tiny/config"
	"github.com/go-chi/chi/v5"
)

func NewServer(
	config *config.Config,

) http.Handler {
	// 1. Initialize the router
	r := chi.NewRouter()
	// 2. Wire the routes and inject dependencies
	router(r)

	// 3. Wrap with Middlewares (Logger, Auth, etc.)

	return r
}
