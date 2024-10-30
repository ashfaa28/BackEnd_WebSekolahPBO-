package entity

import (
	"time"

	"gorm.io/gorm"
)

type Siswa struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Nama           string         `json:"nama"`
	NISN           string         `json:"nisn" gorm:"unique"`
	JenisKelamin   string         `json:"jenkel"`
	Email          string         `json:"email"`
	Usia           string         `json:"usia"`
	Alamat         string         `json:"alamat"`
	NamaAyah       string         `json:"nama_ayah"`
	NamaIbu        string         `json:"nama_ibu"`
	NoTelpOrangTua string         `json:"no_telp_orang_tua"`
	AsalSekolah    string         `json:"asal_sekolah"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
