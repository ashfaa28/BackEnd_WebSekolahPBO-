package entity

import (
	"gorm.io/gorm"
)

type Berita struct {
	gorm.Model
	Judul     string `json:"judul"`
	Gambar    string `json:"gambar"`
	IsiBerita string `json:"isi_berita"`
}
