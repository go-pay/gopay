package alipay

import (
	"encoding/json"
	"fmt"

	"github.com/iGoogle-ink/gopay"
)

// alipay.fund.trans.uni.transfer(单笔转账接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.uni.transfer
func (a *Client) FundTransUniTransfer(bm gopay.BodyMap) (aliRsp *FundTransUniTransferResponse, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "payee_info")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.uni.transfer"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransUniTransferResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.account.query(支付宝资金账户资产查询接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.account.query
func (a *Client) FundAccountQuery(bm gopay.BodyMap) (aliRsp *FundAccountQueryResponse, err error) {
	err = bm.CheckEmptyError("alipay_user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.account.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAccountQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.trans.common.query(转账业务单据查询接口)
//	文档地址：https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.common.query
func (a *Client) FundTransCommonQuery(bm gopay.BodyMap) (aliRsp *FundTransCommonQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.common.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransCommonQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.trans.order.query(查询转账订单接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.order.query
func (a *Client) FundTransOrderQuery(bm gopay.BodyMap) (aliRsp *FundTransOrderQueryResponse, err error) {
	// 两个请求参数不能同时为空
	err1 := bm.CheckEmptyError("out_biz_no")
	err2 := bm.CheckEmptyError("order_id")
	if err1 != nil && err2 != nil {
		return nil, fmt.Errorf("out_biz_no,order_id : Both cannot be empty at some time")
	}

	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.order.query"); err != nil {
		return nil, err
	}

	aliRsp = new(FundTransOrderQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.trans.refund(资金退回接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.refund
func (a *Client) FundTransRefund(bm gopay.BodyMap) (aliRsp *FundTransRefundResponse, err error) {
	err = bm.CheckEmptyError("order_id", "out_request_no", "refund_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransRefundResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.auth.order.freeze(资金授权冻结接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.auth.order.freeze
func (a *Client) FundAuthOrderFreeze(bm gopay.BodyMap) (aliRsp *FundAuthOrderFreezeResponse, err error) {
	err = bm.CheckEmptyError("auth_code", "auth_code_type", "out_order_no", "out_request_no", "order_title", "amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.auth.order.freeze"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOrderFreezeResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.auth.order.voucher.create(资金授权发码接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.auth.order.voucher.create
func (a *Client) FundAuthOrderVoucherCreate(bm gopay.BodyMap) (aliRsp *FundAuthOrderVoucherCreateResponse, err error) {
	err = bm.CheckEmptyError("out_order_no", "out_request_no", "order_title", "amount", "product_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.auth.order.voucher.create"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOrderVoucherCreateResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.auth.order.app.freeze(线上资金授权冻结接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.auth.order.app.freeze
func (a *Client) FundAuthOrderAppFreeze(bm gopay.BodyMap) (aliRsp *FundAuthOrderAppFreezeResponse, err error) {
	err = bm.CheckEmptyError("out_order_no", "out_request_no", "order_title", "amount", "product_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.auth.order.app.freeze"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOrderAppFreezeResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.auth.order.unfreeze(资金授权解冻接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.auth.order.unfreeze
func (a *Client) FundAuthOrderUnfreeze(bm gopay.BodyMap) (aliRsp *FundAuthOrderUnfreezeResponse, err error) {
	err = bm.CheckEmptyError("auth_no", "out_request_no", "amount", "remark")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.auth.order.unfreeze"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOrderUnfreezeResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.auth.operation.detail.query(资金授权操作查询接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.auth.operation.detail.query
func (a *Client) FundAuthOperationDetailQuery(bm gopay.BodyMap) (aliRsp *FundAuthOperationDetailQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.auth.operation.detail.query"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOperationDetailQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.auth.operation.cancel(资金授权撤销接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.auth.operation.cancel
func (a *Client) FundAuthOperationCancel(bm gopay.BodyMap) (aliRsp *FundAuthOperationCancelResponse, err error) {
	err = bm.CheckEmptyError("remark")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.auth.operation.cancel"); err != nil {
		return nil, err
	}
	aliRsp = new(FundAuthOperationCancelResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

// alipay.fund.trans.app.pay(现金红包无线支付接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.app.pay
func (a *Client) FundTransAppPay(bm gopay.BodyMap) (aliRsp *FundTransAppPayResponse, err error) {
	err = bm.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "biz_scene")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.fund.trans.app.pay"); err != nil {
		return nil, err
	}
	aliRsp = new(FundTransAppPayResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}
