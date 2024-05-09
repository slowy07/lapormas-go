package converter

import (
	"github.com/wiscaksono/lapormas-go/internal/entity"
	"github.com/wiscaksono/lapormas-go/internal/model"
)

func RoleToResponse(role *[]entity.Role) *model.RoleResponse {
	return &model.RoleResponse{
		List: roleToRoleIDName(role),
	}
}

func roleToRoleIDName(role *[]entity.Role) []model.RoleIDName {
	var list []model.RoleIDName

	for _, r := range *role {
		list = append(list, model.RoleIDName{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	return list
}
