# cmbpay — 招商银行商户聚合支付 Go SDK

基于《招商银行商户聚合支付接口 V3.3.6》实现的 Go SDK，封装了报文组装、国密
（SM2/SM3/SM4）加验签、敏感信息加密数字信封、报文头 APP ID 校验、HTTP 调用与
异步通知验签等全部协议细节，让业务代码只需关注接口参数本身。

## 目录

- [特性](#特性)
- [安装](#安装)
- [快速开始](#快速开始)
- [配置说明](#配置说明)
- [示例程序](#示例程序)
- [完整支付闭环](#完整支付闭环)
- [敏感信息加密](#敏感信息加密)
- [异步通知](#异步通知)
- [调用未封装的接口](#调用未封装的接口)
- [错误处理](#错误处理)
- [已封装的类型化接口](#已封装的类型化接口)
- [注意事项](#注意事项)
- [常见问题](#常见问题)
- [测试](#测试)

## 特性

- **协议封装**：外层报文（`biz_content` / `encoding` / `signMethod` / `version` /
  `encryptKey` / `sign`）自动组装与解析（接口文档 2.2）。
- **国密加验签**：SM2withSM3 裸签名（USER_ID `1234567812345678`），请求加签、
  同步响应验签、异步通知验签（接口文档 2.4.1）。
- **敏感信息加密**：SM4 加密字段 + SM2 数字信封（`encryptKey`），一次请求多字段共用
  同一对称密钥（接口文档 2.4.4）。
- **报文头校验**：`appid` / `timestamp` / `apisign`（MD5）自动填充（接口文档 2.4.3）。
- **类型化接口**：收款码申请、支付查询、退款、退款查询、关单、付款码收款、撤销、
  微信统一下单等常用接口提供强类型方法；其余接口可用通用 `Execute` 调用。
- **异步通知**：`ParseNotify` 完成 form-data 验签并结构化，`NotifySuccessBody` /
  `NotifyFailBody` 生成标准应答（接口文档 2.5、4.3）。
- **错误处理**：`APIError` 区分通信层（`returnCode`）与业务层（`respCode`）错误，
  并提供 `IsSystemError()` 辅助识别需查询确认的 `SYSTERM_ERROR`（接口文档 2.2）。
- **并发安全**：`Client` 可在多个 goroutine 间共享。

## 安装

```bash
go get github.com/go-pay/gopay/cmbpay
```

要求 Go 1.18+（使用了泛型 `any` 等）。依赖国密库 `github.com/tjfoc/gmsm`，
由 Go Modules 按需拉取（第三方公共包不纳入版本库，`vendor/` 已在 `.gitignore` 中忽略）。
如需离线构建可自行执行 `go mod vendor`。

## 快速开始

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-pay/gopay/cmbpay"
)

func main() {
	client, err := cmbpay.NewClient(cmbpay.Config{
		Host:               cmbpay.HostProd, // 联调环境用 cmbpay.HostUAT
		AppID:              "<APP ID>",
		AppSecret:          "<APP SECRET>",
		MerID:              "<商户号>",
		PrivateKeyHex:      "<商户 SM2 私钥 HEX，64 字符>",
		CMBPublicKeyBase64: "<招行 SM2 公钥 Base64>",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 收款码申请（金额单位：分）
	resp, err := client.QrCodeApply(context.Background(), &cmbpay.QrCodeApplyReq{
		MerID:     client.MerID(),
		OrderID:   "your-unique-order-id", // 商户下唯一
		UserID:    "cashier-id",
		NotifyURL: "https://your.domain/cmb/notify",
		TxnAmt:    "1", // 1 分
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("二维码:", resp.QrCode, "平台订单号:", resp.CmbOrderID)
}
```

## 配置说明

`cmbpay.Config` 字段：

| 字段 | 必填 | 说明 |
|------|:---:|------|
| `Host` | 是 | 服务器地址。生产 `cmbpay.HostProd`，联调 `cmbpay.HostUAT` |
| `AppID` | 是 | 聚合收单平台分配的 APP ID（报文头身份标识） |
| `AppSecret` | 是 | 聚合收单平台分配的 APP SECRET（用于报文头加签，不上送） |
| `MerID` | 是 | 招行商户号 |
| `PrivateKeyHex` | 是 | 商户 SM2 私钥，64 字符 HEX，用于报文体加签 |
| `CMBPublicKeyBase64` | 是 | 招行 SM2 公钥，Base64（ASN.1），用于报文体/通知验签 |
| `HTTPClient` | 否 | 自定义 `*http.Client`；为 nil 时使用带 30s 超时的默认客户端 |

> 联调环境的示例 AppID / AppSecret / 商户私钥 / 招行公钥见接口文档附录 1，
> 本仓库示例在未设置环境变量时会回退到这些参数，可直接体验。

## 示例程序

`example/` 目录提供了 3 个可独立运行的示例，均通过 `example/config` 读取配置
（优先环境变量，回退到联调示例参数）：

```bash
# 完整支付闭环：收款码申请 → 轮询查询 → 关单/退款 → 退款查询 + 付款码加密收款
go run ./example/payment

# 异步通知接收服务（验签 + 幂等 + 标准应答），监听 :8080
go run ./example/notifyserver

# 通用方式调用未封装的接口（银联云闪付、数字人民币统一下单等）
go run ./example/generic
```

指定真实商户参数运行：

```bash
export CMB_ENV=prod
export CMB_APPID=... CMB_APPSECRET=... CMB_MERID=...
export CMB_PRIVKEY=...   # 64 字符 HEX
export CMB_CMBPUBKEY=... # Base64
export CMB_NOTIFYURL=https://your.domain/cmb/notify
go run ./example/payment
```

## 完整支付闭环

```go
ctx := context.Background()
const cashier = "N003109945"
orderID := "your-unique-order-id"

// 1. 收款码申请
qr, err := client.QrCodeApply(ctx, &cmbpay.QrCodeApplyReq{
	MerID: client.MerID(), OrderID: orderID, UserID: cashier,
	NotifyURL: "https://your.domain/cmb/notify", TxnAmt: "1",
})

// 2. 支付结果查询（建议 15s 后首查，之后每 5-10s 一次；遇 SYSTERM_ERROR 继续查）
st, err := client.OrderQuery(ctx, &cmbpay.OrderQueryReq{
	MerID: client.MerID(), OrderID: orderID, UserID: cashier,
})
switch st.TradeState {
case cmbpay.TradeStateSuccess: // "S" 交易成功
case cmbpay.TradeStateProcess: // "P" 进行中，继续轮询
	// ...
}

// 3. 未支付则关单，可用新订单号重新发起
client.Close(ctx, &cmbpay.CloseReq{MerID: client.MerID(), OrderID: orderID, UserID: cashier})

// 4. 已支付则退款 + 退款查询
client.Refund(ctx, &cmbpay.RefundReq{
	MerID: client.MerID(), OrderID: "refund-unique-id",
	CmbOrderID: st.CmbOrderID, UserID: cashier, RefundAmt: "1",
})
client.RefundQuery(ctx, &cmbpay.RefundQueryReq{
	MerID: client.MerID(), OrderID: "refund-unique-id", UserID: cashier,
})
```

完整可运行版本见 [`example/payment/main.go`](example/payment/main.go)。

## 敏感信息加密

对 `terminalInfo`、实名支付信息等需加密的字段，先用 `Encryptor` 加密，再用带
`...Encrypted` 后缀的方法（或 `ExecuteEncrypted`）发送，SDK 会自动生成数字信封
（`encryptKey`）。同一次请求的多个敏感字段共用同一个 `Encryptor`：

```go
enc, _ := cmbpay.NewEncryptor() // 生成随机 SM4 密钥
cipher, _ := enc.Encrypt(`{"location":"+37.12/-121.213","mobile_country_cd":"460"}`)

resp, err := client.PayEncrypted(ctx, &cmbpay.PayReq{
	MerID:               client.MerID(),
	OrderID:             "order-id",
	UserID:              "cashier-id",
	AuthCode:            "134650123456789012",
	TxnAmt:              "1",
	EncryptTerminalInfo: cipher, // 填入密文
}, enc)
```

## 异步通知

招行以 form-data POST 推送支付/退款结果通知。处理要点：**验签 → 幂等判重 →
更新业务 → 返回标准应答**。

```go
func notifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := client.ParseNotify(r) // 自动验签并结构化
	if err != nil {
		w.Write(cmbpay.NotifyFailBody("verify failed")) // 返回 FAIL，招行会重试
		return
	}

	// 幂等：同一通知可能重复投递，处理前先查本地状态
	if alreadyProcessed(data.OrderID) {
		w.Write(cmbpay.NotifySuccessBody())
		return
	}

	if data.IsPaySuccess() {
		// 更新本地订单为已支付、发货等
	}

	w.Write(cmbpay.NotifySuccessBody()) // {"returnCode":"SUCCESS","respCode":"SUCCESS"}
}
```

未使用 `net/http` 时可用 `client.ParseNotifyValues(url.Values)`。完整可运行版本见
[`example/notifyserver/main.go`](example/notifyserver/main.go)。

## 调用未封装的接口

SDK 为常用接口提供了类型化方法；其余接口（数字人民币、微信支付分、支付宝 APP/WAP、
银联云闪付等，路径常量见 [`endpoints.go`](endpoints.go)）可用通用方法调用，请求体可用
`map` 或自定义结构体：

```go
// 请求体用结构体，响应解析到结构体
type ecnyResp struct {
	CmbOrderID string `json:"cmbOrderId"`
	QrCode     string `json:"qrCode"`
}
var out ecnyResp
err := client.Execute(ctx, cmbpay.PathEcnyUnifiedOrder, map[string]string{
	"merId":     client.MerID(),
	"orderId":   "unique-id",
	"userId":    "cashier-id",
	"notifyUrl": "https://your.domain/cmb/notify",
	"txnAmt":    "1",
}, &out)

// 含敏感字段时用 ExecuteEncrypted
err = client.ExecuteEncrypted(ctx, cmbpay.PathXxx, reqWithCipher, enc, &out)
```

完整可运行版本见 [`example/generic/main.go`](example/generic/main.go)。

## 错误处理

招行返回分通信层（`returnCode`）与业务层（`respCode`）两级。当任一为 FAIL 时，
SDK 返回 `*cmbpay.APIError`：

```go
resp, err := client.OrderQuery(ctx, req)
if err != nil {
	var apiErr *cmbpay.APIError
	if errors.As(err, &apiErr) {
		switch {
		case apiErr.IsSystemError():
			// errCode=SYSTERM_ERROR：交易结果不明确，需继续发起查询确认（接口文档 2.2）
		case apiErr.IsCommFail():
			// returnCode=FAIL：报文不合规范（字段超长、非法字符、签名错误等）
		default:
			// 业务失败，查看 apiErr.ErrCode / apiErr.RespMsg
		}
	} else {
		// 网络错误、验签失败、序列化错误等
	}
}
```

## 已封装的类型化接口

带 `Encrypted` 后缀的方法用于请求含需加密敏感字段（如 `encryptTerminalInfo`、
`encryptIdentity`）的场景，会自动生成数字信封。

| 方法 | 说明 | 文档章节 |
|------|------|----------|
| `QrCodeApply` / `QrCodeApplyEncrypted` | 收款码申请 | 4.1 |
| `OrderQuery` | 支付结果查询 | 4.2 |
| `Refund` | 退款申请 | 4.4 |
| `RefundQuery` | 退款结果查询 | 4.5 |
| `Close` | 关闭订单 | 4.7 |
| `Pay` / `PayEncrypted` | 付款码收款 | 4.8 |
| `Cancel` | 付款码支付撤销 | 4.9 |
| `OnlinePay` | 微信统一下单 | 4.10 |
| `ServPay` / `ServPayEncrypted` | 服务窗支付 | 4.11 |
| `ZfbNative` / `ZfbNativeEncrypted` | 支付宝 native 码支付 | 4.12 |
| `BillRecord` | 对账单下载地址获取，和下单/查单接口签名规则不一样 | - |
| `OrderQrCodeApply` / `OrderQrCodeApplyEncrypted` | 订单二维码申请 | 4.14 |
| `MiniAppOrder` / `MiniAppOrderEncrypted` | 微信小程序下单 | 4.15 |
| `CloudPay` | 银联云闪付 | 4.16 |
| `EcnyUnifiedOrder` | 数字人民币统一下单 | 4.17 |
| `EcnyUnifiedPayment` | 数字人民币统一支付 | 4.18 |
| `EcnySubwalletPay` | 数字人民币子钱包支付 | 4.19 |
| `EcnyContractSubwalletPay` | 数字人民币子钱包支付-带合约 | 4.20 |
| `EcnyContractUnifiedOrder` | 数字人民币统一下单-带合约 | 4.21 |
| `PapQuery` | 微信委托代扣查询 | 4.23 |
| `ZfbApp` / `ZfbAppEncrypted` | 支付宝 APP 支付 | 4.24 |
| `ZfbWap` / `ZfbWapEncrypted` | 支付宝手机网站支付 | 4.25 |
| `OpenIDQuery` | 微信授权码查询 openid | 4.37 |

其余接口（微信委托代扣下单、微信支付分系列、智能合约分账、先享后付等）路径见
[`endpoints.go`](endpoints.go)，通过 `Execute` / `ExecuteEncrypted` 调用。

## 注意事项

- **金额单位为分**，字符串传输，不含小数点（接口文档 2.3）。
- **商户订单号需全局唯一**；重发同一笔支付请使用原订单号避免重复支付；已支付/已关单/
  已撤销的订单号不能重新发起支付。
- **必须实现查询机制**：网络异常导致未收到通知时，商户须主动查询直至结果明确，
  否则风险自负（接口文档 2.2、4.2）。
- **`SYSTERM_ERROR` 不等于失败**：支付/退款/查询遇到该错误时交易结果不明确，
  应继续查询确认，而非直接判失败（接口文档 2.2 第 9、10 条）。
- **通知需幂等**：同一通知可能重复投递，处理前请加锁并检查业务状态；处理成功后
  必须返回 `SUCCESS`，否则招行按 `0/15/15/30/180/1800...` 秒的策略重试（接口文档 2.5）。
- **回调地址需提前报备**：招行到商户回调地址的网络开通约需 2-3 个工作日。
- **测试环境二维码**需将 `https://qr.95516.com` 替换为
  `http://payment-uat.cs.cmburl.cn` 再生成图片（接口文档附录 1）。

## 常见问题

**Q：验签失败怎么排查？**
A：确认 `CMBPublicKeyBase64` 是招行公钥（用于验签）而非商户自己的公钥；确认私钥为
32 字节、64 字符 HEX。同步响应验签中，SDK 已按接口文档 2.4.1.2 对 `/` 做 `\/` 转义。

**Q：如何切换联调 / 生产环境？**
A：仅需改 `Config.Host` 为 `cmbpay.HostUAT` 或 `cmbpay.HostProd`，其余不变。

**Q：请求超时时间如何设置？**
A：两种方式——传入自定义 `Config.HTTPClient`，或在调用方法时传入带 `context.WithTimeout`
的 `ctx`（推荐，可按请求粒度控制）。

**Q：SDK 覆盖了全部接口吗？**
A：常用接口提供类型化方法，全部接口路径（含数字人民币、微信支付分、支付宝系列等）
均在 `endpoints.go` 中，可用 `Execute` / `ExecuteEncrypted` 调用。

## 测试

```bash
go test ./...
```