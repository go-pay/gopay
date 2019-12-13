package gopay

const (
	NULL       = ""
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
	SUCCESS    = "SUCCESS"
	FAIL       = "FAIL"
	OK         = "OK"
	Version    = "2.0.3"
)

type ReturnMessage struct {
	ReturnCode string `json:"return_code"`
	ReturnMsg  string `json:"return_msg"`
}
