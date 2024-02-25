package gateway

import (
	"net/http"

	"github.com/Dhruv1993/app/services"
)

type Service struct {
	Services []services.Service
	Handler  http.Handler
}

func New(services []services.Service, urlPrefix string) Service {
	r := initRoutes(services, urlPrefix)
	return Service{services, r}
}
