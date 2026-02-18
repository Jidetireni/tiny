package main

import (
	"github.com/Jidetireni/tiny/internals/shorten"
	"github.com/go-chi/chi/v5"
)

func router(
	r *chi.Mux,
	shortenService *shorten.Service,
) {

	r.Route("/api/v1", func(r chi.Router) {
		r.Use()
		r.Post("/shorten", shorten.HandleShortenURL(shortenService))
	})

}
