package entity

import "github.com/google/uuid"

type Role struct {
	ID uuid.UUID `gorm:"primary_key;unique;type:uuid;column:id;default:uuid_generate_v4()"`

	Name RoleCode `gorm:"column:name;unique;not null;type:(enum('Manager', 'Supervisor', 'Engineer', 'Regular'));default:'Regular'"`

	Accounts []Account `gorm:"foreignKey:RoleID;references:ID"`
}

func (U *Role) TableName() string {
	return "roles"
}

type RoleCode string

const (
	Manager    RoleCode = "Manager"
	Supervisor RoleCode = "Supervisor"
	Engineer   RoleCode = "Engineer"
	Regular    RoleCode = "Regular"
)
