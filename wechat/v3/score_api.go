package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
)

// 创建支付分订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_14.shtml
func (c *ClientV3) V3ScoreOrderCreate(bm gopay.BodyMap) (wxRsp *ScoreOrderCreateRsp, err error) {
	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", c.Appid)
	}
	authorization, err := c.authorization(MethodPost, v3ScoreOrderCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ScoreOrderCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderCreateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderCreate)
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

// 查询支付分订单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_15.shtml
func (c *ClientV3) V3ScoreOrderQuery(orderNoType OrderNoType, orderNo, serviceId, appid string) (wxRsp *ScoreOrderQueryRsp, err error) {
	var uri string
	switch orderNoType {
	case OutTradeNo:
		uri = v3ScoreOrderQuery + "?appid=" + appid + "&out_order_no=" + orderNo + "&service_id=" + serviceId
	case QueryId:
		uri = v3ScoreOrderQuery + "?appid=" + appid + "&query_id=" + orderNo + "&service_id=" + serviceId
	default:
		return nil, errors.New("unsupported order number type")
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ScoreOrderQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ScoreOrderQuery)
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
