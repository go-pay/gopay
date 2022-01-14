package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// Deprecated
// 支付宝已不再支持
// alipay.data.bill.balance.query(支付宝商家账户当前余额查询)
//	文档地址：https://opendocs.alipay.com/apis/api_15/alipay.data.bill.balance.query
func (a *Client) DataBillBalanceQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillBalanceQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.bill.balance.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillBalanceQueryResponse)
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

// alipay.data.dataservice.bill.downloadurl.query(查询对账单下载地址)
//	文档地址：https://opendocs.alipay.com/apis/api_15/alipay.data.dataservice.bill.downloadurl.query
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
