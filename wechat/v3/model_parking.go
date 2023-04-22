package wechat


type VehicleParkingQueryRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *VehicleParkingQuery `json:"response,omitempty"`
	Error    string               `json:"-"`
}

type VehicleParkingInRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *VehicleParkingIn `json:"response,omitempty"`
	Error    string            `json:"-"`
}

type VehicleParkingFeeRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *VehicleParkingFee `json:"response,omitempty"`
	Error    string             `json:"-"`
}

type VehicleParkingOrderRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *VehicleParkingOrder `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// =========================================================分割=========================================================


type VehicleParkingQuery struct {
	PlateNumber     string `json:"plate_number"` // 车牌号，仅包括省份+车牌，不包括特殊字符。
	PlateColor      string `json:"plate_color"`  // 车牌颜色，BLUE：蓝色，GREEN：绿色，YELLOW：黄色，BLACK：黑色，WHITE：白色，LIMEGREEN：黄绿色
	ServiceOpenTime string `json:"service_open_time,omitempty"`
	Openid          string `json:"openid"`
	ServiceState    string `json:"service_state"`
}

type VehicleParkingIn struct {
	Id           string `json:"id"`
	OutParkingNo string `json:"out_parking_no"` // 商户入场id
	PlateNumber  string `json:"plate_number"`   // 车牌号，仅包括省份+车牌，不包括特殊字符。
	PlateColor   string `json:"plate_color"`    // 车牌颜色，BLUE：蓝色，GREEN：绿色，YELLOW：黄色，BLACK：黑色，WHITE：白色，LIMEGREEN：黄绿色
	StartTime    string `json:"start_time"`     // 入场时间
	ParkingName  string `json:"parking_name"`   // 所在停车位车场的名称
	FreeDuration int    `json:"free_duration"`  // 停车场的免费停车时长，单位为秒
	State        string `json:"state"`
	BlockReason  string `json:"block_reason"`
}

type VehicleParkingFee struct {
	Appid                 string             `json:"appid"`
	SubAppid              string             `json:"sub_appid,omitempty"`
	SpMchid               string             `json:"sp_mchid"`
	SubMchid              string             `json:"sub_mchid,omitempty"`
	Description           string             `json:"description"`
	CreateTime            string             `json:"create_time"`
	OutTradeNo            string             `json:"out_trade_no"`
	TransactionId         string             `json:"transaction_id"`
	TradeState            string             `json:"trade_state"`
	TradeStateDescription string             `json:"trade_state_description"`
	SuccessTime           string             `json:"success_time"`
	BankType              string             `json:"bank_type"`
	UserRepaid            string             `json:"user_repaid"`
	Attach                string             `json:"attach"`
	TradeScene            string             `json:"trade_scene"`
	ParkingInfo           *ParkingInfo       `json:"parking_info"`
	Payer                 *Payer             `json:"payer"`
	Amount                *Amount            `json:"amount"`
	PromotionDetail       []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type ParkingInfo struct {
	ParkingId        string `json:"parking_id"`        // 停车入场id
	PlateNumber      string `json:"plate_number"`      // 车牌号，仅包括省份+车牌，不包括特殊字符。
	PlateColor       string `json:"plate_color"`       // 车牌颜色，BLUE：蓝色，GREEN：绿色，YELLOW：黄色，BLACK：黑色，WHITE：白色，LIMEGREEN：黄绿色
	StartTime        string `json:"start_time"`        // 入场时间
	EndTime          string `json:"end_time"`          // 出场时间
	ParkingName      string `json:"parking_name"`      // 所在停车位车场的名称
	ChargingDuration int    `json:"charging_duration"` // 计费的时间长，单位为秒
	DeviceId         string `json:"device_id"`         // 停车场设备id
}

type VehicleParkingOrder struct {
	Appid                 string             `json:"appid"`
	SubAppid              string             `json:"sub_appid,omitempty"`
	SpMchid               string             `json:"sp_mchid"`
	SubMchid              string             `json:"sub_mchid,omitempty"`
	Description           string             `json:"description"`
	CreateTime            string             `json:"create_time"`
	OutTradeNo            string             `json:"out_trade_no"`
	TransactionId         string             `json:"transaction_id"`
	TradeState            string             `json:"trade_state"`
	TradeStateDescription string             `json:"trade_state_description"`
	SuccessTime           string             `json:"success_time"`
	BankType              string             `json:"bank_type"`
	UserRepaid            string             `json:"user_repaid"`
	Attach                string             `json:"attach"`
	TradeScene            string             `json:"trade_scene"`
	ParkingInfo           *ParkingInfo       `json:"parking_info"`
	Payer                 *Payer             `json:"payer"`
	Amount                *Amount            `json:"amount"`
	PromotionDetail       []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}
