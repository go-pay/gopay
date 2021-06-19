package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 商圈积分同步API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_2.shtml
func (c *ClientV3) V3BusinessPointsSync(url string) (wxRsp *EmptyRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("url", url)
	authorization, err := c.authorization(MethodPost, v3BusinessPointsSync, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3BusinessPointsSync, authorization)
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

// 商圈积分授权查询API
//	Code = 0 is success
// 	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_4.shtml
func (c *ClientV3) V3BusinessAuthPointsQuery(appid, openid string) (*BusinessAuthPointsQueryRsp, error) {
	if appid == "" {
		appid = c.Appid
	}
	uri := fmt.Sprintf(v3BusinessAuthPointsQuery, openid) + "?appid=" + appid
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &BusinessAuthPointsQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusinessAuthPointsQuery)
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
