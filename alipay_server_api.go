package gopay

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"log"
	"net/http"
	"reflect"
	"time"
)

//解析支付宝支付完成后的Notify信息
func ParseAliPayNotifyResult(req *http.Request) (notifyRsp *AliPayNotifyRequest, err error) {
	notifyRsp = new(AliPayNotifyRequest)
	notifyRsp.NotifyTime = req.FormValue("notify_time")
	notifyRsp.NotifyType = req.FormValue("notify_type")
	notifyRsp.NotifyId = req.FormValue("notify_id")
	notifyRsp.AppId = req.FormValue("app_id")
	notifyRsp.Charset = req.FormValue("charset")
	notifyRsp.Version = req.FormValue("version")
	notifyRsp.SignType = req.FormValue("sign_type")
	notifyRsp.Sign = req.FormValue("sign")
	notifyRsp.AuthAppId = req.FormValue("auth_app_id")
	notifyRsp.TradeNo = req.FormValue("trade_no")
	notifyRsp.OutTradeNo = req.FormValue("out_trade_no")
	notifyRsp.OutBizNo = req.FormValue("out_biz_no")
	notifyRsp.BuyerId = req.FormValue("buyer_id")
	notifyRsp.BuyerLogonId = req.FormValue("buyer_logon_id")
	notifyRsp.SellerId = req.FormValue("seller_id")
	notifyRsp.SellerEmail = req.FormValue("seller_email")
	notifyRsp.TradeStatus = req.FormValue("trade_status")
	notifyRsp.TotalAmount = req.FormValue("total_amount")
	notifyRsp.ReceiptAmount = req.FormValue("receipt_amount")
	notifyRsp.InvoiceAmount = req.FormValue("invoice_amount")
	notifyRsp.BuyerPayAmount = req.FormValue("buyer_pay_amount")
	notifyRsp.PointAmount = req.FormValue("point_amount")
	notifyRsp.RefundFee = req.FormValue("refund_fee")
	notifyRsp.Subject = req.FormValue("subject")
	notifyRsp.Body = req.FormValue("body")
	notifyRsp.GmtCreate = req.FormValue("gmt_create")
	notifyRsp.GmtPayment = req.FormValue("gmt_payment")
	notifyRsp.GmtRefund = req.FormValue("gmt_refund")
	notifyRsp.GmtClose = req.FormValue("gmt_close")
	billList := req.FormValue("fund_bill_list")
	//log.Println("billList:", billList)
	if billList != null {
		bills := make([]fundBillListInfo, 0)
		err = json.Unmarshal([]byte(billList), &bills)
		if err != nil {
			return nil, err
		}
		notifyRsp.FundBillList = bills
	} else {
		notifyRsp.FundBillList = nil
	}
	notifyRsp.PassbackParams = req.FormValue("passback_params")
	detailList := req.FormValue("voucher_detail_list")
	//log.Println("detailList:", detailList)
	if detailList != null {
		details := make([]voucherDetailListInfo, 0)
		err = json.Unmarshal([]byte(detailList), &details)
		if err != nil {
			return nil, err
		}
		notifyRsp.VoucherDetailList = details
	} else {
		notifyRsp.VoucherDetailList = nil
	}
	return notifyRsp, err
}

//解密支付宝开放数据
//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
//    sessionKey:会话密钥
//    beanPtr:需要解析到的结构体指针
func DecryptAliPayOpenDataToStruct(encryptedData, secretKey string, beanPtr interface{}) (err error) {
	//验证参数类型
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入参数类型必须是以指针形式")
	}
	//验证interface{}类型
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
	}
	aesKey, _ := base64.StdEncoding.DecodeString(secretKey)
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	secretData, _ := base64.StdEncoding.DecodeString(encryptedData)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}
	if len(secretData)%len(aesKey) != 0 {
		return errors.New("encryptedData is error")
	}

	blockMode := cipher.NewCBCDecrypter(block, ivKey)
	originData := make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	originData = PKCS5UnPadding(originData)
	//fmt.Println("originDataStr:", string(originData))
	//解析
	err = json.Unmarshal(originData, beanPtr)
	if err != nil {
		return err
	}
	return nil
}

//支付通知的签名验证和参数签名后的Sign
//    aliPayPublicKey：支付宝公钥
//    notifyRsp：利用 gopay.ParseAliPayNotifyResult() 得到的结构体
//    返回参数ok：是否验证通过
//    返回参数err：错误信息
func VerifyAliPayResultSign(aliPayPublicKey string, notifyRsp *AliPayNotifyRequest) (ok bool, err error) {
	body := make(BodyMap)
	body.Set("notify_time", notifyRsp.NotifyTime)
	body.Set("notify_type", notifyRsp.NotifyType)
	body.Set("notify_id", notifyRsp.NotifyId)
	body.Set("app_id", notifyRsp.AppId)
	body.Set("charset", notifyRsp.Charset)
	body.Set("version", notifyRsp.Version)
	//body.Set("sign", notifyRsp.Sign)          //验签时去掉
	//body.Set("sign_type", notifyRsp.SignType) //验签时去掉
	body.Set("auth_app_id", notifyRsp.AuthAppId)
	body.Set("trade_no", notifyRsp.TradeNo)
	body.Set("out_trade_no", notifyRsp.OutTradeNo)
	body.Set("out_biz_no", notifyRsp.OutBizNo)
	body.Set("buyer_id", notifyRsp.BuyerId)
	body.Set("buyer_logon_id", notifyRsp.BuyerLogonId)
	body.Set("seller_id", notifyRsp.SellerId)
	body.Set("seller_email", notifyRsp.SellerEmail)
	body.Set("trade_status", notifyRsp.TradeStatus)
	body.Set("total_amount", notifyRsp.TotalAmount)
	body.Set("receipt_amount", notifyRsp.ReceiptAmount)
	body.Set("invoice_amount", notifyRsp.InvoiceAmount)
	body.Set("buyer_pay_amount", notifyRsp.BuyerPayAmount)
	body.Set("point_amount", notifyRsp.PointAmount)
	body.Set("refund_fee", notifyRsp.RefundFee)
	body.Set("subject", notifyRsp.Subject)
	body.Set("body", notifyRsp.Body)
	body.Set("gmt_create", notifyRsp.GmtCreate)
	body.Set("gmt_payment", notifyRsp.GmtPayment)
	body.Set("gmt_refund", notifyRsp.GmtRefund)
	body.Set("gmt_close", notifyRsp.GmtClose)
	body.Set("fund_bill_list", jsonToString(notifyRsp.FundBillList))
	body.Set("passback_params", notifyRsp.PassbackParams)
	body.Set("voucher_detail_list", jsonToString(notifyRsp.VoucherDetailList))

	newBody := make(BodyMap)
	for k, v := range body {
		if v != null {
			newBody.Set(k, v)
		}
	}

	pKey := FormatAliPayPublicKey(aliPayPublicKey)
	signData := sortAliPaySignParams(newBody)

	//log.Println("签名字符串：", signData)
	err = verifyAliPaySign(signData, notifyRsp.Sign, notifyRsp.SignType, pKey)
	if err != nil {
		return false, err
	}
	return true, nil
}

func jsonToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	bs, err := json.Marshal(v)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}
	s := string(bs)
	if s == "null" {
		return ""
	}
	return s
}

//格式化秘钥
func FormatPrivateKey(privateKey string) (pKey string) {
	buffer := new(bytes.Buffer)
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
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

//格式化秘钥
func FormatAliPayPublicKey(publicKey string) (pKey string) {
	buffer := new(bytes.Buffer)
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
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

func verifyAliPaySign(signData, sign, signType, aliPayPublicKey string) (err error) {
	var (
		h     hash.Hash
		hashs crypto.Hash
	)
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	//解析秘钥
	block, _ := pem.Decode([]byte(aliPayPublicKey))
	if block == nil {
		return errors.New("支付宝公钥Decode错误")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println("x509.ParsePKIXPublicKey:", err)
		return err
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return errors.New("支付宝公钥转换错误")
	}
	//判断签名方式
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

//换取授权访问令牌（默认使用utf-8，RSA2）
//    appId：应用ID
//    privateKey：应用私钥
//    grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
//    codeOrToken：支付宝授权码或refresh_token
func AlipaySystemOauthToken(appId, privateKey, grantType, codeOrToken string) (rsp *AliPaySystemOauthTokenResponse, err error) {
	var bs []byte
	body := make(BodyMap)
	if "authorization_code" == grantType {
		body.Set("grant_type", "authorization_code")
		body.Set("code", codeOrToken)
	} else if "refresh_token" == grantType {
		body.Set("grant_type", "refresh_token")
		body.Set("refresh_token", codeOrToken)
	} else {
		body.Set("grant_type", "authorization_code")
		body.Set("code", codeOrToken)
	}
	bs, err = doAliPay(appId, privateKey, body, "alipay.system.oauth.token")
	if err != nil {
		return nil, err
	}
	//fmt.Println("bs:", string(bs))
	rsp = new(AliPaySystemOauthTokenResponse)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.AliPaySystemOauthTokenResponse.AccessToken != "" {
		return rsp, nil
	} else {

		return
	}
}

//向支付宝发送请求
func doAliPay(appId, privateKey string, body BodyMap, method string) (bytes []byte, err error) {
	//===============生成参数===================
	//pubBody := make(BodyMap)
	body.Set("app_id", appId)
	body.Set("method", method)
	body.Set("format", "JSON")
	body.Set("charset", "utf-8")
	body.Set("sign_type", "RSA2")
	body.Set("timestamp", time.Now().Format(TimeLayout))
	body.Set("version", "1.0")
	//pubBody.Set("biz_content", string(bodyStr))
	//===============获取签名===================
	pKey := FormatPrivateKey(privateKey)
	sign, err := getRsaSign(body, "RSA2", pKey)
	if err != nil {
		return nil, err
	}
	body.Set("sign", sign)
	//fmt.Println("rsaSign:", sign)
	//===============发起请求===================
	urlParam := FormatAliPayURLParam(body)
	//fmt.Println("urlParam:", urlParam)

	var url string
	agent := HttpAgent()
	//正式环境
	url = zfb_base_url
	//fmt.Println(url)
	agent.Post(url)
	_, bs, errs := agent.
		Type("form-data").
		SendString(urlParam).
		EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return bs, nil
}
