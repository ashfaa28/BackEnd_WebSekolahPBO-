package req

type PrestasiReq struct {
	Gambar string `json:"gambar" validate:"required"`
}
