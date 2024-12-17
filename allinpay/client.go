package allinpay

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"sync"

	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/crypto/xrsa"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/util"
)

type Client struct {
	orgId      string          // 集团/代理编号 可为空
	CusId      string          // 实际交易商户号
	AppId      string          // 平台分配的APPID
	SignType   string          // 签名类型
	isProd     bool            // 是否正式环境
	privateKey *rsa.PrivateKey // 商户的RSA私钥
	publicKey  *rsa.PublicKey  // 通联的公钥
	hc         *xhttp.Client
	mu         sync.Mutex
	sha1Hash   hash.Hash
}

// NewClient 初始化通联客户端
// cusId：实际交易商户号
// appid：平台分配的APPID
// privateKey：商户的RSA私钥
// publicKey：通联的公钥
// isProd：是否是正式环境
func NewClient(cusId, appId, privateKey, publicKey string, isProd bool) (*Client, error) {
	prk, err := xpem.DecodePrivateKey([]byte(xrsa.FormatAlipayPrivateKey(privateKey)))
	if err != nil {
		return nil, err
	}
	puk, err := xpem.DecodePublicKey([]byte(xrsa.FormatAlipayPublicKey(publicKey)))
	if err != nil {
		return nil, err
	}
	return &Client{
		CusId:      cusId,
		AppId:      appId,
		SignType:   RSA,
		isProd:     isProd,
		privateKey: prk,
		publicKey:  puk,
		hc:         xhttp.NewClient(),
		sha1Hash:   sha1.New(),
	}, nil
}

// SetOrgId 集团/代理商商户号（因orgid非必填）因此单开方法
func (c *Client) SetOrgId(id string) *Client {
	c.orgId = id
	return c
}

// getRsaSign 获取签名字符串
func (c *Client) getRsaSign(bm gopay.BodyMap, signType string, privateKey *rsa.PrivateKey) (sign string, err error) {
	var (
		hashs          crypto.Hash
		encryptedBytes []byte
	)
	switch signType {
	case RSA:
		hashs = crypto.SHA1
	case SM2:
		return "", errors.New("暂不支持SM2加密")
	default:
		hashs = crypto.SHA1
	}
	signParams := bm.EncodeAliPaySignParams()
	c.mu.Lock()
	defer func() {
		c.sha1Hash.Reset()
		c.mu.Unlock()
	}()
	c.sha1Hash.Write([]byte(signParams))
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, privateKey, hashs, c.sha1Hash.Sum(nil)); err != nil {
		return gopay.NULL, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

// pubParamsHandle 公共参数处理
func (c *Client) pubParamsHandle(bm gopay.BodyMap) (param string, err error) {
	bm.Set("cusid", c.CusId).
		Set("appid", c.AppId).
		Set("signtype", c.SignType)
	//集团/代理商商户号
	if c.orgId != gopay.NULL {
		bm.Set("orgid", c.orgId)
	}
	// version
	if version := bm.GetString("version"); version == gopay.NULL {
		bm.Set("version", "11")
	}
	bm.Set("randomstr", util.RandomString(20))

	sign, err := c.getRsaSign(bm, bm.GetString("signtype"), c.privateKey)
	if err != nil {
		return "", fmt.Errorf("GetRsaSign Error: %w", err)
	}
	bm.Set("sign", sign)
	param = bm.EncodeURLParams()
	return
}

// doPost 发起请求
func (c *Client) doPost(ctx context.Context, path string, bm gopay.BodyMap) (bs []byte, err error) {
	param, err := c.pubParamsHandle(bm)
	if err != nil {
		return nil, err
	}
	url := baseUrl
	if !c.isProd {
		url = sandboxBaseUrl
	}
	res, bs, err := c.hc.Req(xhttp.TypeFormData).Post(url + path).SendString(param).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// SetHttpClient 设置自定义的xhttp.Client
func (c *Client) SetHttpClient(client *xhttp.Client) {
	c.hc = client
}
