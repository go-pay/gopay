package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 发起批量转账API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter3_1.shtml
func (c *ClientV3) V3Transfer(bm gopay.BodyMap) (*TransferRsp, error) {
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}

	authorization, err := c.authorization(MethodPost, v3Transfer, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3Transfer, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Transfer)
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

// 微信批次单号查询批次单API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter3_2.shtml
func (c *ClientV3) V3TransferQuery(batchId string, bm gopay.BodyMap) (*TransferQueryRsp, error) {
	url := fmt.Sprintf(v3TransferQuery, batchId)
	bm.Remove("batch_id")
	uri := url + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferQuery)
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

// 微信明细单号查询明细单API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter3_3.shtml
func (c *ClientV3) V3TransferDetailQuery(batchId, detailId string) (*TransferDetailQueryRsp, error) {
	url := fmt.Sprintf(v3TransferDetailQuery, batchId, detailId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferDetailQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferDetailQuery)
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

// 商家批次单号查询批次单API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter3_4.shtml
func (c *ClientV3) V3TransferMerchantQuery(outBatchNo string, bm gopay.BodyMap) (*TransferMerchantQueryRsp, error) {
	url := fmt.Sprintf(v3TransferMerchantQuery, outBatchNo)
	bm.Remove("out_batch_no")
	uri := url + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferMerchantQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferMerchantQuery)
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

// 商家明细单号查询明细单API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter3_5.shtml
func (c *ClientV3) V3TransferMerchantDetailQuery(outBatchNo, outDetailNo string) (*TransferMerchantDetailQueryRsp, error) {
	url := fmt.Sprintf(v3TransferMerchantDetailQuery, outBatchNo, outDetailNo)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferMerchantDetailQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferMerchantDetailQuery)
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

// 转账电子回单申请受理API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter4_1.shtml
func (c *ClientV3) V3TransferReceipt(outBatchNo string) (*TransferReceiptRsp, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_batch_no", outBatchNo)

	authorization, err := c.authorization(MethodPost, v3TransferReceipt, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3TransferReceipt, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferReceiptRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferReceipt)
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

// 查询转账电子回单API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter4_2.shtml
func (c *ClientV3) V3TransferReceiptQuery(outBatchNo string) (*TransferReceiptQueryRsp, error) {
	url := fmt.Sprintf(v3TransferReceiptQuery, outBatchNo)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferReceiptQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferReceiptQuery)
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

// 转账明细电子回单受理API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter4_4.shtml
func (c *ClientV3) V3TransferDetailReceipt(bm gopay.BodyMap) (*TransferDetailReceiptRsp, error) {
	authorization, err := c.authorization(MethodPost, v3TransferDetailReceipt, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3TransferDetailReceipt, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferDetailReceiptRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferDetailReceipt)
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

// 查询转账明细电子回单受理结果API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter4_5.shtml
func (c *ClientV3) V3TransferDetailReceiptQuery(bm gopay.BodyMap) (*TransferDetailReceiptQueryRsp, error) {
	uri := v3TransferDetailReceiptQuery + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferDetailReceiptQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferDetailReceiptQuery)
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
