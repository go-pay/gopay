package alipay

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/util"
)

var defaultRequestIdFunc = &requestIdFunc{}

type requestIdFunc struct{}

func (d *requestIdFunc) RequestId() string {
	return fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix())
}

func (a *ClientV3) doPost(ctx context.Context, bm gopay.BodyMap, uri, authorization string) (res *http.Response, bs []byte, err error) {
	var url = v3BaseUrlCh + uri
	if !a.IsProd {
		url = v3SandboxBaseUrl + uri
	}
	req := a.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, a.requestIdFunc.RequestId())
	req.Header.Add(HeaderSdkVersion, "gopay/"+gopay.Version)
	if a.AppAuthToken != "" {
		req.Header.Add(HeaderAppAuthToken, a.AppAuthToken)
	}
	req.Header.Add("Accept", "application/json")
	if a.DebugSwitch == gopay.DebugOn {
		a.logger.Debugf("Alipay_V3_Url: %s", url)
		a.logger.Debugf("Alipay_V3_Req_Body: %s", bm.JsonBody())
		a.logger.Debugf("Alipay_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Post(url).SendBodyMap(bm).EndBytesForAlipayV3(ctx)
	if err != nil {
		return nil, nil, err
	}

	if a.DebugSwitch == gopay.DebugOn {
		a.logger.Debugf("Alipay_V3_Response: %d > %s", res.StatusCode, string(bs))
		a.logger.Debugf("Alipay_V3_Rsp_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (a *ClientV3) doProdPostFile(ctx context.Context, bm gopay.BodyMap, uri, authorization string) (res *http.Response, bs []byte, err error) {
	var url = v3BaseUrlCh + uri
	req := a.hc.Req(xhttp.TypeMultipartFormData)
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, a.requestIdFunc.RequestId())
	req.Header.Add(HeaderSdkVersion, "gopay/"+gopay.Version)
	req.Header.Add("Accept", "application/json")
	if a.DebugSwitch == gopay.DebugOn {
		a.logger.Debugf("Alipay_V3_Url: %s", url)
		a.logger.Debugf("Alipay_V3_Req_Body: %s", bm.JsonBody())
		a.logger.Debugf("Alipay_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Post(url).SendMultipartBodyMap(bm).EndBytesForAlipayV3(ctx)
	if err != nil {
		return nil, nil, err
	}
	if a.DebugSwitch == gopay.DebugOn {
		a.logger.Debugf("Alipay_V3_Response: %d > %s", res.StatusCode, string(bs))
		a.logger.Debugf("Alipay_V3_Rsp_Headers: %#v", res.Header)
	}
	return res, bs, nil
}
