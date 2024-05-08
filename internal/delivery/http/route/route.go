package route

import (
	"github.com/go-chi/chi/v5"
	https "github.com/wiscaksono/lapormas-go/internal/delivery/http"
)

type RouteConfig struct {
	App            *chi.Mux
	UserController *https.UserController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/", c.UserController.Register)
	c.App.Get("/user", c.UserController.Current)
}
