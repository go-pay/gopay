package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// ant.merchant.expand.shop.modify(修改蚂蚁店铺)
// 文档地址：https://opendocs.alipay.com/apis/014tmb
func (a *Client) AntMerchantShopModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopModifyRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopModifyRsp)
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

// ant.merchant.expand.shop.create(蚂蚁店铺创建)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.create
func (a *Client) AntMerchantShopCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopCreateRsp, err error) {
	err = bm.CheckEmptyError("business_address", "shop_category", "store_id", "shop_type", "ip_role_id", "shop_name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.create"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopCreateRsp)
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

// ant.merchant.expand.shop.consult(蚂蚁店铺创建咨询)
// 文档地址：https://opendocs.alipay.com/apis/014yig
func (a *Client) AntMerchantShopConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopConsultRsp, err error) {
	err = bm.CheckEmptyError("business_address", "shop_category", "store_id", "shop_type", "ip_role_id", "shop_name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopConsultRsp)
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

// ant.merchant.expand.order.query(商户申请单查询)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.order.query
func (a *Client) AntMerchantOrderQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantOrderQueryRsp, err error) {
	err = bm.CheckEmptyError("order_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.order.query"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantOrderQueryRsp)
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

// ant.merchant.expand.shop.query(店铺查询接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.query
func (a *Client) AntMerchantShopQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.query"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopQueryRsp)
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

// ant.merchant.expand.shop.page.query(店铺分页查询接口)
// 文档地址：https://opendocs.alipay.com/open/04fgwq
func (a *Client) AntMerchantShopPageQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopPageQueryRsp, err error) {
	err = bm.CheckEmptyError("ip_role_id", "page_num", "page_size")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.page.query"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopPageQueryRsp)
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

// ant.merchant.expand.shop.close(蚂蚁店铺关闭)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.close
func (a *Client) AntMerchantShopClose(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopCloseRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.close"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopCloseRsp)
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

// ant.merchant.expand.indirect.image.upload(图片上传)
// bm参数中 image_content 可不传，file为必传参数
// 文档地址：https://opendocs.alipay.com/open/04fgwt
func (a *Client) AntMerchantExpandIndirectImageUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantExpandIndirectImageUploadRsp, err error) {
	err = bm.CheckEmptyError("image_type", "image_content")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.FileUploadRequest(ctx, bm, "ant.merchant.expand.indirect.image.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantExpandIndirectImageUploadRsp)
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

// ant.merchant.expand.mcc.query(商户mcc信息查询)
// 文档地址：https://opendocs.alipay.com/open/04fgwu
func (a *Client) AntMerchantExpandMccQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantExpandMccQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.mcc.query"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantExpandMccQueryRsp)
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

// ant.merchant.expand.shop.receiptaccount.save(店铺增加收单账号)
// 文档地址：https://opendocs.alipay.com/open/54b69b89_ant.merchant.expand.shop.receiptaccount.save
func (a *Client) AntMerchantExpandShopReceiptAccountSave(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantExpandShopReceiptAccountSaveRsp, err error) {
	err = bm.CheckEmptyError("shop_id", "receipt_account_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.receiptaccount.save"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantExpandShopReceiptAccountSaveRsp)
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
