package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 创建投诉通知回调地址API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_2.shtml
func (c *ClientV3) V3ComplaintNotifyUrlCreate(url string) (wxRsp *ComplaintNotifyUrlRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("url", url)
	authorization, err := c.authorization(MethodPost, v3ComplaintNotifyUrlCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3ComplaintNotifyUrlCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintNotifyUrlRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ComplaintNotifyUrl)
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

// 查询投诉通知回调地址API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_3.shtml
func (c *ClientV3) V3ComplaintNotifyUrlQuery() (wxRsp *ComplaintNotifyUrlRsp, err error) {
	authorization, err := c.authorization(MethodGet, v3ComplaintNotifyUrlQuery, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(v3ComplaintNotifyUrlQuery, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintNotifyUrlRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ComplaintNotifyUrl)
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

// 更新投诉通知回调地址API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_4.shtml
func (c *ClientV3) V3ComplaintNotifyUrlUpdate(url string) (wxRsp *ComplaintNotifyUrlRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("url", url)
	authorization, err := c.authorization(MethodPut, v3ComplaintNotifyUrlUpdate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPut(bm, v3ComplaintNotifyUrlUpdate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintNotifyUrlRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ComplaintNotifyUrl)
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

// 删除投诉通知回调地址API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_5.shtml
func (c *ClientV3) V3ComplaintNotifyUrlDelete() (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodDelete, v3ComplaintNotifyUrlDelete, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdDelete(nil, v3ComplaintNotifyUrlDelete, authorization)
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

// 商户上传反馈图片API
//	注意：图片不能超过2MB
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_10.shtml
func (c *ClientV3) V3ComplaintUploadImage(fileName, fileSha256 string, img *util.File) (wxRsp *MediaUploadRsp, err error) {
	bmFile := make(gopay.BodyMap)
	bmFile.Set("filename", fileName).Set("sha256", fileSha256)
	authorization, err := c.authorization(MethodPost, v3ComplaintUploadImage, bmFile)
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.SetBodyMap("meta", func(bm gopay.BodyMap) {
		bm.Set("filename", fileName).Set("sha256", fileSha256)
	}).SetFormFile("file", img)
	res, si, bs, err := c.doProdPostFile(bm, v3ComplaintUploadImage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MediaUploadRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MediaUpload)
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

// 查询投诉单列表API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_11.shtml
func (c *ClientV3) V3ComplaintList(beginDate, endDate, complaintedMchid string, limit, offset int) (wxRsp *ComplaintListRsp, err error) {
	if limit == 0 {
		limit = 50
	}
	uri := v3ComplaintList + "?begin_date=" + beginDate + "&end_date=" + endDate + "&limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
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

// 查询投诉协商历史API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_12.shtml
func (c *ClientV3) V3ComplaintNegotiationHistory(complaintId string, limit, offset int) (wxRsp *ComplaintNegotiationHistoryRsp, err error) {
	if limit == 0 {
		limit = 50
	}
	uri := fmt.Sprintf(v3ComplaintNegotiationHistory, complaintId) + "?limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ComplaintNegotiationHistoryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ComplaintNegotiationHistory)
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

// 提交回复API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_14.shtml
func (c *ClientV3) V3ComplaintResponse(complaintId, complaintedMchid, content string, mediaIds []string) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ComplaintResponse, complaintId)
	bm := make(gopay.BodyMap)
	bm.Set("complainted_mchid", complaintedMchid).
		Set("response_content", content).
		Set("response_images", mediaIds)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
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

// 反馈处理完成API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_15.shtml
func (c *ClientV3) V3ComplaintComplete(complaintId, complaintedMchid string) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3ComplaintComplete, complaintId)
	bm := make(gopay.BodyMap)
	bm.Set("complainted_mchid", complaintedMchid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
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
