package byte

/*
{
  "err_no": 0,
  "err_tips": "success"
}
*/

type SuccessRsp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
}
