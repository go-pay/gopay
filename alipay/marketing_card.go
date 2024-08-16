package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.marketing.card.template.create(会员卡模板创建)
// 文档地址：https://opendocs.alipay.com/open/b2854ad3_alipay.marketing.card.template.create
func (a *Client) MarketingCardTemplateCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCardTemplateCreateRsp, err error) {
	err = bm.CheckEmptyError("request_id", "template_style_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.card.template.create"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCardTemplateCreateRsp)
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

// alipay.marketing.card.template.modify(会员卡模板修改)
// 文档地址：https://opendocs.alipay.com/open/e3227c82_alipay.marketing.card.template.modify
func (a *Client) MarketingCardTemplateModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCardTemplateModifyRsp, err error) {
	err = bm.CheckEmptyError("request_id", "template_id", "template_style_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.card.template.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCardTemplateModifyRsp)
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

// alipay.marketing.card.template.query(会员卡模板查询接口)
// 文档地址：https://opendocs.alipay.com/open/690f3d16_alipay.marketing.card.template.query
func (a *Client) MarketingCardTemplateQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCardTemplateQueryRsp, err error) {
	err = bm.CheckEmptyError("template_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.card.template.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCardTemplateQueryRsp)
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

// alipay.marketing.card.update(会员卡更新)
// 文档地址：https://opendocs.alipay.com/open/89b55b6d_alipay.marketing.card.update
func (a *Client) MarketingCardUpdate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCardUpdateRsp, err error) {
	err = bm.CheckEmptyError("target_card_no", "target_card_no_type", "occur_time", "card_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.card.update"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCardUpdateRsp)
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

// alipay.marketing.card.query(会员卡查询)
// 文档地址：https://opendocs.alipay.com/open/023c20c1_alipay.marketing.card.query
func (a *Client) MarketingCardQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCardQueryRsp, err error) {
	err = bm.CheckEmptyError("target_card_no", "target_card_no_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.card.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCardQueryRsp)
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

// alipay.marketing.card.delete(会员卡删卡)
// 文档地址：https://opendocs.alipay.com/open/8efddab3_alipay.marketing.card.delete
func (a *Client) MarketingCardDelete(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCardDeleteRsp, err error) {
	err = bm.CheckEmptyError("out_serial_no", "target_card_no", "target_card_no_type", "reason_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.card.delete"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCardDeleteRsp)
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

// alipay.marketing.card.message.notify(会员卡消息通知)
// 文档地址：https://opendocs.alipay.com/open/4c052993_alipay.marketing.card.message.notify
func (a *Client) MarketingCardMessageNotify(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCardMessageNotifyRsp, err error) {
	err = bm.CheckEmptyError("target_card_no", "target_card_no_type", "occur_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.card.message.notify"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCardMessageNotifyRsp)
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

// alipay.marketing.card.formtemplate.set(会员卡开卡表单模板配置)
// 文档地址：https://opendocs.alipay.com/open/78c84d3f_alipay.marketing.card.formtemplate.set
func (a *Client) MarketingCardFormTemplateSet(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCardFormTemplateSetRsp, err error) {
	err = bm.CheckEmptyError("template_id", "fields")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.card.formtemplate.set"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCardFormTemplateSetRsp)
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

// alipay.offline.material.image.upload(上传门店照片和视频接口)
// 文档地址：https://opendocs.alipay.com/open/0af852ff_alipay.offline.material.image.upload
func (a *Client) OfflineMaterialImageUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *OfflineMaterialImageUploadRsp, err error) {
	err = bm.CheckEmptyError("image_type", "image_name", "image_content")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.FileUploadRequest(ctx, bm, "alipay.marketing.material.image.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(OfflineMaterialImageUploadRsp)
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
