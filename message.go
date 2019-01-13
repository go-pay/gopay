package gopay

type errorMessage struct {
	ErrorCode int    `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}

type successMessage struct {
	Msg string `json:"msg"`
}
