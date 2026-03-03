package shorten

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Jidetireni/tiny/pkg/httpio"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func HandleShortenURL(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ShortenRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpio.WriteError(w, httpio.BadRequest("malformed JSON body"))
			return
		}

		if err := validate.Struct(req); err != nil {
			var validationErrors validator.ValidationErrors
			if errors.As(err, &validationErrors) {
				httpio.WriteError(w, httpio.BadRequest(validationErrors[0].Translate(nil)))
				return
			}
			httpio.WriteError(w, httpio.BadRequest("invalid request body"))
			return
		}

		shortCode, err := svc.Shorten(r.Context(), req.LongURL)
		if err != nil {
			httpio.WriteError(w, err)
			return
		}

		httpio.WriteJSON(w, http.StatusCreated, map[string]string{
			"short_code": shortCode,
		})
	}
}
