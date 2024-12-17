package lakala

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

// Client lakala
type Client struct {
	ctx            context.Context   // 上下文
	PartnerCode    string            // partner_code:商户编码，由4~6位大写字母或数字构成
	credentialCode string            // credential_code:系统为商户分配的开发校验码，请妥善保管，不要在公开场合泄露
	IsProd         bool              // 是否生产环境
	DebugSwitch    gopay.DebugSwitch // 调试开关，是否打印日志
	logger         xlog.XLogger
	hc             *xhttp.Client
	sha256Hash     hash.Hash
	mu             sync.Mutex
}

// NewClient 初始化lakala户端
// partnerCode: 商户编码，由4~6位大写字母或数字构成
// credentialCode: 系统为商户分配的开发校验码，请妥善保管，不要在公开场合泄露
// isProd: 是否生产环境
func NewClient(partnerCode, credentialCode string, isProd bool) (client *Client, err error) {
	if partnerCode == gopay.NULL || credentialCode == gopay.NULL {
		return nil, gopay.MissLakalaInitParamErr
	}
	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)
	client = &Client{
		ctx:            context.Background(),
		PartnerCode:    partnerCode,
		credentialCode: credentialCode,
		IsProd:         isProd,
		DebugSwitch:    gopay.DebugOff,
		logger:         logger,
		hc:             xhttp.NewClient(),
		sha256Hash:     sha256.New(),
	}
	return client, nil
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}

// SetHttpClient 设置自定义的xhttp.Client
func (c *Client) SetHttpClient(client *xhttp.Client) {
	if client != nil {
		c.hc = client
	}
}

func (c *Client) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		c.logger = logger
	}
}

// 公共参数处理 Query Params
func (c *Client) pubParamsHandle() (param string, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("time", time.Now().UnixMilli())
	bm.Set("nonce_str", util.RandomString(20))
	sign, err := c.getRsaSign(bm)
	if err != nil {
		return "", fmt.Errorf("GetRsaSign Error: %w", err)
	}
	bm.Set("sign", sign)
	param = bm.EncodeURLParams()
	return
}

// 验证签名
func VerifySign(notifyReq *NotifyRequest, partnerCode string, credentialCode string) (err error) {
	validStr := fmt.Sprintf("%v&%v&%v&%v", partnerCode, notifyReq.Time, notifyReq.NonceStr, credentialCode)
	h := sha256.New()
	h.Write([]byte(validStr))
	validSign := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	if notifyReq.Sign != validSign {
		return fmt.Errorf("签名验证失败")
	}
	return
}

// getRsaSign 获取签名字符串
func (c *Client) getRsaSign(bm gopay.BodyMap) (sign string, err error) {
	var (
		partnerCode    = c.PartnerCode
		ts             = bm.Get("time")
		nonceStr       = bm.Get("nonce_str")
		credentialCode = c.credentialCode
	)
	if ts == "" || nonceStr == "" {
		return "", fmt.Errorf("签名缺少必要的参数")
	}
	validStr := fmt.Sprintf("%v&%v&%v&%v", partnerCode, ts, nonceStr, credentialCode)
	c.mu.Lock()
	defer func() {
		c.sha256Hash.Reset()
		c.mu.Unlock()
	}()
	c.sha256Hash.Write([]byte(validStr))
	sign = strings.ToLower(hex.EncodeToString(c.sha256Hash.Sum(nil)))
	return
}

// PUT 发起请求
func (c *Client) doPut(ctx context.Context, path string, bm gopay.BodyMap) (bs []byte, err error) {
	var url = baseUrlProd + path
	param, err := c.pubParamsHandle()
	if err != nil {
		return nil, err
	}
	req := c.hc.Req()
	req.Header.Add("Accept", "application/json")
	uri := url + "?" + param
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Lakala_Url: %s", uri)
		c.logger.Debugf("Lakala_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Lakala_Req_Headers: %#v", req.Header)
	}
	res, bs, err := req.Put(uri).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// PUT 发起请求
func (c *Client) doPost(ctx context.Context, path string, bm gopay.BodyMap) (bs []byte, err error) {
	var url = baseUrlProd + path
	param, err := c.pubParamsHandle()
	if err != nil {
		return nil, err
	}
	req := c.hc.Req()
	req.Header.Add("Accept", "application/json")
	uri := url + "?" + param
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Lakala_Url: %s", uri)
		c.logger.Debugf("Lakala_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Lakala_Req_Headers: %#v", req.Header)
	}
	res, bs, err := req.Post(uri).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// GET 发起请求
func (c *Client) doGet(ctx context.Context, path, queryParams string) (bs []byte, err error) {
	var url = baseUrlProd + path
	param, err := c.pubParamsHandle()
	if err != nil {
		return nil, err
	}
	if queryParams != "" {
		param = param + "&" + queryParams
	}
	req := c.hc.Req()
	req.Header.Add("Accept", "application/json")
	uri := url + "?" + param
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Lakala_Url: %s", uri)
		c.logger.Debugf("Lakala_Req_Headers: %#v", req.Header)
	}
	res, bs, err := req.Get(uri).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil

}
