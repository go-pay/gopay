package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
)

// 请求分账API
//	微信会在接到请求后立刻返回请求接收结果，分账结果需要自行调用查询接口来获取
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_1.shtml
func (c *ClientV3) V3ProfitShareOrder(bm gopay.BodyMap) (*ProfitSharingOrderRsp, error) {
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}

	authorization, err := c.authorization(MethodPost, v3ProfitSharingOrders, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitSharingOrders, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitSharingOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitSharingOrderResponse)
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

// 查询分账结果,所有的分账请求，都必须走该查询接口来确定一笔分账的最终处理结果
// 	tradeNo:必填，商户自己的分账单号
// 	transId:必填，微信订单号
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_2.shtml
func (c *ClientV3) V3ProfitShareOrderQuery(tradeNo, transId string) (*ProfitSharingOrderResultRsp, error) {
	url := fmt.Sprintf(`%s/%s?transaction_id=%s`, v3ProfitSharingQuery, tradeNo, transId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitSharingOrderResultRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitSharingOrderResultResponse)
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

// 解冻剩余资金,不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给特约商户.
//	微信会在接到请求后立刻返回请求接收结果，结果需要自行调用查询接口来获取
// 	tradeNo:必填，商户自己的分账单号
// 	transId:必填，微信订单号
// 	description:必填，分账的原因描述，分账账单中需要体现
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_5.shtml
func (c *ClientV3) V3ProfitShareOrderUnfreeze(tradeNo, transId, description string) (*ProfitSharingOrderUnfreezeRsp, error) {
	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", transId).Set("out_order_no", tradeNo).Set("description", description)

	authorization, err := c.authorization(MethodPost, v3ProfitSharingUnfreeze, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitSharingUnfreeze, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitSharingOrderUnfreezeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitSharingOrderUnfreezeResponse)
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

// 查询剩余待分金额
// 	transId:必填，微信订单号
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_6.shtml
func (c *ClientV3) V3ProfitShareOrderUnsplitQuery(transId string) (*ProfitSharingUnsplitAmountRsp, error) {
	url := fmt.Sprintf(v3ProfitSharingUnsplitAmount, transId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitSharingUnsplitAmountRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitSharingUnsplitAmountResponse)
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

// 新增分账接收方
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_8.shtml
func (c *ClientV3) V3ProfitShareAddReceivers(bm gopay.BodyMap) (*ProfitSharingAddReceiverRsp, error) {
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}

	authorization, err := c.authorization(MethodPost, v3ProfitSharingAddReceiver, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitSharingAddReceiver, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitSharingAddReceiverRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitSharingAddReceiverResponse)
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

// 删除分账接收方,商户发起删除分账接收方请求。删除后，不支持将分账方商户结算后的资金，分到该分账接收方
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml
func (c *ClientV3) V3ProfitShareDeleteReceiver(bm gopay.BodyMap) (*ProfitSharingDeleteReceiverRsp, error) {
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}

	authorization, err := c.authorization(MethodPost, v3ProfitSharingDeleteReceiver, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ProfitSharingDeleteReceiver, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &ProfitSharingDeleteReceiverRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ProfitSharingDeleteReceiverResponse)
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
