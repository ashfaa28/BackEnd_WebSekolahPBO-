package req

type BeritaReq struct {
	IsiBerita string `json:"isi_berita" validate:"required"`
}
