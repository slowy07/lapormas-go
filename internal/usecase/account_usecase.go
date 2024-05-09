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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	AccountRepository *repository.AccountRepository
}

func NewAccountUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, accountRepository *repository.AccountRepository) *AccountUseCase {
	return &AccountUseCase{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		AccountRepository: accountRepository,
	}
}

func (c *AccountUseCase) Create(ctx context.Context, request *model.RegisterAccountRequest) (*model.AccountResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	exist := new(entity.Account)
	if check := c.AccountRepository.FindByNoPegawai(tx, exist, request.NoPegawai); check == nil {
		c.Log.Warnf("Account with NoPegawai already exist : %+v", request.NoPegawai)
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	acc := &entity.Account{
		NoPegawai: request.NoPegawai,
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Jabatan:   request.Jabatan,
		Password:  string(password),
		// Role: request.
	}

	if err := c.AccountRepository.Create(tx, acc); err != nil {
		c.Log.Warnf("Failed create account to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.AccountToResponse(acc), nil
}

func (c *AccountUseCase) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	return nil, nil
}
