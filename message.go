package gopay

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
	OK      = "OK"
)

type ReturnMessage struct {
	ReturnCode string `json:"return_code"`
	ReturnMsg  string `json:"return_msg"`
}
