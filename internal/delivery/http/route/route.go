package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiscaksono/lapormas-go/internal/delivery/http"
)

type RouteConfig struct {
	App *fiber.App

	UserController    *http.UserController
	AccountController *http.AccountController
	RoleController    *http.RoleController

	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/register", c.UserController.Register)
	c.App.Post("/login", c.UserController.Login)
	c.App.Post("/account/register", c.AccountController.Register)
	c.App.Post("/account/login", c.AccountController.Login)
	c.App.Get("/roles", c.RoleController.List)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/logout", c.UserController.Logout)
	c.App.Get("/user", c.UserController.Current)
}
