package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/lapormas-go/internal/delivery/http"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/register", c.UserController.Register)
	c.App.Post("/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/logout", c.UserController.Logout)
	c.App.Get("/user", c.UserController.Current)
}
