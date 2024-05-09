package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/wiscaksono/lapormas-go/internal/entity"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Repository[entity.Role]
	Log *logrus.Logger
}

func NewRoleRepository(log *logrus.Logger) *RoleRepository {
	return &RoleRepository{
		Log: log,
	}
}

func (r *RoleRepository) List(db *gorm.DB, role *[]entity.Role) error {
	return db.Find(role).Error
}
