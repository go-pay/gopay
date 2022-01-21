package alipay

import "fmt"

type ErrorResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

func (ersp ErrorResponse) HasBizError() bool {
	return ersp.Code != "10000"
}

func (ersp ErrorResponse) Error() string {
	return fmt.Sprintf(`{"code": "%s","msg": "%s","sub_code": "%s","sub_msg": "%s"}`, ersp.Code, ersp.Msg, ersp.SubCode, ersp.SubMsg)
}
