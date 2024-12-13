package entity

import (
	"go-clean-architecture/pkg"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"type:varchar(36);primary_key;not null" json:"id"`
	Username  string    `gorm:"type:varchar(100);not null" json:"username"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = pkg.GenerateXID()
	return
}
