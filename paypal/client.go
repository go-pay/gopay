package paypal

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
)

// Client PayPal支付客户端
type Client struct {
	Clientid    string
	Secret      string
	Appid       string
	AccessToken string
	ExpiresIn   int
	bodySize    int // http response body size(MB), default is 10MB
	IsProd      bool
	ctx         context.Context
	DebugSwitch gopay.DebugSwitch
}

// NewClient 初始化PayPal支付客户端
func NewClient(clientid, secret string, isProd bool) (client *Client, err error) {
	if clientid == util.NULL || secret == util.NULL {
		return nil, gopay.MissPayPalInitParamErr
	}
	client = &Client{
		Clientid:    clientid,
		Secret:      secret,
		IsProd:      isProd,
		ctx:         context.Background(),
		DebugSwitch: gopay.DebugOff,
	}
	_, err = client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.bodySize = sizeMB
	}
}

func (c *Client) doPayPalGet(ctx context.Context, uri string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + uri
	if !c.IsProd {
		url = baseUrlSandbox + uri
	}
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Url: %s", url)
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalPost(ctx context.Context, bm gopay.BodyMap, path string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + path
	if !c.IsProd {
		url = baseUrlSandbox + path
	}
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalPatch(ctx context.Context, patchs []*Patch, path string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + path
	if !c.IsProd {
		url = baseUrlSandbox + path
	}
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == gopay.DebugOn {
		jb, _ := json.Marshal(patchs)
		xlog.Debugf("PayPal_RequestBody: %s", string(jb))
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Patch(url).SendStruct(patchs).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}
