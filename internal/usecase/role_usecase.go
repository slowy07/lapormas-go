package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wiscaksono/lapormas-go/internal/entity"
	"github.com/wiscaksono/lapormas-go/internal/model"
	"github.com/wiscaksono/lapormas-go/internal/model/converter"
	"github.com/wiscaksono/lapormas-go/internal/repository"
	"gorm.io/gorm"
)

type RoleUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	RoleRepository *repository.RoleRepository
}

func NewRoleUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, roleRepository *repository.RoleRepository) *RoleUseCase {
	return &RoleUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		RoleRepository: roleRepository,
	}
}

func (c *RoleUseCase) List(ctx context.Context) (*model.RoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	role := new([]entity.Role)

	if err := c.RoleRepository.List(tx, role); err != nil {
		c.Log.Warnf("Failed to list role : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.RoleToResponse(role), nil
}
