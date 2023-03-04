package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 图片上传API
// 注意：图片不能超过2MB
// Code = 0 is success
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter2_1_1.shtml
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter2_1_1.shtml
func (c *ClientV3) V3MediaUploadImage(ctx context.Context, fileName, fileSha256 string, img *util.File) (wxRsp *MediaUploadRsp, err error) {
	bmFile := make(gopay.BodyMap)
	bmFile.Set("filename", fileName).Set("sha256", fileSha256)
	authorization, err := c.authorization(MethodPost, v3MediaUploadImage, bmFile)
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.SetBodyMap("meta", func(bm gopay.BodyMap) {
		bm.Set("filename", fileName).Set("sha256", fileSha256)
	}).SetFormFile("file", img)
	res, si, bs, err := c.doProdPostFile(ctx, bm, v3MediaUploadImage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MediaUploadRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MediaUpload)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 视频上传API
// 注意：媒体视频只支持avi、wmv、mpeg、mp4、mov、mkv、flv、f4v、m4v、rmvb格式，文件大小不能超过5M。
// Code = 0 is success
// 商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter2_1_2.shtml
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter2_1_2.shtml
func (c *ClientV3) V3MediaUploadVideo(ctx context.Context, fileName, fileSha256 string, img *util.File) (wxRsp *MediaUploadRsp, err error) {
	bmFile := make(gopay.BodyMap)
	bmFile.Set("filename", fileName).Set("sha256", fileSha256)
	authorization, err := c.authorization(MethodPost, v3MediaUploadVideo, bmFile)
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.SetBodyMap("meta", func(bm gopay.BodyMap) {
		bm.Set("filename", fileName).Set("sha256", fileSha256)
	}).SetFormFile("file", img)
	res, si, bs, err := c.doProdPostFile(ctx, bm, v3MediaUploadVideo, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MediaUploadRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MediaUpload)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
