package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wiscaksono/lapormas-go/internal/model"
	"github.com/wiscaksono/lapormas-go/internal/usecase"
)

type AccountController struct {
	Log     *logrus.Logger
	UseCase *usecase.AccountUseCase
}

func NewAccountController(useCase *usecase.AccountUseCase, logger *logrus.Logger) *AccountController {
	return &AccountController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *AccountController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterAccountRequest)

	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AccountResponse]{
		Success: true, Data: response, Message: "Account has been registered",
	})
}

func (c *AccountController) Current(ctx *fiber.Ctx) error {
	return nil
}

func (c *AccountController) Login(ctx *fiber.Ctx) error {
	return nil
}
