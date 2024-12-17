package saobei

type RspBase struct {
	ReturnCode string `json:"return_code"` //响应码：01成功 ，02失败，响应码仅代表通信状态，不代表业务结果
	ReturnMsg  string `json:"return_msg"`  //返回信息提示，“退款成功”、“订单不存在”等
	KeySign    string `json:"key_sign"`    //签名串《2.4签名算法》 签名测试页
	ResultCode string `json:"result_code"` //业务结果：01成功 ，02失败
}

// MiniPayRsp 小程序支付响应
type MiniPayRsp struct {
	RspBase
	PayType       string `json:"pay_type"`       //支付方式，010微信，020支付宝
	MerchantName  string `json:"merchant_name"`  //商户名称
	MerchantNo    string `json:"merchant_no"`    //商户号
	TerminalId    string `json:"terminal_id"`    //终端号
	DeviceNo      string `json:"device_no"`      //商户终端设备号(商户自定义，如门店编号),必须在平台已配置过
	TerminalTrace string `json:"terminal_trace"` //终端流水号，商户系统的订单号，系统原样返回
	TerminalTime  string `json:"terminal_time"`  //终端交易时间，yyyyMMddHHmmss，全局统一时间格式，系统原样返回

	TotalFee   string `json:"total_fee"`    //金额，单位分
	OutTradeNo string `json:"out_trade_no"` //平台唯一订单号
	AppId      string `json:"appId"`        //微信小程序支付返回字段，公众号id
	TimeStamp  string `json:"timeStamp"`    //微信小程序支付返回字段，时间戳，示例：1414561699，标准北京时间，时区为东八区，自1970年1月1日 0点0分0秒以来的秒数。注意：部分系统取到的值为毫秒级，需要转换成秒(10位数字)。
	NonceStr   string `json:"nonceStr"`     //微信小程序支付返回字段，随机字符串
	PackageStr string `json:"package_str"`  //微信小程序支付返回字段，订单详情扩展字符串，示例：prepay_id=123456789，统一下单接口返回的prepay_id参数值，提交格式如：prepay_id=
	SignType   string `json:"signType"`     //微信小程序支付返回字段，签名方式，示例：MD5,RSA
	PaySign    string `json:"paySign"`      //微信小程序支付返回字段，签名
	AliTradeNo string `json:"ali_trade_no"` //支付宝小程序支付返回字段用于调起支付宝小程序
}

// BarcodePayRsp 付款码支付(扫码支付)响应
type BarcodePayRsp struct {
	RspBase
	PayType       string `json:"pay_type"`       //支付方式，010微信，020支付宝
	MerchantName  string `json:"merchant_name"`  //商户名称
	MerchantNo    string `json:"merchant_no"`    //商户号
	TerminalId    string `json:"terminal_id"`    //终端号
	DeviceNo      string `json:"device_no"`      //商户终端设备号(商户自定义，如门店编号),必须在平台已配置过
	TerminalTrace string `json:"terminal_trace"` //终端流水号，商户系统的订单号，系统原样返回
	TerminalTime  string `json:"terminal_time"`  //终端交易时间，yyyyMMddHHmmss，全局统一时间格式，系统原样返回

	TotalFee            string `json:"total_fee"`             //金额，单位分
	BuyerPayFee         string `json:"buyer_pay_fee"`         // 买家实付金额（分）pay_ver为202时返回
	PlatformDiscountFee string `json:"platform_discount_fee"` // 平台优惠金额（分）pay_ver为202时返回
	MerchantDiscountFee string `json:"merchant_discount_fee"` // 商家优惠金额（分）pay_ver为202时返回
	EndTime             string `json:"end_time"`              // 支付完成时间，yyyyMMddHHmmss，全局统一时间格式
	OutTradeNo          string `json:"out_trade_no"`          //平台唯一订单号
	ChannelTradeNo      string `json:"channel_trade_no"`      //通道订单号，微信订单号、支付宝订单号等
	ChannelOrderNo      string `json:"channel_order_no"`      //银行渠道订单号，微信支付时显示在支付成功页面的条码，可用作扫码查询和扫码退款时匹配
	UserId              string `json:"user_id"`               //付款方用户id，“微信openid”、“支付宝账户”
	Attach              string `json:"attach"`                //附加数据,原样返回
	ReceiptFee          string `json:"receipt_fee"`           //商家应结算金额,单位分
	BankType            string `json:"bank_type"`             //银行类型，采用字符串类型的银行标识
	PromotionDetail     string `json:"promotion_detail"`      //官方营销详情,pay_ver=202时返回. 本交易支付时使用的所有优惠券信息 ，单品优惠功能字段，详情见
	OrderBody           string `json:"order_body"`            //订单标题描述
	SubOpenid           string `json:"sub_openid"`            //微信子商户sub_appid对应的用户标识
}

// QueryRsp 支付查询
type QueryRsp struct {
	RspBase
	PayType       string `json:"pay_type"`       //支付方式，010微信，020支付宝
	MerchantName  string `json:"merchant_name"`  //商户名称
	MerchantNo    string `json:"merchant_no"`    //商户号
	TerminalId    string `json:"terminal_id"`    //终端号
	DeviceNo      string `json:"device_no"`      //商户终端设备号(商户自定义，如门店编号),必须在平台已配置过
	TerminalTrace string `json:"terminal_trace"` //终端流水号，商户系统的订单号，系统原样返回
	TerminalTime  string `json:"terminal_time"`  //终端交易时间，yyyyMMddHHmmss，全局统一时间格式，系统原样返回

	TotalFee            string `json:"total_fee"`             //金额，单位分
	BuyerPayFee         string `json:"buyer_pay_fee"`         //买家实付金额（分）pay_ver为202时返回
	PlatformDiscountFee string `json:"platform_discount_fee"` //平台优惠金额（分）pay_ver为202时返回
	MerchantDiscountFee string `json:"merchant_discount_fee"` //商家优惠金额（分）pay_ver为202时返回
	SubOpenid           string `json:"sub_openid"`            //微信子商户sub_appid对应的用户标识
	OrderBody           string `json:"order_body"`            //订单标题描述
	EndTime             string `json:"end_time"`              //支付完成时间，yyyyMMddHHmmss，全局统一时间格式
	OutTradeNo          string `json:"out_trade_no"`          //平台唯一订单号
	TradeState          string `json:"trade_state"`           //交易订单状态，SUCCESS支付成功，REFUND转入退款，NOTPAY未支付，CLOSED已关闭，USERPAYING用户支付中，REVOKED已撤销，NOPAY未支付支付超时，PAYERROR支付失败
	ChannelTradeNo      string `json:"channel_trade_no"`      //通道订单号，微信订单号、支付宝订单号等
	ChannelOrderNo      string `json:"channel_order_no"`      //银行渠道订单号，微信支付时显示在支付成功页面的条码，可用作扫码查询和扫码退款时匹配
	UserId              string `json:"user_id"`               //付款方用户id，“微信openid”、“支付宝账户”
	Attach              string `json:"attach"`                //附加数据,原样返回
	ReceiptFee          string `json:"receipt_fee"`           //商家应结算金额,单位分
	PayTrace            string `json:"pay_trace"`             //当前支付终端流水号
	PayTime             string `json:"pay_time"`              //当前支付终端交易时间，yyyyMMddHHmmss，全局统一时间格式
	BankType            string `json:"bank_type"`             //银行类型，采用字符串类型的银行标识
	PromotionDetail     string `json:"promotion_detail"`      //官方营销详情,pay_ver=202时返回. 本交易支付时使用的所有优惠券信息 ，单品优惠功能字段，详情见
}

// RefundRsp 退款申请
type RefundRsp struct {
	RspBase
	PayType       string `json:"pay_type"`       //支付方式，010微信，020支付宝
	MerchantName  string `json:"merchant_name"`  //商户名称
	MerchantNo    string `json:"merchant_no"`    //商户号
	TerminalId    string `json:"terminal_id"`    //终端号
	DeviceNo      string `json:"device_no"`      //商户终端设备号(商户自定义，如门店编号),必须在平台已配置过
	TerminalTrace string `json:"terminal_trace"` //终端流水号，商户系统的订单号，系统原样返回
	TerminalTime  string `json:"terminal_time"`  //终端交易时间，yyyyMMddHHmmss，全局统一时间格式，系统原样返回

	RefundFee                 string `json:"refund_fee"`                   //退款金额，单位分
	RefundReceiptFee          string `json:"refund_receipt_fee"`           //退商家应结算金额,单位分
	RefundBuyerPayFee         string `json:"refund_buyer_pay_fee"`         //退买家实付金额（分）
	RefundPlatformDiscountFee string `json:"refund_platform_discount_fee"` //退平台优惠金额（分）
	RefundMerchantDiscountFee string `json:"refund_merchant_discount_fee"` //退商家优惠金额（分）
	RefundPromotionDetail     string `json:"refund_promotion_detail"`      //退优惠明细，详情见
	EndTime                   string `json:"end_time"`                     //退款完成时间，yyyyMMddHHmmss，全局统一时间格式
	OutTradeNo                string `json:"out_trade_no"`                 //平台唯一订单号
	OutRefundNo               string `json:"out_refund_no"`                //平台唯一退款单号
}

// QueryRefundRsp 退款订单查询
type QueryRefundRsp struct {
	RspBase
	PayType       string `json:"pay_type"`       //支付方式，010微信，020支付宝
	MerchantName  string `json:"merchant_name"`  //商户名称
	MerchantNo    string `json:"merchant_no"`    //商户号
	TerminalId    string `json:"terminal_id"`    //终端号
	DeviceNo      string `json:"device_no"`      //商户终端设备号(商户自定义，如门店编号),必须在平台已配置过
	TerminalTrace string `json:"terminal_trace"` //终端流水号，商户系统的订单号，系统原样返回
	TerminalTime  string `json:"terminal_time"`  //终端交易时间，yyyyMMddHHmmss，全局统一时间格式，系统原样返回

	RefundFee                 string `json:"refund_fee"`                   //退款金额，单位分
	RefundReceiptFee          string `json:"refund_receipt_fee"`           //退商家应结算金额,单位分
	RefundBuyerPayFee         string `json:"refund_buyer_pay_fee"`         //退买家实付金额（分）
	RefundPlatformDiscountFee string `json:"refund_platform_discount_fee"` //退平台优惠金额（分）
	RefundMerchantDiscountFee string `json:"refund_merchant_discount_fee"` //退商家优惠金额（分）
	RefundPromotionDetail     string `json:"refund_promotion_detail"`      //退优惠明细，详情见
	EndTime                   string `json:"end_time"`                     //退款完成时间，yyyyMMddHHmmss，全局统一时间格式
	OutRefundNo               string `json:"out_refund_no"`                //平台唯一退款单号
	OutTradeNo                string `json:"out_trade_no"`                 //平台唯一订单号
	TradeState                string `json:"trade_state"`                  //交易订单状态，SUCCESS支付成功，REFUND转入退款，NOTPAY未支付，CLOSED已关闭，USERPAYING用户支付中，REVOKED已撤销，NOPAY未支付支付超时，PAYERROR支付失败
	ChannelTradeNo            string `json:"channel_trade_no"`             //通道订单号，微信订单号、支付宝订单号等
	ChannelOrderNo            string `json:"channel_order_no"`             //银行渠道订单号，微信支付时显示在支付成功页面的条码，可用作扫码查询和扫码退款时匹配
	UserId                    string `json:"user_id"`                      //退款方用户id，“微信openid”、“支付宝账户”、“qq号”等
	Attach                    string `json:"attach"`                       //附加数据,原样返回
	PayTrace                  string `json:"pay_trace"`                    //退款终端流水号
	PayTime                   string `json:"pay_time"`                     //退款终端交易时间
}
