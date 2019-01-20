package gopay

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

type ReturnMessage struct {
	ReturnCode string `json:"return_code"`
	ReturnMsg  string `json:"return_msg"`
}
