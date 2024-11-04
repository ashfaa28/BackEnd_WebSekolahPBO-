package req

type BeritaReq struct {
	Judul     string `json:"judul" validate:"required"`
	IsiBerita string `json:"isi_berita" validate:"required"`
}
