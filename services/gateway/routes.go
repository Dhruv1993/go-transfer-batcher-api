package gateway

import (
	"github.com/Dhruv1993/app/services"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func initRoutes(services []services.Service, URLPrefix string) http.Handler {
	r := chi.NewRouter()
	// Use the Logger middleware to log request details.
	r.Use(middleware.Logger)
	for _, svc := range services {
		svc.InitRoutes(r, URLPrefix)
	}

	return r
}
