package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// datadigital.fincloud.generalsaas.ocr.server.detect(服务端OCR)
// 文档地址：https://opendocs.alipay.com/open/05ut8h
func (a *Client) OcrServerDetect(ctx context.Context, bm gopay.BodyMap) (aliRsp *OcrServerDetectRsp, err error) {
	err = bm.CheckEmptyError("ocr_type", "outer_order_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.ocr.server.detect"); err != nil {
		return nil, err
	}
	aliRsp = new(OcrServerDetectRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// datadigital.fincloud.generalsaas.ocr.mobile.initialize(App端OCR初始化)
// 文档地址：https://opendocs.alipay.com/open/043ksf
func (a *Client) OcrMobileInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *OcrMobileInitializeRsp, err error) {
	err = bm.CheckEmptyError("outer_order_no", "biz_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.ocr.mobile.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(OcrMobileInitializeRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// datadigital.fincloud.generalsaas.ocr.common.detect(文字识别OCR)
// 文档地址：https://opendocs.alipay.com/open/0776c2cb_datadigital.fincloud.generalsaas.ocr.common.detect
func (a *Client) OcrCommonDetect(ctx context.Context, bm gopay.BodyMap) (aliRsp *OcrCommonDetectRsp, err error) {
	err = bm.CheckEmptyError("ocr_type", "outer_order_no", "file_content")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "datadigital.fincloud.generalsaas.ocr.common.detect"); err != nil {
		return nil, err
	}
	aliRsp = new(OcrCommonDetectRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
