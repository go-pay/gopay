//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/13 14:42
//==================================
package gopay

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"hash"
	"log"
)

type aliPayPublicBody struct {
	AppId      string  `json:"app_id"`      //支付宝分配给开发者的应用ID
	Method     string  `json:"method"`      //接口名称
	Format     string  `json:"format"`      //仅支持 JSON
	ReturnUrl  string  `json:"return_url"`  //HTTP/HTTPS开头字符串
	Charset    string  `json:"charset"`     //请求使用的编码格式，如utf-8,gbk,gb2312等，推荐使用 utf-8
	SignType   string  `json:"sign_type"`   //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
	Sign       string  `json:"sign"`        //商户请求参数的签名串
	Timestamp  string  `json:"timestamp"`   //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version    string  `json:"version"`     //调用的接口版本，固定为：1.0
	NotifyUrl  string  `json:"notify_url"`  //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
	BizContent BodyMap `json:"biz_content"` //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}

//设置支付后的ReturnUrl
func (this *aliPayClient) SetReturnUrl(url string) (client *aliPayClient) {
	this.ReturnUrl = url
	return this
}

//设置支付宝服务器主动通知商户服务器里指定的页面http/https路径。
func (this *aliPayClient) SetNotifyUrl(url string) (client *aliPayClient) {
	this.NotifyUrl = url
	return this
}

//设置编码格式，如utf-8,gbk,gb2312等，推荐使用 utf-8
func (this *aliPayClient) SetCharset(charset string) (client *aliPayClient) {
	if charset == null {
		this.Charset = "utf-8"
	} else {
		this.Charset = charset
	}
	return this
}

//设置签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
func (this *aliPayClient) SetSignType(signType string) (client *aliPayClient) {
	if signType == null {
		this.SignType = "RSA2"
	} else {
		this.SignType = signType
	}
	return this
}

//app_id=2014072300007148&biz_content={"button"}&charset=GBK&method=alipay.mobile.public.menu.add&sign_type=RSA2&timestamp=2014-07-24 03:07:50&version=1.0
func getRsaSign(appId, signType, charset, method, timestamp, privateKey string, bodyMap BodyMap) (sign string, err error) {
	var (
		h              hash.Hash
		key            *rsa.PrivateKey
		buffer         *bytes.Buffer
		bodyStr        []byte
		encryptedBytes []byte
	)
	bodyStr, err = json.Marshal(bodyMap)
	if err != nil {
		log.Println("json.Marshal:", err)
		return "", err
	}
	//log.Println("privateKey:", privateKey)
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("秘钥错误")
	}
	log.Println(block.Type)
	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println("x509.ParsePKCS1PrivateKey:", err)
		return "", err
	}

	switch signType {
	case "RSA":
		h = sha1.New()
	case "RSA2":
		h = sha256.New()
	}
	buffer = new(bytes.Buffer)
	buffer.WriteString("app_id=")
	buffer.WriteString(appId)
	buffer.WriteString("&biz_content=")
	buffer.WriteString(string(bodyStr))
	buffer.WriteString("&charset=")
	buffer.WriteString(charset)
	buffer.WriteString("&method=")
	buffer.WriteString(method)
	buffer.WriteString("&sign_type=")
	buffer.WriteString(signType)
	buffer.WriteString("&timestamp=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&version=1.0")
	log.Println("参数拼接：", buffer.String())

	_, err = h.Write(buffer.Bytes())
	if err != nil {
		return "", err
	}

	encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, h.Sum(nil))
	if err != nil {
		log.Println("rsa.SignPKCS1v15:", err)
		return "", err
	}
	secretData := base64.StdEncoding.EncodeToString(encryptedBytes)
	return secretData, nil
}
