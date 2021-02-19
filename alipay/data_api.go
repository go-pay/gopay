package alipay

import (
	"encoding/json"
	"fmt"

	"github.com/iGoogle-ink/gopay"
)

// Deprecated
// alipay.data.bill.balance.query(支付宝商家账户当前余额查询)
//	文档地址：https://opendocs.alipay.com/apis/api_15/alipay.data.bill.balance.query
func (a *Client) DataBillBalanceQuery(bm gopay.BodyMap) (aliRsp *DataBillBalanceQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.data.bill.balance.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillBalanceQueryResponse)
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

// alipay.data.dataservice.bill.downloadurl.query(查询对账单下载地址)
//	文档地址：https://opendocs.alipay.com/apis/api_15/alipay.data.dataservice.bill.downloadurl.query
func (a *Client) DataBillDownloadUrlQuery(bm gopay.BodyMap) (aliRsp *DataBillDownloadUrlQueryResponse, err error) {
	err = bm.CheckEmptyError("bill_type", "bill_date")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.data.dataservice.bill.downloadurl.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillDownloadUrlQueryResponse)
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
