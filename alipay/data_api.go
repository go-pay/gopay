package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.data.bill.balance.query(支付宝商家账户当前余额查询)
// 文档地址：https://opendocs.alipay.com/open/2acb3c34_alipay.data.bill.balance.query
func (a *Client) DataBillBalanceQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillBalanceQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.bill.balance.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillBalanceQueryResponse)
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

// alipay.data.bill.accountlog.query(支付宝商家账户账务明细查询)
// 文档地址：https://opendocs.alipay.com/open/dae3ac99_alipay.data.bill.accountlog.query
func (a *Client) DataBillAccountLogQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillAccountLogQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.bill.accountlog.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillAccountLogQueryResponse)
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

// alipay.data.bill.ereceipt.apply(申请电子回单(incubating))
// 文档地址：https://opendocs.alipay.com/open/1aad1956_alipay.data.bill.ereceipt.apply
func (a *Client) DataBillEreceiptApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillEreceiptApplyRsp, err error) {
	err = bm.CheckEmptyError("type", "key")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.bill.ereceipt.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillEreceiptApplyRsp)
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

// alipay.data.bill.ereceipt.query(查询电子回单状态(incubating))
// 文档地址：https://opendocs.alipay.com/open/30b94a2f_alipay.data.bill.ereceipt.query
func (a *Client) DataBillEreceiptQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillEreceiptQueryRsp, err error) {
	err = bm.CheckEmptyError("file_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.bill.ereceipt.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillEreceiptQueryRsp)
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

// alipay.data.dataservice.bill.downloadurl.query(查询对账单下载地址)
// 文档地址：https://opendocs.alipay.com/open/02e7gr
func (a *Client) DataBillDownloadUrlQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillDownloadUrlQueryResponse, err error) {
	err = bm.CheckEmptyError("bill_type", "bill_date")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.dataservice.bill.downloadurl.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillDownloadUrlQueryResponse)
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
