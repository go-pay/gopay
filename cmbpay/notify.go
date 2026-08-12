package cmbpay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// NotifyData 是支付/退款结果异步通知的业务数据（详见接口文档 4.3、4.6）。
// 招行以 form-data 形式推送，字段与支付结果查询大体一致。
// 支付通知和退款通知共用此结构体，未使用的字段保持零值。
type NotifyData struct {
	// 公共字段
	MerID        string `json:"merId"`
	OrderID      string `json:"orderId"`
	CmbOrderID   string `json:"cmbOrderId"`
	UserID       string `json:"userId,omitempty"`
	CurrencyCode string `json:"currencyCode"`
	PayType      string `json:"payType"`
	TxnTime      string `json:"txnTime"`
	EndDate      string `json:"endDate,omitempty"`
	EndTime      string `json:"endTime,omitempty"`
	MchReserved  string `json:"mchReserved,omitempty"`

	// 支付通知字段
	TxnAmt              string `json:"txnAmt,omitempty"`
	DscAmt              string `json:"dscAmt,omitempty"`
	CouponInfo          string `json:"couponInfo,omitempty"`
	OrderOrigAmt        string `json:"orderOrigAmt,omitempty"`
	OrderCouponAmt      string `json:"orderCouponAmt,omitempty"`
	PromotionDetail     string `json:"promotionDetail,omitempty"`
	OpenID              string `json:"openId,omitempty"`
	OriOpenID           string `json:"oriOpenId,omitempty"`
	PayBank             string `json:"payBank,omitempty"`
	ThirdOrderID        string `json:"thirdOrderId,omitempty"`
	BuyerLogonID        string `json:"buyerLogonId,omitempty"`
	PayChannel          string `json:"payChannel,omitempty"`
	TradeState          string `json:"tradeState,omitempty"`
	ContractResp        string `json:"contractResp,omitempty"`
	DebtorAgentID       string `json:"debtorAgentId,omitempty"`
	DebtorAgentName     string `json:"debtorAgentName,omitempty"`
	EcnyPromotionDetail string `json:"ecnyPromotionDetail,omitempty"`

	// 退款通知字段
	RefundAmt          string `json:"refundAmt,omitempty"`
	RefundDscAmt       string `json:"refundDscAmt,omitempty"`
	RefundOrigAmt      string `json:"refundOrigAmt,omitempty"`
	RefundCouponAmt    string `json:"refundCouponAmt,omitempty"`
	IssAddnData        string `json:"issAddnData,omitempty"`
	RefundDetail       string `json:"refundDetail,omitempty"`
	ContractRefundResp string `json:"contractRefundResp,omitempty"`
	RefundState        string `json:"refundState,omitempty"`

	// Raw 保存通知的全部原始字段（已 url_decode），便于读取本 SDK 未显式列出的扩展字段。
	Raw map[string]string `json:"-"`
}

// ParseNotify 解析并验签招行异步通知（详见接口文档 2.4.1.2 ② 异步通知验签）。
//
// 通知为 form-data 格式：除 sign 外的所有字段先做 url_decode，再字典排序拼接为
// 待验签字符串，使用招行公钥、SM2withSM3、base64 解码后的 sign 值进行验签。
// 校验通过后返回结构化的 NotifyData。
//
// 典型用法（HTTP handler 中）：
//
//	data, err := client.ParseNotify(r)
//	if err != nil {
//	    body, _ := client.NotifyFailBody("verify")
//	    w.Write(body)
//	    return
//	}
//	// ... 幂等处理业务 ...
//	body, err := client.NotifySuccessBody()
//	if err != nil { /* 加签失败，告警 */ return }
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(body)
func (c *Client) ParseNotify(r *http.Request) (*NotifyData, error) {
	if err := r.ParseForm(); err != nil {
		return nil, fmt.Errorf("cmbpay: 解析通知表单失败: %w", err)
	}
	form := make(map[string]string, len(r.PostForm))
	for k := range r.PostForm {
		form[k] = r.PostForm.Get(k)
	}
	return c.parseNotifyForm(form)
}

// ParseNotifyValues 从已解析的表单值（url.Values）验签并解析通知，
// 适用于非 net/http 场景（如网关已完成表单解析）。
func (c *Client) ParseNotifyValues(values url.Values) (*NotifyData, error) {
	form := make(map[string]string, len(values))
	for k := range values {
		form[k] = values.Get(k)
	}
	return c.parseNotifyForm(form)
}

// ParseNotifyBytes 从原始请求体字节验签并解析通知。
// 适用于需要先读取原始 body 做日志记录、再解析通知的场景。
//
// 典型用法（Gin handler 中）：
//
//	rawBody, _ := io.ReadAll(c.Request.Body)
//	log.Info("招行通知原始参数", "body="+string(rawBody))
//	data, err := client.ParseNotifyBytes(rawBody)
func (c *Client) ParseNotifyBytes(body []byte) (*NotifyData, error) {
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 解析通知表单失败: %w", err)
	}
	return c.ParseNotifyValues(values)
}

func (c *Client) parseNotifyForm(form map[string]string) (*NotifyData, error) {
	sign, ok := form["sign"]
	if !ok || sign == "" {
		return nil, fmt.Errorf("cmbpay: 通知缺少 sign 字段")
	}

	// 异步通知验签（接口文档 2.4.1.2 ②）：
	// 除 sign 外的所有字段（含 biz_content 整体）直接参与验签，
	// 值已由表单解析完成 url_decode，按字典排序拼接为待验签字符串。
	params := make(map[string]string, len(form))
	for k, v := range form {
		if k == "sign" {
			continue
		}
		params[k] = v
	}
	if err := verifySM2(c.cmbPub, buildSignString(params), sign); err != nil {
		return nil, err
	}

	// 验签通过后，展开 biz_content 信封以提取业务字段（接口文档 4.3）。
	if bizContent, hasEnvelope := params["biz_content"]; hasEnvelope {
		var bizFields map[string]string
		if err := json.Unmarshal([]byte(bizContent), &bizFields); err != nil {
			return nil, fmt.Errorf("cmbpay: biz_content JSON 解析失败: %w", err)
		}
		for k, v := range bizFields {
			params[k] = v
		}
		delete(params, "biz_content")
	}

	// 结构化映射：借助 JSON 往返把 map 填入 NotifyData。
	b, _ := json.Marshal(params)
	var data NotifyData
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, fmt.Errorf("cmbpay: 通知字段映射失败: %w", err)
	}
	data.Raw = params
	return &data, nil
}

// NotifySuccessBody 返回商户应答招行通知的成功报文（JSON），包含 version、encoding、
// signMethod 与 SM2 签名。商户处理成功或幂等判定已处理后须返回该报文，招行据此不再重复通知
// （详见接口文档 2.3 第 3 条、2.5 未知交易通知机制）。
func (c *Client) NotifySuccessBody() ([]byte, error) {
	return c.notifyBody(ResultSuccess, ResultSuccess, "")
}

// NotifyFailBody 返回商户应答招行通知的失败报文（JSON），包含 version、encoding、
// signMethod 与 SM2 签名。招行会按重试策略再次通知。
func (c *Client) NotifyFailBody(respMsg string) ([]byte, error) {
	return c.notifyBody(ResultSuccess, ResultFail, respMsg)
}

// notifyBody 构造商户应答报文：填充 version/encoding/signMethod 等固定字段，
// 对除 sign 外的全部字段做 SM2 签名并将签名值写入 sign。
//
// 加签失败时返回错误而非返回未加签的报文 —— 未加签的应答会被招行判定为无效，
// 商户应当据此告警，而不是让应答静默失效。
func (c *Client) notifyBody(returnCode, respCode, respMsg string) ([]byte, error) {
	m := map[string]string{
		"returnCode": returnCode,
		"respCode":   respCode,
		"version":    Version,
		"encoding":   Encoding,
		"signMethod": SignMethodSM2,
	}
	if respMsg != "" {
		m["respMsg"] = respMsg
	}
	sign, err := signSM2(c.priv, buildSignString(m))
	if err != nil {
		return nil, err
	}
	m["sign"] = sign
	b, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 应答报文序列化失败: %w", err)
	}
	return b, nil
}

// IsPaySuccess 报告该通知是否为支付成功（tradeState=S，详见接口文档 4.2.3）。
//
// 注意：仅当通知中明确带有 tradeState=S 时才返回 true。若通知未携带 tradeState
// （例如退款通知，或招行后续调整了通知字段），此处一律返回 false，商户应调用
// OrderQuery 查询订单以确认真实状态，避免把状态不明的通知误判为支付成功。
func (n *NotifyData) IsPaySuccess() bool {
	return strings.EqualFold(n.TradeState, TradeStateSuccess)
}
