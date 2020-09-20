/*
	QQ 现金红包
	文档：https://qpay.qq.com/buss/wiki/221/1219
*/

package qq

import (
	"encoding/xml"
	"fmt"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
)

// SendCashRed 创建现金红包
//	// 注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	文档：https://qpay.qq.com/buss/wiki/221/1220
func (q *Client) SendCashRed(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (qqRsp *SendCashRedResponse, err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("charset", "nonce_str", "mch_billno", "mch_name", "re_openid",
		"total_amount", "total_num", "wishing", "act_name", "icon_id", "min_value", "max_value")
	if err != nil {
		return nil, err
	}
	tlsConfig, err := q.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQRed(bm, createCashRed, tlsConfig)
	if err != nil {
		return nil, err
	}
	qqRsp = new(SendCashRedResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}

// DownloadRedListFile 对账单下载
//	注意：data类型为int类型，例如：date=20200909，2020年9月9日
//	文档：https://qpay.qq.com/buss/wiki/221/1224
func (q *Client) DownloadRedListFile(bm gopay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("date")
	if err != nil {
		return gopay.NULL, err
	}
	bs, err := q.doQQGet(bm, redFileDown, SignType_MD5)
	if err != nil {
		return gotil.NULL, err
	}
	return string(bs), nil
}

// QueryRedInfo 查询红包详情
//	文档：https://qpay.qq.com/buss/wiki/221/2174
func (q *Client) QueryRedInfo(bm gopay.BodyMap) (qqRsp *QueryRedInfoResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "listid")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQRed(bm, queryRedInfo, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(QueryRedInfoResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return qqRsp, nil
}
