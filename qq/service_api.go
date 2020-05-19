package qq

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/iGoogle-ink/gopay"
)

// 向QQ发送Post请求，对于本库未提供的微信API，可自行实现，通过此方法发送请求
//    bm：请求参数的BodyMap
//    url：完整url地址，例如：https://qpay.qq.com/cgi-bin/pay/qpay_unified_order.cgi
//    tlsConfig：tls配置，如无需证书请求，传nil
func (q *Client) PostRequest(bm gopay.BodyMap, url string, tlsConfig *tls.Config) (bs []byte, err error) {
	return q.doQQ(bm, url, tlsConfig)
}

// 解析QQ支付异步通知的结果到BodyMap
//    req：*http.Request
//    返回参数bm：Notify请求的参数
//    返回参数err：错误信息
func ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	bs, err := ioutil.ReadAll(io.LimitReader(req.Body, int64(3<<20))) // default 3MB change the size you want;
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll：%w", err)
	}
	bm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &bm); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return
}

// 解析QQ支付异步通知的参数
//    req：*http.Request
//    返回参数notifyReq：Notify请求的参数
//    返回参数err：错误信息
func ParseNotify(req *http.Request) (notifyReq *NotifyRequest, err error) {
	notifyReq = new(NotifyRequest)
	if err = xml.NewDecoder(req.Body).Decode(notifyReq); err != nil {
		return nil, fmt.Errorf("xml.NewDecoder.Decode：%w", err)
	}
	return
}

// QQ同步返回参数验签或异步通知参数验签
//    ApiKey：API秘钥值
//    signType：签名类型（调用API方法时填写的类型）
//    bean：微信同步返回的结构体 qqRsp 或 异步通知解析的结构体 notifyReq
//    返回参数ok：是否验签通过
//    返回参数err：错误信息
func VerifySign(apiKey, signType string, bean interface{}) (ok bool, err error) {
	if apiKey == gopay.NULL || signType == gopay.NULL {
		return false, errors.New("apiKey or signType can not null")
	}
	if bean == nil {
		return false, errors.New("bean is nil")
	}
	kind := reflect.ValueOf(bean).Kind()
	if kind == reflect.Map {
		bm := bean.(gopay.BodyMap)
		bodySign := bm.Get("sign")
		bm.Remove("sign")
		return getReleaseSign(apiKey, signType, bm) == bodySign, nil
	}

	bs, err := json.Marshal(bean)
	if err != nil {
		return false, fmt.Errorf("json.Marshal(%s)：%w", string(bs), err)
	}
	bm := make(gopay.BodyMap)
	if err = json.Unmarshal(bs, &bm); err != nil {
		return false, fmt.Errorf("json.Marshal(%s)：%w", string(bs), err)
	}
	bodySign := bm.Get("sign")
	bm.Remove("sign")
	return getReleaseSign(apiKey, signType, bm) == bodySign, nil
}

type NotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// 返回数据给QQ
func (w *NotifyResponse) ToXmlString() (xmlStr string) {
	var buffer strings.Builder
	buffer.WriteString("<xml><return_code>")
	buffer.WriteString(w.ReturnCode)
	buffer.WriteString("</return_code>")
	if w.ReturnMsg != gopay.NULL {
		buffer.WriteString("<return_msg>")
		buffer.WriteString(w.ReturnMsg)
		buffer.WriteString("</return_msg>")
	}
	buffer.WriteString("</xml>")

	xmlStr = buffer.String()
	return
}
