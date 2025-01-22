package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
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

	wxRsp := &TransferBillsRsp{Code: Success, SignInfo: si, Response: &TransferBills{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
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

	wxRsp := &TransferBillsCancelRsp{Code: Success, SignInfo: si, Response: &TransferBillsCancel{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
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

	wxRsp := &TransferBillsMerchantQueryRsp{Code: Success, SignInfo: si, Response: &TransferBillsMerchantQuery{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
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
	wxRsp := &TransferBillsQueryRsp{Code: Success, SignInfo: si, Response: &TransferBillsQuery{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
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
	wxRsp := &TransferElecsignMerchantRsp{Code: Success, SignInfo: si, Response: &TransferElecsignMerchant{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
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
	wxRsp := &TransferElecsignRsp{Code: Success, SignInfo: si, Response: &TransferElecsign{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 微信单号查询电子回单
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
	wxRsp := &TransferElecsignQueryRsp{Code: Success, SignInfo: si, Response: &TransferElecsignQuery{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户单号查询电子回单
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
	wxRsp := &TransferElecsignMerchantQueryRsp{Code: Success, SignInfo: si, Response: &TransferElecsignMerchantQuery{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}
