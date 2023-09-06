package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 获取对私银行卡号开户银行
// 注意：accountNo 需此方法加密：client.V3EncryptText()
// Code = 0 is success
func (c *ClientV3) V3BankSearchBank(ctx context.Context, accountNo string) (wxRsp *BankSearchBankRsp, err error) {
	uri := v3BankSearchBank + "?account_number=" + url.QueryEscape(accountNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchBankRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchBank)
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

// 查询支持个人业务的银行列表
// Code = 0 is success
func (c *ClientV3) V3BankSearchPersonalList(ctx context.Context, limit, offset int) (wxRsp *BankSearchPersonalListRsp, err error) {
	if limit == 0 {
		limit = 20
	}
	uri := v3BankSearchPersonalList + "?limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchPersonalListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchList)
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

// 查询支持对公业务的银行列表
// Code = 0 is success
func (c *ClientV3) V3BankSearchCorporateList(ctx context.Context, limit, offset int) (wxRsp *BankSearchCorporateListRsp, err error) {
	if limit == 0 {
		limit = 20
	}
	uri := v3BankSearchCorporateList + "?limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchCorporateListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchList)
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

// 查询省份列表
// Code = 0 is success
func (c *ClientV3) V3BankSearchProvinceList(ctx context.Context) (wxRsp *BankSearchProvinceListRsp, err error) {
	authorization, err := c.authorization(MethodGet, v3BankSearchProvinceList, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, v3BankSearchProvinceList, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchProvinceListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchProvince)
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

// 查询城市列表
// Code = 0 is success
func (c *ClientV3) V3BankSearchCityList(ctx context.Context, provinceCode int) (wxRsp *BankSearchCityListRsp, err error) {
	url := fmt.Sprintf(v3BankSearchCityList, provinceCode)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchCityListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchCity)
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

// 查询支行列表
// Code = 0 is success
func (c *ClientV3) V3BankSearchBranchList(ctx context.Context, bankAliasCode string, cityCode, limit, offset int) (wxRsp *BankSearchBranchListRsp, err error) {
	if limit == 0 {
		limit = 20
	}
	uri := fmt.Sprintf(v3BankSearchBranchList, bankAliasCode) + "?city_code=" + util.Int2String(cityCode) + "&limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchBranchListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchBranch)
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
