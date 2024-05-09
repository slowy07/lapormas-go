package entity

import "github.com/google/uuid"

type Account struct {
	ID uuid.UUID `gorm:"primary_key;unique;type:uuid;column:id;default:uuid_generate_v4()"`

	NoPegawai string    `gorm:"column:no_pegawai;unique;not null"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;unique;not null"`
	Phone     string    `gorm:"column:phone;unique;not null"`
	Jabatan   string    `gorm:"column:jabatan;not null"`
	Password  string    `gorm:"column:password;not null"`
	Image     string    `gorm:"column:image"`
	LastSeen  int64     `gorm:"column:last_seen"`
	RoleID    uuid.UUID `gorm:"column:role_id;type:uuid;not null"`
	Role      Role      `gorm:"foreignKey:RoleID;references:ID"`

	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (U *Account) TableName() string {
	return "accounts"
}
