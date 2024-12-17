package alipay

import (
	"fmt"
)

// BizErr 用于判断支付宝的业务逻辑是否有错误
type BizErr struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

// bizErrCheck 检查业务码是否为10000 否则返回一个BizErr
func bizErrCheck(errRsp ErrorResponse) error {
	if errRsp.Code != "10000" {
		return &BizErr{
			Code:    errRsp.Code,
			Msg:     errRsp.Msg,
			SubCode: errRsp.SubCode,
			SubMsg:  errRsp.SubMsg,
		}
	}
	return nil
}

// bizErrCheckTradePay 检查业务码是否为10000、10003，否则返回一个BizErr
func bizErrCheckTradePay(errRsp ErrorResponse) error {
	if errRsp.Code != "10000" && errRsp.Code != "10003" {
		return &BizErr{
			Code:    errRsp.Code,
			Msg:     errRsp.Msg,
			SubCode: errRsp.SubCode,
			SubMsg:  errRsp.SubMsg,
		}
	}
	return nil
}

func (e *BizErr) Error() string {
	return fmt.Sprintf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, e.Code, e.Msg, e.SubCode, e.SubMsg)
}

func IsBizError(err error) (*BizErr, bool) {
	if bizErr, ok := err.(*BizErr); ok {
		return bizErr, true
	}
	return nil, false
}
