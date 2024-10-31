package req

type SiswaReq struct {
	Nama           string `json:"nama" validate:"required"`
	NISN           string `json:"nisn" validate:"required"`
	JenisKelamin   string `json:"jenkel" validate:"required"` // Added required tag
	Email          string `json:"email" validate:"required,email"`
	Usia           string `json:"usia" validate:"required"`
	Alamat         string `json:"alamat" validate:"required"`
	NamaAyah       string `json:"nama_ayah" validate:"required"`         // Added required tag
	NamaIbu        string `json:"nama_ibu" validate:"required"`          // Added required tag
	NoTelpOrangTua string `json:"no_telp_orang_tua" validate:"required"` // Added required tag
	AsalSekolah    string `json:"asal_sekolah" validate:"required"`      // Added required tag
	Jurusan        string `json:"jurusan" validate:"required"`
}
