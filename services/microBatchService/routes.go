package microBatchService

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// InitRoutes will have all the defined route paths.For this example the url path is {URL_PREFIX}/transfer
func (s Service) InitRoutes(r chi.Router, urlPrefix string) {
	r.Route(fmt.Sprintf("%s/%s", urlPrefix, ""), func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Route("/transfer", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Method(http.MethodPost, "/", s.TaskHandler())
				})
			})
		})
	})
}
