package allinpay

import (
	"fmt"
)

// BizErr 用于判断通联的业务逻辑是否有错误
type BizErr struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// bizErrCheck 检查返回码是否为SUCCESS 否则返回一个BizErr
func bizErrCheck(resp RspBase) error {
	if resp.RetCode != "SUCCESS" {
		return &BizErr{
			Code: resp.RetCode,
			Msg:  resp.RetMsg,
		}
	}
	return nil
}

func (e *BizErr) Error() string {
	return fmt.Sprintf(`{"code":"%s","msg":"%s"}`, e.Code, e.Msg)
}
