# GoPay 项目说明

## 项目概述

GoPay 是一个 Go 语言的支付 SDK 项目，提供微信、支付宝、QQ、通联支付、拉卡拉、PayPal、扫呗、Apple 支付的统一接口。

- **仓库**: https://github.com/go-pay/gopay
- **Go 版本**: 1.24.0
- **许可证**: Apache 2.0

## 项目结构

```
gopay/
├── alipay/          # 支付宝支付 SDK
├── wechat/          # 微信支付 SDK
├── apple/           # Apple 支付校验
├── paypal/          # PayPal 支付 SDK
├── qq/              # QQ 支付 SDK
├── allinpay/        # 通联支付 SDK
├── lakala/          # 拉卡拉支付 SDK
├── saobei/          # 扫呗支付 SDK
├── pkg/             # 公共包
├── doc/             # 各支付方式文档
└── examples/        # 示例代码
```

## 开发约定

### 代码风格

1. **遵循 Go 标准规范**
   - 使用 `gofmt` 格式化代码
   - 遵循 Go 命名约定（驼峰命名）
   - 导出的函数、类型、常量首字母大写

2. **测试文件**
   - 每个模块都有对应的 `*_test.go` 文件
   - 测试文件包含实际使用示例
   - 建议使用正式环境 1 分钱测试法

3. **日志处理**
   - 使用项目依赖的 `github.com/go-pay/xlog` 包
   - 支持自定义 Logger（实现 `xlog.XLogger` interface）

### 支付平台版本

- **支付宝**: 优先使用 V3 版本
- **微信**: 优先使用 V3 版本（V2 版本不推荐）
- 其他平台参考 `doc/` 目录下的文档

### 重要注意事项

1. **安全性优先**
   - 处理支付相关代码时特别注意安全性
   - 不要在代码中硬编码敏感信息（密钥、证书等）
   - 注意参数验证和签名校验

2. **向后兼容**
   - 这是一个公开的 SDK，修改时注意 API 兼容性
   - 重大变更需要在 `release_note.md` 中记录

3. **文档同步**
   - 新增功能需要更新对应的 `doc/*.md` 文档
   - 保持代码示例和文档的一致性

## 依赖管理

项目使用 Go Modules，主要依赖：
- `github.com/go-pay/crypto` - 加密工具
- `github.com/go-pay/xlog` - 日志工具
- `github.com/go-pay/util` - 通用工具
- `golang.org/x/crypto` - Go 官方加密库

## 参考资源

- 各支付方式文档: `doc/` 目录
- 测试用例: 各模块的 `*_test.go` 文件
- 示例项目: https://github.com/go-pay/gopay-platform
