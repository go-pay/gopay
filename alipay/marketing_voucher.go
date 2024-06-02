package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.marketing.material.image.upload(营销图片资源上传接口)
// bm参数中 file_content 可不传，file为必传参数
// 文档地址：https://opendocs.alipay.com/open/389b24b6_alipay.marketing.material.image.upload
func (a *Client) MarketingMaterialImageUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingMaterialImageUploadRsp, err error) {
	if err = bm.CheckEmptyError("file_key", "file_content"); err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.FileUploadRequest(ctx, bm, "alipay.marketing.material.image.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingMaterialImageUploadRsp)
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

// alipay.marketing.activity.voucher.create(创建支付券)
// 文档地址：https://opendocs.alipay.com/open/049d65
func (a *Client) MarketingActivityVoucherCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityVoucherCreateRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "activity_base_info", "voucher_send_mode_info", "voucher_deduct_info", "merchant_access_mode",
		"voucher_available_scope_info", "voucher_use_rule_info", "voucher_customer_guide_info", "voucher_display_pattern_info", "voucher_budget_supply_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.voucher.create"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityVoucherCreateRsp)
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

// alipay.marketing.activity.voucher.publish(激活支付券)
// 文档地址：https://opendocs.alipay.com/open/049d66
func (a *Client) MarketingActivityVoucherPublish(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityVoucherPublishRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "out_biz_no", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.voucher.publish"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityVoucherPublishRsp)
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

// alipay.marketing.activity.voucher.query(查询支付券详情)
// 文档地址：https://opendocs.alipay.com/open/049d6g
func (a *Client) MarketingActivityVoucherQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityVoucherQueryRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.voucher.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityVoucherQueryRsp)
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

// alipay.marketing.activity.voucher.modify(修改支付券基本信息)
// 文档地址：https://opendocs.alipay.com/open/049d67
func (a *Client) MarketingActivityVoucherModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityVoucherModifyRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "merchant_access_mode", "activity_base_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.voucher.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityVoucherModifyRsp)
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

// alipay.marketing.activity.voucher.append(追加支付券预算)
// 文档地址：https://opendocs.alipay.com/open/049d68
func (a *Client) MarketingActivityVoucherAppend(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityVoucherAppendRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "voucher_quantity", "out_biz_no", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.voucher.append"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityVoucherAppendRsp)
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

// alipay.marketing.activity.voucher.stop(停止支付券)
// 文档地址：https://opendocs.alipay.com/open/049d69
func (a *Client) MarketingActivityVoucherStop(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityVoucherStopRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "out_biz_no", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.voucher.stop"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityVoucherStopRsp)
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
