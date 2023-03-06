package allinpay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

const (
	payPath    = "/unitorder/pay"
	scanQrPath = "/unitorder/scanqrpay"
	queryPath  = "/tranx/query"
	refundPath = "/tranx/refund"
	cancelPath = "/tranx/cancel"
)

type (
	PayResp struct {
		RespBase
		Trxid     string `json:"trxid"`
		ChnlTrxId string `json:"chnltrxid"`
		Reqsn     string `json:"reqsn"`
		RandomStr string `json:"randomstr"`
		TrxStatus string `json:"trxstatus"`
		FinTime   string `json:"fintime"`
		ErrMsg    string `json:"errmsg"`
		PayInfo   string `json:"payinfo"`
	}
	ScanPayResp struct {
		RespBase
		Trxid     string `json:"trxid"`
		ChnlTrxId string `json:"chnltrxid"`
		Reqsn     string `json:"reqsn"`
		TrxStatus string `json:"trxstatus"`
		Acct      string `json:"acct"`
		TrxCode   string `json:"trxcode"`
		FinTime   string `json:"fintime"`
		ErrMsg    string `json:"errmsg"`
		RandomStr string `json:"randomstr"`
		InitAmt   string `json:"initamt"`
		TrxAmt    string `json:"trxamt"`
		Fee       string `json:"fee"`
		Cmid      string `json:"cmid"`
		Chnlid    string `json:"chnlid"`
		ChnlData  string `json:"chnldata"`
		AcctType  string `json:"accttype"`
	}
	RefundResp struct {
		RespBase
		Trxid     string `json:"trxid"`
		Reqsn     string `json:"reqsn"`
		TrxStatus string `json:"trxstatus"`
		FinTime   string `json:"fintime"`
		ErrMsg    string `json:"errmsg"`
		Fee       string `json:"fee"`
		TrxCode   string `json:"trxCode"`
		RandomStr string `json:"randomstr"`
		ChnlTrxId string `json:"chnltrxid"`
	}
)

// Pay 统一支付接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=88
func (c *Client) Pay(ctx context.Context, bm gopay.BodyMap) (rsp *PayResp, err error) {
	err = bm.CheckEmptyError("reqsn", "paytype")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, payPath, bm); err != nil {
		return nil, err
	}
	rsp = new(PayResp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RespBase); err != nil {
		return nil, err
	}
	return rsp, c.verifySign(bs)
}

// ScanPay 统一扫码接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=89
func (c *Client) ScanPay(ctx context.Context, bm gopay.BodyMap) (rsp *ScanPayResp, err error) {
	err = bm.CheckEmptyError("reqsn", "authcode", "terminfo")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, scanQrPath, bm); err != nil {
		return nil, err
	}
	rsp = new(ScanPayResp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RespBase); err != nil {
		return nil, err
	}
	return rsp, c.verifySign(bs)
}

// Query 统一查询接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=836
func (c *Client) Query(ctx context.Context, orderType string, no string) (rsp *ScanPayResp, err error) {
	bm := gopay.BodyMap{}
	switch orderType {
	case OrderTypeReqSN:
		bm.Set("reqsn", no)
	case OrderTypeTrxId:
		bm.Set("trxid", no)
	}
	var bs []byte
	if bs, err = c.doPost(ctx, queryPath, bm); err != nil {
		return nil, err
	}
	rsp = new(ScanPayResp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RespBase); err != nil {
		return nil, err
	}
	return rsp, nil
} // Refund 统一退款接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=838
func (c *Client) Refund(ctx context.Context, bm gopay.BodyMap) (rsp *RefundResp, err error) {
	err = bm.CheckEmptyError("reqsn", "trxamt")
	if err != nil {
		return nil, err
	}
	if bm.GetString("oldreqsn") == util.NULL && bm.GetString("oldtrxid") == util.NULL {
		return nil, fmt.Errorf("[%w], %v", gopay.MissParamErr, "oldreqsn和oldtrxid必填其一")
	}
	var bs []byte
	if bs, err = c.doPost(ctx, refundPath, bm); err != nil {
		return nil, err
	}
	rsp = new(RefundResp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RespBase); err != nil {
		return nil, err
	}
	return rsp, nil
}

// Cancel 统一撤销接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=837
func (c *Client) Cancel(ctx context.Context, bm gopay.BodyMap) (rsp *RefundResp, err error) {
	err = bm.CheckEmptyError("reqsn", "trxamt")
	if err != nil {
		return nil, err
	}
	if bm.GetString("oldreqsn") == util.NULL && bm.GetString("oldtrxid") == util.NULL {
		return nil, fmt.Errorf("[%w], %v", gopay.MissParamErr, "oldreqsn和oldtrxid必填其一")
	}
	var bs []byte
	if bs, err = c.doPost(ctx, cancelPath, bm); err != nil {
		return nil, err
	}
	rsp = new(RefundResp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RespBase); err != nil {
		return nil, err
	}
	return rsp, nil
}
