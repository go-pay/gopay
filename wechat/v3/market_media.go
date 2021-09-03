package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 图片上传(营销专用)API
//	注意：图片不能超过2MB
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_0_1.shtml
func (c *ClientV3) v3FavorMediaUploadImage(fileName, fileSha256 string, img *util.File) (wxRsp *MarketMediaUploadRsp, err error) {
	bmFile := make(gopay.BodyMap)
	bmFile.Set("filename", fileName).Set("sha256", fileSha256)
	authorization, err := c.authorization(MethodPost, v3FavorMediaUploadImage, bmFile)
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.SetBodyMap("meta", func(bm gopay.BodyMap) {
		bm.Set("filename", fileName).Set("sha256", fileSha256)
	}).SetFormFile("file", img)
	res, si, bs, err := c.doProdPostFile(bm, v3FavorMediaUploadImage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MarketMediaUploadRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MarketMediaUpload)
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
