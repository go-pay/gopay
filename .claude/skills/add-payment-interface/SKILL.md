---
name: add-payment-interface
description: Add new payment interfaces to GoPay project following standard workflow - analyze docs, implement code, update documentation and version records. Use when adding WeChat Pay, Alipay or other payment platform APIs.
---

# Add Payment Interface

Standardized workflow for adding new payment platform interfaces to the GoPay project.

## When to Activate

- Adding new WeChat Pay v3 interfaces
- Adding new Alipay v3 interfaces
- Adding interfaces for other payment platforms
- Updating existing payment interfaces with new capabilities

## Standard Workflow

### Step 1: Analyze Interface Documentation

**Objective:** Extract complete interface specifications from official documentation.

**Actions:**
1. Use `WebFetch` to retrieve official documentation
2. Use browser automation (`mcp__browser-automation__*`) if interactive navigation needed
3. Extract:
   - HTTP method (GET/POST/PUT/DELETE/PATCH)
   - URL path (e.g., `/v3/med-ins/orders`)
   - Request parameters (name, type, required, description)
   - Response parameters (name, type, description)

**Example Output:**
```
Interface: 医保自费混合收款下单
Method: POST
Path: /v3/med-ins/orders
Request: mix_pay_type, order_type, appid, openid, ...
Response: mix_trade_no, mix_pay_status, ...
```

### Step 2: Implement Interface Code

**Objective:** Implement interface following project conventions.

#### 2.1 Add API Path Constants

**File:** `wechat/v3/constant.go` or `alipay/v3/constant.go`

**Naming:** `v3` + ModuleName + ActionName

**Example:**
```go
// 医保支付
v3MedInsOrder             = "/v3/med-ins/orders"                 // 医保自费混合收款下单 POST
v3MedInsOrderQueryByMixNo = "/v3/med-ins/orders/mix-trade-no/%s" // mix_trade_no 查询 GET
v3MedInsOrderQueryByOutNo = "/v3/med-ins/orders/out-trade-no/%s" // out_trade_no 查询 GET
```

#### 2.2 Define Data Models

**File:** Create or update `model_*.go` (e.g., `model_medins.go`)

**Required structures:**
```go
// Response wrapper
type XxxRsp struct {
    Code        int         `json:"-"`
    SignInfo    *SignInfo   `json:"-"`
    Response    *Xxx        `json:"response,omitempty"`
    ErrResponse ErrResponse `json:"err_response,omitempty"`
    Error       string      `json:"-"`
}

// Response data
type Xxx struct {
    Field1 string `json:"field1"` // 字段说明
    Field2 int    `json:"field2"` // 字段说明
    // ... more fields
}
```

#### 2.3 Implement Interface Methods

**File:** Create or update module file (e.g., `medins.go`)

**Method naming:** `V3` + ModuleName + ActionName

**Implementation pattern:**
```go
func (c *ClientV3) V3XxxMethod(ctx context.Context, bm gopay.BodyMap) (wxRsp *XxxRsp, err error) {
    // 1. Generate authorization
    authorization, err := c.authorization(MethodPost, v3ApiPath, bm)
    if err != nil {
        return nil, err
    }
    
    // 2. Send request (use doProdPost for POST, doProdGet for GET)
    res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPath, authorization)
    if err != nil {
        return nil, err
    }
    
    // 3. Initialize response
    wxRsp = &XxxRsp{Code: Success, SignInfo: si, Response: new(Xxx)}
    if res.StatusCode != http.StatusOK {
        wxRsp.Code = res.StatusCode
        wxRsp.Error = string(bs)
        _ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
        return wxRsp, nil
    }
    
    // 4. Parse response
    if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
        return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
    }
    
    // 5. Verify signature
    return wxRsp, c.verifySyncSign(si)
}
```

**Key points:**
- POST requests: use `doProdPost`
- GET requests: use `doProdGet`
- Path parameters: use `fmt.Sprintf` to format URL
- Always verify signature with `verifySyncSign`

### Step 3: Update Documentation

**Objective:** Document new interfaces in project docs.

#### 3.1 Update API Documentation

**File:** `doc/wechat_v3.md` or `doc/alipay_v3.md`

**Location:** Appendix section, grouped by functionality

**Format:**
```markdown
* <font color='#07C160' size='4'>Module Name</font>
    * Interface description: `client.V3MethodName()`
    * Interface description: `client.V3MethodName2()`
```

**Example:**
```markdown
* <font color='#07C160' size='4'>医保支付</font>
    * 医保自费混合收款下单：`client.V3MedInsOrder()`
    * 使用医保自费混合订单号查看下单结果：`client.V3MedInsOrderQueryByMixNo()`
    * 使用商户订单号查看下单结果：`client.V3MedInsOrderQueryByOutNo()`
```

### Step 4: Update Version Records

**Objective:** Record changes in version history.

#### 4.1 Update Version Number

**File:** `constant.go` (root directory)

**Action:** Increment patch version (e.g., `v1.5.116` → `v1.5.117`)

#### 4.2 Update Release Notes

**File:** `release_note.md`

**Action:** Add new version section at the top

**Format:**
```markdown
## 版本号：v1.5.xxx

* 修改记录：
  * 平台名：新增 功能模块 相关接口。
    * client.MethodName()，接口说明。
    * client.MethodName2()，接口说明。
```

### Step 5: Commit Changes

**Objective:** Stage all modified files for commit.

**Actions:**
```bash
# Add all related files
git add constant.go
git add release_note.md
git add doc/wechat_v3.md  # or corresponding platform doc
git add wechat/v3/constant.go  # or corresponding platform
git add wechat/v3/new_file.go
git add wechat/v3/model_new_file.go

# Check status
git status

# Add any linter-formatted files
git add <formatted_files>
```

## Implementation Example

**Case:** WeChat Pay v3 Medical Insurance Payment

**Documentation:** `https://pay.weixin.qq.com/doc/v3/merchant/4016781466`

**Interfaces identified:**
1. Create medical insurance order - POST `/v3/med-ins/orders`
2. Query by mix trade no - GET `/v3/med-ins/orders/mix-trade-no/{mix_trade_no}`
3. Query by merchant order no - GET `/v3/med-ins/orders/out-trade-no/{out_trade_no}`

**Files modified:**
- `constant.go` - version bump to v1.5.117
- `release_note.md` - added v1.5.117 changelog
- `doc/wechat_v3.md` - added medical insurance section
- `wechat/v3/constant.go` - added 3 API path constants
- `wechat/v3/medins.go` - new file with 3 methods
- `wechat/v3/model_medins.go` - new file with data structures

## Quality Checklist

Before completing:
- [ ] Code compiles without errors (`go build ./wechat/v3/...`)
- [ ] All API paths added to constants
- [ ] Data models include all documented fields
- [ ] Methods follow existing implementation patterns
- [ ] Error handling is complete
- [ ] Signature verification is included
- [ ] API documentation is updated
- [ ] Version number is incremented
- [ ] Release notes are updated
- [ ] All files are staged for commit

## Reference Files

- **Constants:** `wechat/v3/constant.go`, `alipay/v3/constant.go`
- **Models:** `wechat/v3/model_*.go`, `alipay/v3/model_*.go`
- **Implementations:** `wechat/v3/*.go`, `alipay/v3/*.go`
- **Documentation:** `doc/wechat_v3.md`, `doc/alipay_v3.md`
- **Version:** `constant.go`, `release_note.md`

## Tools Used

- **WebFetch** - Retrieve documentation content
- **Browser automation** - Navigate interactive documentation
- **Read** - Reference existing implementations
- **Write/Edit** - Create or modify code files
- **Bash** - Execute git commands and build verification
