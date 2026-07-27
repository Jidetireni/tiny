package redirect

import (
	"net/http"

	"github.com/Jidetireni/tiny/pkg/httpio"
	"github.com/go-chi/chi/v5"
)

func HandleRedirectURL(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortCode := chi.URLParam(r, "code")
		if shortCode == "" {
			httpio.WriteError(w, httpio.BadRequest(""))
			return
		}

		longURL, err := svc.Redirect(r.Context(), shortCode)
		if err != nil {
			httpio.WriteError(w, err)
			return
		}
		http.Redirect(w, r, longURL, http.StatusFound)
	}
}
