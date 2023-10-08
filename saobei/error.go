package saobei

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
	if resp.ReturnCode != "01" {
		return &BizErr{
			Code: resp.ReturnCode,
			Msg:  resp.ReturnMsg,
		}
	}
	//if resp.ResultCode != "01" {
	//	return &BizErr{
	//		Code: resp.ResultCode,
	//		Msg:  resp.ReturnMsg,
	//	}
	//}
	return nil
}

func (e *BizErr) Error() string {
	return fmt.Sprintf(`[%s]%s`, e.Code, e.Msg)
}
