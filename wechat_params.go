package go_pay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
	NecessaryParams: 必传参数
	DeviceInfo: 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	Openid: 用户标识: trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识
*/
type WechatParams struct {
	NecessaryParams WechatParamsNecessary
	DeviceInfo      string `xml:"device_info"`
	Openid          string `xml:"openid"`
}

/*
	NonceStr: 随机字符串，长度要求在32位以内（如不写，go-pay将为你随机生成）
	Body: 商品简单描述
	OutTradeNo: 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一
	TotalFee: 订单总金额，单位为分
	SpbillCreateIp: 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP
	NotifyUrl: 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数
	TradeType: 交易类型：（JSAPI--JSAPI支付（或小程序支付）、NATIVE--Native支付、APP--app支付，MWEB--H5支付，不同trade_type决定了调起支付的方式）
*/
type WechatParamsNecessary struct {
	Appid          string `xml:"appid"`
	MchId          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	Body           string `xml:"body"`
	OutTradeNo     string `xml:"out_trade_no"`
	TotalFee       int    `xml:"total_fee"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`
}

//获取Sign签名
func getSign(secretKey string, params *WechatParams) string {

	paramMap := make(map[string]string, 0)
	paramMap["appid"] = params.NecessaryParams.Appid
	paramMap["mch_id"] = params.NecessaryParams.MchId
	paramMap["nonce_str"] = params.NecessaryParams.NonceStr
	paramMap["body"] = params.NecessaryParams.Body
	paramMap["out_trade_no"] = params.NecessaryParams.OutTradeNo
	paramMap["total_fee"] = strconv.Itoa(params.NecessaryParams.TotalFee)
	paramMap["spbill_create_ip"] = params.NecessaryParams.SpbillCreateIp
	paramMap["notify_url"] = params.NecessaryParams.NotifyUrl
	paramMap["trade_type"] = params.NecessaryParams.TradeType

	if params.Openid != "" {
		paramMap["openid"] = params.Openid
	}
	if params.DeviceInfo != "" {
		paramMap["openid"] = params.DeviceInfo
	}
	signStr := getSignString(secretKey, paramMap)
	fmt.Println("signStr:", signStr)
	hash := md5.New()
	hash.Write([]byte(signStr))
	md5Sign := hash.Sum(nil)
	sign := strings.ToUpper(hex.EncodeToString(md5Sign))
	return sign
}

//获取排好序的Key
func getSignString(secretKey string, paramMap map[string]string) string {
	keyList := make([]string, 0)
	for k := range paramMap {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	stringA := ""
	for _, k := range keyList {
		stringA += k + "=" + paramMap[k] + "&"
	}
	stringA += "key=" + secretKey
	return stringA
}

func generateXml(m map[string]string) string {
	xml := "<xml>"
	for k, v := range m {
		xml += fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xml += "</xml>"
	return xml
}
