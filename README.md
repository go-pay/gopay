<div align=center><img width="240" height="240" alt="Logo was Loading Faild!" src="https://raw.githubusercontent.com/go-pay/gopay/main/logo.png"/></div>

# GoPay

### 微信、支付宝、PayPal、QQ 的 Golang 版本SDK

[![Github](https://img.shields.io/github/followers/iGoogle-ink?label=Follow&style=social)](https://github.com/iGoogle-ink)
[![Github](https://img.shields.io/github/forks/go-pay/gopay?label=Fork&style=social)](https://github.com/go-pay/gopay/fork)

[![Golang](https://img.shields.io/badge/golang-1.16-brightgreen.svg)](https://golang.google.cn)
[![GoDoc](https://img.shields.io/badge/doc-pkg.go.dev-informational.svg)](https://pkg.go.dev/github.com/go-pay/gopay)
[![Drone CI](https://cloud.drone.io/api/badges/go-pay/gopay/status.svg)](https://cloud.drone.io/go-pay/gopay)
[![GitHub Release](https://img.shields.io/github/v/release/go-pay/gopay)](https://github.com/go-pay/gopay/releases)
[![License](https://img.shields.io/github/license/go-pay/gopay)](https://www.apache.org/licenses/LICENSE-2.0)
[![GoDoc](https://tokei.rs/b1/github.com/go-pay/gopay?category=lines)](https://github.com/go-pay/gopay)

---

- #### 近期计划：

> 将 gopay 库中，非支付相关的一些接口方法独立出去另外的 sdk 库，在 go-pay 组织下新建 `wechat-sdk` 和 `alipay-sdk` 两个项目，分别实现各个平台相关接口方法，优先进行 `wechat-sdk` 开发。

> 微信小程序或公众号相关接口方法：已从 `微信v2` 移步替换成 `github.com/go-pay/wechat-sdk`

<br>

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

* #### [Alipay](https://github.com/go-pay/gopay/blob/main/doc/alipay.md)
* #### [Wechat](https://github.com/go-pay/gopay/blob/main/doc/wechat_v3.md)
* #### [QQ](https://github.com/go-pay/gopay/blob/main/doc/qq.md)
* #### [Paypal](https://github.com/go-pay/gopay/blob/main/doc/paypal.md)
* #### [Apple](https://github.com/go-pay/gopay/blob/main/doc/apple.md)

---

<br>

# 三、其他说明

* 各支付方式接入，请仔细查看 `xxx_test.go` 使用方式
    * `gopay/wechat/v3/client_test.go`
    * `gopay/alipay/client_test.go`
    * `gopay/qq/client_test.go`
    * `gopay/paypal/client_test.go`
    * `gopay/apple/verify_test.go`
    * 或 examples
* 有问题请加QQ群（加群验证答案：gopay），或加微信好友拉群。在此，非常感谢提出宝贵意见和反馈问题的同志们！
* 开发过程中，请尽量使用正式环境，1分钱测试法！

QQ群：
<img width="280" height="280" src="https://raw.githubusercontent.com/go-pay/gopay/main/qq_gopay.png"/>
加微信拉群：
<img width="280" height="280" src="https://raw.githubusercontent.com/go-pay/gopay/main/wechat_jerry.png"/>

---

<br>

## 赞赏多少是您的心意，感谢支持！

微信赞赏码： <img width="200" height="200" src="https://raw.githubusercontent.com/go-pay/gopay/main/zanshang.png"/>
支付宝赞助码： <img width="200" height="200" src="https://raw.githubusercontent.com/go-pay/gopay/main/zanshang_zfb.png"/>

## License

```
Copyright 2019 Jerry

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
