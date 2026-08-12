package cmbpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-pay/gopay"
)

// 对账文件接口服务器地址（详见接口文档 - 聚合对账文件接口）
// 注意：对账文件接口的主机路径与普通支付接口不同
const (
	// BillCheckHostST 对账文件接口 - 联调环境
	BillCheckHostST = "https://api.cmburl.cn:8065/bill/chk"
	// BillCheckHostUAT 对账文件接口 - UAT 环境
	BillCheckHostUAT = "https://api.cmburl.cn:8065/bill/check"
	// BillCheckHostPRD 对账文件接口 - 生产环境
	BillCheckHostPRD = "https://api.cmbchina.com/bill/check"
)

// 对账文件接口路径
const (
	// PathBillRecord 对账文件下载接口
	PathBillRecord = "/adapter"
)

// 对账文件接口响应码
const (
	// BillRecordCodeSuccess 账单文件已生成
	BillRecordCodeSuccess = "SUC000000"
	// BillRecordCodeNotGenerated 账单还未生成
	BillRecordCodeNotGenerated1 = "PMS05V0"
	// BillRecordCodeNotGenerated2 账单还未生成（另一种）
	BillRecordCodeNotGenerated2 = "PMS0RV0"
	// BillRecordCodeNoBill 当日无账单
	BillRecordCodeNoBill1 = "PMS05V1"
	// BillRecordCodeNoBill2 当日无账单（另一种）
	BillRecordCodeNoBill2 = "PMS0RV1"
	// BillRecordCodeEmpty
	BillRecordCodeBillEmpty = "PMS0RT2"
)

// 清分状态
const (
	// SettleStatusAll 全部清分
	SettleStatusAll = "S"
	// SettleStatusNone 未清分
	SettleStatusNone = "F"
	// SettleStatusPartial 部分清分
	SettleStatusPartial = "T"
)

// BillRecordReq 对账文件下载请求参数
type BillRecordReq struct {
	MerchantNo string `json:"merchantNo"` // 商户号
	BillDate   string `json:"billDate"`   // 账单日期，格式：YYYY-MM-DD
	BillType   string `json:"billType"`   // 账单类型，如 JH_JZ
}

// BillRecordResp 对账文件下载响应
type BillRecordResp struct {
	Code      string          `json:"code"`      // 响应码
	Data      *BillRecordData `json:"data"`      // 响应数据
	Message   string          `json:"message"`   // 响应消息
	Timestamp string          `json:"timestamp"` // 时间戳
}

// BillRecordData 对账文件下载响应数据
type BillRecordData struct {
	MerchantNo   string `json:"merchantNo"`   // 商户号
	BillType     string `json:"billType"`     // 账单类型
	DownloadUrl  string `json:"downloadUrl"`  // 下载链接
	SettleStatus string `json:"settleStatus"` // 清分状态：S-全部清分，F-未清分，T-部分清分
	BillDate     string `json:"billDate"`     // 账单日期
}

// BillRecord 查询对账文件下载地址（聚合对账文件接口）。
//
// 该接口与普通支付接口的签名/验签方式不同：
//   - 请求体为 {"requestBody": {...}}，而非 biz_content 信封格式
//   - 不需要 SM2 报文体签名
//   - 不需要 SM2 响应验签
//   - 响应格式为 {code, data, message, timestamp}，而非 {returnCode, respCode, biz_content, sign}
//
// billCheckHost 为对账接口专用主机地址（如 BillCheckHostPRD）；传空串则取
// Config.BillCheckHost。两者皆为空时直接返回错误 —— 对账接口的主机地址已包含
// /bill/check 等路径，退化为支付接口的 Config.Host 只会拼出错误的 URL。
//
// 建议在每日 10:35 后查询。
func (c *Client) BillRecord(ctx context.Context, billCheckHost string, req *BillRecordReq) (*BillRecordResp, error) {
	// 确定主机地址
	host := billCheckHost
	if host == "" {
		host = c.cfg.BillCheckHost
	}
	if host == "" {
		return nil, fmt.Errorf("cmbpay: 对账文件接口主机地址为空，请传入 billCheckHost 或配置 Config.BillCheckHost")
	}
	if req == nil {
		return nil, fmt.Errorf("cmbpay: 对账文件下载请求参数不能为空")
	}
	url := strings.TrimRight(host, "/") + PathBillRecord

	// 构建请求体：{"requestBody": {...}}
	wrapper := struct {
		RequestBody *BillRecordReq `json:"requestBody"`
	}{RequestBody: req}

	bodyBytes, err := json.Marshal(wrapper)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 对账请求序列化失败: %w", err)
	}

	// 构建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 构造对账请求失败: %w", err)
	}

	// 设置请求头（对账接口专用 header 格式）
	ts := strconv.FormatInt(time.Now().Unix(), 10)

	// 计算 sign：对请求体进行 SM2withSM3 签名（Base64 格式）
	signValue, err := signSM2(c.priv, string(bodyBytes))
	if err != nil {
		return nil, err
	}

	// 计算 apisign：对 "appid=value&secret=value&sign=value&timestamp=value" 拼接字符串进行 SM2withSM3 签名
	// 使用 HEX 编码的 R||S 格式（64字节），对应 Java 的 signHexBySm3WithSm2 方法
	signStr := fmt.Sprintf("appid=%s&secret=%s&sign=%s&timestamp=%s",
		c.cfg.AppID, c.cfg.AppSecret, signValue, ts)
	apisign, err := signSM2Hex(c.priv, signStr)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("timestamp", ts)
	httpReq.Header.Set("appid", c.cfg.AppID)
	httpReq.Header.Set("apisign", apisign) // SM2withSM3 签名（HEX 编码的 R||S 格式）
	httpReq.Header.Set("sign", signValue)  // SM2withSM3 签名（Base64 格式）
	httpReq.Header.Set("verify", "SM3withSM2")
	httpReq.Header.Set("channel", "AP")
	httpReq.Header.Set("funcCode", "BILLRECORD_GET_FORAPI")
	httpReq.Header.Set("sysCode", "AP")

	if c.cfg.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("CMBPay_BillCheck_Url: %s", url)
		c.logger.Debugf("CMBPay_BillCheck_Request: %s", bodyBytes)
	}

	// 发送请求（复用 Client 的 HTTP 客户端）
	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 对账请求发送失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 读取对账响应失败: %w", err)
	}
	if c.cfg.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("CMBPay_BillCheck_Response: %s", respBody)
	}
	// 非 200 时响应体通常不是约定的 JSON 结构，需先拦截，避免把网关错误页
	// 解析成一个 Code 为空的“成功”响应。
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cmbpay: 对账接口 HTTP 状态码 %d，响应: %s", resp.StatusCode, truncate(string(respBody), 512))
	}

	// 解析响应（对账接口响应格式与普通接口不同，直接解析）
	var billResp BillRecordResp
	if err = json.Unmarshal(respBody, &billResp); err != nil {
		return nil, fmt.Errorf("cmbpay: 对账响应解析失败: %w，原文: %s", err, truncate(string(respBody), 512))
	}

	// 检查响应码
	switch billResp.Code {
	case BillRecordCodeSuccess:
		return &billResp, nil
	case BillRecordCodeBillEmpty:
		return nil, fmt.Errorf("cmbpay: 账单文件为空: %s - %s", billResp.Code, billResp.Message)
	case BillRecordCodeNotGenerated1, BillRecordCodeNotGenerated2:
		return nil, fmt.Errorf("cmbpay: 账单还未生成: %s - %s", billResp.Code, billResp.Message)
	case BillRecordCodeNoBill1, BillRecordCodeNoBill2:
		return nil, fmt.Errorf("cmbpay: 当日无账单: %s - %s", billResp.Code, billResp.Message)
	default:
		return nil, fmt.Errorf("cmbpay: 对账文件查询失败: %s - %s", billResp.Code, billResp.Message)
	}
}

// IsBillReady 检查账单是否已生成（BillRecordCodeSuccess）
func (resp *BillRecordResp) IsBillReady() bool {
	return resp.Code == BillRecordCodeSuccess
}

// IsBillNotGenerated 检查账单是否还未生成
func (resp *BillRecordResp) IsBillNotGenerated() bool {
	return resp.Code == BillRecordCodeNotGenerated1 || resp.Code == BillRecordCodeNotGenerated2
}

// IsNoBill 检查当日是否无账单
func (resp *BillRecordResp) IsNoBill() bool {
	return resp.Code == BillRecordCodeNoBill1 || resp.Code == BillRecordCodeNoBill2
}

// GetSettleStatusDesc 获取清分状态描述
func (d *BillRecordData) GetSettleStatusDesc() string {
	switch d.SettleStatus {
	case SettleStatusAll:
		return "全部清分"
	case SettleStatusNone:
		return "未清分"
	case SettleStatusPartial:
		return "部分清分"
	default:
		return "未知状态"
	}
}
