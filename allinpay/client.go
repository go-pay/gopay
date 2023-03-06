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
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/pkg/xrsa"
	"hash"
)

type Client struct {
	orgId      string
	CusId      string
	AppId      string
	SignType   string
	isProd     bool
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey // 支付宝证书公钥内容 alipayCertPublicKey_RSA2.crt
}

func NewClient(cusId, appId, privateKey, publicKey string) (*Client, error) {

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
		isProd:     true,
		privateKey: prk,
		publicKey:  puk,
	}, nil
}
func (c *Client) SetOrgId(id string) *Client {
	c.orgId = id
	return c
}

func (c *Client) SwitchEnv(b bool) *Client {
	c.isProd = b
	return c
}

func (c *Client) getRsaSign(bm gopay.BodyMap, signType string, privateKey *rsa.PrivateKey) (sign string, err error) {
	var (
		h              hash.Hash
		hashs          crypto.Hash
		encryptedBytes []byte
	)

	switch signType {
	case RSA:
		h = sha1.New()
		hashs = crypto.SHA1
	case SM2:
		return "", errors.New("暂不支持SM2加密")
	default:
		h = sha1.New()
		hashs = crypto.SHA1
	}
	signParams := bm.EncodeAliPaySignParams()
	if _, err = h.Write([]byte(signParams)); err != nil {
		return
	}
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, privateKey, hashs, h.Sum(nil)); err != nil {
		return util.NULL, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

// 公共参数处理
func (c *Client) pubParamsHandle(bm gopay.BodyMap) (param string, err error) {
	bm.Set("cusid", c.CusId).
		Set("appid", c.AppId).
		Set("signtype", c.SignType)

	if c.orgId != util.NULL {
		bm.Set("orgid", c.orgId)
	}
	// version
	if version := bm.GetString("version"); version == util.NULL {
		bm.Set("version", "11")
	}
	bm.Set("randomstr", util.RandomString(20))

	sign, err := c.getRsaSign(bm, bm.GetString("sign_type"), c.privateKey)
	if err != nil {
		return "", fmt.Errorf("GetRsaSign Error: %w", err)
	}
	bm.Set("sign", sign)
	param = bm.EncodeURLParams()

	return
}

func (c *Client) doPost(ctx context.Context, path string, bm gopay.BodyMap) (bs []byte, err error) {

	param, err := c.pubParamsHandle(bm)
	if err != nil {
		return nil, err
	}
	httpClient := xhttp.NewClient()
	url := baseUrl
	if !c.isProd {
		url = sandboxBaseUrl
	}
	res, bs, err := httpClient.Type(xhttp.TypeForm).Post(url + path).SendString(param).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}
