package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 请求分账API
//	微信会在接到请求后立刻返回请求接收结果，分账结果需要自行调用查询接口来获取
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_1.shtml
func (c *ClientV3) V3ProfitShareOrder(bm gopay.BodyMap) (*ProfitShareOrderRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareOrder, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitShareOrder, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareOrder)
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

// 查询分账结果API
//	Code = 0 is success
// 	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_2.shtml
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_2.shtml
func (c *ClientV3) V3ProfitShareOrderQuery(orderNo string, bm gopay.BodyMap) (*ProfitShareOrderQueryRsp, error) {
	uri := fmt.Sprintf(v3ProfitShareQuery, orderNo) + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareOrderQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareOrderQuery)
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

// 请求分账回退API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_3.shtml
func (c *ClientV3) V3ProfitShareReturn(bm gopay.BodyMap) (*ProfitShareReturnRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareReturn, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitShareReturn, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareReturnRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareReturn)
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

// 查询分账回退结果API
//	Code = 0 is success
// 	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_4.shtml
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_4.shtml
func (c *ClientV3) V3ProfitShareReturnResult(returnNo string, bm gopay.BodyMap) (*ProfitShareReturnResultRsp, error) {
	uri := fmt.Sprintf(v3ProfitShareReturnResult, returnNo) + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareReturnResultRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareReturnResult)
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

// 解冻剩余资金API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_5.shtml
// 服务商模式会比直连商户模式多sub_mchid字段，改为gopay.BodyMap更灵活
func (c *ClientV3) V3ProfitShareOrderUnfreeze(bm gopay.BodyMap) (*ProfitShareOrderUnfreezeRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareUnfreeze, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitShareUnfreeze, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareOrderUnfreezeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareOrderUnfreeze)
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

// 查询剩余待分金额API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_6.shtml
func (c *ClientV3) V3ProfitShareUnsplitAmount(transId string) (*ProfitShareUnsplitAmountRsp, error) {
	url := fmt.Sprintf(v3ProfitShareUnsplitAmount, transId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareUnsplitAmountRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareUnsplitAmount)
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

// 新增分账接收方API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_8.shtml
func (c *ClientV3) V3ProfitShareAddReceiver(bm gopay.BodyMap) (*ProfitShareAddReceiverRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareAddReceiver, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitShareAddReceiver, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareAddReceiverRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareAddReceiver)
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

// 删除分账接收方API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml
func (c *ClientV3) V3ProfitShareDeleteReceiver(bm gopay.BodyMap) (*ProfitShareDeleteReceiverRsp, error) {
	authorization, err := c.authorization(MethodPost, v3ProfitShareDeleteReceiver, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitShareDeleteReceiver, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareDeleteReceiverRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareDeleteReceiver)
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

// 查询最大分账比例API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_7.shtml
func (c *ClientV3) V3ProfitShareMerchantConfigs(subMchId string) (*ProfitShareMerchantConfigsRsp, error) {
	uri := fmt.Sprintf(v3ProfitShareMerchantConfigs, subMchId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareMerchantConfigsRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareMerchantConfigs)
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

// 申请分账账单
//	Code = 0 is success
// 直连商户：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_11.shtml
// 服务商：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_11.shtml
func (c *ClientV3) V3ProfitShareBills(billDate string, bm gopay.BodyMap) (*ProfitShareBillsRsp, error) {
	uri := fmt.Sprintf(v3ProfitShareBills, billDate) + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitShareBillsRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitShareBills)
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
