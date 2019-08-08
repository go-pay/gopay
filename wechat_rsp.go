package gopay

type WeChatUnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
	TradeType  string `xml:"trade_type,omitempty"`
	PrepayId   string `xml:"prepay_id,omitempty"`
	CodeUrl    string `xml:"code_url,omitempty"`
	MwebUrl    string `xml:"mweb_url,omitempty"`
}

type WeChatQueryOrderResponse struct {
	ReturnCode         string `xml:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty"`
	Openid             string `xml:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty"`
	TradeState         string `xml:"trade_state,omitempty"`
	BankType           string `xml:"bank_type,omitempty"`
	TotalFee           int    `xml:"total_fee,omitempty"`
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty"`
	CashFee            int    `xml:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty"`
	CouponFee          int    `xml:"coupon_fee,omitempty"`
	CouponCount        int    `xml:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty"`
	CouponFee0         int    `xml:"coupon_fee_0,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty"`
	TradeStateDesc     string `xml:"trade_state_desc,omitempty"`
}

type WeChatCloseOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
}

type WeChatReverseResponse struct {
	ReturnCode string `xml:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
	Recall     string `xml:"recall,omitempty"`
}

type WeChatRefundResponse struct {
	ReturnCode          string `xml:"return_code,omitempty"`
	ReturnMsg           string `xml:"return_msg,omitempty"`
	ResultCode          string `xml:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty"`
	Appid               string `xml:"appid,omitempty"`
	MchId               string `xml:"mch_id,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty"`
	TransactionId       string `xml:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty"`
	RefundFee           int    `xml:"refund_fee,omitempty"`
	SettlementRefundFee int    `xml:"settlement_refund_fee,omitempty"`
	TotalFee            int    `xml:"total_fee,omitempty"`
	SettlementTotalFee  int    `xml:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty"`
	CashFee             int    `xml:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty"`
	CashRefundFee       int    `xml:"cash_refund_fee,omitempty"`
	CouponType0         string `xml:"coupon_type_0,omitempty"`
	CouponRefundFee     int    `xml:"coupon_refund_fee,omitempty"`
	CouponRefundFee0    int    `xml:"coupon_refund_fee_0,omitempty"`
	CouponRefundCount   int    `xml:"coupon_refund_count,omitempty"`
	CouponRefundId0     string `xml:"coupon_refund_id_0,omitempty"`
}

type WeChatQueryRefundResponse struct {
	ReturnCode           string `xml:"return_code,omitempty"`
	ReturnMsg            string `xml:"return_msg,omitempty"`
	ResultCode           string `xml:"result_code,omitempty"`
	ErrCode              string `xml:"err_code,omitempty"`
	ErrCodeDes           string `xml:"err_code_des,omitempty"`
	Appid                string `xml:"appid,omitempty"`
	MchId                string `xml:"mch_id,omitempty"`
	NonceStr             string `xml:"nonce_str,omitempty"`
	Sign                 string `xml:"sign,omitempty"`
	TotalRefundCount     int    `xml:"total_refund_count,omitempty"`
	TransactionId        string `xml:"transaction_id,omitempty"`
	OutTradeNo           string `xml:"out_trade_no,omitempty"`
	TotalFee             int    `xml:"total_fee,omitempty"`
	SettlementTotalFee   int    `xml:"settlement_total_fee,omitempty"`
	FeeType              string `xml:"fee_type,omitempty"`
	CashFee              int    `xml:"cash_fee,omitempty"`
	RefundCount          int    `xml:"refund_count,omitempty"`
	OutRefundNo0         string `xml:"out_refund_no_0,omitempty"`
	RefundId0            string `xml:"refund_id_0,omitempty"`
	RefundChannel0       string `xml:"refund_channel_0,omitempty"`
	RefundFee0           int    `xml:"refund_fee_0,omitempty"`
	SettlementRefundFee0 int    `xml:"settlement_refund_fee_0,omitempty"`
	CouponType00         string `xml:"coupon_type_0_0,omitempty"`
	CouponRefundFee0     int    `xml:"coupon_refund_fee_0,omitempty"`
	CouponRefundCount0   int    `xml:"coupon_refund_count_0,omitempty"`
	CouponRefundId00     string `xml:"coupon_refund_id_0_0,omitempty"`
	CouponRefundFee00    int    `xml:"coupon_refund_fee_0_0,omitempty"`
	RefundStatus0        string `xml:"refund_status_0,omitempty"`
	RefundAccount0       string `xml:"refund_account_0,omitempty"`
	RefundRecvAccout0    string `xml:"refund_recv_accout_0,omitempty"`
	RefundSuccessTime0   string `xml:"refund_success_time_0,omitempty"`
}

type WeChatMicropayResponse struct {
	ReturnCode         string `xml:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty"`
	Openid             string `xml:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty"`
	BankType           string `xml:"bank_type,omitempty"`
	FeeType            string `xml:"fee_type,omitempty"`
	TotalFee           int    `xml:"total_fee,omitempty"`
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty"`
	CouponFee          int    `xml:"coupon_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty"`
	CashFee            int    `xml:"cash_fee,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty"`
	PromotionDetail    string `xml:"promotion_detail,omitempty"`
}

type getSignKeyResponse struct {
	ReturnCode     string `xml:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty"`
	Retmsg         string `xml:"retmsg,omitempty"`
	Retcode        string `xml:"retcode,omitempty"`
	MchId          string `xml:"mch_id,omitempty"`
	SandboxSignkey string `xml:"sandbox_signkey,omitempty"`
}

type WeChatNotifyRequest struct {
	ReturnCode         string `xml:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty"`
	ResultCode         string `xml:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty"`
	Appid              string `xml:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty"`
	SignType           string `xml:"sign_type,omitempty"`
	Openid             string `xml:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty"`
	BankType           string `xml:"bank_type,omitempty"`
	TotalFee           int    `xml:"total_fee,omitempty"`
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty"`
	CashFee            int    `xml:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty"`
	CouponFee          int    `xml:"coupon_fee,omitempty"`
	CouponCount        int    `xml:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty"`
	CouponFee0         int    `xml:"coupon_fee_0,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty"`
}

type Code2SessionRsp struct {
	SessionKey string `json:"session_key"` //会话密钥
	ExpiresIn  int    `json:"expires_in"`  //SessionKey超时时间（秒）
	Openid     string `json:"openid"`      //用户唯一标识
	Unionid    string `json:"unionid"`     //用户在开放平台的唯一标识符
	Errcode    int    `json:"errcode"`     //错误码
	Errmsg     string `json:"errmsg"`      //错误信息
}

type PaidUnionId struct {
	Unionid string `json:"unionid"` //用户在开放平台的唯一标识符
	Errcode int    `json:"errcode"` //错误码
	Errmsg  string `json:"errmsg"`  //错误信息
}

type AccessToken struct {
	AccessToken string `json:"access_token"` //获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   //SessionKey超时时间（秒）
	Errcode     int    `json:"errcode"`      //错误码
	Errmsg      string `json:"errmsg"`       //错误信息
}

type WeChatUserInfo struct {
	Subscribe      int    `json:"subscribe"`       //用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Openid         string `json:"openid"`          //用户唯一标识
	Nickname       string `json:"nickname"`        //用户的昵称
	Sex            int    `json:"sex"`             //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language       string `json:"language"`        //用户的语言，简体中文为zh_CN
	City           string `json:"city"`            //用户所在城市
	Province       string `json:"province"`        //用户所在省份
	Country        string `json:"country"`         //用户所在国家
	Headimgurl     string `json:"headimgurl"`      //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime  int    `json:"subscribe_time"`  //用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid        string `json:"unionid"`         //只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string `json:"remark"`          //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid        int    `json:"groupid"`         //用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int  `json:"tagid_list"`      //用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene"` //返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他
	QrScene        int    `json:"qr_scene"`        //二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str"`    //二维码扫码场景描述（开发者自定义）
	Errcode        int    `json:"errcode"`         //错误码
	Errmsg         string `json:"errmsg"`          //错误信息
}

type WeChatUserPhone struct {
	PhoneNumber     string        `json:"phoneNumber"`
	PurePhoneNumber string        `json:"purePhoneNumber"`
	CountryCode     string        `json:"countryCode"`
	Watermark       watermarkInfo `json:"watermark"`
}

type watermarkInfo struct {
	Appid     string `json:"appid"`
	Timestamp int    `json:"timestamp"`
}

type OpenIdByAuthCodeRsp struct {
	ReturnCode string `xml:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	Openid     string `xml:"openid,omitempty"` //用户唯一标识
}
