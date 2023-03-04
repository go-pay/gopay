package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// alipay.fund.trans.uni.transfer(单笔转账接口)
// 文档地址：https://opendocs.alipay.com/open/02byuo
func (a *Client) FundTransUniTransfer(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransUniTransferResponse, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "payee_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.trans.uni.transfer"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransUniTransferResponse)
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

// alipay.fund.account.query(支付宝资金账户资产查询接口)
// 文档地址：https://opendocs.alipay.com/open/02byuq
func (a *Client) FundAccountQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAccountQueryResponse, err error) {
	err = bm.CheckEmptyError("alipay_user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.account.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAccountQueryResponse)
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

// alipay.fund.trans.common.query(转账业务单据查询接口)
// 文档地址：https://opendocs.alipay.com/open/02byup
func (a *Client) FundTransCommonQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransCommonQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.trans.common.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransCommonQueryResponse)
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

// alipay.fund.trans.order.query(查询转账订单接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.order.query
func (a *Client) FundTransOrderQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransOrderQueryResponse, err error) {
	// 两个请求参数不能同时为空
	err1 := bm.CheckEmptyError("out_biz_no")
	err2 := bm.CheckEmptyError("order_id")
	if err1 != nil && err2 != nil {
		return nil, fmt.Errorf("out_biz_no,order_id : Both cannot be empty at some time")
	}

	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.trans.order.query"); err != nil {
		return nil, err
	}

	aliRsp = new(FundTransOrderQueryResponse)
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

// alipay.fund.trans.refund(资金退回接口)
// 文档地址: https://opendocs.alipay.com/open/02byvd
func (a *Client) FundTransRefund(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransRefundResponse, err error) {
	err = bm.CheckEmptyError("order_id", "out_request_no", "refund_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.trans.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransRefundResponse)
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

// alipay.fund.auth.order.freeze(资金授权冻结接口)
// 文档地址: https://opendocs.alipay.com/open/02fkb9
func (a *Client) FundAuthOrderFreeze(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOrderFreezeResponse, err error) {
	err = bm.CheckEmptyError("auth_code", "auth_code_type", "out_order_no", "out_request_no", "order_title", "amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.auth.order.freeze"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOrderFreezeResponse)
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

// alipay.fund.auth.order.voucher.create(资金授权发码接口)
// 文档地址: https://opendocs.alipay.com/open/02fit5
func (a *Client) FundAuthOrderVoucherCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOrderVoucherCreateResponse, err error) {
	err = bm.CheckEmptyError("out_order_no", "out_request_no", "order_title", "amount", "product_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.auth.order.voucher.create"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOrderVoucherCreateResponse)
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

// alipay.fund.auth.order.app.freeze(线上资金授权冻结接口)
// 文档地址: https://opendocs.alipay.com/open/02f912
func (a *Client) FundAuthOrderAppFreeze(ctx context.Context, bm gopay.BodyMap) (payParam string, err error) {
	err = bm.CheckEmptyError("out_order_no", "out_request_no", "order_title", "amount", "product_code")
	if err != nil {
		return util.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.auth.order.app.freeze"); err != nil {
		return "", err
	}
	payParam = string(bs)
	return payParam, nil
}

// alipay.fund.auth.order.unfreeze(资金授权解冻接口)
// 文档地址: https://opendocs.alipay.com/open/02fkbc
func (a *Client) FundAuthOrderUnfreeze(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOrderUnfreezeResponse, err error) {
	err = bm.CheckEmptyError("auth_no", "out_request_no", "amount", "remark")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.auth.order.unfreeze"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOrderUnfreezeResponse)
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

// alipay.fund.auth.operation.detail.query(资金授权操作查询接口)
// 文档地址: https://opendocs.alipay.com/open/02fkbd
func (a *Client) FundAuthOperationDetailQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOperationDetailQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.auth.operation.detail.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOperationDetailQueryResponse)
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

// alipay.fund.auth.operation.cancel(资金授权撤销接口)
// 文档地址: https://opendocs.alipay.com/open/02fkbb
func (a *Client) FundAuthOperationCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundAuthOperationCancelResponse, err error) {
	err = bm.CheckEmptyError("remark")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.auth.operation.cancel"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOperationCancelResponse)
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

// alipay.fund.batch.create(批次下单接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.batch.create
func (a *Client) FundBatchCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundBatchCreateResponse, err error) {
	err = bm.CheckEmptyError("out_batch_no", "product_code", "biz_scene", "order_title", "total_trans_amount", "total_count", "trans_order_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.batch.create"); err != nil {
		return nil, err
	}
	aliRsp = new(FundBatchCreateResponse)
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

// alipay.fund.batch.close(批量转账关单接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.batch.close
func (a *Client) FundBatchClose(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundBatchCloseResponse, err error) {
	err = bm.CheckEmptyError("biz_scene")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.batch.close"); err != nil {
		return nil, err
	}
	aliRsp = new(FundBatchCloseResponse)
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

// alipay.fund.batch.detail.query(批量转账明细查询接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.batch.detail.query
func (a *Client) FundBatchDetailQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundBatchDetailQueryResponse, err error) {
	err = bm.CheckEmptyError("biz_scene")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.batch.detail.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundBatchDetailQueryResponse)
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

// alipay.fund.trans.app.pay(现金红包无线支付接口)
// 文档地址: https://opendocs.alipay.com/open/03rbyf
func (a *Client) FundTransAppPay(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransAppPayResponse, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "biz_scene")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.trans.app.pay"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransAppPayResponse)
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

// alipay.fund.trans.payee.bind.query(资金收款账号绑定关系查询)
// 文档地址: https://opendocs.alipay.com/apis/020tl1
func (a *Client) FundTransPayeeBindQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransPayeeBindQueryRsp, err error) {
	err = bm.CheckEmptyError("identity", "identity_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.trans.payee.bind.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransPayeeBindQueryRsp)
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

// alipay.fund.trans.page.pay(资金转账页面支付接口)
// 文档地址: https://opendocs.alipay.com/open/03rbye
func (a *Client) FundTransPagePay(ctx context.Context, bm gopay.BodyMap) (aliRsp *FundTransPagePayRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "biz_scene")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.fund.trans.page.pay"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransPagePayRsp)
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
