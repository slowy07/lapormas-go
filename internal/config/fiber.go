package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/wiscaksono/lapormas-go/internal/model"
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: NewErrorHandler(),
		AppName:      config.GetString("app.name"),
		Prefork:      config.GetBool("web.prefork"),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(model.WebErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}
}
