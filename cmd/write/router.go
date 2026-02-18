package main

import (
	"github.com/Jidetireni/tiny/pkg/zookeeper"
	"github.com/go-chi/chi/v5"
)

func router(
	r *chi.Mux,
	zookeeper *zookeeper.Zookeeper,
) {

	r.Route("/api/v1", func(r chi.Router) {
		r.Use()
		r.Get("/health", healthHandler())
	})

}
