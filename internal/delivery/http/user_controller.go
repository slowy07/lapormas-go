package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/wiscaksono/lapormas-go/internal/model"
	"github.com/wiscaksono/lapormas-go/internal/usecase"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	body := new(model.RegisterUserRequest)

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (c *UserController) Current(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
