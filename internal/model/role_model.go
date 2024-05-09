package model

import (
	"github.com/google/uuid"
	"github.com/wiscaksono/lapormas-go/internal/entity"
)

type RoleResponse struct {
	ID uuid.UUID `json:"id,omitempty"`

	Name entity.RoleCode `json:"role,omitempty"`
}
