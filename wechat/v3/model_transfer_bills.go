package wechat

// 商家转账 user_recv_style.type 枚举：用户收款样式（2025-05 新增功能）。
// 文档：https://pay.weixin.qq.com/doc/v3/merchant/4012716434.md
// 使用方式：bm.Set("user_recv_style", gopay.BodyMap{"type": TransferUserRecvStyleRedPacket})
const (
	TransferUserRecvStyleConfirmPage = "CONFIRM_PAGE" // 收款确认页样式（默认）
	TransferUserRecvStyleRedPacket   = "RED_PACKET"   // 红包样式（单笔金额 ≤ 200 元）
)

// 发起转账 Rsp
type TransferBillsRsp struct {
	Code        int            `json:"-"`
	SignInfo    *SignInfo      `json:"-"`
	Response    *TransferBills `json:"response,omitempty"`
	ErrResponse ErrResponse    `json:"err_response,omitempty"`
	Error       string         `json:"-"`
}

// 发起撤销转账 Rsp
type TransferBillsCancelRsp struct {
	Code        int                  `json:"-"`
	SignInfo    *SignInfo            `json:"-"`
	Response    *TransferBillsCancel `json:"response,omitempty"`
	ErrResponse ErrResponse          `json:"err_response,omitempty"`
	Error       string               `json:"-"`
}

// 商户单号查询转账单 Rsp
type TransferBillsQueryRsp struct {
	Code        int                 `json:"-"`
	SignInfo    *SignInfo           `json:"-"`
	Response    *TransferBillsQuery `json:"response,omitempty"`
	ErrResponse ErrResponse         `json:"err_response,omitempty"`
	Error       string              `json:"-"`
}

// 微信单号查询转账单 Rsp
type TransferBillsMerchantQueryRsp struct {
	Code        int                         `json:"-"`
	SignInfo    *SignInfo                   `json:"-"`
	Response    *TransferBillsMerchantQuery `json:"response,omitempty"`
	ErrResponse ErrResponse                 `json:"err_response,omitempty"`
	Error       string                      `json:"-"`
}

// 商户单号申请电子回单 Rsp
type TransferElecsignMerchantRsp struct {
	Code        int                       `json:"-"`
	SignInfo    *SignInfo                 `json:"-"`
	Response    *TransferElecsignMerchant `json:"response,omitempty"`
	ErrResponse ErrResponse               `json:"err_response,omitempty"`
	Error       string                    `json:"-"`
}

// 商户单号查询电子回单 Rsp
type TransferElecsignMerchantQueryRsp struct {
	Code        int                            `json:"-"`
	SignInfo    *SignInfo                      `json:"-"`
	Response    *TransferElecsignMerchantQuery `json:"response,omitempty"`
	ErrResponse ErrResponse                    `json:"err_response,omitempty"`
	Error       string                         `json:"-"`
}

// 微信单号申请电子回单 Rsp
type TransferElecsignRsp struct {
	Code        int               `json:"-"`
	SignInfo    *SignInfo         `json:"-"`
	Response    *TransferElecsign `json:"response,omitempty"`
	ErrResponse ErrResponse       `json:"err_response,omitempty"`
	Error       string            `json:"-"`
}

// 微信单号申请电子回单 Rsp
type TransferElecsignQueryRsp struct {
	Code        int                    `json:"-"`
	SignInfo    *SignInfo              `json:"-"`
	Response    *TransferElecsignQuery `json:"response,omitempty"`
	ErrResponse ErrResponse            `json:"err_response,omitempty"`
	Error       string                 `json:"-"`
}

// =========================================================分割=========================================================

type TransferBills struct {
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	CreateTime     string `json:"create_time"`
	State          string `json:"state"`
	FailReason     string `json:"fail_reason"`
	PackageInfo    string `json:"package_info"`
}

type TransferBillsCancel struct {
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	State          string `json:"state"`
	UpdateTime     string `json:"update_time"`
}

type TransferBillsMerchantQuery struct {
	Mchid          string `json:"mch_id"`
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	Appid          string `json:"appid"`
	State          string `json:"state"`
	TransferAmount int    `json:"transfer_amount"`
	TransferRemark string `json:"transfer_remark"`
	FailReason     string `json:"fail_reason"`
	Openid         string `json:"openid"`
	UserName       string `json:"user_name"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
}

type TransferBillsQuery struct {
	Mchid          string `json:"mch_id"`
	OutBillNo      string `json:"out_bill_no"`
	TransferBillNo string `json:"transfer_bill_no"`
	Appid          string `json:"appid"`
	State          string `json:"state"`
	TransferAmount int    `json:"transfer_amount"`
	TransferRemark string `json:"transfer_remark"`
	FailReason     string `json:"fail_reason"`
	Openid         string `json:"openid"`
	UserName       string `json:"user_name"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
}

type TransferElecsignMerchant struct {
	State      string `json:"state"`
	CreateTime string `json:"create_time"`
}

type TransferElecsignMerchantQuery struct {
	State       string `json:"state"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadURL string `json:"download_url"`
	FailReason  string `json:"fail_reason"`
}

type TransferElecsign struct {
	State      string `json:"state"`
	CreateTime string `json:"create_time"`
}

type TransferElecsignQuery struct {
	State       string `json:"state"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadURL string `json:"download_url"`
	FailReason  string `json:"fail_reason"`
}

// ===================== 用户授权免确认收款（2025-05 新增） =====================

// 发起转账并完成免确认收款授权 Rsp
type TransferPreTransferWithAuthRsp struct {
	Code        int                          `json:"-"`
	SignInfo    *SignInfo                    `json:"-"`
	Response    *TransferPreTransferWithAuth `json:"response,omitempty"`
	ErrResponse ErrResponse                  `json:"err_response,omitempty"`
	Error       string                       `json:"-"`
}

type TransferPreTransferWithAuth struct {
	OutBillNo          string `json:"out_bill_no"`
	TransferBillNo     string `json:"transfer_bill_no"`
	CreateTime         string `json:"create_time"`
	State              string `json:"state"`
	PackageInfo        string `json:"package_info"` // 调起授权页所需的 package 串
	UserDisplayName    string `json:"user_display_name"`
	OutAuthorizationNo string `json:"out_authorization_no"`
}

// 发起免确认收款授权 Rsp
type TransferUserConfirmAuthRsp struct {
	Code        int                      `json:"-"`
	SignInfo    *SignInfo                `json:"-"`
	Response    *TransferUserConfirmAuth `json:"response,omitempty"`
	ErrResponse ErrResponse              `json:"err_response,omitempty"`
	Error       string                   `json:"-"`
}

type TransferUserConfirmAuth struct {
	OutAuthorizationNo string `json:"out_authorization_no"`
	State              string `json:"state"`
	CreateTime         string `json:"create_time"`
	PackageInfo        string `json:"package_info"` // 调起授权页所需的 package 串
}

// 商户单号查询授权结果 Rsp
type TransferUserConfirmAuthQryRsp struct {
	Code        int                         `json:"-"`
	SignInfo    *SignInfo                   `json:"-"`
	Response    *TransferUserConfirmAuthQry `json:"response,omitempty"`
	ErrResponse ErrResponse                 `json:"err_response,omitempty"`
	Error       string                      `json:"-"`
}

type TransferUserConfirmAuthQry struct {
	OutAuthorizationNo string                 `json:"out_authorization_no"`
	Appid              string                 `json:"appid"`
	Openid             string                 `json:"openid,omitempty"`
	UserDisplayName    string                 `json:"user_display_name,omitempty"`
	AuthorizationId    string                 `json:"authorization_id,omitempty"`
	State              string                 `json:"state"`
	AuthorizeTime      string                 `json:"authorize_time,omitempty"`
	CloseInfo          *TransferAuthCloseInfo `json:"close_info,omitempty"`
	TransferSceneId    string                 `json:"transfer_scene_id,omitempty"`
	UserRecvPerception string                 `json:"user_recv_perception,omitempty"`
	CreateTime         string                 `json:"create_time,omitempty"`
	PackageInfo        string                 `json:"package_info,omitempty"` // 当 query 参数 is_display_authorization=true 时返回
}

// 解除免确认收款授权 Rsp
type TransferUserConfirmAuthCloseRsp struct {
	Code        int                           `json:"-"`
	SignInfo    *SignInfo                     `json:"-"`
	Response    *TransferUserConfirmAuthClose `json:"response,omitempty"`
	ErrResponse ErrResponse                   `json:"err_response,omitempty"`
	Error       string                        `json:"-"`
}

type TransferUserConfirmAuthClose struct {
	OutAuthorizationNo string                 `json:"out_authorization_no"`
	Appid              string                 `json:"appid"`
	Openid             string                 `json:"openid,omitempty"`
	UserDisplayName    string                 `json:"user_display_name,omitempty"`
	AuthorizationId    string                 `json:"authorization_id,omitempty"`
	State              string                 `json:"state"`
	AuthorizeTime      string                 `json:"authorize_time,omitempty"`
	CloseInfo          *TransferAuthCloseInfo `json:"close_info,omitempty"`
	TransferSceneId    string                 `json:"transfer_scene_id,omitempty"`
	UserRecvPerception string                 `json:"user_recv_perception,omitempty"`
	CreateTime         string                 `json:"create_time,omitempty"`
}

// 授权关闭信息。
type TransferAuthCloseInfo struct {
	CloseTime   string `json:"close_time,omitempty"`
	CloseReason string `json:"close_reason,omitempty"`
}
