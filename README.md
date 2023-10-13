<div align=center><img width="240" height="240" alt="Logo was Loading Faild!" src="logo.png"/></div>

# GoPay

### 微信、支付宝、QQ、通联支付、拉卡拉、PayPal 的 Golang 版本SDK

[![Github](https://img.shields.io/github/followers/iGoogle-ink?label=Follow&style=social)](https://github.com/iGoogle-ink)
[![Github](https://img.shields.io/github/forks/go-pay/gopay?label=Fork&style=social)](https://github.com/go-pay/gopay/fork)

[![Golang](https://img.shields.io/badge/golang-1.18-brightgreen.svg)](https://golang.google.cn)
[![GoDoc](https://img.shields.io/badge/doc-pkg.go.dev-informational.svg)](https://pkg.go.dev/github.com/go-pay/gopay)
[![Go](https://github.com/go-pay/gopay/actions/workflows/go.yml/badge.svg)](https://github.com/go-pay/gopay/actions/workflows/go.yml)
[![GitHub Release](https://img.shields.io/github/v/release/go-pay/gopay)](https://github.com/go-pay/gopay/releases)
[![License](https://img.shields.io/github/license/go-pay/gopay)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-pay/gopay)](https://goreportcard.com/report/github.com/go-pay/gopay)

---

# 一、安装

```bash
go get -u github.com/go-pay/gopay
```

#### 查看 GoPay 版本

  [版本更新记录](https://github.com/go-pay/gopay/blob/main/release_note.txt)

```go
import (
    "github.com/go-pay/gopay"
    "github.com/go-pay/gopay/pkg/xlog"
)

func main() {
    xlog.Info("GoPay Version: ", gopay.Version)
}
```

---

<br>

# 二、文档目录

> ### 点击查看不同支付方式的使用文档。方便的话，请留下您认可的小星星，十分感谢！

* #### [支付宝支付](https://github.com/go-pay/gopay/blob/main/doc/alipay.md)
* #### [微信支付](https://github.com/go-pay/gopay/blob/main/doc/wechat_v3.md)
* #### [QQ支付](https://github.com/go-pay/gopay/blob/main/doc/qq.md)
* #### [通联支付](https://github.com/go-pay/gopay/blob/main/doc/allinpay.md)
* #### [拉卡拉支付](https://github.com/go-pay/gopay/blob/main/doc/lakala.md)
* #### [Paypal支付](https://github.com/go-pay/gopay/blob/main/doc/paypal.md)
* #### [Apple支付校验](https://github.com/go-pay/gopay/blob/main/doc/apple.md)
* #### [扫呗支付](https://github.com/go-pay/gopay/blob/main/doc/saobei.md)

---

<br>

# 三、其他说明

* 如需自定义Log输出，请调用以下方法设置自定义Logger，实现 `xlog.XLogger` 接口即可。
    * `xlog.SetDebugLog()`
    * `xlog.SetInfoLog()`
    * `xlog.SetWarnLog()`
    * `xlog.SetErrLog()`
* 各支付方式接入，请仔细查看 `xxx_test.go` 使用方式
    * `gopay/wechat/v3/client_test.go`
    * `gopay/alipay/client_test.go`
    * `gopay/qq/client_test.go`
    * `gopay/allinpay/client_test.go`
    * `gopay/lakala/client_test.go`
    * `gopay/paypal/client_test.go`
    * `gopay/apple/verify_test.go`
    * 或 examples
* 接入gopay示例项目(可参考接入使用方式)：[gopay-platform](https://github.com/go-pay/gopay-platform)
* 有问题请加QQ群(加群验证答案：gopay) 或 加微信好友(备注：gopay开发)拉群。在此，非常感谢提出宝贵意见和反馈问题的同志们！
* 开发过程中，请尽量使用正式环境，1分钱测试法！
* 业余时间接各类业务外包项目开发(前端+后端，架构设计->系统开发->部署运营)，如需要加微信联系。

QQ群：
<img width="280" height="280" src=".github/qq_gopay.png"/>
加微信拉群：
<img width="280" height="280" src=".github/wechat_jerry.png"/>

---

<br>

## 赞赏多少是您的心意，感谢支持！

微信赞赏码： <img width="200" height="200" src=".github/zanshang.png"/>
支付宝赞助码： <img width="200" height="200" src=".github/zanshang_zfb.png"/>

---

<br>

## 鸣谢

> [GoLand](https://www.jetbrains.com/go/?from=gopay) A Go IDE with extended support for JavaScript, TypeScript, and Databases。

特别感谢 [JetBrains](https://www.jetbrains.com/?from=gopay) 为开源项目提供免费的 [GoLand](https://www.jetbrains.com/go/?from=gopay) 等 IDE 的授权  
[<img src=".github/jetbrains-main-logo.png" width="200"/>](https://www.jetbrains.com/?from=gopay)