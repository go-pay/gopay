package gopay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

type requestBody map[string]string

//获取参数
func (w requestBody) Get(key string) string {
	if w == nil {
		return ""
	}
	ws := w[key]
	return ws
}

//设置参数
func (w requestBody) Set(key string, value string) {
	w[key] = value
}

//删除参数
func (w requestBody) Remove(key string) {
	delete(w, key)
}

// 必选.
//    Appid: 微信支付分配的公众账号ID（企业号corpid即为此appId）.
//    MchId: 微信支付分配的商户号.
//    NonceStr: 随机字符串，长度要求在32位以内（如不写，go-pay将为你随机生成）.
//    Body: 商品简单描述.
//    OutTradeNo: 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一.
//    TotalFee: 订单总金额，单位为分.
//    SpbillCreateIp: 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP.
//    NotifyUrl: 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数.
//    TradeType: 交易类型：（JSAPI--JSAPI支付（或小程序支付）、NATIVE--Native支付、APP--app支付，MWEB--H5支付，不同trade_type决定了调起支付的方式）.
//
// 非必选.
//    DeviceInfo: 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB".
//    SignType: 签名类型，默认为MD5，支持HMAC-SHA256和MD5.
//    Detail: 商品详细描述，对于使用单品优惠的商户，字段必须按照规范上传.
//    Attach: 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用.
//    FeeType: 符合ISO 4217标准的三位字母代码，默认人民币：CNY.
//    TimeStart: 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010.
//    TimeExpire: 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。订单失效时间是针对订单号而言的，由于在请求支付的时候有一个必传参数prepay_id只有两小时的有效期，所以在重入时间超过2小时的时候需要重新请求下单接口获取新的prepay_id.
//    GoodsTag: 订单优惠标记，使用代金券或立减优惠功能时需要的参数.
//    ProductId: trade_type=NATIVE时，此参数必传。此参数为二维码中包含的商品ID，商户自行定义.
//    LimitPay: 上传此参数 no_credit 可限制用户不能使用信用卡支付.
//    Openid: 用户标识: trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识.
//    Receipt: Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效.
//    SceneInfo: 该字段常用于线下活动时的场景信息上报，支持上报实际门店信息，商户也可以按需求自己上报相关信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }}.
type WeChatPayParams struct {
	NonceStr       string `xml:"nonce_str"`
	Body           string `xml:"body"`
	OutTradeNo     string `xml:"out_trade_no"`
	TotalFee       int    `xml:"total_fee"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`

	DeviceInfo string     `xml:"device_info"`
	SignType   string     `xml:"sign_type"`
	Detail     string     `xml:"detail"`
	Attach     string     `xml:"attach"`
	FeeType    string     `xml:"fee_type"`
	TimeStart  string     `xml:"time_start"`
	TimeExpire string     `xml:"time_expire"`
	GoodsTag   string     `xml:"goods_tag"`
	ProductId  string     `xml:"product_id"`
	LimitPay   string     `xml:"limit_pay"`
	Openid     string     `xml:"openid"`
	Receipt    string     `xml:"receipt"`
	SceneInfo  *StoreInfo `xml:"scene_info"`
}

//StoreInfo: SceneInfo 的字段信息
type StoreInfo struct {
	Id       string `json:"id"`        // 门店唯一标识
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 门店所在地行政区划码，详细见《最新县及县以上行政区划代码》
	Address  string `json:"address"`   // 门店详细地址
}

//获取请求支付的参数
func (w *WeChatPayParams) getRequestBody(appId, mchId string, params *WeChatPayParams) requestBody {
	reqs := make(requestBody)
	reqs.Set("appid", appId)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", params.NonceStr)
	reqs.Set("body", params.Body)
	reqs.Set("out_trade_no", params.OutTradeNo)
	reqs.Set("total_fee", strconv.Itoa(params.TotalFee))
	reqs.Set("spbill_create_ip", params.SpbillCreateIp)
	reqs.Set("notify_url", params.NotifyUrl)
	reqs.Set("trade_type", params.TradeType)

	if params.DeviceInfo != "" {
		reqs.Set("device_info", params.DeviceInfo)
	}
	if params.SignType != "" {
		reqs.Set("sign_type", params.SignType)
	} else {
		reqs.Set("sign_type", "MD5")
	}
	if params.Detail != "" {
		reqs.Set("detail", params.Detail)
	}
	if params.Attach != "" {
		reqs.Set("attach", params.Attach)
	}
	if params.FeeType != "" {
		reqs.Set("fee_type", params.FeeType)
	}
	if params.TimeStart != "" {
		reqs.Set("time_start", params.TimeStart)
	}
	if params.TimeExpire != "" {
		reqs.Set("time_expire", params.TimeExpire)
	}
	if params.GoodsTag != "" {
		reqs.Set("goods_tag", params.GoodsTag)
	}
	if params.ProductId != "" {
		reqs.Set("product_id", params.ProductId)
	}
	if params.LimitPay != "" {
		reqs.Set("limit_pay", params.LimitPay)
	}
	if params.Openid != "" {
		reqs.Set("openid", params.Openid)
	}
	if params.Receipt != "" {
		reqs.Set("receipt", params.Receipt)
	}
	if params.SceneInfo != nil {
		marshal, _ := json.Marshal(params.SceneInfo)
		reqs.Set("scene_info", string(marshal))
	}
	return reqs
}

//获取Sign签名和请求支付的参数
func getSignAndRequestBody(appId, mchId string, secretKey string, param *WeChatPayParams) (sign string, reqs requestBody) {

	reqs = param.getRequestBody(appId, mchId, param)

	signStr := getSignString(secretKey, reqs)
	//fmt.Println("signStr:", signStr)
	var hashSign []byte
	if param.SignType == WX_SignType_MD5 {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := hmac.New(sha256.New, []byte(secretKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	sign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//获取根据Key排序后的请求参数字符串
func getSignString(secretKey string, wxReq requestBody) string {
	keyList := make([]string, 0)
	for k := range wxReq {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(wxReq[k])
		buffer.WriteString("&")
	}
	buffer.WriteString("key=")
	buffer.WriteString(secretKey)
	return buffer.String()
}

//获取SanboxKey
func getSanBoxSignKey() {
	//url:= "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"

}
