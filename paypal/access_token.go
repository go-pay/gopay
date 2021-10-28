package paypal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
)

// 获取AccessToken（Get an access token）
//	文档：https://developer.paypal.com/docs/api/reference/get-an-access-token
func (c *Client) GetAccessToken() (token *AccessToken, err error) {
	var (
		baseUrl = baseUrlProd
		url     string
	)
	if !c.IsProd {
		baseUrl = baseUrlSandbox
	}
	url = baseUrl + getAccessToken
	// Authorization
	authHeader := AuthorizationPrefixBasic + base64.StdEncoding.EncodeToString([]byte(c.Clientid+":"+c.Secret))
	// Request
	httpClient := xhttp.NewClient()
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	// Body
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "client_credentials")
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	res, bs, err := httpClient.Type(xhttp.TypeForm).Post(url).SendBodyMap(bm).EndBytes(c.ctx)
	if err != nil {
		return nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	token = new(AccessToken)
	if err = json.Unmarshal(bs, token); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	c.Appid = token.Appid
	c.AccessToken = token.AccessToken
	c.ExpiresIn = token.ExpiresIn
	return token, nil
}
