package paypal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/util/retry"
)

func (c *Client) goAuthRefreshToken() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			c.logger.Errorf("paypal_goAuthRefreshToken: panic recovered: %s\n%s", r, buf)
		}
	}()
	for {
		time.Sleep(time.Duration(c.ExpiresIn/2) * time.Second)
		err := retry.Retry(func() error {
			_, err := c.GetAccessToken()
			if err != nil {
				return err
			}
			return nil
		}, 3, time.Second)
		if err != nil {
			c.logger.Errorf("PayPal GetAccessToken Error: %s", err.Error())
		}
	}
}

// =====================================================================================================================

// 获取AccessToken（Get an access token）
// 文档：https://developer.paypal.com/docs/api/reference/get-an-access-token
func (c *Client) GetAccessToken() (token *AccessToken, err error) {
	var (
		baseUrl = c.baseUrlProd
		url     string
	)
	if !c.IsProd {
		baseUrl = c.baseUrlSandbox
	}
	url = baseUrl + getAccessToken
	// Authorization
	authHeader := AuthorizationPrefixBasic + base64.StdEncoding.EncodeToString([]byte(c.Clientid+":"+c.Secret))
	req := c.hc.Req(xhttp.TypeFormData)
	req.Header.Add(HeaderAuthorization, authHeader)
	req.Header.Add("Accept", "*/*")
	// Body
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "client_credentials")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Url: %s", url)
		c.logger.Debugf("PayPal_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("PayPal_Req_Headers: %#v", req.Header)
	}
	res, bs, err := req.Post(url).SendBodyMap(bm).EndBytes(c.ctx)
	if err != nil {
		return nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("PayPal_Rsp_Headers: %#v", res.Header)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	token = new(AccessToken)
	if err = json.Unmarshal(bs, token); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s): %w", string(bs), err)
	}
	c.Appid = token.Appid
	c.AccessToken = token.AccessToken
	c.ExpiresIn = token.ExpiresIn
	return token, nil
}
