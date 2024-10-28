// File: models/entity/berita.go
package entity

import "gorm.io/gorm"

type Berita struct {
	gorm.Model
	Gambar    string `json:"gambar"`
	IsiBerita string `json:"isi_berita"`
}
