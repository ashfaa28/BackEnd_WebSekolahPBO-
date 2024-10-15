package req

type SiswaReq struct {
	Nama           string `json:"nama" validate:"required"`
	NISN           string `json:"nisn" validate:"required"`
	NoTelp         string `json:"NoTelp" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Usia           string `json:"usia" validate:"required"`
	Alamat         string `json:"alamat" validate:"required"`
	NamaOrangTua   string `json:"namaOrangTua" validate:"required"`
	NoTelpOrangTua string `json:"noTelpOrangTua" validate:"required"`
	AsalSekolah    string `json:"asalSekolah" validate:"required"`
}
