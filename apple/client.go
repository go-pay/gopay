package apple

import (
	"context"
	"crypto/ecdsa"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
)

// Client AppleClient
type Client struct {
	iss        string // Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
	bid        string // Your app’s bundle ID (Ex: “com.example.testbundleid2021”)
	kid        string // Your private key ID from App Store Connect (Ex: 2X9R4HXF34)
	isProd     bool   // 是否是正式环境
	privateKey *ecdsa.PrivateKey
	hc         *xhttp.Client
}

// NewClient 初始化Apple客户端
// iss：issuer ID
// bid：bundle ID
// kid：private key ID
// privateKey：私钥文件读取后的字符串内容
// isProd：是否是正式环境
func NewClient(iss, bid, kid, privateKey string, isProd bool) (client *Client, err error) {
	if iss == util.NULL || bid == util.NULL || kid == util.NULL || privateKey == util.NULL {
		return nil, gopay.MissAppleInitParamErr
	}
	ecPrivateKey, err := ParseECPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	//ecPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//if err != nil {
	//	panic(err)
	//}
	client = &Client{
		iss:        iss,
		bid:        bid,
		kid:        kid,
		privateKey: ecPrivateKey,
		isProd:     isProd,
		hc:         xhttp.NewClient(),
	}
	return client, nil
}

func (c *Client) doRequestGet(ctx context.Context, path string) (res *http.Response, bs []byte, err error) {
	uri := hostUrl + path
	if !c.isProd {
		uri = sandBoxHostUrl + path
	}
	token, err := c.generatingToken()
	if err != nil {
		return nil, nil, err
	}
	req := c.hc.Req()
	req.Header.Set("Authorization", "Bearer "+token)
	res, bs, err = req.Get(uri).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	return res, bs, nil
}

func (c *Client) doRequestPost(ctx context.Context, path string, bm gopay.BodyMap) (res *http.Response, bs []byte, err error) {
	uri := hostUrl + path
	if !c.isProd {
		uri = sandBoxHostUrl + path
	}
	token, err := c.generatingToken()
	if err != nil {
		return nil, nil, err
	}
	req := c.hc.Req()
	req.Header.Set("Authorization", "Bearer "+token)
	res, bs, err = req.Post(uri).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	return res, bs, nil
}

func (c *Client) doRequestPut(ctx context.Context, path string, bm gopay.BodyMap) (res *http.Response, bs []byte, err error) {
	uri := hostUrl + path
	if !c.isProd {
		uri = sandBoxHostUrl + path
	}
	token, err := c.generatingToken()
	if err != nil {
		return nil, nil, err
	}
	req := c.hc.Req()
	req.Header.Set("Authorization", "Bearer "+token)
	res, bs, err = req.Put(uri).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	return res, bs, nil
}
