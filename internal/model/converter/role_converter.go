package converter

import (
	"github.com/wiscaksono/lapormas-go/internal/entity"
	"github.com/wiscaksono/lapormas-go/internal/model"
)

func RoleToResponse(role *entity.Role) *model.RoleResponse {
	return &model.RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
}
