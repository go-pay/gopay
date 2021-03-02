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

// APP下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_1.shtml
func (c *ClientV3) V3TransactionApp(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	if bm != nil {
		if bm.GetString("appid") == util.NULL {
			bm.Set("appid", c.Appid)
		}
		if bm.GetString("mchid") == util.NULL {
			bm.Set("mchid", c.Mchid)
		}
	}
	authorization, err := c.authorization(MethodPost, v3ApiPayApp, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ApiPayApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
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

// JSAPI/小程序下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_2.shtml
func (c *ClientV3) V3TransactionJsapi(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	if bm != nil {
		if bm.GetString("appid") == util.NULL {
			bm.Set("appid", c.Appid)
		}
		if bm.GetString("mchid") == util.NULL {
			bm.Set("mchid", c.Mchid)
		}
	}
	authorization, err := c.authorization(MethodPost, v3ApiJsapi, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ApiJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
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

// Native下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_3.shtml
func (c *ClientV3) V3TransactionNative(bm gopay.BodyMap) (wxRsp *NativeRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	if bm != nil {
		if bm.GetString("appid") == util.NULL {
			bm.Set("appid", c.Appid)
		}
		if bm.GetString("mchid") == util.NULL {
			bm.Set("mchid", c.Mchid)
		}
	}
	authorization, err := c.authorization(MethodPost, v3ApiNative, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ApiNative, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &NativeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Native)
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

// H5下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_4.shtml
func (c *ClientV3) V3TransactionH5(bm gopay.BodyMap) (wxRsp *H5Rsp, err error) {
	ts := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	if bm != nil {
		if bm.GetString("appid") == util.NULL {
			bm.Set("appid", c.Appid)
		}
		if bm.GetString("mchid") == util.NULL {
			bm.Set("mchid", c.Mchid)
		}
	}
	authorization, err := c.authorization(MethodPost, v3ApiH5, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ApiH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(H5Url)
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

// 查询订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_5.shtml
func (c *ClientV3) V3TransactionQueryOrder(orderNoType OrderNoType, orderNo string) (wxRsp *QueryOrderRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		uri      string
	)
	switch orderNoType {
	case TransactionId:
		uri = fmt.Sprintf(v3ApiQueryOrderTransactionId, orderNo) + "?mchid=" + c.Mchid
	case OutTradeNo:
		uri = fmt.Sprintf(v3ApiQueryOrderOutTradeNo, orderNo) + "?mchid=" + c.Mchid
	default:
		return nil, errors.New("unsupported order number type")
	}
	authorization, err := c.authorization(MethodGet, uri, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &QueryOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(QueryOrder)
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

// 关闭订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_6.shtml
func (c *ClientV3) V3TransactionCloseOrder(tradeNo string) (wxRsp *CloseOrderRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		url      = fmt.Sprintf(v3ApiCloseOrder, tradeNo)
	)
	bm := make(gopay.BodyMap)
	bm.Set("mchid", c.Mchid)
	authorization, err := c.authorization(MethodPost, url, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &CloseOrderRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

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

// 合单APP下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_1.shtml
func (c *ClientV3) V3CombineTransactionApp(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	if bm != nil {
		if bm.GetString("combine_appid") == util.NULL {
			bm.Set("combine_appid", c.Appid)
		}
		if bm.GetString("combine_mchid") == util.NULL {
			bm.Set("combine_mchid", c.Mchid)
		}
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayApp, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3CombinePayApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
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

// 合单H5下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_2.shtml
func (c *ClientV3) V3CombineTransactionH5(bm gopay.BodyMap) (wxRsp *H5Rsp, err error) {
	ts := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	if bm != nil {
		if bm.GetString("combine_appid") == util.NULL {
			bm.Set("combine_appid", c.Appid)
		}
		if bm.GetString("combine_mchid") == util.NULL {
			bm.Set("combine_mchid", c.Mchid)
		}
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayH5, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3CombinePayH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(H5Url)
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

// 合单JSAPI/小程序下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_3.shtml
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_4.shtml
func (c *ClientV3) V3CombineTransactionJsapi(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	if bm != nil {
		if bm.GetString("combine_appid") == util.NULL {
			bm.Set("combine_appid", c.Appid)
		}
		if bm.GetString("combine_mchid") == util.NULL {
			bm.Set("combine_mchid", c.Mchid)
		}
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayJsapi, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3CombinePayJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
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

// 合单Native下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_5.shtml
func (c *ClientV3) V3CombineTransactionNative(bm gopay.BodyMap) (wxRsp *NativeRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := util.GetRandomString(32)
	if bm != nil {
		if bm.GetString("combine_appid") == util.NULL {
			bm.Set("combine_appid", c.Appid)
		}
		if bm.GetString("combine_mchid") == util.NULL {
			bm.Set("combine_mchid", c.Mchid)
		}
	}
	authorization, err := c.authorization(MethodPost, v3CombineNative, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3CombineNative, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &NativeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Native)
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

// 合单查询订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_11.shtml
func (c *ClientV3) V3CombineTransactionQueryOrder(traderNo string) (wxRsp *CombineQueryOrderRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		uri      string
	)
	uri = fmt.Sprintf(v3CombineQuery, traderNo)
	authorization, err := c.authorization(MethodGet, uri, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &CombineQueryOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(CombineQueryOrder)
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

// 合单关闭订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_12.shtml
func (c *ClientV3) V3CombineTransactionCloseOrder(tradeNo string, bm gopay.BodyMap) (wxRsp *CloseOrderRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		url      = fmt.Sprintf(v3CombineClose, tradeNo)
	)
	if bm != nil {
		if bm.GetString("combine_appid") == util.NULL {
			bm.Set("combine_appid", c.Appid)
		}
	}
	authorization, err := c.authorization(MethodPost, url, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &CloseOrderRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 申请退款API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_14.shtml
func (c *ClientV3) V3Refund(bm gopay.BodyMap) (wxRsp *RefundRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		url      = v3Refund
	)
	authorization, err := c.authorization(MethodPost, url, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundOrderResponse)
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

// 查询单笔退款API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_10.shtml
func (c *ClientV3) V3RefundQuery(outRefundNo string) (wxRsp *RefundQueryRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = util.GetRandomString(32)
		uri      = fmt.Sprintf(v3RefundQuery, outRefundNo)
	)
	authorization, err := c.authorization(MethodGet, uri, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundQueryResponse)
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
