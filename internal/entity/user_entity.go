package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"primary_key;unique;type:uuid;column:id;default:uuid_generate_v4()"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Token     string    `gorm:"column:token"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (U *User) TableName() string {
	return "users"
}
