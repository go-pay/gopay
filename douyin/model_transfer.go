package douyin

// TransferSceneReportInfo 转账场景报备信息
type TransferSceneReportInfo struct {
	InfoType    string `json:"info_type"`
	InfoContent string `json:"info_content"`
}

// TransferBill 转账单信息（创建/查询/通知 共用结构）
// 创建响应仅含：OutBillNo、TransferBillNo、State、CreateTime
// 查询/通知响应额外含：MchId、Appid、TransferAmount、TransferRemark、FailReason、Openid、UserName、UpdateTime
// 注意：通知里 TransferAmount 是字符串（例如 "100"），故此处保留为 string 以兼容两种来源
type TransferBill struct {
	MchId          string `json:"mch_id,omitempty"`
	Appid          string `json:"appid,omitempty"`
	OutBillNo      string `json:"out_bill_no,omitempty"`
	TransferBillNo string `json:"transfer_bill_no,omitempty"`
	State          string `json:"state,omitempty"`           // ACCEPTED / TRANSFERING / SUCCESS / FAIL
	TransferAmount int    `json:"transfer_amount,omitempty"` // 查询接口返回 int（单位：分）
	TransferRemark string `json:"transfer_remark,omitempty"`
	FailReason     string `json:"fail_reason,omitempty"`
	Openid         string `json:"openid,omitempty"`
	UserName       string `json:"user_name,omitempty"`
	CreateTime     string `json:"create_time,omitempty"`
	UpdateTime     string `json:"update_time,omitempty"`
}

// TransferRsp 转账 / 转账查询 响应壳
type TransferRsp struct {
	Code        int           `json:"-"`
	SignInfo    *SignInfo     `json:"-"`
	Response    *TransferBill `json:"-"`
	ErrResponse ErrResponse   `json:"-"`
	Error       string        `json:"-"`
}

// DecryptTransferResult 转账结果通知（event_type=TRANSFER.SUCCESS）解密后的明文
// 抖音文档中 transfer_amount 在通知里为字符串，故此结构保留为 string
type DecryptTransferResult struct {
	MchId          string `json:"mch_id,omitempty"`
	OutBillNo      string `json:"out_bill_no,omitempty"`
	TransferBillNo string `json:"transfer_bill_no,omitempty"`
	State          string `json:"state,omitempty"`
	TransferAmount string `json:"transfer_amount,omitempty"`
	Openid         string `json:"openid,omitempty"`
	CreateTime     string `json:"create_time,omitempty"`
	UpdateTime     string `json:"update_time,omitempty"`
}
