package wechat

// 发起批量转账 Rsp
type TransferRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Transfer `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// 微信批次单号查询批次单 Rsp
type TransferQueryRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *TransferQuery `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 微信批次单号查询批次单（服务商） Rsp
type PartnerTransferQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *PartnerTransferQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 微信明细单号查询明细单 Rsp
type TransferDetailRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *TransferDetailQuery `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 微信明细单号查询明细单（服务商） Rsp
type PartnerTransferDetailRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *PartnerTransferDetail `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 商家批次单号查询批次单 Rsp
type TransferMerchantQueryRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *TransferMerchantQuery `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 商家批次单号查询批次单（服务商） Rsp
type PartnerTransferMerchantQueryRsp struct {
	Code     int                           `json:"-"`
	SignInfo *SignInfo                     `json:"-"`
	Response *PartnerTransferMerchantQuery `json:"response,omitempty"`
	Error    string                        `json:"-"`
}

// 商家明细单号查询明细单 Rsp
type TransferMerchantDetailRsp struct {
	Code     int                     `json:"-"`
	SignInfo *SignInfo               `json:"-"`
	Response *TransferMerchantDetail `json:"response,omitempty"`
	Error    string                  `json:"-"`
}

// 商家明细单号查询明细单（服务商） Rsp
type PartnerTransferMerchantDetailRsp struct {
	Code     int                            `json:"-"`
	SignInfo *SignInfo                      `json:"-"`
	Response *PartnerTransferMerchantDetail `json:"response,omitempty"`
	Error    string                         `json:"-"`
}

// 转账电子回单申请受理 Rsp
type TransferReceiptRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *TransferReceipt `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询转账电子回单 Rsp
type TransferReceiptQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *TransferReceiptQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 转账明细电子回单受理 Rsp
type TransferDetailReceiptRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *TransferDetailReceipt `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 查询转账明细电子回单受理结果 Rsp
type TransferDetailReceiptQueryRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *TransferDetailReceiptQuery `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// =========================================================分割=========================================================

type Transfer struct {
	OutBatchNo string `json:"out_batch_no"` // 商户系统内部的商家批次单号
	BatchId    string `json:"batch_id"`     // 微信批次单号，微信商家转账系统返回的唯一标识
	CreateTime string `json:"create_time"`  // 批次受理成功时返回
}

type TransferQuery struct {
	TransferBatch      *TransferBatch    `json:"transfer_batch"`                 // 转账批次单基本信息
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferBatch struct {
	Mchid         string `json:"mchid"`                    // 微信支付分配的商户号
	OutBatchNo    string `json:"out_batch_no"`             // 商户系统内部的商家批次单号
	BatchId       string `json:"batch_id"`                 // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid         string `json:"appid"`                    // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	BatchStatus   string `json:"batch_status"`             // 批次状态
	BatchType     string `json:"batch_type"`               // 批次类型
	BatchName     string `json:"batch_name"`               // 该笔批量转账的名称
	BatchRemark   string `json:"batch_remark"`             // 转账说明，UTF8编码，最多允许32个字符
	CloseReason   string `json:"close_reason,omitempty"`   // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount   int    `json:"total_amount"`             // 转账金额单位为分
	TotalNum      int    `json:"total_num"`                // 一个转账批次单最多发起三千笔转账
	CreateTime    string `json:"create_time,omitempty"`    // 批次受理成功时返回
	UpdateTime    string `json:"update_time,omitempty"`    // 批次最近一次状态变更的时间
	SuccessAmount int    `json:"success_amount,omitempty"` // 转账成功的金额，单位为分
	SuccessNum    int    `json:"success_num,omitempty"`    // 转账成功的笔数
	FailAmount    int    `json:"fail_amount,omitempty"`    // 转账失败的金额，单位为分
	FailNum       int    `json:"fail_num,omitempty"`       // 转账失败的笔数
}

type TransferDetail struct {
	DetailId     string `json:"detail_id"`     // 微信明细单号
	OutDetailNo  string `json:"out_detail_no"` // 商家明细单号
	DetailStatus string `json:"detail_status"` // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
}

type PartnerTransferQuery struct {
	SpMchid            string            `json:"sp_mchid"`                       // 微信支付分配的服务商商户号
	SubMchid           string            `json:"sub_mchid"`                      // 微信支付分配的特约商户号
	OutBatchNo         string            `json:"out_batch_no"`                   // 商户系统内部的商家批次单号
	BatchId            string            `json:"batch_id"`                       // 微信批次单号，微信商家转账系统返回的唯一标识
	SpAppid            string            `json:"sp_appid,omitempty"`             // 微信分配的服务商商户公众账号Id，特约商户授权类型为FUND_AUTHORIZATION_TYPE时才有该字段
	SubAppid           string            `json:"sub_appid"`                      // 微信分配的特约商户公众账号Id。特约商户appid
	BatchStatus        string            `json:"batch_status"`                   // 批次状态
	BatchType          string            `json:"batch_type"`                     // 批次类型
	AuthorizationType  string            `json:"authorization_type"`             // 特约商户授权类型
	BatchName          string            `json:"batch_name"`                     // 该笔批量转账的名称
	BatchRemark        string            `json:"batch_remark"`                   // 转账说明，UTF8编码，最多允许32个字符
	CloseReason        string            `json:"close_reason,omitempty"`         // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount        int               `json:"total_amount"`                   // 转账金额单位为分
	TotalNum           int               `json:"total_num"`                      // 一个转账批次单最多发起三千笔转账
	CreateTime         string            `json:"create_time,omitempty"`          // 批次受理成功时返回
	UpdateTime         string            `json:"update_time,omitempty"`          // 批次最近一次状态变更的时间
	SuccessAmount      int               `json:"success_amount,omitempty"`       // 转账成功的金额，单位为分
	SuccessNum         int               `json:"success_num,omitempty"`          // 转账成功的笔数
	FailAmount         int               `json:"fail_amount,omitempty"`          // 转账失败的金额，单位为分
	FailNum            int               `json:"fail_num,omitempty"`             // 转账失败的笔数
	TransferPurpose    string            `json:"transfer_purpose"`               // 批量转账用途
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferDetailQuery struct {
	Mchid          string `json:"mchid"`                 // 微信支付分配的商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	UserName       string `json:"user_name"`             // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type PartnerTransferDetail struct {
	SpMchid        string `json:"sp_mchid"`              // 微信支付分配的服务商商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	Username       string `json:"username"`              // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type TransferMerchantQuery struct {
	TransferBatch      *TransferBatch    `json:"transfer_batch"`                 // 转账批次单基本信息
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
	Offset             int               `json:"offset,omitempty"`               // 该次请求资源（转账明细单）的起始位置
	Limit              int               `json:"limit,omitempty"`                // 该次请求可返回的最大资源（转账明细单）条数
}

type PartnerTransferMerchantQuery struct {
	SpMchid            string            `json:"sp_mchid"`                       // 微信支付分配的服务商商户号
	SubMchid           string            `json:"sub_mchid"`                      // 微信支付分配的特约商户号
	OutBatchNo         string            `json:"out_batch_no"`                   // 商户系统内部的商家批次单号
	BatchId            string            `json:"batch_id"`                       // 微信批次单号，微信商家转账系统返回的唯一标识
	SpAppid            string            `json:"sp_appid,omitempty"`             // 微信分配的服务商商户公众账号Id，特约商户授权类型为FUND_AUTHORIZATION_TYPE时才有该字段
	SubAppid           string            `json:"sub_appid"`                      // 微信分配的特约商户公众账号Id。特约商户appid
	BatchStatus        string            `json:"batch_status"`                   // 批次状态
	BatchType          string            `json:"batch_type"`                     // 批次类型
	AuthorizationType  string            `json:"authorization_type"`             // 特约商户授权类型
	BatchName          string            `json:"batch_name"`                     // 该笔批量转账的名称
	BatchRemark        string            `json:"batch_remark"`                   // 转账说明，UTF8编码，最多允许32个字符
	CloseReason        string            `json:"close_reason,omitempty"`         // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount        int               `json:"total_amount"`                   // 转账金额单位为分
	TotalNum           int               `json:"total_num"`                      // 一个转账批次单最多发起三千笔转账
	CreateTime         string            `json:"create_time,omitempty"`          // 批次受理成功时返回
	UpdateTime         string            `json:"update_time,omitempty"`          // 批次最近一次状态变更的时间
	SuccessAmount      int               `json:"success_amount,omitempty"`       // 转账成功的金额，单位为分
	SuccessNum         int               `json:"success_num,omitempty"`          // 转账成功的笔数
	FailAmount         int               `json:"fail_amount,omitempty"`          // 转账失败的金额，单位为分
	FailNum            int               `json:"fail_num,omitempty"`             // 转账失败的笔数
	TransferPurpose    string            `json:"transfer_purpose"`               // 批量转账用途
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferMerchantDetail struct {
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	UserName       string `json:"user_name"`             // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type PartnerTransferMerchantDetail struct {
	SpMchid        string `json:"sp_mchid"`              // 微信支付分配的服务商商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	Username       string `json:"username"`              // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type TransferReceipt struct {
	OutBatchNo      string `json:"out_batch_no"`               // 商户系统内部的商家批次单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
	CreateTime      string `json:"create_time,omitempty"`      // 电子签章单创建时间
	UpdateTime      string `json:"update_time,omitempty"`      // 电子签章单最近一次状态变更的时间
}

type TransferReceiptQuery struct {
	OutBatchNo      string `json:"out_batch_no"`               // 商户系统内部的商家批次单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
	CreateTime      string `json:"create_time,omitempty"`      // 电子签章单创建时间
	UpdateTime      string `json:"update_time,omitempty"`      // 电子签章单最近一次状态变更的时间
}

type TransferDetailReceipt struct {
	AcceptType      string `json:"accept_type"`                // 电子回单受理类型
	OutBatchNo      string `json:"out_batch_no,omitempty"`     // 商户系统内部的商家批次单号
	OutDetailNo     string `json:"out_detail_no"`              // 商家明细单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
}

type TransferDetailReceiptQuery struct {
	AcceptType      string `json:"accept_type"`                // 电子回单受理类型
	OutBatchNo      string `json:"out_batch_no,omitempty"`     // 商户系统内部的商家批次单号
	OutDetailNo     string `json:"out_detail_no"`              // 商家明细单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
}
