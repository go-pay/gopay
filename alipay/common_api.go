package alipay

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/go-pay/gopay"
	xaes "github.com/go-pay/gopay/pkg/aes"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/pkg/xrsa"
)

// 格式化请求URL参数
func FormatURLParam(body gopay.BodyMap) (urlParam string) {
	v := url.Values{}
	for key, value := range body {
		v.Add(key, value.(string))
	}
	return v.Encode()
}

// DecryptOpenDataToStruct 解密支付宝开放数据到 结构体
// encryptedData:包括敏感数据在内的完整用户信息的加密数据
// secretKey:AES密钥，支付宝管理平台配置
// beanPtr:需要解析到的结构体指针
// 文档：https://opendocs.alipay.com/common/02mse3
func DecryptOpenDataToStruct(encryptedData, secretKey string, beanPtr any) (err error) {
	if encryptedData == util.NULL || secretKey == util.NULL {
		return errors.New("encryptedData or secretKey is null")
	}
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入参数类型必须是以指针形式")
	}
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入any必须是结构体")
	}
	var (
		block      cipher.Block
		blockMode  cipher.BlockMode
		originData []byte
	)
	aesKey, _ := base64.StdEncoding.DecodeString(secretKey)
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	secretData, _ := base64.StdEncoding.DecodeString(encryptedData)
	if block, err = aes.NewCipher(aesKey); err != nil {
		return fmt.Errorf("aes.NewCipher：%w", err)
	}
	if len(secretData)%len(aesKey) != 0 {
		return errors.New("encryptedData is error")
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = xaes.PKCS5UnPadding(originData)
	}
	if err = json.Unmarshal(originData, beanPtr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s)：%w", string(originData), err)
	}
	return nil
}

// DecryptOpenDataToBodyMap 解密支付宝开放数据到 BodyMap
// encryptedData:包括敏感数据在内的完整用户信息的加密数据
// secretKey:AES密钥，支付宝管理平台配置
// 文档：https://opendocs.alipay.com/common/02mse3
func DecryptOpenDataToBodyMap(encryptedData, secretKey string) (bm gopay.BodyMap, err error) {
	if encryptedData == util.NULL || secretKey == util.NULL {
		return nil, errors.New("encryptedData or secretKey is null")
	}
	var (
		aesKey, originData []byte
		ivKey              = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		block              cipher.Block
		blockMode          cipher.BlockMode
	)
	aesKey, _ = base64.StdEncoding.DecodeString(secretKey)
	secretData, _ := base64.StdEncoding.DecodeString(encryptedData)
	if block, err = aes.NewCipher(aesKey); err != nil {
		return nil, fmt.Errorf("aes.NewCipher：%w", err)
	}
	if len(secretData)%len(aesKey) != 0 {
		return nil, errors.New("encryptedData is error")
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = xaes.PKCS5UnPadding(originData)
	}
	bm = make(gopay.BodyMap)
	if err = json.Unmarshal(originData, &bm); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(originData), err)
	}
	return
}

// SystemOauthToken 换取授权访问令牌（默认使用utf-8，RSA2）
// appId：应用ID
// privateKey：应用私钥
// grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
// codeOrToken：支付宝授权码或refresh_token
// signType：签名方式 RSA 或 RSA2，默认 RSA2
// appAuthToken：可选参数，三方授权令牌
// 文档：https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token
func SystemOauthToken(ctx context.Context, appId string, privateKey, grantType, codeOrToken, signType string, appAuthToken ...string) (rsp *SystemOauthTokenResponse, err error) {
	key := xrsa.FormatAlipayPrivateKey(privateKey)
	aat := ""
	if len(appAuthToken) > 0 {
		aat = appAuthToken[0]
	}
	priKey, err := xpem.DecodePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	var bs []byte
	bm := make(gopay.BodyMap)
	switch grantType {
	case "authorization_code":
		bm.Set("grant_type", "authorization_code")
		bm.Set("code", codeOrToken)
	case "refresh_token":
		bm.Set("grant_type", "refresh_token")
		bm.Set("refresh_token", codeOrToken)
	default:
		bm.Set("grant_type", "authorization_code")
		bm.Set("code", codeOrToken)
	}
	if bs, err = systemOauthToken(ctx, appId, priKey, bm, "alipay.system.oauth.token", true, signType, aat); err != nil {
		return
	}
	rsp = new(SystemOauthTokenResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if (rsp.Response == nil) || (rsp.Response != nil && rsp.Response.AccessToken == "") {
		return nil, errors.New("response is nil or access_token is NULL")
	}
	return
}

// systemOauthToken 向支付宝发送请求
func systemOauthToken(ctx context.Context, appId string, privateKey *rsa.PrivateKey, bm gopay.BodyMap, method string, isProd bool, signType, appAuthToken string) (bs []byte, err error) {
	bm.Set("app_id", appId)
	bm.Set("method", method)
	bm.Set("format", "JSON")
	bm.Set("charset", "utf-8")
	if signType == util.NULL {
		bm.Set("sign_type", RSA2)
	} else {
		bm.Set("sign_type", signType)
	}
	bm.Set("timestamp", time.Now().Format(util.TimeLayout))
	bm.Set("version", "1.0")
	if appAuthToken != util.NULL {
		bm.Set("app_auth_token", appAuthToken)
	}
	var (
		sign    string
		baseUrl = baseUrlUtf8
	)
	if sign, err = GetRsaSign(bm, bm.GetString("sign_type"), privateKey); err != nil {
		return nil, err
	}
	bm.Set("sign", sign)
	if !isProd {
		baseUrl = sandboxBaseUrlUtf8
	}
	_, bs, err = xhttp.NewClient().Req(xhttp.TypeFormData).Post(baseUrl).SendString(bm.EncodeURLParams()).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// monitor.heartbeat.syn(验签接口)
// appId：应用ID
// privateKey：应用私钥，支持PKCS1和PKCS8
// signType：签名方式 alipay.RSA 或 alipay.RSA2，默认 RSA2
// bizContent：验签时该参数不做任何处理，{任意值}，此参数具体看文档
// 文档：https://opendocs.alipay.com/apis/api_9/monitor.heartbeat.syn
func MonitorHeartbeatSyn(ctx context.Context, appId string, privateKey, signType, bizContent string) (aliRsp *MonitorHeartbeatSynResponse, err error) {
	key := xrsa.FormatAlipayPrivateKey(privateKey)
	priKey, err := xpem.DecodePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	var bs []byte
	bm := make(gopay.BodyMap)
	bm.Set("biz_content", bizContent)
	bm.Set("app_id", appId)
	bm.Set("method", "monitor.heartbeat.syn")
	bm.Set("format", "JSON")
	bm.Set("charset", "utf-8")
	if signType == util.NULL {
		bm.Set("sign_type", RSA2)
	} else {
		bm.Set("sign_type", signType)
	}
	bm.Set("timestamp", time.Now().Format(util.TimeLayout))
	bm.Set("version", "1.0")

	sign, err := GetRsaSign(bm, bm.GetString("sign_type"), priKey)
	if err != nil {
		return nil, err
	}
	bm.Set("sign", sign)

	_, bs, err = xhttp.NewClient().Req(xhttp.TypeFormData).Post(baseUrlUtf8).SendString(bm.EncodeURLParams()).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	aliRsp = new(MonitorHeartbeatSynResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	return aliRsp, nil
}
