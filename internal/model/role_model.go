package model

import (
	"github.com/google/uuid"
	"github.com/wiscaksono/lapormas-go/internal/entity"
)

type RoleResponse struct {
	List []RoleIDName `json:"list,omitempty"`
}

type RoleIDName struct {
	ID uuid.UUID `json:"id,omitempty"`

	Name entity.RoleCode `json:"name,omitempty"`
}
