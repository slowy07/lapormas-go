package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wiscaksono/lapormas-go/internal/model"
	"github.com/wiscaksono/lapormas-go/internal/usecase"
)

type RoleController struct {
	Log     *logrus.Logger
	UseCase *usecase.RoleUseCase
}

func NewRoleController(useCase *usecase.RoleUseCase, logger *logrus.Logger) *RoleController {
	return &RoleController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *RoleController) List(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to list role : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.RoleResponse]{
		Success: true, Data: response, Message: "Success get roles",
	})
}
