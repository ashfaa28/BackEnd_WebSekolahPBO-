package req

type SiswaReq struct {
	Nama           string `json:"nama" validate:"required"`
	NISN           string `json:"nisn" validate:"required"`
	NoTelp         string `json:"no_telp" validate:"required"` // Perbaikan "NoTelp" menjadi "no_telp"
	Email          string `json:"email" validate:"required,email"`
	Usia           string `json:"usia" validate:"required"`
	Alamat         string `json:"alamat" validate:"required"`
	NamaAyah       string `json:"nama_ayah" validate:"required"` // Hilangkan spasi pada `json` dan gunakan snake_case
	NamaIbu        string `json:"nama_ibu" validate:"required"`  // Hilangkan spasi pada `json` dan gunakan snake_case
	NoTelpOrangTua string `json:"no_telp_orang_tua" validate:"required"`
	AsalSekolah    string `json:"asal_sekolah" validate:"required"`
}
