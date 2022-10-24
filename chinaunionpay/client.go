package chinaunionpay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*BaseClient

	// 服务商配置信息 商户级别的
	AppId        string // 应用ID 实际交易的商户号
	AppKey       string // 签名秘钥
	MerchantCode string // 服务商商户号
	isProEnv bool // 是否生产环境
}

func NewClient(requestId, appId, appKey string, isProEnv bool) *Client {
	c := &Client{
		BaseClient: NewBaseClient(requestId),
		AppId:      appId,
		AppKey:     appKey,
		isProEnv:   isProEnv,
	}
	return c
}

func (c *Client) buildRequestBody(req interface{}) ([]byte, error) {
	return json.Marshal(req)
}

// 请求API地址
func (c *Client) requestApi(reqObj interface{}, addr string) ([]byte, error) {
	body, err := c.buildRequestBody(reqObj)
	if err != nil {
		return nil, err
	}

	c.SetRequestBody(body)

	// 发起请求
	req, err := http.NewRequest("POST", addr, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	authorization, err := Sign(body, c.AppId, c.AppKey)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("format", Format)
	req.Header.Set("charset", Charset)

	resp, err := c.HttpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.SetResponseBody(respBody)
	return respBody, nil
}

// 包装的通用请求和解析
func (c *Client) doRequest(reqObj, respObj interface{}, addr string) error {
	respBody, err := c.requestApi(reqObj, addr)
	if err != nil {
		return err
	}
	return c.decodeRespObj(respBody, respObj)
}

// 解析通用响应结构体 并且验证网关和业务状态码
func (c *Client) decodeRespObj(respBody []byte, respObj interface{}) error {
	// 解析应答结构体
	if err := json.Unmarshal(respBody, respObj); err != nil {
		return err
	}

	// 解析通用应答字段，判断网关状态码和业务状态码
	commonResp := new(CommonResponseParams)
	if err := json.Unmarshal(respBody, commonResp); err != nil {
		return err
	}

	// 判断网关状态码
	if commonResp.ErrCode != GateWaySuccess {
		return c.GetGateWayError(commonResp.ErrCode, commonResp.ErrInfo)
	}

	return nil
}

// 聚合反扫（商家扫用户）
func (c *Client) PrepayMchScanUser(req *MchScanRequest) (respObj *MchScanResponse, err error) {
	respObj = new(MchScanResponse)

	if c.isProEnv {
		err = c.doRequest(req, respObj, MchScanApi)
	} else {
		err = c.doRequest(req, respObj, MchScanApiBeta)
	}

	return respObj, err
}

// 支付查询
func (c *Client) QueryOrder(req *QueryRequest) (respObj *QueryResponse, err error) {
	respObj = new(QueryResponse)

	if c.isProEnv {
		err = c.doRequest(req, respObj, QueryApi)
	} else {
		err = c.doRequest(req, respObj, QueryApiBeta)
	}

	return respObj, err
}

// 交易退款
func (c *Client) Refund(req *RefundRequest) (respObj *RefundResponse, err error) {
	respObj = new(RefundResponse)

	if c.isProEnv {
		err = c.doRequest(req, respObj, RefundApi)
	} else {
		err = c.doRequest(req, respObj, RefundApiBeta)
	}

	return respObj, err
}

// 交易退款
func (c *Client) Cancel(req *CancelRequest) (respObj *CancelResponse, err error) {
	respObj = new(CancelResponse)

	if c.isProEnv {
		err = c.doRequest(req, respObj, CancelApi)
	} else {
		err = c.doRequest(req, respObj, CancelApiBeta)
	}

	return respObj, err
}

// 交易退款查询
func (c *Client) QueryRefund(req *QueryRefundRequest) (respObj *QueryRefundResponse, err error) {
	respObj = new(QueryRefundResponse)

	if c.isProEnv {
		err = c.doRequest(req, respObj, QueryRefundApi)
	} else {
		err = c.doRequest(req, respObj, QueryRefundApiBeta)
	}

	return respObj, err
}
