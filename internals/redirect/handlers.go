package redirect

import (
	"net/http"

	"github.com/Jidetireni/tiny/pkg/httpio"
)

func HandleRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: extract short code from URL, call svc.Redirect(ctx, shortCode)
		_ = httpio.WriteError
	}
}
