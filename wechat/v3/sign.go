package wechat

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// V3VerifySign 微信V3 版本验签
//	wxPkContent：微信平台证书公钥内容，通过client.GetPlatformCerts() 获取
func V3VerifySign(timestamp, nonce, signBody, sign, wxPkContent string) (err error) {
	var (
		block     *pem.Block
		pubKey    *x509.Certificate
		publicKey *rsa.PublicKey
		ok        bool
	)
	str := timestamp + "\n" + nonce + "\n" + signBody + "\n"
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	if block, _ = pem.Decode([]byte(wxPkContent)); block == nil {
		return errors.New("parse wechat platform public key error")
	}
	if pubKey, err = x509.ParseCertificate(block.Bytes); err != nil {
		return fmt.Errorf("x509.ParseCertificate：%+v", err)
	}
	if publicKey, ok = pubKey.PublicKey.(*rsa.PublicKey); !ok {
		return errors.New("convert wechat platform public to rsa.PublicKey error")
	}
	h := sha256.New()
	h.Write([]byte(str))
	if err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("verify sign failed: %+v", err)
	}
	return nil
}

// PaySignOfJSAPI 获取 JSAPI paySign
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml
func (c *ClientV3) PaySignOfJSAPI(prepayid string) (jsapi *JSAPIPayParams, err error) {
	ts := util.Int642String(time.Now().Unix())
	nonceStr := util.GetRandomString(32)
	prepayId := "prepay_id=" + prepayid

	_str := c.Appid + "\n" + ts + "\n" + nonceStr + "\n" + prepayId + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}

	jsapi = &JSAPIPayParams{
		AppId:     c.Appid,
		TimeStamp: ts,
		NonceStr:  nonceStr,
		Package:   prepayId,
		SignType:  SignTypeRSA,
		PaySign:   sign,
	}
	return jsapi, nil
}

// PaySignOfApp 获取 App paySign
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_4.shtml
func (c *ClientV3) PaySignOfApp(prepayid string) (app *AppPayParams, err error) {
	ts := util.Int642String(time.Now().Unix())
	nonceStr := util.GetRandomString(32)
	prepayId := prepayid

	_str := c.Appid + "\n" + ts + "\n" + nonceStr + "\n" + prepayId + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}

	app = &AppPayParams{
		Appid:     c.Appid,
		Partnerid: c.Mchid,
		Prepayid:  prepayid,
		Package:   "Sign=WXPay",
		Noncestr:  nonceStr,
		Timestamp: ts,
		PaySign:   sign,
	}
	return app, nil
}

// PaySignOfApplet 获取 小程序 paySign
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_4.shtml
func (c *ClientV3) PaySignOfApplet(prepayid string) (applet *AppletParams, err error) {
	ts := util.Int642String(time.Now().Unix())
	nonceStr := util.GetRandomString(32)
	prepayId := "prepay_id=" + prepayid

	_str := c.Appid + "\n" + ts + "\n" + nonceStr + "\n" + prepayId + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}

	applet = &AppletParams{
		AppId:     c.Appid,
		TimeStamp: ts,
		NonceStr:  nonceStr,
		Package:   prepayId,
		SignType:  SignTypeRSA,
		PaySign:   sign,
	}
	return applet, nil
}

// v3 鉴权请求Header
func (c *ClientV3) authorization(method, path string, bm gopay.BodyMap) (string, error) {
	var (
		jb        = ""
		timestamp = time.Now().Unix()
		nonceStr  = util.GetRandomString(32)
	)
	if bm != nil {
		jb = bm.JsonBody()
	}
	ts := util.Int642String(timestamp)
	_str := method + "\n" + path + "\n" + ts + "\n" + nonceStr + "\n" + jb + "\n"
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
		return "", fmt.Errorf("rsa.SignPKCS1v15(),err:%+v", err)
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

// 自动同步请求验签
func (c *ClientV3) verifySyncSign(si *SignInfo) (err error) {
	if c.autoSign {
		if si != nil {
			var (
				block     *pem.Block
				pubKey    *x509.Certificate
				publicKey *rsa.PublicKey
				ok        bool
			)
			str := si.HeaderTimestamp + "\n" + si.HeaderNonce + "\n" + si.SignBody + "\n"
			signBytes, _ := base64.StdEncoding.DecodeString(si.HeaderSignature)

			if block, _ = pem.Decode(c.wxPkContent); block == nil {
				return errors.New("parse wechat platform public key error")
			}
			if pubKey, err = x509.ParseCertificate(block.Bytes); err != nil {
				return fmt.Errorf("x509.ParseCertificate：%+v", err)
			}
			if publicKey, ok = pubKey.PublicKey.(*rsa.PublicKey); !ok {
				return errors.New("convert wechat platform public to rsa.PublicKey error")
			}
			h := sha256.New()
			h.Write([]byte(str))
			if err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
				return fmt.Errorf("verify sign failed: %+v", err)
			}
			return nil
		}
		return errors.New("auto verify sign, bug SignInfo is nil")
	}
	return nil
}
