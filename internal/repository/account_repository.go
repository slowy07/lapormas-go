package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/wiscaksono/lapormas-go/internal/entity"
	"gorm.io/gorm"
)

type AccountRepository struct {
	Repository[entity.Account]
	Log *logrus.Logger
}

func NewAccountRepository(log *logrus.Logger) *AccountRepository {
	return &AccountRepository{
		Log: log,
	}
}

func (r *AccountRepository) FindByNoPegawai(db *gorm.DB, acc *entity.Account, NoPegawai string) error {
	return db.Where("no_pegawai = ?", NoPegawai).First(acc).Error
}
