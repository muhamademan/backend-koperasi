package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	NIK       string         `json:"nik" form:"nik" validate:"required" gorm:"not null"`
	Name      string         `json:"name" form:"name" validate:"required" gorm:"not null"`
	Email     string         `json:"email" form:"email" validate:"required,email" gorm:"not null"`
	Password  string         `json:"password" validate:"required" gorm:"not null"`
	Address   string         `json:"address" form:"address" gorm:"not null"`
	Phone     string         `json:"phone" form:"phone" validate:"lte=12" gorm:"not null"`
	Role      string         `json:"role" form:"role" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
