package req

type EkskulReq struct {
	NamaEkskul string `json:"nama_ekskul" validate:"required"`
}
