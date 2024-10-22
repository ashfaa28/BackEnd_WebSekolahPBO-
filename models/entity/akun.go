package entity

import (
	"time"

	"gorm.io/gorm"
)

type Akun struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserName    string         `json:"username" gorm:"unique"`
	Password    string         `json:"password"`
	AsalSekolah string         `json:"asal_sekolah"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
