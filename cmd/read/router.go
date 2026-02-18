package main

import "github.com/go-chi/chi/v5"

func router(
	r *chi.Mux,
) {

	r.Route("/api/v1", func(r chi.Router) {
		r.Use()
		r.Get("/health", healthHandler)
	})

}
