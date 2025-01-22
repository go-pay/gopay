package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// 发起转账
func (c *ClientV3) V3TransferBills(ctx context.Context, bm gopay.BodyMap) (*TransferBillsRsp, error) {
	authorization, err := c.authorization(MethodPost, V3TransferBills, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, V3TransferBills, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferBillsRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferBills)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 撤销转账
func (c *ClientV3) V3TransferBillsCancel(ctx context.Context, outBillNo string) (*TransferBillsCancelRsp, error) {
	uri := fmt.Sprintf(V3TransferBillsCancel, outBillNo)
	authorization, err := c.authorization(MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, nil, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferBillsCancelRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferBillsCancel)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户单号查询转账单
func (c *ClientV3) V3TransferBillsMerchantQuery(ctx context.Context, outBillNo string) (*TransferBillsMerchantQueryRsp, error) {
	uri := fmt.Sprintf(V3TransferBillsMerchantQuery, outBillNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferBillsMerchantQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferBillsMerchantQuery)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 微信单号查询转账单
func (c *ClientV3) V3TransferBillsQuery(ctx context.Context, transferBillNo string) (*TransferBillsQueryRsp, error) {
	uri := fmt.Sprintf(V3TransferBillsQuery, transferBillNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferBillsQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferBillsQuery)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户单号申请电子回单
func (c *ClientV3) V3TransferElecsignMerchant(ctx context.Context, bm gopay.BodyMap) (*TransferElecsignMerchantRsp, error) {
	authorization, err := c.authorization(MethodPost, V3TransferElecsignMerchant, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, V3TransferElecsignMerchant, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferElecsignMerchantRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferElecsignMerchant)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 微信单号申请电子回单
func (c *ClientV3) V3TransferElecsign(ctx context.Context, bm gopay.BodyMap) (*TransferElecsignRsp, error) {
	authorization, err := c.authorization(MethodPost, V3TransferElecsign, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, V3TransferElecsign, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferElecsignRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferElecsign)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户单号查询电子回单
func (c *ClientV3) V3TransferElecsignQuery(ctx context.Context, transferBillNo string) (*TransferElecsignQueryRsp, error) {
	uri := fmt.Sprintf(V3TransferElecsignQuery, transferBillNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferElecsignQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferElecsignQuery)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 微信单号查询电子回单
func (c *ClientV3) V3TransferElecsignMerchantQuery(ctx context.Context, transferBillNo string) (*TransferElecsignMerchantQueryRsp, error) {
	uri := fmt.Sprintf(V3TransferElecsignMerchantQuery, transferBillNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &TransferElecsignMerchantQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferElecsignMerchantQuery)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
