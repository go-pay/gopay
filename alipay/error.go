package alipay

import "fmt"

// BizErr 用于判断支付宝的业务逻辑是否有错误
type BizErr struct {
	Code    string
	Msg     string
	SubCode string
	SubMsg  string
}

// bizErrCheck 检查业务码是否为10000 否则返回一个BizErr
func bizErrCheck(errRsp ErrorResponse) error {
	if errRsp.Code != "10000" {
		return &BizErr{
			Code:    errRsp.SubCode,
			Msg:     errRsp.Msg,
			SubCode: errRsp.SubCode,
			SubMsg:  errRsp.SubMsg,
		}
	}
	return nil
}

func (e *BizErr) Error() string {
	return fmt.Sprintf(`{"code": "%s","msg": "%s","sub_code": "%s","sub_msg": "%s"}`, e.Code, e.Msg, e.SubCode, e.SubMsg)
}

func AsBizError(err error) *BizErr {
	if bizerr, ok := err.(*BizErr); ok && bizerr != nil {
		return bizerr
	}
	return nil
}
