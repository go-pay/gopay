package alipay

type MerchantItemFileUploadRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	MaterialId  string `json:"material_id"`
	MaterialKey string `json:"material_key"`
}
