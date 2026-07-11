package cmbpay

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tjfoc/gmsm/sm2"
)

// Config 保存初始化 Client 所需的商户配置。
type Config struct {
	// Host 服务器地址，如 HostProd 或 HostUAT。
	Host string
	// AppID 聚合收单平台分配的 APP ID（报文头身份标识，详见接口文档 2.4.3）。
	AppID string
	// AppSecret 聚合收单平台分配的 APP SECRET（用于报文头加签，不上送）。
	AppSecret string
	// MerID 招行商户号。
	MerID string
	// PrivateKeyHex 商户 SM2 私钥（64 字符 HEX），用于报文体加签。
	PrivateKeyHex string
	// CMBPublicKeyBase64 招行 SM2 公钥（Base64 ASN.1），用于报文体验签。
	CMBPublicKeyBase64 string
	// HTTPClient 可选的自定义 HTTP 客户端；为 nil 时使用带 30s 超时的默认客户端。
	HTTPClient *http.Client
	// 是否开启生产环境
	IsTest bool
}

var (
	cmbClient *Client
	cmbMu     sync.Mutex
)

// Client 是招行聚合支付接口的客户端。它是并发安全的，可在多个 goroutine 间共享。
type Client struct {
	cfg    Config
	priv   *sm2.PrivateKey
	cmbPub *sm2.PublicKey
	http   *http.Client
}

// NewClient 依据配置创建 Client，并完成密钥加载与校验。
// 每次调用都会返回一个独立实例，调用方可自行管理生命周期。
func NewClient(cfg Config) (*Client, error) {
	if cfg.Host == "" {
		return nil, fmt.Errorf("cmbpay: Host 不能为空")
	}
	if cfg.AppID == "" || cfg.AppSecret == "" {
		return nil, fmt.Errorf("cmbpay: AppID/AppSecret 不能为空")
	}
	priv, err := LoadPrivateKeyHex(cfg.PrivateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 加载商户私钥失败: %w", err)
	}
	pub, err := LoadPublicKeyBase64(cfg.CMBPublicKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 加载招行公钥失败: %w", err)
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	}
	return &Client{cfg: cfg, priv: priv, cmbPub: pub, http: cfg.HTTPClient}, nil
}

// New 返回全局单例 Client。首次调用时初始化，后续调用直接复用。
// 若初始化失败，错误会被返回给调用方，且下次调用会重新尝试初始化。
func New(cfg Config) (*Client, error) {
	if cmbClient != nil {
		return cmbClient, nil
	}
	cmbMu.Lock()
	defer cmbMu.Unlock()
	if cmbClient != nil {
		return cmbClient, nil
	}
	c, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}
	cmbClient = c
	return cmbClient, nil
}

// MerID 返回配置中的商户号，便于装配业务请求时复用。
func (c *Client) MerID() string { return c.cfg.MerID }

// requestEnvelope 是发送给招行的外层报文（详见接口文档 2.2、附录 2）。
type requestEnvelope struct {
	BizContent string `json:"biz_content"`
	Encoding   string `json:"encoding"`
	SignMethod string `json:"signMethod"`
	Version    string `json:"version"`
	EncryptKey string `json:"encryptKey,omitempty"`
	Sign       string `json:"sign"`
}

// Execute 调用指定接口，将 bizContent 序列化为 biz_content 后完成加签、发送、
// 验签，并把成功返回的完整报文反序列化到 out。
//
//   - path 为接口路径常量，如 PathQrCodeApply；
//   - bizContent 为业务请求（结构体或 map），会被 JSON 序列化进 biz_content；
//   - out 为业务响应的接收者指针，需为嵌入 *PubResp 的结构体指针。
//
// 若接口返回 returnCode 或 respCode 为 FAIL，返回 *APIError。
func (c *Client) Execute(ctx context.Context, path string, bizContent any, out any) error {
	return c.execute(ctx, path, bizContent, "", out)
}

// ExecuteEncrypted 与 Execute 类似，但会带上由 enc 生成的数字信封（encryptKey），
// 适用于请求中含有需 SM4 加密的敏感字段的场景（详见接口文档 2.4.4）。
// 敏感字段的密文需由调用方通过 enc.Encrypt 生成后填入 bizContent。
func (c *Client) ExecuteEncrypted(ctx context.Context, path string, bizContent any, enc *Encryptor, out any) error {
	if enc == nil {
		return fmt.Errorf("cmbpay: Encryptor 不能为空")
	}
	env, err := enc.envelope(c.cmbPub)
	if err != nil {
		return err
	}
	return c.execute(ctx, path, bizContent, env, out)
}

func (c *Client) execute(ctx context.Context, path string, bizContent any, encryptKey string, out any) error {
	bizJSON, err := json.Marshal(bizContent)
	if err != nil {
		return fmt.Errorf("cmbpay: biz_content 序列化失败: %w", err)
	}

	// 装配参与签名的参数并计算报文体签名（详见接口文档 2.4.1.1）。
	params := map[string]string{
		"biz_content": string(bizJSON),
		"encoding":    Encoding,
		"signMethod":  SignMethodSM2,
		"version":     Version,
	}
	if encryptKey != "" {
		params["encryptKey"] = encryptKey
	}
	sign, err := signSM2(c.priv, buildSignString(params))
	if err != nil {
		return err
	}

	env := requestEnvelope{
		BizContent: string(bizJSON),
		Encoding:   Encoding,
		SignMethod: SignMethodSM2,
		Version:    Version,
		EncryptKey: encryptKey,
		Sign:       sign,
	}
	body, err := json.Marshal(env)
	if err != nil {
		return fmt.Errorf("cmbpay: 报文序列化失败: %w", err)
	}

	if c.cfg.IsTest {
		log.Println("请求URL：", c.cfg.Host+path)
		log.Println("请求原始：", string(body))
	}

	respBody, err := c.doHTTP(ctx, path, sign, body)
	if c.cfg.IsTest {
		log.Println("请求响应原始：", string(respBody))
		log.Println("请求响应错误：", err)
	}
	if err != nil {
		return err
	}
	return c.handleResponse(respBody, out)
}

// doHTTP 发送 HTTP POST 请求，并填充 APP ID 校验所需的报文头（详见接口文档 2.4.3）。
func (c *Client) doHTTP(ctx context.Context, path, bodySign string, body []byte) ([]byte, error) {
	url := strings.TrimRight(c.cfg.Host, "/") + path
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 构造请求失败: %w", err)
	}
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("appid", c.cfg.AppID)
	req.Header.Set("timestamp", ts)
	req.Header.Set("apisign", headerSign(c.cfg.AppID, c.cfg.AppSecret, bodySign, ts))

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 请求发送失败: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 读取响应失败: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cmbpay: HTTP 状态码 %d，响应: %s", resp.StatusCode, truncate(string(data), 512))
	}
	return data, nil
}

// handleResponse 验签并解析同步返回报文，将完整响应反序列化到 out（详见接口文档 2.4.1.2 ①、4.x 返回参数）。
func (c *Client) handleResponse(body []byte, out any) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(body, &fields); err != nil {
		return fmt.Errorf("cmbpay: 响应报文解析失败: %w，原文: %s", err, truncate(string(body), 512))
	}

	// 提取签名并校验（招行公钥、SM2withSM3）。
	sign, err := rawString(fields["sign"])
	if err != nil {
		return fmt.Errorf("cmbpay: 响应缺少 sign 字段: %w", err)
	}
	params, err := responseSignParams(fields)
	if err != nil {
		return err
	}
	if err := verifySM2(c.cmbPub, buildSignString(params), sign); err != nil {
		return err
	}

	// 先判断通信层，再判断业务层（详见接口文档 2.2 第 8 条）。
	returnCode, _ := rawString(fields["returnCode"])
	if returnCode != ResultSuccess {
		return newAPIError(fields)
	}
	respCode, _ := rawString(fields["respCode"])
	if respCode == ResultFail {
		return newAPIError(fields)
	}

	if out == nil {
		return nil
	}

	// 招行 API 的 biz_content 可能以 JSON 字符串（而非 JSON 对象）形式返回，
	// 需要先将其展开为原始 JSON 再反序列化到目标结构体。
	if raw, ok := fields["biz_content"]; ok && len(raw) > 0 && raw[0] == '"' {
		var bizStr string
		if err := json.Unmarshal(raw, &bizStr); err == nil {
			fields["biz_content"] = json.RawMessage(bizStr)
		}
	}

	// 将修正后的字段重新组装并反序列化到输出结构体（需嵌入 *PubResp）。
	fixed, err := json.Marshal(fields)
	if err != nil {
		return fmt.Errorf("cmbpay: 响应重组失败: %w", err)
	}
	if err := json.Unmarshal(fixed, out); err != nil {
		return fmt.Errorf("cmbpay: 响应反序列化失败: %w", err)
	}
	return nil
}

// responseSignParams 从响应字段中构造待验签参数集：剔除 sign，其余字段的字符串值
// 参与验签。使用 rawString 反序列化后的原始字符串值（详见接口文档 2.4.1.2）。
func responseSignParams(fields map[string]json.RawMessage) (map[string]string, error) {
	params := make(map[string]string, len(fields))
	for k, raw := range fields {
		if k == "sign" {
			continue
		}
		v, err := rawString(raw)
		if err != nil {
			return nil, fmt.Errorf("cmbpay: 响应字段 %s 非字符串: %w", k, err)
		}
		params[k] = v
	}
	return params, nil
}

// newAPIError 从响应字段装配 *APIError。
func newAPIError(fields map[string]json.RawMessage) error {
	returnCode, _ := rawString(fields["returnCode"])
	respCode, _ := rawString(fields["respCode"])
	errCode, _ := rawString(fields["errCode"])
	respMsg, _ := rawString(fields["respMsg"])
	errDescription, _ := rawString(fields["errDescription"])
	return &APIError{
		ReturnCode:     returnCode,
		RespCode:       respCode,
		ErrCode:        errCode,
		RespMsg:        respMsg,
		ErrDescription: errDescription,
	}
}

// rawString 将 json.RawMessage 解析为字符串。
func rawString(raw json.RawMessage) (string, error) {
	if len(raw) == 0 {
		return "", fmt.Errorf("空值")
	}
	var s string
	if err := json.Unmarshal(raw, &s); err != nil {
		return "", err
	}
	return s, nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// Encryptor 承载一次请求中用于敏感字段加密的对称密钥（SM4）。
// 同一次请求的多个敏感字段应使用同一个 Encryptor（详见接口文档 2.4.4）。
type Encryptor struct {
	sm4Key []byte
}

// NewEncryptor 生成一个持有随机 16 字节 SM4 密钥的 Encryptor。
func NewEncryptor() (*Encryptor, error) {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("cmbpay: 生成 SM4 密钥失败: %w", err)
	}
	return &Encryptor{sm4Key: key}, nil
}

// NewEncryptorWithKey 使用调用方指定的 16 字节 SM4 密钥创建 Encryptor。
func NewEncryptorWithKey(sm4Key []byte) (*Encryptor, error) {
	if len(sm4Key) != 16 {
		return nil, fmt.Errorf("cmbpay: SM4 密钥长度应为 16 字节，实际 %d 字节", len(sm4Key))
	}
	k := make([]byte, 16)
	copy(k, sm4Key)
	return &Encryptor{sm4Key: k}, nil
}

// Encrypt 对敏感字段明文做 SM4 加密并 Base64 编码，返回填入 biz_content 的密文。
func (e *Encryptor) Encrypt(plaintext string) (string, error) {
	return encryptSM4(e.sm4Key, []byte(plaintext))
}

// envelope 使用招行公钥生成 encryptKey 数字信封。
func (e *Encryptor) envelope(pub *sm2.PublicKey) (string, error) {
	return makeEnvelope(pub, e.sm4Key)
}
