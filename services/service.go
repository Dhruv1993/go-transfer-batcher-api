package services

import (
	"github.com/go-chi/chi/v5"
)

// Service is an interface which can be reused for different services which will have different routes
type Service interface {
	InitRoutes(r chi.Router, urlPrefix string)
}
