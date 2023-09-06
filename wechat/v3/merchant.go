package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 查询特约商户账户实时余额、查询二级商户账户实时余额
// Code = 0 is success
// 注意：服务商时，bm参数传 nil
func (c *ClientV3) V3EcommerceBalance(ctx context.Context, subMchid string, bm gopay.BodyMap) (*EcommerceBalanceRsp, error) {
	url := fmt.Sprintf(v3EcommerceBalance, subMchid) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &EcommerceBalanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceBalance)
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

// 查询二级商户账户日终余额
// date示例值：2019-08-17
// Code = 0 is success
func (c *ClientV3) V3EcommerceDayBalance(ctx context.Context, subMchid, date string) (*EcommerceBalanceRsp, error) {
	uri := fmt.Sprintf(v3EcommerceDayBalance, subMchid) + "?date=" + date
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &EcommerceBalanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceBalance)
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

// 查询账户实时余额
// Code = 0 is success
func (c *ClientV3) V3MerchantBalance(ctx context.Context, accountType string) (*MerchantBalanceRsp, error) {
	url := fmt.Sprintf(v3MerchantBalance, accountType)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &MerchantBalanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MerchantBalance)
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

// 查询账户日终余额
// date示例值：2019-08-17
// Code = 0 is success
func (c *ClientV3) V3MerchantDayBalance(ctx context.Context, accountType, date string) (*MerchantBalanceRsp, error) {
	uri := fmt.Sprintf(v3MerchantDayBalance, accountType) + "?date=" + date
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &MerchantBalanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MerchantBalance)
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

// 特约商户银行来账查询API
// Code = 0 is success
func (c *ClientV3) V3EcommerceIncomeRecord(ctx context.Context, bm gopay.BodyMap) (*PartnerIncomeRecordRsp, error) {
	uri := v3EcommerceIncomeRecord + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &PartnerIncomeRecordRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerIncomeRecord)
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

// 商户/服务商银行来账查询API
// Code = 0 is success
func (c *ClientV3) V3MerchantIncomeRecord(ctx context.Context, bm gopay.BodyMap) (*MerchantIncomeRecordRsp, error) {
	uri := v3MerchantIncomeRecord + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &MerchantIncomeRecordRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MerchantIncomeRecord)
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
