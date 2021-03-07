package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
)

// 申请交易账单API
//	Code = 0 is success
//	注意：账单日期不可写当天日期
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/bill/chapter3_1.shtml
func (c *ClientV3) V3BillTradeBill(bm gopay.BodyMap) (wxRsp *BillRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		uri      string
	)
	if bm != nil {
		if bm.GetString("bill_date") == util.NULL {
			now := time.Now()
			yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local).Format(util.DateLayout)
			bm.Set("bill_date", yesterday)
		}
	}
	uri = v3ApiTradeBill + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &BillRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TradeBill)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 申请资金账单API
//	Code = 0 is success
//	注意：账单日期不可写当天日期
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/bill/chapter3_2.shtml
func (c *ClientV3) V3BillFundFlowBill(bm gopay.BodyMap) (wxRsp *BillRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		uri      string
	)
	if bm != nil {
		if bm.GetString("bill_date") == util.NULL {
			now := time.Now()
			yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local).Format(util.DateLayout)
			bm.Set("bill_date", yesterday)
		}
	}
	uri = v3ApiFundFlowBill + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &BillRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TradeBill)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 申请二级商户资金账单API
//	Code = 0 is success
//	注意：账单日期不可写当天日期
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/bill/chapter3_2.shtml
func (c *ClientV3) V3BillLevel2FundFlowBill(bm gopay.BodyMap) (wxRsp *Level2FundFlowBillRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		uri      string
	)
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
	uri = v3ApiLevel2FundFlowBill + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &Level2FundFlowBillRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(DownloadBill)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 下载账单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/bill/chapter3_3.shtml
func (c *ClientV3) V3BillDownLoadBill(downloadUrl string) (fileBytes []byte, err error) {
	if downloadUrl == gopay.NULL {
		return nil, errors.New("invalid download url")
	}
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
	)
	split := strings.Split(downloadUrl, ".com")
	if len(split) != 2 {
		return nil, errors.New("invalid download url")
	}
	authorization, err := c.authorization(MethodGet, split[1], nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, _, bs, err := c.doProdGet(split[1], authorization)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}
