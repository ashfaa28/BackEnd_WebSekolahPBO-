package req

type AkunReq struct {
	UserName    string `json:"username"  validate:"required"`
	Password    string `json:"password"  validate:"required"`
	AsalSekolah string `json:"asal_sekolah"  validate:"required"`
}
