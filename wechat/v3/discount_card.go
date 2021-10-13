package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cedarwu/gopay"
)

// 预受理领卡请求API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_3_1.shtml
func (c *ClientV3) V3DiscountCardApply(bm gopay.BodyMap) (wxRsp *DiscountCardApplyRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3CardPre, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3CardPre, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &DiscountCardApplyRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(DiscountCardApply)
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

// 增加用户记录API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_3_2.shtml
func (c *ClientV3) V3DiscountCardAddUser(bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	if err = bm.CheckEmptyError("out_card_code"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(v3CardAddUser, bm.GetString("out_card_code"))
	bm.Remove("out_card_code")
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询先享卡订单API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_3_3.shtml
func (c *ClientV3) V3DiscountCardQuery(outCardCode string) (wxRsp *DiscountCardQueryRsp, err error) {
	url := fmt.Sprintf(v3CardQuery, outCardCode)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &DiscountCardQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(DiscountCardQuery)
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
