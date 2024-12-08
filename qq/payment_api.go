/*
	QQ 支付
	文档：https://qpay.qq.com/buss/doc.shtml
*/

package qq

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-pay/gopay"
)

// ParseNotifyToBodyMap 解析QQ支付异步通知的结果到BodyMap
// req：*http.Request
// 返回参数bm：Notify请求的参数
// 返回参数err：错误信息
func ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	bs, err := io.ReadAll(io.LimitReader(req.Body, int64(3<<20))) // default 3MB change the size you want;
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}
	bm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &bm); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return
}

// Deprecated
// 推荐使用 ParseNotifyToBodyMap
func ParseNotify(req *http.Request) (notifyReq *NotifyRequest, err error) {
	notifyReq = new(NotifyRequest)
	if err = xml.NewDecoder(req.Body).Decode(notifyReq); err != nil {
		return nil, fmt.Errorf("xml.NewDecoder.Decode: %w", err)
	}
	return
}

// VerifySign QQ同步返回参数验签或异步通知参数验签
//
//	ApiKey：API秘钥值
//	signType：签名类型（调用API方法时填写的类型）
//	bean：微信同步返回的结构体 qqRsp 或 异步通知解析的结构体 notifyReq
//	返回参数ok：是否验签通过
//	返回参数err：其他错误信息，不要根据 error 是否为空来判断验签正确与否，需再单独判断返回的 ok
func VerifySign(apiKey, signType string, bean any) (ok bool, err error) {
	if apiKey == gopay.NULL || signType == gopay.NULL {
		return false, errors.New("apiKey or signType can not null")
	}
	if bean == nil {
		return false, errors.New("bean is nil")
	}
	kind := reflect.ValueOf(bean).Kind()
	if kind == reflect.Map {
		bm := bean.(gopay.BodyMap)
		bodySign := bm.GetString("sign")
		bm.Remove("sign")
		return GetReleaseSign(apiKey, signType, bm) == bodySign, nil
	}

	bs, err := json.Marshal(bean)
	if err != nil {
		return false, fmt.Errorf("[%w]: %v, value: %v", gopay.MarshalErr, err, bean)
	}
	bm := make(gopay.BodyMap)
	if err = json.Unmarshal(bs, &bm); err != nil {
		return false, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	bodySign := bm.GetString("sign")
	bm.Remove("sign")
	return GetReleaseSign(apiKey, signType, bm) == bodySign, nil
}

type NotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// ToXmlString 返回数据给QQ
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
