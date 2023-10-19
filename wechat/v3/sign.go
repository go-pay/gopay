package wechat

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"strings"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/pkg/xpem"
)

// Deprecated
// 推荐使用 wechat.V3VerifySignByPK()
func V3VerifySign(timestamp, nonce, signBody, sign, wxPubKeyContent string) (err error) {
	publicKey, err := xpem.DecodePublicKey([]byte(wxPubKeyContent))
	if err != nil {
		return err
	}
	str := timestamp + "\n" + nonce + "\n" + signBody + "\n"
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	h := sha256.New()
	h.Write([]byte(str))
	if err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}

// 推荐直接开启自动同步验签功能
// 微信V3 版本验签（同步）
// wxPublicKey：微信平台证书公钥内容，通过 client.WxPublicKeyMap() 获取，然后根据 signInfo.HeaderSerial 获取相应的公钥
func V3VerifySignByPK(timestamp, nonce, signBody, sign string, wxPublicKey *rsa.PublicKey) (err error) {
	if wxPublicKey == nil || wxPublicKey.N == nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, "wxPublicKey is nil")
	}
	str := timestamp + "\n" + nonce + "\n" + signBody + "\n"
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	h := sha256.New()
	h.Write([]byte(str))
	if err = rsa.VerifyPKCS1v15(wxPublicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}

// PaySignOfJSAPI 获取 JSAPI 支付所需要的参数
// 文档：https://pay.weixin.qq.com/docs/merchant/apis/jsapi-payment/jsapi-transfer-payment.html
func (c *ClientV3) PaySignOfJSAPI(appid, prepayid string) (jsapi *JSAPIPayParams, err error) {
	ts := util.Int642String(time.Now().Unix())
	nonceStr := util.RandomString(32)
	pkg := "prepay_id=" + prepayid

	_str := appid + "\n" + ts + "\n" + nonceStr + "\n" + pkg + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}

	jsapi = &JSAPIPayParams{
		AppId:     appid,
		TimeStamp: ts,
		NonceStr:  nonceStr,
		Package:   pkg,
		SignType:  SignTypeRSA,
		PaySign:   sign,
	}
	return jsapi, nil
}

// PaySignOfApp 获取 App 支付所需要的参数
// 文档：https://pay.weixin.qq.com/docs/merchant/apis/in-app-payment/app-transfer-payment.html
func (c *ClientV3) PaySignOfApp(appid, prepayid string) (app *AppPayParams, err error) {
	ts := util.Int642String(time.Now().Unix())
	nonceStr := util.RandomString(32)

	_str := appid + "\n" + ts + "\n" + nonceStr + "\n" + prepayid + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}

	app = &AppPayParams{
		Appid:     appid,
		Partnerid: c.Mchid,
		Prepayid:  prepayid,
		Package:   "Sign=WXPay",
		Noncestr:  nonceStr,
		Timestamp: ts,
		Sign:      sign,
	}
	return app, nil
}

// PaySignOfApplet 获取 小程序 支付所需要的参数
// 文档：https://pay.weixin.qq.com/docs/merchant/apis/mini-program-payment/mini-transfer-payment.html
func (c *ClientV3) PaySignOfApplet(appid, prepayid string) (applet *AppletParams, err error) {
	jsapi, err := c.PaySignOfJSAPI(appid, prepayid)
	if err != nil {
		return nil, err
	}
	applet = &AppletParams{
		AppId:     jsapi.AppId,
		TimeStamp: jsapi.TimeStamp,
		NonceStr:  jsapi.NonceStr,
		Package:   jsapi.Package,
		SignType:  jsapi.SignType,
		PaySign:   jsapi.PaySign,
	}
	return applet, nil
}

// PaySignOfAppScore 获取 APP调起支付分 接口，query属性中的sign
// 文档：https://pay.weixin.qq.com/docs/merchant/apis/weixin-pay-score/app-confirm.html
func (c *ClientV3) PaySignOfAppScore(mchId, pkg string) (query *APPScoreQuery, err error) {
	var (
		buffer   strings.Builder
		h        hash.Hash
		ts       = util.Int642String(time.Now().Unix())
		nonceStr = util.RandomString(32)
	)
	buffer.WriteString("mch_id=")
	buffer.WriteString(mchId)
	buffer.WriteString("&nonce_str=")
	buffer.WriteString(nonceStr)
	buffer.WriteString("&package=")
	buffer.WriteString(pkg)
	buffer.WriteString("&sign_type=HMAC-SHA256")
	buffer.WriteString("&timestamp=")
	buffer.WriteString(ts)
	buffer.WriteString("&key=")
	buffer.WriteString(string(c.ApiV3Key))

	h = hmac.New(sha256.New, c.ApiV3Key)
	h.Write([]byte(buffer.String()))

	query = &APPScoreQuery{
		MchId:     mchId,
		TimeStamp: ts,
		NonceStr:  nonceStr,
		Package:   pkg,
		SignType:  "HMAC-SHA256",
		Sign:      strings.ToUpper(hex.EncodeToString(h.Sum(nil))),
	}
	return query, nil
}

// PaySignOfJSAPIScore 获取 JSAPI调起支付分 接口，queryString属性中的sign
// 文档：https://pay.weixin.qq.com/docs/merchant/apis/weixin-pay-score/jsapi-confirm.html
func (c *ClientV3) PaySignOfJSAPIScore(mchId, pkg string) (queryString *JSAPIScoreQuery, err error) {
	var (
		buffer   strings.Builder
		h        hash.Hash
		ts       = util.Int642String(time.Now().Unix())
		nonceStr = util.RandomString(32)
	)
	buffer.WriteString("mch_id=")
	buffer.WriteString(mchId)
	buffer.WriteString("&nonce_str=")
	buffer.WriteString(nonceStr)
	buffer.WriteString("&package=")
	buffer.WriteString(pkg)
	buffer.WriteString("&sign_type=HMAC-SHA256")
	buffer.WriteString("&timestamp=")
	buffer.WriteString(ts)
	buffer.WriteString("&key=")
	buffer.WriteString(string(c.ApiV3Key))

	h = hmac.New(sha256.New, c.ApiV3Key)
	h.Write([]byte(buffer.String()))

	queryString = &JSAPIScoreQuery{
		MchId:     mchId,
		TimeStamp: ts,
		NonceStr:  nonceStr,
		Package:   pkg,
		SignType:  "HMAC-SHA256",
		Sign:      strings.ToUpper(hex.EncodeToString(h.Sum(nil))),
	}
	return queryString, nil
}

// PaySignOfAppletScore 获取 小程序调起支付分 接口，extraData属性中的sign
// 文档：https://pay.weixin.qq.com/docs/merchant/apis/weixin-pay-score/applets-confirm.html
func (c *ClientV3) PaySignOfAppletScore(mchId, pkg string) (extraData *AppletScoreExtraData, err error) {
	var (
		buffer   strings.Builder
		h        hash.Hash
		ts       = util.Int642String(time.Now().Unix())
		nonceStr = util.RandomString(32)
	)
	buffer.WriteString("mch_id=")
	buffer.WriteString(mchId)
	buffer.WriteString("&nonce_str=")
	buffer.WriteString(nonceStr)
	buffer.WriteString("&package=")
	buffer.WriteString(pkg)
	buffer.WriteString("&sign_type=HMAC-SHA256")
	buffer.WriteString("&timestamp=")
	buffer.WriteString(ts)
	buffer.WriteString("&key=")
	buffer.WriteString(string(c.ApiV3Key))

	h = hmac.New(sha256.New, c.ApiV3Key)
	h.Write([]byte(buffer.String()))

	extraData = &AppletScoreExtraData{
		MchId:     mchId,
		TimeStamp: ts,
		NonceStr:  nonceStr,
		Package:   pkg,
		SignType:  "HMAC-SHA256",
		Sign:      strings.ToUpper(hex.EncodeToString(h.Sum(nil))),
	}
	return extraData, nil
}

// v3 鉴权请求Header
func (c *ClientV3) authorization(method, path string, bm gopay.BodyMap) (string, error) {
	var (
		jb        = ""
		timestamp = time.Now().Unix()
		nonceStr  = util.RandomString(32)
	)
	if bm != nil {
		jb = bm.JsonBody()
	}
	path = strings.TrimSuffix(path, "?")
	ts := util.Int642String(timestamp)
	_str := method + "\n" + path + "\n" + ts + "\n" + nonceStr + "\n" + jb + "\n"
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_SignString:\n%s", _str)
	}
	sign, err := c.rsaSign(_str)
	if err != nil {
		return "", err
	}
	return Authorization + ` mchid="` + c.Mchid + `",nonce_str="` + nonceStr + `",timestamp="` + ts + `",serial_no="` + c.SerialNo + `",signature="` + sign + `"`, nil
}

func (c *ClientV3) rsaSign(str string) (string, error) {
	if c.privateKey == nil {
		return "", errors.New("privateKey can't be nil")
	}
	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return util.NULL, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

// 自动同步请求验签
func (c *ClientV3) verifySyncSign(si *SignInfo) (err error) {
	if !c.autoSign {
		return nil
	}
	if si == nil {
		return errors.New("auto verify sign, but SignInfo is nil")
	}
	c.rwMu.RLock()
	wxPublicKey, exist := c.SnCertMap[si.HeaderSerial]
	c.rwMu.RUnlock()
	if !exist {
		err = c.AutoVerifySign(false)
		if err != nil {
			return fmt.Errorf("[get all public key err]: %v", err)
		}
		c.rwMu.RLock()
		wxPublicKey, exist = c.SnCertMap[si.HeaderSerial]
		c.rwMu.RUnlock()
		if !exist {
			return errors.New("auto verify sign, but public key not found")
		}
	}
	str := si.HeaderTimestamp + "\n" + si.HeaderNonce + "\n" + si.SignBody + "\n"
	signBytes, _ := base64.StdEncoding.DecodeString(si.HeaderSignature)
	h := sha256.New()
	h.Write([]byte(str))
	if err = rsa.VerifyPKCS1v15(wxPublicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}
