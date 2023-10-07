package allinpay

const (
	// URL
	baseUrl        = ""
	sandboxBaseUrl = "http://test2.lcsw.cn:8117/lcsw"
)

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
	TotalFee      string `json:"total_fee"`      //金额，单位分
	OutTradeNo    string `json:"out_trade_no"`   //平台唯一订单号
	AppId         string `json:"appId"`          //微信小程序支付返回字段，公众号id
	TimeStamp     string `json:"timeStamp"`      //微信小程序支付返回字段，时间戳，示例：1414561699，标准北京时间，时区为东八区，自1970年1月1日 0点0分0秒以来的秒数。注意：部分系统取到的值为毫秒级，需要转换成秒(10位数字)。
	NonceStr      string `json:"nonceStr"`       //微信小程序支付返回字段，随机字符串
	PackageStr    string `json:"package_str"`    //微信小程序支付返回字段，订单详情扩展字符串，示例：prepay_id=123456789，统一下单接口返回的prepay_id参数值，提交格式如：prepay_id=
	SignType      string `json:"signType"`       //微信小程序支付返回字段，签名方式，示例：MD5,RSA
	PaySign       string `json:"paySign"`        //微信小程序支付返回字段，签名
	AliTradeNo    string `json:"ali_trade_no"`   //支付宝小程序支付返回字段用于调起支付宝小程序
}
