# 添加支付接口工作流程

这个 skill 定义了为 GoPay 项目添加新支付接口的标准流程。

## 适用场景

- 添加微信支付新接口
- 添加支付宝新接口
- 添加其他支付平台接口
- 更新现有接口

## 标准工作流程

### 1. 分析接口文档

**目标：** 从官方文档中提取完整的接口信息

**操作步骤：**
- 使用 `WebFetch` 工具获取官方文档内容
- 如果文档需要交互，使用浏览器自动化工具（`mcp__browser-automation__*`）
- 提取以下信息：
  - HTTP 方法（GET/POST/PUT/DELETE/PATCH）
  - URL 路径（如 `/v3/med-ins/orders`）
  - 请求参数结构（字段名、类型、是否必填、说明）
  - 响应参数结构（字段名、类型、说明）

**示例：**
```
接口名称：医保自费混合收款下单
HTTP方法：POST
URL路径：/v3/med-ins/orders
请求参数：mix_pay_type, order_type, appid, openid, ...
响应参数：mix_trade_no, mix_pay_status, ...
```

### 2. 实现接口代码

**目标：** 按照项目规范实现接口代码

**操作步骤：**

#### 2.1 添加 API 路径常量
- 文件位置：`wechat/v3/constant.go` 或 `alipay/v3/constant.go`
- 命名规范：`v3` + 功能模块名 + 操作名
- 示例：
  ```go
  // 医保支付
  v3MedInsOrder             = "/v3/med-ins/orders"                 // 医保自费混合收款下单 POST
  v3MedInsOrderQueryByMixNo = "/v3/med-ins/orders/mix-trade-no/%s" // mix_trade_no 使用医保自费混合订单号查看下单结果 GET
  ```

#### 2.2 定义数据结构
- 文件位置：创建或更新 `model_*.go` 文件（如 `model_medins.go`）
- 结构体命名：
  - 响应结构：`功能名Rsp`（如 `MedInsOrderRsp`）
  - 数据结构：`功能名`（如 `MedInsOrder`）
- 必须包含的字段：
  ```go
  type XxxRsp struct {
      Code        int         `json:"-"`
      SignInfo    *SignInfo   `json:"-"`
      Response    *Xxx        `json:"response,omitempty"`
      ErrResponse ErrResponse `json:"err_response,omitempty"`
      Error       string      `json:"-"`
  }
  ```

#### 2.3 实现接口方法
- 文件位置：创建或更新功能模块文件（如 `medins.go`）
- 方法命名：`V3` + 功能模块名 + 操作名（如 `V3MedInsOrder`）
- 实现模式（参考现有接口）：
  ```go
  func (c *ClientV3) V3XxxMethod(ctx context.Context, bm gopay.BodyMap) (wxRsp *XxxRsp, err error) {
      // 1. 生成签名
      authorization, err := c.authorization(MethodPost, v3ApiPath, bm)
      if err != nil {
          return nil, err
      }
      
      // 2. 发送请求
      res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPath, authorization)
      if err != nil {
          return nil, err
      }
      
      // 3. 处理响应
      wxRsp = &XxxRsp{Code: Success, SignInfo: si, Response: new(Xxx)}
      if res.StatusCode != http.StatusOK {
          wxRsp.Code = res.StatusCode
          wxRsp.Error = string(bs)
          _ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
          return wxRsp, nil
      }
      
      // 4. 解析响应
      if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
          return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
      }
      
      // 5. 验证签名
      return wxRsp, c.verifySyncSign(si)
  }
  ```

**注意事项：**
- POST 请求使用 `doProdPost`
- GET 请求使用 `doProdGet`
- 路径参数使用 `fmt.Sprintf` 格式化
- 必须进行签名验证

### 3. 更新文档

**目标：** 在项目文档中添加新接口说明

**操作步骤：**

#### 3.1 更新 API 文档
- 文件位置：`doc/wechat_v3.md` 或 `doc/alipay_v3.md`
- 添加位置：附录部分，按功能模块分类
- 格式：
  ```markdown
  * <font color='#07C160' size='4'>功能模块名</font>
      * 接口说明：`client.V3MethodName()`
      * 接口说明：`client.V3MethodName2()`
  ```

#### 3.2 示例
```markdown
* <font color='#07C160' size='4'>医保支付</font>
    * 医保自费混合收款下单：`client.V3MedInsOrder()`
    * 使用医保自费混合订单号查看下单结果：`client.V3MedInsOrderQueryByMixNo()`
    * 使用商户订单号查看下单结果：`client.V3MedInsOrderQueryByOutNo()`
```

### 4. 更新版本记录

**目标：** 记录本次更新内容

**操作步骤：**

#### 4.1 更新版本号
- 文件位置：`constant.go`（根目录）
- 修改 `Version` 常量，递增小版本号
- 示例：`v1.5.116` → `v1.5.117`

#### 4.2 更新发布说明
- 文件位置：`release_note.md`
- 在文件顶部添加新版本记录
- 格式：
  ```markdown
  ## 版本号：v1.5.xxx
  
  * 修改记录：
    * 平台名：新增 功能模块 相关接口。
      * client.MethodName()，接口说明。
      * client.MethodName2()，接口说明。
  ```

### 5. 提交代码

**目标：** 将所有修改提交到 Git

**操作步骤：**
```bash
# 1. 添加所有相关文件
git add constant.go
git add release_note.md
git add doc/wechat_v3.md  # 或对应平台文档
git add wechat/v3/constant.go  # 或对应平台
git add wechat/v3/新增文件.go
git add wechat/v3/model_新增文件.go

# 2. 检查状态
git status

# 3. 确保 linter 格式化后的文件也被添加
# 如果有未暂存的修改，再次 git add

# 4. 创建提交（由用户决定）
```

## 实现示例

### 案例：微信 v3 医保支付接口

**背景：** 微信支付新增医保支付能力，需要添加相关接口。

**执行过程：**

1. **分析文档**
   - 文档地址：`https://pay.weixin.qq.com/doc/v3/merchant/4016781466`
   - 提取到 3 个接口：下单、查询（混合订单号）、查询（商户订单号）

2. **实现代码**
   - 新增常量：`wechat/v3/constant.go`
   - 新增模型：`wechat/v3/model_medins.go`
   - 新增实现：`wechat/v3/medins.go`

3. **更新文档**
   - 在 `doc/wechat_v3.md` 添加"医保支付"模块

4. **更新版本**
   - 版本号：`v1.5.116` → `v1.5.117`
   - 在 `release_note.md` 添加版本记录

5. **提交代码**
   - 共修改 6 个文件（包括 linter 格式化）

## 注意事项

1. **代码规范**
   - 遵循项目现有的代码风格
   - 参考同类接口的实现方式
   - 保持命名一致性

2. **错误处理**
   - 必须处理所有可能的错误情况
   - HTTP 状态码检查
   - JSON 解析错误处理
   - 签名验证

3. **文档完整性**
   - API 文档必须更新
   - 版本记录必须更新
   - 注释要清晰准确

4. **测试验证**
   - 代码必须能够编译通过
   - 建议编写测试用例（参考 `*_test.go` 文件）

## 相关文件

- 常量定义：`wechat/v3/constant.go`、`alipay/v3/constant.go`
- 数据模型：`wechat/v3/model_*.go`、`alipay/v3/model_*.go`
- 接口实现：`wechat/v3/*.go`、`alipay/v3/*.go`
- API 文档：`doc/wechat_v3.md`、`doc/alipay_v3.md`
- 版本记录：`release_note.md`
- 版本号：`constant.go`

## 工具使用

- **WebFetch**：获取网页内容并提取信息
- **浏览器自动化**：处理需要交互的文档页面
- **Read**：读取现有代码作为参考
- **Write/Edit**：创建或修改代码文件
- **Bash**：执行 git 命令和编译验证

## 后续维护

当官方文档更新或需要添加新接口时，严格按照此流程执行，确保：
- 代码质量一致
- 文档保持同步
- 版本记录完整
