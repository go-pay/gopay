/*
 微信支付
 文档：https://pay.weixin.qq.com/wiki/doc/api/index.html
*/

package wechat

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	xaes "github.com/go-pay/crypto/aes"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
)

// ParseNotifyToBodyMap 解析微信支付异步通知的结果到BodyMap（推荐）
// req：*http.Request
// 返回参数bm：Notify请求的参数
// 返回参数err：错误信息
func ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	bs, err := io.ReadAll(io.LimitReader(req.Body, int64(3<<20))) // default 3MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
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
	err = xml.NewDecoder(req.Body).Decode(notifyReq)
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("xml.NewDecoder.Decode: %w", err)
	}
	return
}

// ParseRefundNotify 解析微信退款异步通知的参数
// req：*http.Request
// 返回参数notifyReq：Notify请求的参数
// 返回参数err：错误信息
func ParseRefundNotify(req *http.Request) (notifyReq *RefundNotifyRequest, err error) {
	notifyReq = new(RefundNotifyRequest)
	err = xml.NewDecoder(req.Body).Decode(notifyReq)
	defer req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("xml.NewDecoder.Decode: %w", err)
	}
	return
}

// DecryptRefundNotifyReqInfo 解密微信退款异步通知的加密数据
// reqInfo：gopay.ParseRefundNotify() 方法获取的加密数据 req_info
// apiKey：API秘钥值
// 返回参数refundNotify：RefundNotify请求的加密数据
// 返回参数err：错误信息
// 文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10
func DecryptRefundNotifyReqInfo(reqInfo, apiKey string) (refundNotify *RefundNotify, err error) {
	if reqInfo == gopay.NULL || apiKey == gopay.NULL {
		return nil, errors.New("reqInfo or apiKey is null")
	}
	var (
		encryptionB, bs []byte
		block           cipher.Block
		blockSize       int
	)
	if encryptionB, err = base64.StdEncoding.DecodeString(reqInfo); err != nil {
		return nil, err
	}
	h := md5.New()
	h.Write([]byte(apiKey))
	key := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	if len(encryptionB)%aes.BlockSize != 0 {
		return nil, errors.New("encryptedData is error")
	}
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		return nil, err
	}
	blockSize = block.BlockSize()

	err = func(dst, src []byte) error {
		if len(src)%blockSize != 0 {
			return errors.New("crypto/cipher: input not full blocks")
		}
		if len(dst) < len(src) {
			return errors.New("crypto/cipher: output smaller than input")
		}
		for len(src) > 0 {
			block.Decrypt(dst, src[:blockSize])
			src = src[blockSize:]
			dst = dst[blockSize:]
		}
		return nil
	}(encryptionB, encryptionB)
	if err != nil {
		return nil, err
	}

	bs = xaes.PKCS7UnPadding(encryptionB)
	refundNotify = new(RefundNotify)
	if err = xml.Unmarshal(bs, refundNotify); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return
}

type NotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// ToXmlString 返回数据给微信
func (w *NotifyResponse) ToXmlString() (xmlStr string) {
	var buffer strings.Builder
	buffer.WriteString("<xml><return_code><![CDATA[")
	buffer.WriteString(w.ReturnCode)
	buffer.WriteString("]]></return_code>")
	buffer.WriteString("<return_msg><![CDATA[")
	buffer.WriteString(w.ReturnMsg)
	buffer.WriteString("]]></return_msg></xml>")
	xmlStr = buffer.String()
	return
}

// GetOpenIdByAuthCode 授权码查询openid(AccessToken:157字符)
// appId:APPID
// mchId:商户号
// ApiKey:apiKey
// authCode:用户授权码
// nonceStr:随即字符串
// 文档：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9
func GetOpenIdByAuthCode(ctx context.Context, appId, mchId, apiKey, authCode, nonceStr string) (openIdRsp *OpenIdByAuthCodeRsp, err error) {
	var (
		url string
		bm  gopay.BodyMap
	)
	url = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
	bm = make(gopay.BodyMap)
	bm.Set("appid", appId)
	bm.Set("mch_id", mchId)
	bm.Set("auth_code", authCode)
	bm.Set("nonce_str", nonceStr)
	bm.Set("sign", GetReleaseSign(apiKey, SignType_MD5, bm))

	openIdRsp = new(OpenIdByAuthCodeRsp)
	_, err = xhttp.NewClient().Req(xhttp.TypeXML).Post(url).SendString(GenerateXml(bm)).EndStruct(ctx, openIdRsp)
	if err != nil {
		return nil, err
	}
	return openIdRsp, nil
}
