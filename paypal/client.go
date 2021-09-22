package paypal

import (
	"github.com/go-pay/gopay"
)

// Client PayPal支付客
type Client struct {
	Clientid    string
	Secret      string
	Appid       string
	AccessToken string
	ExpiresIn   int
	IsProd      bool
	DebugSwitch gopay.DebugSwitch
}

// NewClient 初始化PayPal支付客户端
func NewClient(clientid, secret string, isProd bool) (client *Client, err error) {
	client = &Client{
		Clientid:    clientid,
		Secret:      secret,
		IsProd:      isProd,
		DebugSwitch: gopay.DebugOff,
	}
	_, err = client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	return client, nil
}

//func (c *Client) doPayPalGet(uri, authorization string) (res *http.Response, bs []byte, err error) {
//	var url = v3BaseUrlCh + uri
//	httpClient := xhttp.NewClient()
//	if c.DebugSwitch == gopay.DebugOn {
//		xlog.Debugf("Wechat_V3_Url: %s", url)
//		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
//	}
//	httpClient.Header.Add(HeaderAuthorization, authorization)
//	httpClient.Header.Add(HeaderSerial, c.wxSerialNo)
//	httpClient.Header.Add("Accept", "*/*")
//	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes()
//	if len(errs) > 0 {
//		return nil, nil, nil, errs[0]
//	}
//	si = &SignInfo{
//		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
//		HeaderNonce:     res.Header.Get(HeaderNonce),
//		HeaderSignature: res.Header.Get(HeaderSignature),
//		HeaderSerial:    res.Header.Get(HeaderSerial),
//		SignBody:        string(bs),
//	}
//	if c.DebugSwitch == gopay.DebugOn {
//		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
//		xlog.Debugf("Wechat_Headers: %#v", res.Header)
//		xlog.Debugf("Wechat_SignInfo: %#v", si)
//	}
//	return res, si, bs, nil
//}

//func (c *Client) doPayPalPost(bm gopay.BodyMap, path, authorization string) (res *http.Response, bs []byte, err error) {
//	var url = v3BaseUrlCh + path
//	httpClient := xhttp.NewClient()
//	if c.DebugSwitch == gopay.DebugOn {
//		jb, _ := json.Marshal(bm)
//		xlog.Debugf("Wechat_V3_RequestBody: %s", jb)
//		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
//	}
//	httpClient.Header.Add(HeaderAuthorization, authorization)
//	httpClient.Header.Add(HeaderSerial, c.wxSerialNo)
//	httpClient.Header.Add("Accept", "*/*")
//	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes()
//	if len(errs) > 0 {
//		return nil, nil, nil, errs[0]
//	}
//	si = &SignInfo{
//		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
//		HeaderNonce:     res.Header.Get(HeaderNonce),
//		HeaderSignature: res.Header.Get(HeaderSignature),
//		HeaderSerial:    res.Header.Get(HeaderSerial),
//		SignBody:        string(bs),
//	}
//	if c.DebugSwitch == gopay.DebugOn {
//		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
//		xlog.Debugf("Wechat_Headers: %#v", res.Header)
//		xlog.Debugf("Wechat_SignInfo: %#v", si)
//	}
//	return res, si, bs, nil
//}

//func (c *Client) doPayPalPatch(bm gopay.BodyMap, path, authorization string) (res *http.Response, bs []byte, err error) {
//	var url = v3BaseUrlCh + path
//	httpClient := xhttp.NewClient()
//	if c.DebugSwitch == gopay.DebugOn {
//		jb, _ := json.Marshal(bm)
//		xlog.Debugf("Wechat_V3_RequestBody: %s", jb)
//		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
//	}
//	httpClient.Header.Add(HeaderAuthorization, authorization)
//	httpClient.Header.Add(HeaderSerial, c.wxSerialNo)
//	httpClient.Header.Add("Accept", "*/*")
//	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Patch(url).SendBodyMap(bm).EndBytes()
//	if len(errs) > 0 {
//		return nil, nil, nil, errs[0]
//	}
//	si = &SignInfo{
//		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
//		HeaderNonce:     res.Header.Get(HeaderNonce),
//		HeaderSignature: res.Header.Get(HeaderSignature),
//		HeaderSerial:    res.Header.Get(HeaderSerial),
//		SignBody:        string(bs),
//	}
//	if c.DebugSwitch == gopay.DebugOn {
//		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
//		xlog.Debugf("Wechat_Headers: %#v", res.Header)
//		xlog.Debugf("Wechat_SignInfo: %#v", si)
//	}
//	return res, si, bs, nil
//}
