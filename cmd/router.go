package main

import (
	"github.com/Jidetireni/tiny/internals/redirect"
	"github.com/Jidetireni/tiny/internals/shorten"
	"github.com/go-chi/chi/v5"
)

func router(
	r *chi.Mux,
	shortenService *shorten.Service,
	redirectService *redirect.Service,
) {

	r.Route("/api/v1", func(r chi.Router) {
		r.Use()
		r.Post("/shorten", shorten.HandleShortenURL(shortenService))

	})

	r.Route("/", func(r chi.Router) {
		r.Get("/{code}", redirect.HandleRedirectURL(redirectService))
	})

}
