package entity

import "gorm.io/gorm"

type Prestasi struct {
	gorm.Model
	Gambar string `json:"gambar"`
}
