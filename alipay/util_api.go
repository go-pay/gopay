package alipay

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.open.app.alipaycert.download(应用支付宝公钥证书下载)
// 文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.app.alipaycert.download
func (a *Client) PublicCertDownload(ctx context.Context, bm gopay.BodyMap) (aliRsp *PublicCertDownloadRsp, err error) {
	err = bm.CheckEmptyError("alipay_cert_sn")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.app.alipaycert.download"); err != nil {
		return nil, err
	}
	aliRsp = new(PublicCertDownloadRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	certBs, err := base64.StdEncoding.DecodeString(aliRsp.Response.AlipayCertContent)
	if err != nil {
		return nil, fmt.Errorf("AlipayCertContent(%s)_DecodeErr:%+v", aliRsp.Response.AlipayCertContent, err)
	}
	aliRsp.Response.AlipayCertContent = string(certBs)
	return aliRsp, nil
}
