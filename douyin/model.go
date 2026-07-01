package douyin

// ErrResponse 抖音支付错误响应
// 抖音错误响应示例：{"code":"PARAM_ERROR","message":"参数错误","detail":{"field":"...","value":"...","issue":"...","location":"..."}}
type ErrResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Detail  *ErrDetail  `json:"detail,omitempty"`
}

type ErrDetail struct {
	Field    string `json:"field,omitempty"`
	Value    string `json:"value,omitempty"`
	Issue    string `json:"issue,omitempty"`
	Location string `json:"location,omitempty"`
}

// SignInfo 用于同步应答/回调验签的信息
type SignInfo struct {
	HeaderTimestamp string `json:"Douyinpay-Timestamp"`
	HeaderNonce     string `json:"Douyinpay-Nonce"`
	HeaderSignature string `json:"Douyinpay-Signature"`
	HeaderSerial    string `json:"Douyinpay-Serial"`
	SignBody        string `json:"sign_body"`
}

// EmptyRsp 无 Response 数据的响应壳
type EmptyRsp struct {
	Code        int         `json:"-"`
	SignInfo    *SignInfo   `json:"-"`
	ErrResponse ErrResponse `json:"-"`
	Error       string      `json:"-"`
}

// ===================== App 端调起参数 =====================

// AppPayParams App 端调起支付所需参数
type AppPayParams struct {
	AppId     string `json:"appid"`
	PartnerId string `json:"partnerid"`
	PrepayId  string `json:"prepayid"`
	Package   string `json:"package"`
	NonceStr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
}

// JSAPIPayParams JSAPI 前端调起支付所需参数
type JSAPIPayParams struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

// ===================== 通用金额与场景结构 =====================

type Amount struct {
	Total    int    `json:"total"`              // 单位：分
	Currency string `json:"currency,omitempty"` // CNY
}

type RefundAmount struct {
	Refund   int    `json:"refund"`             // 退款金额，单位：分
	Total    int    `json:"total"`              // 原订单金额，单位：分
	Currency string `json:"currency,omitempty"` // CNY
}

type SceneInfo struct {
	PayerClientIp string     `json:"payer_client_ip,omitempty"`
	DeviceId      string     `json:"device_id,omitempty"`
	StoreInfo     *StoreInfo `json:"store_info,omitempty"`
}

type StoreInfo struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	AreaCode string `json:"area_code,omitempty"`
	Address  string `json:"address,omitempty"`
}

type SettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}
