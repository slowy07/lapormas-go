package model

import (
	"github.com/google/uuid"
	"github.com/wiscaksono/lapormas-go/internal/entity"
)

type AccountResponse struct {
	ID uuid.UUID `json:"id,omitempty"`

	NoPegawai string          `json:"no_pegawai,omitempty"`
	Name      string          `json:"name,omitempty"`
	Email     string          `json:"email,omitempty"`
	Phone     string          `json:"phone,omitempty"`
	Jabatan   string          `json:"jabatan,omitempty"`
	Password  string          `json:"password,omitempty"`
	Image     string          `json:"image,omitempty"`
	LastSeen  int64           `json:"last_seen,omitempty"`
	Role      entity.RoleCode `json:"role,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

type RegisterAccountRequest struct {
	NoPegawai string    `json:"no_pegawai" validate:"required,max=100"`
	Name      string    `json:"name" validate:"required,max=100"`
	Email     string    `json:"email" validate:"required,email,max=100"`
	Phone     string    `json:"phone" validate:"required,max=100"`
	Jabatan   string    `json:"jabatan" validate:"required,max=100"`
	Password  string    `json:"password" validate:"required,max=100,min=5"`
	RoleID    uuid.UUID `json:"role_id" validate:"required"`
}

type LoginAccountRequest struct {
	Username string `json:"username" validate:"required,max=100,min=5"`
	Password string `json:"password" validate:"required,max=100"`
}
