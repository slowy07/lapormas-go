package converter

import (
	"github.com/wiscaksono/lapormas-go/internal/entity"
	"github.com/wiscaksono/lapormas-go/internal/model"
)

func AccountToResponse(acc *entity.Account) *model.AccountResponse {
	return &model.AccountResponse{
		ID:        acc.ID,
		NoPegawai: acc.NoPegawai,
		Name:      acc.Name,
		Email:     acc.Email,
		Phone:     acc.Phone,
		Jabatan:   acc.Jabatan,
		Image:     acc.Image,
		LastSeen:  acc.LastSeen,
		Role:      acc.Role.Name,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}
}
