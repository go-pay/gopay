package douyin

// ProfitReceiverReq 请求端分账接收方
// 若 Name 需要加密（type=MERCHANT_ID 时必传商户名称；type=PERSONAL_OPENID 时选传个人姓名），
// 请先调用 (c *Client).EncryptText(name) 得到密文后再赋值给 Name。
type ProfitReceiverReq struct {
	Type        string `json:"type"`                  // MERCHANT_ID / PERSONAL_OPENID
	Account     string `json:"account"`               // 商户号 或 个人 openid
	Name        string `json:"name,omitempty"`        // 已加密的接收方名称
	Amount      int    `json:"amount"`                // 分账金额（分）
	Description string `json:"description,omitempty"` // 分账描述
}

// ProfitReceiver 响应/通知中的分账接收方（含结果字段）
type ProfitReceiver struct {
	Amount      int    `json:"amount"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Account     string `json:"account,omitempty"`
	Result      string `json:"result,omitempty"`      // PENDING / SUCCESS / CLOSED
	FailReason  string `json:"fail_reason,omitempty"` // ACCOUNT_ABNORMAL / RECEIVER_HIGH_RISK / PAYER_ACCOUNT_ABNORMAL
	CreateTime  string `json:"create_time,omitempty"`
	FinishTime  string `json:"finish_time,omitempty"`
	DetailId    string `json:"detail_id,omitempty"` // 分账明细单号
}

// Profit 分账单（申请/查询/完结 通用响应）
type Profit struct {
	Mchid             string            `json:"mchid,omitempty"`
	TransactionId     string            `json:"transaction_id,omitempty"`
	OutOrderNo        string            `json:"out_order_no,omitempty"`
	OrderId           string            `json:"order_id,omitempty"`
	State             string            `json:"state,omitempty"` // PROCESSING / FINISHED
	Receivers         []*ProfitReceiver `json:"receivers,omitempty"`
	FinishAmount      int               `json:"finish_amount,omitempty"`
	FinishDescription string            `json:"finish_description,omitempty"`
	SplitFinishTime   string            `json:"split_finish_time,omitempty"`
}

type ProfitRsp struct {
	Code        int         `json:"-"`
	SignInfo    *SignInfo   `json:"-"`
	Response    *Profit     `json:"-"`
	ErrResponse ErrResponse `json:"-"`
	Error       string      `json:"-"`
}

// ProfitReturn 分账回退
type ProfitReturn struct {
	OrderId     string `json:"order_id,omitempty"`
	OutOrderNo  string `json:"out_order_no,omitempty"`
	OutReturnNo string `json:"out_return_no,omitempty"`
	ReturnId    string `json:"return_id,omitempty"`
	ReturnMchid string `json:"return_mchid,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Result      string `json:"result,omitempty"`      // SUCCESS / PROCESSING / FAIL
	FailReason  string `json:"fail_reason,omitempty"` // e.g. BALANCE_NOT_ENOUGH
	CreateTime  string `json:"create_time,omitempty"`
	FinishTime  string `json:"finish_time,omitempty"`
}

type ProfitReturnRsp struct {
	Code        int           `json:"-"`
	SignInfo    *SignInfo     `json:"-"`
	Response    *ProfitReturn `json:"-"`
	ErrResponse ErrResponse   `json:"-"`
	Error       string        `json:"-"`
}

// ProfitBalance 剩余待分账金额
type ProfitBalance struct {
	Mchid         string `json:"mchid,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
	UnsplitAmount int    `json:"unsplit_amount"` // 剩余待分账金额（分）
}

type ProfitBalanceRsp struct {
	Code        int            `json:"-"`
	SignInfo    *SignInfo      `json:"-"`
	Response    *ProfitBalance `json:"-"`
	ErrResponse ErrResponse    `json:"-"`
	Error       string         `json:"-"`
}

// ProfitReceiverInfo 添加/删除分账接收方
type ProfitReceiverInfo struct {
	Mchid        string `json:"mchid,omitempty"`
	Appid        string `json:"appid,omitempty"`
	Type         string `json:"type,omitempty"`
	Account      string `json:"account,omitempty"`
	Name         string `json:"name,omitempty"`          // 已加密的名称（同 ProfitReceiverReq.Name）
	RelationType string `json:"relation_type,omitempty"` // STORE / STAFF / STORE_OWNER / ...
}

type ProfitReceiverRsp struct {
	Code        int                 `json:"-"`
	SignInfo    *SignInfo           `json:"-"`
	Response    *ProfitReceiverInfo `json:"-"`
	ErrResponse ErrResponse         `json:"-"`
	Error       string              `json:"-"`
}

// DecryptProfitResult 分账结果通知（event_type=ASYNC_SPLIT.FINISH）解密后的明文
type DecryptProfitResult struct {
	Mchid             string            `json:"mchid,omitempty"`
	OrderId           string            `json:"order_id,omitempty"`
	OutOrderNo        string            `json:"out_order_no,omitempty"`
	TransactionId     string            `json:"transaction_id,omitempty"`
	State             string            `json:"state,omitempty"`
	Receivers         []*ProfitReceiver `json:"receivers,omitempty"`
	FinishAmount      int               `json:"finish_amount,omitempty"`
	FinishDescription string            `json:"finish_description,omitempty"`
	// 抖音文档中该字段有拼写差异（split_finish_time vs spilt_finish_time），两者均反序列化保留
	SplitFinishTime string `json:"split_finish_time,omitempty"`
	SpiltFinishTime string `json:"spilt_finish_time,omitempty"`
}

// DecryptProfitDynamic 分账动账通知（event_type=SPLIT.SUCCESS）解密后的明文
type DecryptProfitDynamic struct {
	Mchid         string          `json:"mchid,omitempty"`
	TransactionId string          `json:"transaction_id,omitempty"`
	OutOrderNo    string          `json:"out_order_no,omitempty"`
	OrderId       string          `json:"order_id,omitempty"`
	Receiver      *ProfitReceiver `json:"receiver,omitempty"`
	SuccessTime   string          `json:"success_time,omitempty"`
}
