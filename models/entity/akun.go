package entity

import (
	"time"

	"gorm.io/gorm"
)

type Akun struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserName  string         `json:"user_name" gorm:"unique"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
