package qq

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/iGoogle-ink/gopay"
)

type Client struct {
	MchId      string
	ApiKey     string
	CertFile   []byte
	KeyFile    []byte
	Pkcs12File []byte
	IsProd     bool
	mu         sync.RWMutex
}

// 初始化QQ客户端（正式环境）
//    mchId：商户ID
//    ApiKey：API秘钥值
func NewClient(mchId, apiKey string) (client *Client) {
	return &Client{
		MchId:  mchId,
		ApiKey: apiKey,
	}
}

// 提交付款码支付
//    文档地址：https://qpay.qq.com/buss/wiki/43/1157
func (q *Client) MicroPay(bm gopay.BodyMap) (qqRsp *MicroPayResponse, err error) {
	err = bm.CheckEmptyError("sub_mch_id", "nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "device_info", "auth_code")
	if err != nil {
		return nil, err
	}
	if bm.Get("trade_type") == gopay.NULL {
		bm.Set("trade_type", TradeType_MicroPay)
	}
	bs, err := q.doQQ(bm, qqMicroPay, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(MicroPayResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%s", string(bs), err.Error())
	}
	return qqRsp, nil
}

// 向QQ发送请求
func (q *Client) doQQ(bm gopay.BodyMap, url string, tlsConfig *tls.Config) (bs []byte, err error) {
	bm.Set("mch_id", q.MchId)
	if bm.Get("fee_type") == gopay.NULL {
		bm.Set("fee_type", "CNY")
	}

	if bm.Get("sign") == gopay.NULL {
		var sign string
		sign = getReleaseSign(q.ApiKey, bm.Get("sign_type"), bm)
		bm.Set("sign", sign)
	}

	httpClient := gopay.NewHttpClient()
	if tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig)
	}

	res, bs, errs := httpClient.Type(gopay.TypeXML).Post(url).SendString(generateXml(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	if strings.Contains(string(bs), "HTML") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}
