package apple

import "fmt"

// StatusCodeErr 用于判断Apple的status_code错误
type StatusCodeErr struct {
	ErrorCode    int    `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// statusCodeErrCheck 检查状态码是否为非200错误
func statusCodeErrCheck(errRsp StatusCodeErr) error {
	if errRsp.ErrorCode != 0 {
		return &StatusCodeErr{
			ErrorCode:    errRsp.ErrorCode,
			ErrorMessage: errRsp.ErrorMessage,
		}
	}
	return nil
}

func (e *StatusCodeErr) Error() string {
	return fmt.Sprintf(`{"errorCode":"%d","errorMessage":"%s"}`, e.ErrorCode, e.ErrorMessage)
}

func IsStatusCodeError(err error) (*StatusCodeErr, bool) {
	if bizErr, ok := err.(*StatusCodeErr); ok {
		return bizErr, true
	}
	return nil, false
}
