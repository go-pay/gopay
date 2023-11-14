package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
)

// PostAliPayAPISelfV2 支付宝接口自行实现方法
// 注意：biz_content 需要自行通过bm.SetBodyMap()设置，不设置则没有此参数
// 示例：请参考 client_test.go 的 TestClient_PostAliPayAPISelfV2() 方法
func (a *Client) PostAliPayAPISelfV2(ctx context.Context, bm gopay.BodyMap, method string, aliRsp any) (err error) {
	var (
		bs, bodyBs []byte
	)
	// check if there is biz_content
	bz := bm.GetInterface("biz_content")
	if bzBody, ok := bz.(gopay.BodyMap); ok {
		if bodyBs, err = json.Marshal(bzBody); err != nil {
			return fmt.Errorf("json.Marshal(%v)：%w", bzBody, err)
		}
		bm.Set("biz_content", string(bodyBs))
	}

	if bs, err = a.doAliPaySelf(ctx, bm, method); err != nil {
		return err
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return err
	}
	return nil
}

// PostFileAliPayAPISelfV2 用于支付宝带有文件上传的接口自行实现方法
// 注意：最新版本的支付宝接口，对于文件的上传已统一改为通过formData上传
// 请求form格式如下： {file: "fileData", "data": BodyMap{"key": "value"}}
// 其中file为file请求字段名称，data为其他请求参数（key为文件名，value为文件内容）
func (a *Client) PostFileAliPayAPISelfV2(ctx context.Context, bm gopay.BodyMap, method string, aliRsp any) (err error) {
	var (
		url  string
		sign string
	)
	fm := make(gopay.BodyMap)
	for k, v := range bm {
		if _, ok := v.(*util.File); ok {
			fm.Set(k, v)
			bm.Remove(k)
			continue
		}
	}

	bm.Set("method", method)
	// check public parameter
	a.checkPublicParam(bm)
	// check sign, 需要先移除文件字段
	if bm.GetString("sign") == "" {
		sign, err = a.getRsaSign(bm, bm.GetString("sign_type"))
		if err != nil {
			return fmt.Errorf("GetRsaSign Error: %w", err)
		}
		bm.Set("sign", sign)
	}
	// 增加文件字段
	for k, v := range fm {
		bm.Set(k, v)
	}
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Request: %s", bm.JsonBody())
	}
	if a.IsProd {
		url = baseUrlUtf8
	} else {
		url = sandboxBaseUrlUtf8
	}
	res, bs, err := a.hc.Req(xhttp.TypeMultipartFormData).Post(url).
		SendMultipartBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil
	}
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return err
	}
	return nil
}

// 向支付宝发送自定义请求
func (a *Client) doAliPaySelf(ctx context.Context, bm gopay.BodyMap, method string) (bs []byte, err error) {
	var (
		url, sign string
	)
	bm.Set("method", method)
	// check public parameter
	a.checkPublicParam(bm)
	// check sign
	if bm.GetString("sign") == "" {
		sign, err = a.getRsaSign(bm, bm.GetString("sign_type"))
		if err != nil {
			return nil, fmt.Errorf("GetRsaSign Error: %w", err)
		}
		bm.Set("sign", sign)
	}
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Request: %s", bm.JsonBody())
	}
	if a.IsProd {
		url = baseUrlUtf8
	} else {
		url = sandboxBaseUrlUtf8
	}
	res, bs, err := a.hc.Req(xhttp.TypeFormData).Post(url).SendString(bm.EncodeURLParams()).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// 向支付宝发送请求
func (a *Client) doAliPay(ctx context.Context, bm gopay.BodyMap, method string, authToken ...string) (bs []byte, err error) {
	var (
		bizContent, url string
		bodyBs          []byte
	)
	if bm != nil {
		_, has := appAuthTokenInBizContent[method]
		if has {
			if bodyBs, err = json.Marshal(bm); err != nil {
				return nil, fmt.Errorf("json.Marshal：%w", err)
			}
			bizContent = string(bodyBs)
			bm.Remove("app_auth_token")
		} else {
			aat := bm.GetString("app_auth_token")
			bm.Remove("app_auth_token")
			if bodyBs, err = json.Marshal(bm); err != nil {
				return nil, fmt.Errorf("json.Marshal：%w", err)
			}
			bizContent = string(bodyBs)
			bm.Set("app_auth_token", aat)
		}
	}
	// 处理公共参数
	param, err := a.pubParamsHandle(bm, method, bizContent, authToken...)
	if err != nil {
		return nil, err
	}
	switch method {
	case "alipay.trade.app.pay", "alipay.fund.auth.order.app.freeze":
		return []byte(param), nil
	case "alipay.trade.wap.pay", "alipay.trade.page.pay", "alipay.user.certify.open.certify":
		if !a.IsProd {
			return []byte(sandboxBaseUrl + "?" + param), nil
		}
		return []byte(baseUrl + "?" + param), nil
	default:
		url = baseUrlUtf8
		if !a.IsProd {
			url = sandboxBaseUrlUtf8
		}
		res, bs, err := a.hc.Req(xhttp.TypeFormData).Post(url).SendString(param).EndBytes(ctx)
		if err != nil {
			return nil, err
		}
		if a.DebugSwitch == gopay.DebugOn {
			xlog.Debugf("Alipay_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
		}
		if res.StatusCode != 200 {
			return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
		}
		return bs, nil
	}
}

// 向支付宝发送请求
func (a *Client) DoAliPay(ctx context.Context, bm gopay.BodyMap, method string, authToken ...string) (bs []byte, err error) {
	var (
		bizContent, url string
		bodyBs          []byte
	)
	if bm != nil {
		_, has := appAuthTokenInBizContent[method]
		if has {
			if bodyBs, err = json.Marshal(bm); err != nil {
				return nil, fmt.Errorf("json.Marshal：%w", err)
			}
			bizContent = string(bodyBs)
			bm.Remove("app_auth_token")
		} else {
			aat := bm.GetString("app_auth_token")
			bm.Remove("app_auth_token")
			if bodyBs, err = json.Marshal(bm); err != nil {
				return nil, fmt.Errorf("json.Marshal：%w", err)
			}
			bizContent = string(bodyBs)
			bm.Set("app_auth_token", aat)
		}
	}
	// 处理公共参数
	param, err := a.pubParamsHandle(bm, method, bizContent, authToken...)
	if err != nil {
		return nil, err
	}
	switch method {
	case "alipay.trade.app.pay", "alipay.fund.auth.order.app.freeze":
		return []byte(param), nil
	case "alipay.trade.wap.pay", "alipay.trade.page.pay", "alipay.user.certify.open.certify":
		if !a.IsProd {
			return []byte(sandboxBaseUrl + "?" + param), nil
		}
		return []byte(baseUrl + "?" + param), nil
	default:
		url = baseUrlUtf8
		if !a.IsProd {
			url = sandboxBaseUrlUtf8
		}
		res, bs, err := a.hc.Req(xhttp.TypeFormData).Post(url).SendString(param).EndBytes(ctx)
		if err != nil {
			return nil, err
		}
		if a.DebugSwitch == gopay.DebugOn {
			xlog.Debugf("Alipay_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
		}
		if res.StatusCode != 200 {
			return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
		}
		return bs, nil
	}
}

// 保持和官方 SDK 命名方式一致
func (a *Client) PageExecute(ctx context.Context, bm gopay.BodyMap, method string, authToken ...string) (url string, err error) {
	var (
		bizContent string
		bodyBs     []byte
	)
	if bm != nil {
		_, has := appAuthTokenInBizContent[method]
		if has {
			if bodyBs, err = json.Marshal(bm); err != nil {
				return "", fmt.Errorf("json.Marshal：%w", err)
			}
			bizContent = string(bodyBs)
			bm.Remove("app_auth_token")
		} else {
			aat := bm.GetString("app_auth_token")
			bm.Remove("app_auth_token")
			if bodyBs, err = json.Marshal(bm); err != nil {
				return "", fmt.Errorf("json.Marshal：%w", err)
			}
			bizContent = string(bodyBs)
			bm.Set("app_auth_token", aat)
		}
	}
	// 处理公共参数
	param, err := a.pubParamsHandle(bm, method, bizContent, authToken...)
	if err != nil {
		return "", err
	}

	if !a.IsProd {
		return sandboxBaseUrl + "?" + param, nil
	}
	return baseUrl + "?" + param, nil
}

// 文件上传
func (a *Client) FileRequest(ctx context.Context, bm gopay.BodyMap, file *util.File, method string) (bs []byte, err error) {
	var (
		bodyStr string
		bodyBs  []byte
		aat     string
	)
	if bm != nil {
		aat = bm.GetString("app_auth_token")
		bm.Remove("app_auth_token")
		if bodyBs, err = json.Marshal(bm); err != nil {
			return nil, fmt.Errorf("json.Marshal：%w", err)
		}
		bodyStr = string(bodyBs)
	}
	pubBody := make(gopay.BodyMap)
	pubBody.Set("app_id", a.AppId).
		Set("method", method).
		Set("format", "JSON").
		Set("charset", a.Charset).
		Set("sign_type", a.SignType).
		Set("version", "1.0").
		Set("scene", "SYNC_ORDER").
		Set("timestamp", time.Now().Format(util.TimeLayout))

	if a.AppCertSN != util.NULL {
		pubBody.Set("app_cert_sn", a.AppCertSN)
	}
	if a.AliPayRootCertSN != util.NULL {
		pubBody.Set("alipay_root_cert_sn", a.AliPayRootCertSN)
	}
	if a.ReturnUrl != util.NULL {
		pubBody.Set("return_url", a.ReturnUrl)
	}
	if a.location != nil {
		pubBody.Set("timestamp", time.Now().In(a.location).Format(util.TimeLayout))
	}
	if a.NotifyUrl != util.NULL { //如果返回url为空，传过来的返回url不为空
		//fmt.Println("url不为空？", a.NotifyUrl)
		pubBody.Set("notify_url", a.NotifyUrl)
	}
	//fmt.Println("notify,", pubBody.JsonBody())
	if a.AppAuthToken != util.NULL {
		pubBody.Set("app_auth_token", a.AppAuthToken)
	}
	if aat != util.NULL {
		pubBody.Set("app_auth_token", aat)
	}
	if bodyStr != util.NULL {
		pubBody.Set("biz_content", bodyStr)
	}
	sign, err := a.getRsaSign(pubBody, pubBody.GetString("sign_type"))
	if err != nil {
		return nil, fmt.Errorf("GetRsaSign Error: %w", err)
	}
	//pubBody.Set("file_content", file.Content)
	pubBody.Set("sign", sign)
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Request: %s", pubBody.JsonBody())
	}
	param := pubBody.EncodeURLParams()
	url := baseUrlUtf8 + "&" + param
	bm.Reset()
	bm.SetFormFile("file_content", file)
	res, bs, err := a.hc.Req(xhttp.TypeMultipartFormData).Post(url).
		SendMultipartBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}
