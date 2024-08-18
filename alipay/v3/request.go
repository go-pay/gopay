package alipay

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
)

func (c *ClientV3) doPost(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	req := c.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Post(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}
