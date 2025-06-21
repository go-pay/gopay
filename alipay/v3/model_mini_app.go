package alipay

type OpenMiniVersionAuditApplyRsp struct {
	StatusCode  int         `json:"status_code"`
	ErrResponse ErrResponse `json:"-"`

	SpeedUp     string `json:"speed_up"`
	SpeedUpMemo string `json:"speed_up_memo"`
}
