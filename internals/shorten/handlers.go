package shorten

import (
	"net/http"
)

func HandleShortenURL(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: decode request, call svc.Shorten(longURL)
	}
}
