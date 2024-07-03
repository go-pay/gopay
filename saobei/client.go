package saobei

import (
	"context"
	"crypto/md5"
	"fmt"
	"hash"
	"sync"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/xlog"
)

type Client struct {
	instNo      string //商户系统机构号inst_no
	key         string // 商户系统令牌
	merchantNo  string // 支付系统：商户号
	terminalId  string // 支付系统：商户号终端号
	accessToken string // 支付系统： 令牌
	isProd      bool   // 是否正式环境
	payVer      string //版本号 当前201
	serviceId   string //接口类型，当前类型015
	hc          *xhttp.Client
	mu          sync.Mutex
	md5Hash     hash.Hash
}

// NewClient 初始化扫呗客户端
// instNo      string //商户系统机构号inst_no
// key         string // 商户系统令牌
// merchantNo  string // 支付系统：商户号
// terminalId  string // 支付系统：商户号终端号
// accessToken string // 支付系统： 令牌
// isProd：是否是正式环境
func NewClient(instNo, key, merchantNo, terminalId, accessToken string, isProd bool) (*Client, error) {
	return &Client{
		instNo:      instNo,
		key:         key,
		merchantNo:  merchantNo,
		terminalId:  terminalId,
		accessToken: accessToken,
		isProd:      isProd,
		hc:          xhttp.NewClient(),
		md5Hash:     md5.New(),
		payVer:      "201",
		serviceId:   "015",
	}, nil
}

// pubParamsHandle 公共参数处理
func (c *Client) pubParamsHandle(bm gopay.BodyMap) gopay.BodyMap {
	if ver := bm.GetString("pay_ver"); ver == gopay.NULL {
		bm.Set("pay_ver", c.payVer)
	}
	if v := bm.GetString("service_id"); v == gopay.NULL {
		bm.Set("service_id", c.serviceId)
	}
	if v := bm.GetString("merchant_no"); v == gopay.NULL {
		bm.Set("merchant_no", c.merchantNo)
	}
	if v := bm.GetString("terminal_id"); v == gopay.NULL {
		bm.Set("terminal_id", c.terminalId)
	}
	sign := c.getRsaSign(bm)
	bm.Set("key_sign", sign)
	return bm
}

// doPost 发起请求
func (c *Client) doPost(ctx context.Context, path string, bm gopay.BodyMap) (bs []byte, err error) {
	param := c.pubParamsHandle(bm)
	xlog.Debugf("saobeiParam:%+v", param.JsonBody())
	url := baseUrl
	if !c.isProd {
		url = sandboxBaseUrl
	}
	res, bs, err := c.hc.Req(xhttp.TypeJSON).Post(url + path).SendBodyMap(param).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}
