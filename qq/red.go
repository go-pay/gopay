/*
	QQ 现金红包
	文档：https://qpay.qq.com/buss/wiki/221/1219
*/

package qq

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/go-pay/gopay"
)

// SendCashRed 创建现金红包
// 注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
// 文档：https://qpay.qq.com/buss/wiki/221/1220
func (q *Client) SendCashRed(ctx context.Context, bm gopay.BodyMap) (qqRsp *SendCashRedResponse, err error) {
	err = bm.CheckEmptyError("charset", "nonce_str", "mch_billno", "mch_name", "re_openid",
		"total_amount", "total_num", "wishing", "act_name", "icon_id", "min_value", "max_value")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQRed(ctx, bm, createCashRed)
	if err != nil {
		return nil, err
	}
	qqRsp = new(SendCashRedResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// DownloadRedListFile 对账单下载
//
//	注意：data类型为int类型，例如：date=20200909，2020年9月9日
//	文档：https://qpay.qq.com/buss/wiki/221/1224
func (q *Client) DownloadRedListFile(ctx context.Context, bm gopay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("date")
	if err != nil {
		return gopay.NULL, err
	}
	bs, err := q.doQQGet(ctx, bm, redFileDown, SignType_MD5)
	if err != nil {
		return gopay.NULL, err
	}
	return string(bs), nil
}

// QueryRedInfo 查询红包详情
//
//	文档：https://qpay.qq.com/buss/wiki/221/2174
func (q *Client) QueryRedInfo(ctx context.Context, bm gopay.BodyMap) (qqRsp *QueryRedInfoResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "listid")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQRed(ctx, bm, queryRedInfo)
	if err != nil {
		return nil, err
	}
	qqRsp = new(QueryRedInfoResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}
