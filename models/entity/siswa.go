package entity

import (
	"time"

	"gorm.io/gorm"
)

type Siswa struct {
	ID             uint   `json : "id" gorm:primarykey`
	Nama           string `json : "nama"`
	NISN           string `json : "NISN"`
	NoTelp         string `json : "NoTelp"`
	Email          string `json : "email"`
	Usia           string `json : "usia"`
	Alamat         string `json : "alamat"`
	NamaOrangTua   string `json : "namaOrangTua"`
	NoTelpOrangTua string `json : "noTelpOrangTua"`
	AsalSekolah    string `json : "asalSekolah"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
