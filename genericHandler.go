package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/martian/log"
)

// HandlerFunc is reusable type of function
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	if err != nil {
		w.Header().Set("Content-Type", "application/problem+json")
		if err != nil {
			log.Errorf("failed to marshal error data", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(
				fmt.Sprintf(
					`{"status": %d, "code": "%s", "title": "%s"}`,
					http.StatusInternalServerError,
					ErrCodeInternal,
					err.Error(),
				)))
			return
		}
	}
}

// DecodeJSON is a helper function that automatically checks for any error while decoding the JSON and returns Error
// struct for convenience
func DecodeJSON(r io.Reader, v interface{}) error {
	if err := render.DecodeJSON(r, v); err != nil {
		log.Errorf("failed decoding json", "error", err)
		return BadRequest(err)
	}
	return nil
}
