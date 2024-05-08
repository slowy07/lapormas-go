package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

func NewChi(config *viper.Viper) *chi.Mux {
	app := chi.NewRouter()

	app.Use(middleware.Logger)
	app.Use(middleware.RealIP)
	app.Use(middleware.RequestID)
	app.Use(middleware.Recoverer)

	return app
}
