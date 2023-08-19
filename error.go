package gopay

import "errors"

var (
	MissWechatInitParamErr = errors.New("missing wechat init parameter")
	MissAlipayInitParamErr = errors.New("missing alipay init parameter")
	MissPayPalInitParamErr = errors.New("missing paypal init parameter")
	MissAppleInitParamErr  = errors.New("missing apple init parameter")
	MissLakalaInitParamErr = errors.New("missing lakala init parameter")
	MissParamErr           = errors.New("missing required parameter")
	MarshalErr             = errors.New("marshal error")
	UnmarshalErr           = errors.New("unmarshal error")
	SignatureErr           = errors.New("signature error")
	VerifySignatureErr     = errors.New("verify signature error")
	CertNotMatchErr        = errors.New("cert not match error")
	GetSignDataErr         = errors.New("get signature data error")
)
