package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iGoogle-ink/gopay/pkg/util"
)

// 查询投诉单列表API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_11.shtml
func (c *ClientV3) V3ComplaintList(beginDate, endDate, complaintedMchid string, limit, offset int) (wxRsp *ComplaintListRsp, err error) {
	uri := v3ComplaintList + "?begin_date=" + beginDate + "&end_date=" + endDate
	if limit != 0 {
		uri += "&limit=" + util.Int2String(limit)
	}
	if offset != 0 {
		uri += "&offset=" + util.Int2String(offset)
	}
	if complaintedMchid != "" {
		uri += "&complainted_mchid=" + complaintedMchid
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ComplaintList)
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

// 查询投诉单详情API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_13.shtml
func (c *ClientV3) V3ComplaintDetail(complaintId string) (wxRsp *ComplaintDetailRsp, err error) {
	url := fmt.Sprintf(v3ComplaintDetail, complaintId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ComplaintDetail)
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
