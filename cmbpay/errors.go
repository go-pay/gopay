package cmbpay

import "fmt"

// 协议层固定常量（详见接口文档 2.2 协议规定）。
const (
	// Version 报文版本号，固定为 0.0.1。
	Version = "0.0.1"
	// Encoding 字符编码，固定为 UTF-8。
	Encoding = "UTF-8"
	// SignMethodSM2 签名方法，02 表示 SM2 国密算法。
	SignMethodSM2 = "02"

	// ResultSuccess 通信/业务成功标识。
	ResultSuccess = "SUCCESS"
	// ResultFail 通信/业务失败标识。
	ResultFail = "FAIL"

	// ErrSystemError 平台内部未知错误，交易结果不明确，需商户发起查询确认
	// （详见接口文档 2.2 第 9、10 条）。
	ErrSystemError = "SYSTERM_ERROR"
)

// APIError 表示招行聚合支付接口返回的错误。
//
// 招行的返回分两层：
//   - ReturnCode 为通信层标识（SUCCESS/FAIL），FAIL 表示报文不合规范（字段超长、
//     非法字符、签名错误等）；
//   - RespCode 为业务层标识（SUCCESS/FAIL），FAIL 表示业务受理失败。
//
// 当 ReturnCode 或 RespCode 为 FAIL 时，ErrCode / RespMsg 给出具体原因。
type APIError struct {
	// ReturnCode 通信层返回码。
	ReturnCode string
	// RespCode 业务层响应码（仅在 ReturnCode 为 SUCCESS 时才有意义）。
	RespCode string
	// ErrCode 具体错误码，取值参见接口文档附录 3。
	ErrCode string
	// RespMsg 错误的详细描述信息。
	RespMsg string
	// ErrDescription 错误信息详情
	ErrDescription string
}

// Error 实现 error 接口。
func (e *APIError) Error() string {
	if e.ErrDescription != "" {
		return fmt.Sprintf("cmbpay: returnCode=%s respCode=%s errCode=%s respMsg=%s errDescription=%s",
			e.ReturnCode, e.RespCode, e.ErrCode, e.RespMsg, e.ErrDescription)
	}
	return fmt.Sprintf("cmbpay: returnCode=%s respCode=%s errCode=%s respMsg=%s",
		e.ReturnCode, e.RespCode, e.ErrCode, e.RespMsg)
}

// IsSystemError 报告该错误是否为平台内部未知错误（errCode=SYSTERM_ERROR）。
// 对于支付/退款/查询类交易，遇到该错误应发起查询直至交易结果明确，
// 而非直接判定为失败（详见接口文档 2.2 第 9、10 条）。
func (e *APIError) IsSystemError() bool {
	return e.ErrCode == ErrSystemError
}

// IsCommFail 报告通信层是否失败（ReturnCode=FAIL）。
func (e *APIError) IsCommFail() bool {
	return e.ReturnCode == ResultFail
}
