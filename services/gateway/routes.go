package gateway

import (
	"net/http"

	"github.com/Dhruv1993/app/services"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Your API Title
// @version 1.0
// @description Your API Description
// @host localhost:5001
// @BasePath /v1
func initRoutes(services []services.Service, URLPrefix string) http.Handler {
	r := chi.NewRouter()
	// Enable CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://*","https://"}, // used by swagger
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of the major browsers
	})
	// Use the Logger middleware to log request details.
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware.Handler)
	
	r.Get("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to api.."))
	})
	
	r.Get("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
    ))
	for _, svc := range services {
		svc.InitRoutes(r, URLPrefix)
	}

	return r
}
