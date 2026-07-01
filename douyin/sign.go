package douyin

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/util/convert"
)

// authorization 构造抖音支付 HTTP 请求头 Authorization
// 签名串 5 行：{method}\n{path}\n{timestamp}\n{nonceStr}\n{body}\n
func (c *Client) authorization(method, path string, bm gopay.BodyMap) (string, error) {
	var (
		jb        = ""
		timestamp = time.Now().Unix()
		nonceStr  = util.RandomString(32)
	)
	if bm != nil {
		jb = bm.JsonBody()
	}
	path = strings.TrimSuffix(path, "?")
	ts := convert.Int64ToString(timestamp)
	_str := method + "\n" + path + "\n" + ts + "\n" + nonceStr + "\n" + jb + "\n"
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Douyin_SignString:\n%s", _str)
	}
	sign, err := c.rsaSign(_str)
	if err != nil {
		return "", err
	}
	return Authorization + ` mchid="` + c.Mchid +
		`",nonce_str="` + nonceStr +
		`",timestamp="` + ts +
		`",serial_no="` + c.SerialNo +
		`",signature="` + sign + `"`, nil
}

// rsaSign 使用商户私钥进行 SHA256withRSA 签名并 Base64 编码
func (c *Client) rsaSign(str string) (string, error) {
	if c.privateKey == nil {
		return gopay.NULL, errors.New("privateKey can't be nil")
	}
	sum256 := sha256.Sum256([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, sum256[:])
	if err != nil {
		return gopay.NULL, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

// VerifySignByPK 抖音支付同步/回调应答验签（外部调用入口）
// 验签串 3 行：{timestamp}\n{nonce}\n{body}\n
func VerifySignByPK(timestamp, nonce, signBody, sign string, publicKey *rsa.PublicKey) error {
	if publicKey == nil || publicKey.N == nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, "publicKey is nil")
	}
	str := timestamp + "\n" + nonce + "\n" + signBody + "\n"
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	sum256 := sha256.Sum256([]byte(str))
	if err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, sum256[:], signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}

// verifySyncSign 同步应答自动验签（仅在 autoSign=true 且 SignInfo 非空时执行）
// 在验签前会先校验响应 Header 时间戳与本地时间的偏差（可通过 c.RespTimestampWindow 调整或关闭）
func (c *Client) verifySyncSign(si *SignInfo) error {
	if !c.autoSign {
		return nil
	}
	if si == nil {
		return errors.New("auto verify sign, but SignInfo is nil")
	}
	if c.RespTimestampWindow > 0 {
		ts, err := strconv.ParseInt(si.HeaderTimestamp, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid Douyinpay-Timestamp: %q, err: %v", si.HeaderTimestamp, err)
		}
		diff := time.Now().Unix() - ts
		if diff < 0 {
			diff = -diff
		}
		if diff >= c.RespTimestampWindow {
			return fmt.Errorf("response timestamp expired: diff=%ds, window=%ds", diff, c.RespTimestampWindow)
		}
	}
	pubKey, ok := c.getPlatformKey(si.HeaderSerial)
	if !ok {
		return fmt.Errorf("auto verify sign, but public key of serial(%s) not found, please call SetPlatformCert() first", si.HeaderSerial)
	}
	return VerifySignByPK(si.HeaderTimestamp, si.HeaderNonce, si.SignBody, si.HeaderSignature, pubKey)
}

// PaySignOfApp 生成 App 端调起抖音支付所需要的参数
// 签名串 4 行：{appid}\n{timestamp}\n{nonceStr}\n{prepayId}\n
func (c *Client) PaySignOfApp(appid, prepayId string) (params *AppPayParams, err error) {
	ts := convert.Int64ToString(time.Now().Unix())
	nonceStr := util.RandomString(32)
	_str := appid + "\n" + ts + "\n" + nonceStr + "\n" + prepayId + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}
	params = &AppPayParams{
		AppId:     appid,
		PartnerId: c.Mchid,
		PrepayId:  prepayId,
		Package:   AppPackage,
		NonceStr:  nonceStr,
		Timestamp: ts,
		Sign:      sign,
	}
	return params, nil
}

// PaySignOfJSAPI 生成 JSAPI 前端调起支付所需要的参数
// 签名串 4 行：{appId}\n{timeStamp}\n{nonceStr}\n{package(prepay_id=xxx)}\n
func (c *Client) PaySignOfJSAPI(appid, prepayId string) (params *JSAPIPayParams, err error) {
	ts := convert.Int64ToString(time.Now().Unix())
	nonceStr := util.RandomString(32)
	pkg := "prepay_id=" + prepayId
	_str := appid + "\n" + ts + "\n" + nonceStr + "\n" + pkg + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}
	params = &JSAPIPayParams{
		AppId:     appid,
		TimeStamp: ts,
		NonceStr:  nonceStr,
		Package:   pkg,
		SignType:  SignTypeRSA,
		PaySign:   sign,
	}
	return params, nil
}
