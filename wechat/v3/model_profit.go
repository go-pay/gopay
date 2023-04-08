package wechat

// 请求分账 Rsp
type ProfitShareOrderRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ProfitShareOrder `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询分账结果 Rsp
type ProfitShareOrderQueryRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *ProfitShareOrderQuery `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 请求分账回退 Rsp
type ProfitShareReturnRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *ProfitShareReturn `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 请求分账回退 Rsp
type ProfitShareReturnResultRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *ProfitShareReturnResult `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 解冻剩余资金 Rsp
type ProfitShareOrderUnfreezeRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *ProfitShareOrderUnfreeze `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 查询剩余待分金额 Rsp
type ProfitShareUnsplitAmountRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *ProfitShareUnsplitAmount `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 添加分账接收方 Rsp
type ProfitShareAddReceiverRsp struct {
	Code     int                     `json:"-"`
	SignInfo *SignInfo               `json:"-"`
	Response *ProfitShareAddReceiver `json:"response,omitempty"`
	Error    string                  `json:"-"`
}

// 删除分账接收方 Rsp
type ProfitShareDeleteReceiverRsp struct {
	Code     int                        `json:"-"`
	SignInfo *SignInfo                  `json:"-"`
	Response *ProfitShareDeleteReceiver `json:"response,omitempty"`
	Error    string                     `json:"-"`
}

type ProfitShareMerchantConfigsRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *ProfitShareMerchantConfigs `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 申请分账账单 Rsp
type ProfitShareBillsRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ProfitShareBills `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// =========================================================分割=========================================================

type ProfitShareOrder struct {
	SubMchid      string                   `json:"sub_mchid"`           // 二级商户号
	TransactionId string                   `json:"transaction_id"`      // 微信订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户分账单号
	OrderId       string                   `json:"order_id"`            // 微信分账单号
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareOrderQuery struct {
	SubMchid      string                   `json:"sub_mchid,omitempty"` // 子商户号，即分账的出资商户号【服务商模式】
	TransactionId string                   `json:"transaction_id"`      // 微信订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户分账单号
	OrderId       string                   `json:"order_id"`            // 微信分账单号
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareReturn struct {
	SubMchid    string `json:"sub_mchid"`             // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	OrderId     string `json:"order_id"`              // 微信分账单号，微信系统返回的唯一标识
	OutOrderNo  string `json:"out_order_no"`          // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutReturnNo string `json:"out_return_no"`         // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	ReturnId    string `json:"return_id"`             // 微信分账回退单号，微信系统返回的唯一标识
	ReturnMchid string `json:"return_mchid"`          // 只能对原分账请求中成功分给商户接收方进行回退
	Amount      int    `json:"amount"`                // 需要从分账接收方回退的金额，单位为分，只能为整数
	Description string `json:"description"`           // 分账回退的原因描述
	Result      string `json:"result"`                // 回退结果: PROCESSING：处理中，SUCCESS：已成功，FAILED：已失败
	FailReason  string `json:"fail_reason,omitempty"` // 失败原因: ACCOUNT_ABNORMAL : 分账接收方账户异常，TIME_OUT_CLOSED : 超时关单
	CreateTime  string `json:"create_time"`           // 分账回退创建时间
	FinishTime  string `json:"finish_time"`           // 分账回退完成时间
}

type ProfitShareReturnResult struct {
	SubMchid    string `json:"sub_mchid,omitempty"`   // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	OrderId     string `json:"order_id"`              // 微信分账单号，微信系统返回的唯一标识
	OutOrderNo  string `json:"out_order_no"`          // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutReturnNo string `json:"out_return_no"`         // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	ReturnId    string `json:"return_id"`             // 微信分账回退单号，微信系统返回的唯一标识
	ReturnMchid string `json:"return_mchid"`          // 只能对原分账请求中成功分给商户接收方进行回退
	Amount      int    `json:"amount"`                // 需要从分账接收方回退的金额，单位为分，只能为整数
	Description string `json:"description"`           // 分账回退的原因描述
	Result      string `json:"result"`                // 回退结果: PROCESSING：处理中，SUCCESS：已成功，FAILED：已失败
	FailReason  string `json:"fail_reason,omitempty"` // 失败原因: ACCOUNT_ABNORMAL : 分账接收方账户异常，TIME_OUT_CLOSED : 超时关单
	CreateTime  string `json:"create_time"`           // 分账回退创建时间
	FinishTime  string `json:"finish_time"`           // 分账回退完成时间
}

type ProfitShareOrderUnfreeze struct {
	SubMchid      string                   `json:"sub_mchid,omitempty"` // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	TransactionId string                   `json:"transaction_id"`      // 微信支付订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OrderId       string                   `json:"order_id"`            // 微信分账单号，微信系统返回的唯一标识
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareUnsplitAmount struct {
	TransactionId string `json:"transaction_id"` // 微信支付订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额，整数，单位为分
}

// 分账接收方
type ProfitSharingReceiver struct {
	Amount      int    `json:"amount"`      // 分账金额
	Description string `json:"description"` // 分账描述
	Type        string `json:"type"`        // 分账接收方类型
	Account     string `json:"account"`     // 分账接收方帐号
	Result      string `json:"result"`      // 分账结果,PENDING：待分账,SUCCESS：分账成功,CLOSED：已关闭
	FailReason  string `json:"fail_reason"` // 分账失败原因ACCOUNT_ABNORMAL : 分账接收账户异常、NO_RELATION : 分账关系已解除、RECEIVER_HIGH_RISK : 高风险接收方、RECEIVER_REAL_NAME_NOT_VERIFIED : 接收方未实名、NO_AUTH : 分账权限已解除
	CreateTime  string `json:"create_time"` // 分账创建时间,遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
	FinishTime  string `json:"finish_time"` // 分账完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
	DetailId    string `json:"detail_id"`   // 分账明细单号, 微信分账明细单号，每笔分账业务执行的明细单号，可与资金账单对账使用
}

type ProfitShareAddReceiver struct {
	SubMchid       string `json:"sub_mchid,omitempty"`       // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	Type           string `json:"type"`                      // 分账接收方类型MERCHANT_ID：商户ID,PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account        string `json:"account"`                   // 分账接收方帐号
	Name           string `json:"name,omitempty"`            // 分账接收方类型是MERCHANT_ID时，是商户全称（必传），当商户是小微商户或个体户时，是开户人姓名 分账接收方类型是PERSONAL_OPENID时，是个人姓名（选传，传则校验）
	RelationType   string `json:"relation_type"`             // 商户与接收方的关系。STORE：门店STAFF：员工	STORE_OWNER：店主	PARTNER：合作伙伴	HEADQUARTER：总部	BRAND：品牌方	DISTRIBUTOR：分销商	USER：用户	SUPPLIER： 供应商	CUSTOM：自定义
	CustomRelation string `json:"custom_relation,omitempty"` // 子商户与接收方具体的关系，本字段最多10个字。当字段relation_type的值为CUSTOM时，本字段必填;当字段relation_type的值不为CUSTOM时，本字段无需填写。
}

type ProfitShareDeleteReceiver struct {
	SubMchid string `json:"sub_mchid,omitempty"` // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	Type     string `json:"type"`                // 分账接收方类型MERCHANT_ID：商户ID,PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account  string `json:"account"`             // 分账接收方帐号
}

type ProfitShareMerchantConfigs struct {
	SubMchId string `json:"sub_mchid"` // 子商户号
	MaxRatio int    `json:"max_ratio"` // 最大分账比例 (单位万分比，比如2000表示20%)
}

type ProfitShareBills struct {
	DownloadUrl string `json:"download_url"` // 下载地址	原始账单（gzip需要解压缩）的摘要值，用于校验文件的完整性。
	HashType    string `json:"hash_type"`    // 哈希类型 原始账单（gzip需要解压缩）的摘要算法，用于校验文件的完整性
	HashValue   string `json:"hash_value"`   // 哈希值	供下一步请求账单文件的下载地址，该地址30s内有效。
}
