package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiscaksono/lapormas-go/internal/delivery/http"
	"github.com/wiscaksono/lapormas-go/internal/delivery/http/route"
	"github.com/wiscaksono/lapormas-go/internal/repository"
	"github.com/wiscaksono/lapormas-go/internal/usecase"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *chi.Mux
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// Setup repository
	userRepository := repository.NewUserRepository(config.Log)

	// Setup usecase
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)

	// Setup controller
	userController := http.NewUserController(userUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}

	routeConfig.Setup()
}
