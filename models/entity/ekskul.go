package entity

import (
	"gorm.io/gorm"
)

type Ekskul struct {
	gorm.Model
	NamaEkskul string `json:"nama_ekskul"`
	Gambar     string `json:"gambar"`
}
