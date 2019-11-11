package gopay

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// 允许进行 sn 提取的证书签名算法
var allowSignatureAlgorithm map[string]bool = map[string]bool{
	"MD2-RSA":       true,
	"MD5-RSA":       true,
	"SHA1-RSA":      true,
	"SHA256-RSA":    true,
	"SHA384-RSA":    true,
	"SHA512-RSA":    true,
	"SHA256-RSAPSS": true,
	"SHA384-RSAPSS": true,
	"SHA512-RSAPSS": true,
}

// ParseAliPayNotifyResult 解析支付宝支付完成后的Notify信息
func ParseAliPayNotifyResult(req *http.Request) (notifyReq *AliPayNotifyRequest, err error) {
	notifyReq = new(AliPayNotifyRequest)
	notifyReq.NotifyTime = req.FormValue("notify_time")
	notifyReq.NotifyType = req.FormValue("notify_type")
	notifyReq.NotifyId = req.FormValue("notify_id")
	notifyReq.AppId = req.FormValue("app_id")
	notifyReq.Charset = req.FormValue("charset")
	notifyReq.Version = req.FormValue("version")
	notifyReq.SignType = req.FormValue("sign_type")
	notifyReq.Sign = req.FormValue("sign")
	notifyReq.AuthAppId = req.FormValue("auth_app_id")
	notifyReq.TradeNo = req.FormValue("trade_no")
	notifyReq.OutTradeNo = req.FormValue("out_trade_no")
	notifyReq.OutBizNo = req.FormValue("out_biz_no")
	notifyReq.BuyerId = req.FormValue("buyer_id")
	notifyReq.BuyerLogonId = req.FormValue("buyer_logon_id")
	notifyReq.SellerId = req.FormValue("seller_id")
	notifyReq.SellerEmail = req.FormValue("seller_email")
	notifyReq.TradeStatus = req.FormValue("trade_status")
	notifyReq.TotalAmount = req.FormValue("total_amount")
	notifyReq.ReceiptAmount = req.FormValue("receipt_amount")
	notifyReq.InvoiceAmount = req.FormValue("invoice_amount")
	notifyReq.BuyerPayAmount = req.FormValue("buyer_pay_amount")
	notifyReq.PointAmount = req.FormValue("point_amount")
	notifyReq.RefundFee = req.FormValue("refund_fee")
	notifyReq.Subject = req.FormValue("subject")
	notifyReq.Body = req.FormValue("body")
	notifyReq.GmtCreate = req.FormValue("gmt_create")
	notifyReq.GmtPayment = req.FormValue("gmt_payment")
	notifyReq.GmtRefund = req.FormValue("gmt_refund")
	notifyReq.GmtClose = req.FormValue("gmt_close")
	billList := req.FormValue("fund_bill_list")
	if billList != null {
		bills := make([]fundBillListInfo, 0)
		if err = json.Unmarshal([]byte(billList), &bills); err != nil {
			return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
		}
		notifyReq.FundBillList = bills
	} else {
		notifyReq.FundBillList = nil
	}
	notifyReq.PassbackParams = req.FormValue("passback_params")
	detailList := req.FormValue("voucher_detail_list")
	if detailList != null {
		details := make([]voucherDetailListInfo, 0)
		if err = json.Unmarshal([]byte(detailList), &details); err != nil {
			return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
		}
		notifyReq.VoucherDetailList = details
	} else {
		notifyReq.VoucherDetailList = nil
	}
	return
}

/*
Q：使用公钥证书签名方式下，为什么开放平台网关的响应报文需要携带支付宝公钥证书SN（alipay_cert_sn）？
**
A：开发者上传自己的应用公钥证书后，开放平台会为开发者应用自动签发支付宝公钥证书供开发者下载，用来对开放平台网关响应报文做验签。

但是支付宝公钥证书可能因证书到期或者变更CA签发机构等原因，可能会重新签发证书。在重新签发前，开放平台会在门户上提前提醒开发者支付宝应用公钥证书变更时间。

但为避免开发者因未能及时感知支付宝公钥证书变更而导致验签失败，开放平台提供了一种支付宝公钥证书无感知升级机制，具体流程如下：
1）开放平台网关在响应报文中会多返回支付宝公钥证书SN
2）开放平台网关提供根据SN下载对应支付宝公钥证书的API接口
3）开发者在验签过程中，先比较本地使用的支付宝公钥证书SN与开放平台网关响应中SN是否一致。若不一致，可调用支付宝公钥证书下载接口下载对应SN的支付宝公钥证书。
4）对下载的支付宝公钥证书执行证书链校验，若校验通过，则用该证书验签。

基于该机制可实现支付宝公钥证书变更时开发者无感知，当前开放平台提供的SDK已基于该机制实现对应功能。若开发者未通过SDK接入，须自行实现该功能。
*/

// VerifyAliPaySign 支付宝同步返回验签或异步通知验签
//    注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
//    aliPayPublicKey：支付宝公钥
//    bean： 同步返回验签时，此参数为 aliRsp.SignData ；异步通知验签时，此参数为异步通知解析的结构体 notifyReq
//    syncSign：同步返回验签时，此参数必传，即：aliRsp.Sign ；异步通知验签时，不传此参数，否则会出错。
//    返回参数ok：是否验签通过
//    返回参数err：错误信息
//    验签文档：https://docs.open.alipay.com/200/106120
func VerifyAliPaySign(aliPayPublicKey string, bean interface{}, syncSign ...string) (ok bool, err error) {
	if bean == nil {
		return false, errors.New("bean is nil")
	}
	var (
		bodySign     string
		bodySignType string
		pKey         string
		signData     string
		bm           BodyMap
		bs           []byte
	)
	if len(syncSign) > 0 {
		bodySign = syncSign[0]
		bodySignType = "RSA2"
		signData = bean.(string)
		goto Verify
	}
	if bs, err = json.Marshal(bean); err != nil {
		return false, fmt.Errorf("json.Marshal：%v", err.Error())
	}
	bm = make(BodyMap)
	if err = json.Unmarshal(bs, &bm); err != nil {
		return false, fmt.Errorf("json.Unmarshal：%v", err.Error())
	}
	bodySign = bm.Get("sign")
	bodySignType = bm.Get("sign_type")
	bm.Remove("sign")
	bm.Remove("sign_type")
	signData = bm.EncodeAliPaySignParams()
Verify:
	pKey = FormatAliPayPublicKey(aliPayPublicKey)
	if err = verifyAliPaySign(signData, bodySign, bodySignType, pKey); err != nil {
		return false, err
	}
	return true, nil
}

func verifyAliPaySign(signData, sign, signType, aliPayPublicKey string) (err error) {
	var (
		h         hash.Hash
		hashs     crypto.Hash
		block     *pem.Block
		pubKey    interface{}
		publicKey *rsa.PublicKey
		ok        bool
	)
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	if block, _ = pem.Decode([]byte(aliPayPublicKey)); block == nil {
		return errors.New("支付宝公钥Decode错误")
	}
	if pubKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return fmt.Errorf("x509.ParsePKIXPublicKey：%v", err.Error())
	}
	if publicKey, ok = pubKey.(*rsa.PublicKey); !ok {
		return errors.New("支付宝公钥转换错误")
	}
	switch signType {
	case "RSA":
		hashs = crypto.SHA1
	case "RSA2":
		hashs = crypto.SHA256
	default:
		hashs = crypto.SHA256
	}
	h = hashs.New()
	h.Write([]byte(signData))
	return rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes)
}

// FormatPrivateKey 格式化 普通应用秘钥
func FormatPrivateKey(privateKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

// FormatAliPayPublicKey 格式化 普通支付宝公钥
func FormatAliPayPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

// GetCertSN 获取证书序列号SN
//    certPath：X.509证书文件路径(appCertPublicKey.crt、alipayRootCert.crt、alipayCertPublicKey_RSA2)
//    返回 sn：证书序列号(app_cert_sn、alipay_root_cert_sn、alipay_cert_sn)
//    返回 err：error 信息
func GetCertSN(certPath string) (sn string, err error) {
	var (
		certData           []byte
		certs              []*x509.Certificate
		name, serialNumber string
		h                  hash.Hash
	)
	certData, err = ioutil.ReadFile(certPath)
	if err != nil {
		return "", err
	}
	strs := strings.Split(string(certData), "-----END CERTIFICATE-----")
	for i := 0; i < len(strs); i++ {
		if strs[i] == "" {
			continue
		}
		if block, _ := pem.Decode([]byte(strs[i] + "-----END CERTIFICATE-----")); block != nil {
			if certs, err = x509.ParseCertificates(block.Bytes); err != nil {
				continue
			}
			if !allowSignatureAlgorithm[certs[0].SignatureAlgorithm.String()] {
				continue
			}
			name = certs[0].Issuer.String()
			serialNumber = certs[0].SerialNumber.String()
			h = md5.New()
			h.Write([]byte(name))
			h.Write([]byte(serialNumber))
			if sn == "" {
				sn += hex.EncodeToString(h.Sum(nil))
			} else {
				sn += "_"
				sn += hex.EncodeToString(h.Sum(nil))
			}
		}
	}
	if sn == "" {
		return "", errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// DecryptAliPayOpenDataToStruct 解密支付宝开放数据到 结构体
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    secretKey:AES密钥，支付宝管理平台配置
//    beanPtr:需要解析到的结构体指针
//    文档：https://docs.alipay.com/mini/introduce/aes
//    文档：https://docs.open.alipay.com/common/104567
func DecryptAliPayOpenDataToStruct(encryptedData, secretKey string, beanPtr interface{}) (err error) {
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入参数类型必须是以指针形式")
	}
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
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
		return fmt.Errorf("aes.NewCipher：%v", err.Error())
	}
	if len(secretData)%len(aesKey) != 0 {
		return errors.New("encryptedData is error")
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = PKCS5UnPadding(originData)
	}
	if err = json.Unmarshal(originData, beanPtr); err != nil {
		return fmt.Errorf("json.Unmarshal：%v", err.Error())
	}
	return nil
}

// DecryptAliPayOpenDataToBodyMap 解密支付宝开放数据到 BodyMap
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    secretKey:AES密钥，支付宝管理平台配置
//    文档：https://docs.alipay.com/mini/introduce/aes
//    文档：https://docs.open.alipay.com/common/104567
func DecryptAliPayOpenDataToBodyMap(encryptedData, secretKey string) (bm BodyMap, err error) {
	var (
		aesKey, originData []byte
		ivKey              = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		block              cipher.Block
		blockMode          cipher.BlockMode
	)
	aesKey, _ = base64.StdEncoding.DecodeString(secretKey)
	secretData, _ := base64.StdEncoding.DecodeString(encryptedData)
	if block, err = aes.NewCipher(aesKey); err != nil {
		return nil, fmt.Errorf("aes.NewCipher：%v", err.Error())
	}
	if len(secretData)%len(aesKey) != 0 {
		return nil, errors.New("encryptedData is error")
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = PKCS5UnPadding(originData)
	}
	bm = make(BodyMap)
	if err = json.Unmarshal(originData, &bm); err != nil {
		return nil, fmt.Errorf("json.Unmarshal：%v", err.Error())
	}
	return
}

// AliPaySystemOauthToken 换取授权访问令牌（默认使用utf-8，RSA2）
//    appId：应用ID
//    PrivateKey：应用私钥
//    grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
//    codeOrToken：支付宝授权码或refresh_token
//    文档：https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func AliPaySystemOauthToken(appId, privateKey, grantType, codeOrToken string) (rsp *AliPaySystemOauthTokenResponse, err error) {
	var bs []byte
	bm := make(BodyMap)
	if "authorization_code" == grantType {
		bm.Set("grant_type", "authorization_code")
		bm.Set("code", codeOrToken)
	} else if "refresh_token" == grantType {
		bm.Set("grant_type", "refresh_token")
		bm.Set("refresh_token", codeOrToken)
	} else {
		bm.Set("grant_type", "authorization_code")
		bm.Set("code", codeOrToken)
	}
	if bs, err = aliPaySystemOauthToken(appId, privateKey, bm, "alipay.system.oauth.token", true); err != nil {
		return
	}
	rsp = new(AliPaySystemOauthTokenResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal：%v", err.Error())
	}
	if rsp.Response.AccessToken == "" {
		return nil, errors.New("access_token is null")
	}
	return
}

// aliPaySystemOauthToken 向支付宝发送请求
func aliPaySystemOauthToken(appId, privateKey string, body BodyMap, method string, isProd bool) (bytes []byte, err error) {
	body.Set("app_id", appId)
	body.Set("method", method)
	body.Set("format", "JSON")
	body.Set("charset", "utf-8")
	body.Set("sign_type", "RSA2")
	body.Set("timestamp", time.Now().Format(TimeLayout))
	body.Set("version", "1.0")
	var (
		sign, url string
		errs      []error
	)
	pKey := FormatPrivateKey(privateKey)
	if sign, err = getRsaSign(body, "RSA2", pKey); err != nil {
		return
	}
	body.Set("sign", sign)
	agent := HttpAgent()
	if !isProd {
		url = zfbSandboxBaseUrlUtf8
	} else {
		url = zfbBaseUrlUtf8
	}
	if _, bytes, errs = agent.Post(url).Type("form-data").SendString(FormatAliPayURLParam(body)).EndBytes(); len(errs) > 0 {
		return nil, errs[0]
	}
	return
}
