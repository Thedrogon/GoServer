package routes

import (
	"github.com/Thedrogon/v1/Goserver/internal/app"
	"github.com/go-chi/chi"
)

func SetRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health",app.HealthCheck)
	return r
}