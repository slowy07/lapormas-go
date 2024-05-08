package route

import (
	"github.com/gofiber/fiber/v2"
	https "github.com/wiscaksono/lapormas-go/internal/delivery/http"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *https.UserController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/register", c.UserController.Register)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Post("/user", c.UserController.Current)
}
