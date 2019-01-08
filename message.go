package go_pay

type ErrorMessage struct {
	ErrorCode int    `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}

type SuccessMessage struct {
	Msg string `json:"msg"`
}
