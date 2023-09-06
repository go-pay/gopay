package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 申请交易账单API
// 注意：如 bill_date 为空，默认查前一天的
// Code = 0 is success
func (c *ClientV3) V3BillTradeBill(ctx context.Context, bm gopay.BodyMap) (wxRsp *BillRsp, err error) {
	if bm != nil {
		if bm.GetString("bill_date") == util.NULL {
			now := time.Now()
			yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local).Format(util.DateLayout)
			bm.Set("bill_date", yesterday)
		}
	}
	uri := v3TradeBill + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &BillRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TradeBill)
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

// 申请资金账单API
// 注意：如 bill_date 为空，默认查前一天的
// Code = 0 is success
func (c *ClientV3) V3BillFundFlowBill(ctx context.Context, bm gopay.BodyMap) (wxRsp *BillRsp, err error) {
	if bm != nil {
		if bm.GetString("bill_date") == util.NULL {
			now := time.Now()
			yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local).Format(util.DateLayout)
			bm.Set("bill_date", yesterday)
		}
	}
	uri := v3FundFlowBill + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &BillRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TradeBill)
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

// 申请特约商户资金账单API
// 注意：如 bill_date 为空，默认查前一天的
// Code = 0 is success
func (c *ClientV3) V3BillEcommerceFundFlowBill(ctx context.Context, bm gopay.BodyMap) (wxRsp *EcommerceFundFlowBillRsp, err error) {
	if bm != nil {
		if bm.GetString("bill_date") == util.NULL {
			now := time.Now()
			yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local).Format(util.DateLayout)
			bm.Set("bill_date", yesterday)
		}
		if bm.GetString("account_type") == util.NULL {
			bm.Set("account_type", "ALL")
		}
		if bm.GetString("algorithm") == util.NULL {
			bm.Set("algorithm", "AEAD_AES_256_GCM")
		}
	}
	uri := v3EcommerceFundFlowBill + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &EcommerceFundFlowBillRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(DownloadBill)
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

// 申请单个子商户资金账单API
// 注意：如 bill_date 为空，默认查前一天的
// Code = 0 is success
func (c *ClientV3) V3BillSubFundFlowBill(ctx context.Context, bm gopay.BodyMap) (wxRsp *BillRsp, err error) {
	if bm != nil {
		if bm.GetString("bill_date") == util.NULL {
			now := time.Now()
			yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local).Format(util.DateLayout)
			bm.Set("bill_date", yesterday)
		}
	}
	uri := v3SubFundFlowBill + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &BillRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TradeBill)
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

// 下载账单API
// Code = 0 is success
func (c *ClientV3) V3BillDownLoadBill(ctx context.Context, downloadUrl string) (fileBytes []byte, err error) {
	if downloadUrl == gopay.NULL {
		return nil, errors.New("invalid download url")
	}
	split := strings.Split(downloadUrl, ".com")
	if len(split) != 2 {
		return nil, errors.New("invalid download url")
	}
	authorization, err := c.authorization(MethodGet, split[1], nil)
	if err != nil {
		return nil, err
	}
	res, _, bs, err := c.doProdGet(ctx, split[1], authorization)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}
