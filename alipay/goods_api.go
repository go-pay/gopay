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

// alipay.merchant.item.file.upload(商品文件上传接口)
//	文档地址：https://opendocs.alipay.com/apis/api_4/alipay.merchant.item.file.upload
func (a *Client) MerchantItemFileUpload(ctx context.Context, file *util.File) (aliRsp *MerchantItemFileUploadRsp, err error) {
	// todo: finish
	bm := make(gopay.BodyMap)
	bm.Set("scene", "SYNC_ORDER") //素材固定值

	var bs []byte
	if bs, err = a.FileRequest(ctx, bm, file, "alipay.merchant.item.file.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(MerchantItemFileUploadRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

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
	//fmt.Println("bodyStr的信息是", bodyStr)
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
	sign, err := GetRsaSign(pubBody, pubBody.GetString("sign_type"), a.privateKey)
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
	httpClient := xhttp.NewClient()
	res, bs, err := httpClient.Type(xhttp.TypeMultipartFormData).Post(url).
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
