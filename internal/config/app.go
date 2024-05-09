package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/wiscaksono/lapormas-go/internal/delivery/http"
	"github.com/wiscaksono/lapormas-go/internal/delivery/http/middleware"
	"github.com/wiscaksono/lapormas-go/internal/delivery/http/route"
	"github.com/wiscaksono/lapormas-go/internal/repository"
	"github.com/wiscaksono/lapormas-go/internal/usecase"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// Setup repository
	userRepository := repository.NewUserRepository(config.Log)
	accountRepository := repository.NewAccountRepository(config.Log)
	roleRepository := repository.NewRoleRepository(config.Log)

	// Setup usecase
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	accountUseCase := usecase.NewAccountUseCase(config.DB, config.Log, config.Validate, accountRepository)
	roleUseCase := usecase.NewRoleUseCase(config.DB, config.Log, config.Validate, roleRepository)

	// Setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	accountController := http.NewAccountController(accountUseCase, config.Log)
	roleController := http.NewRoleController(roleUseCase, config.Log)

	// Setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App: config.App,

		UserController:    userController,
		AccountController: accountController,
		RoleController:    roleController,

		AuthMiddleware: authMiddleware,
	}

	routeConfig.Setup()
}
