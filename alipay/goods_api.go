package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.merchant.item.file.upload(商品文件上传接口)
// 文档地址：https://opendocs.alipay.com/apis/api_4/alipay.merchant.item.file.upload
func (a *Client) MerchantItemFileUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *MerchantItemFileUploadRsp, err error) {
	err = bm.CheckEmptyError("scene", "file_content")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.FileUploadRequest(ctx, bm, "alipay.merchant.item.file.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(MerchantItemFileUploadRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
