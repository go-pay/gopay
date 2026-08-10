package alipay

// ---------------------------------------------------------------------------
// 支付宝行业发票（正向开票）相关常量
// 文档: https://opendocs.alipay.com/mini/031f9074_alipay.commerce.ec.industryinvoice.invoiceapply.create?pathHash=a4e5340c
// ---------------------------------------------------------------------------

const (
	// 发票类型
	IndustryInvoiceTypeBlue = "BLUE" // 蓝票
	IndustryInvoiceTypeRed  = "RED"  // 红票（红冲）

	// 发票种类（invoice_kind）
	IndustryInvoiceKindElectronicGeneral = "ALL_ELECTRONIC_GENERAL" // 数电普通发票
	IndustryInvoiceKindElectronicSpecial = "ALL_ELECTRONIC_SPECIAL" // 数电专用发票
)

// ---------------------------------------------------------------------------
// alipay.commerce.ec.industryinvoice.invoiceapply.create
// 创建发票申请（正向开票，支持蓝票和红票）
// 文档: https://opendocs.alipay.com/mini/031f9074_alipay.commerce.ec.industryinvoice.invoiceapply.create?pathHash=a4e5340c
// ---------------------------------------------------------------------------

// IndustryInvoiceApplyCreateRsp 创建发票申请响应
type IndustryInvoiceApplyCreateRsp struct {
	Response     *IndustryInvoiceApplyCreate `json:"alipay_commerce_ec_industryinvoice_invoiceapply_create_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

// IndustryInvoiceApplyCreate 创建发票申请响应数据
type IndustryInvoiceApplyCreate struct {
	ErrorResponse
	InvoiceApplyId string `json:"invoice_apply_id,omitempty"` // 支付宝开票申请ID
}

// ---------------------------------------------------------------------------
// alipay.commerce.ec.industryinvoice.invoiceapply.query
// 查询发票申请详情
// 文档: https://opendocs.alipay.com/mini/45168e10_alipay.commerce.ec.industryinvoice.invoiceapply.query?scene=common&pathHash=9e5a0b67
// ---------------------------------------------------------------------------

// IndustryInvoiceApplyQueryRsp 查询发票申请响应
type IndustryInvoiceApplyQueryRsp struct {
	Response     *IndustryInvoiceApplyQuery `json:"alipay_commerce_ec_industryinvoice_invoiceapply_query_response"`
	AlipayCertSn string                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                     `json:"-"`
	Sign         string                     `json:"sign"`
}

// IndustryInvoiceApplyQuery 查询发票申请响应数据
type IndustryInvoiceApplyQuery struct {
	ErrorResponse
	InvoiceApplyId           string                                `json:"invoice_apply_id,omitempty"`            // 支付宝开票申请ID
	OuterApplyId             string                                `json:"outer_apply_id,omitempty"`              // 外部开票申请ID
	TradeList                []IndustryInvoiceTradeInfo            `json:"trade_list,omitempty"`                  // 交易信息列表
	TaxNo                    string                                `json:"tax_no,omitempty"`                      // 企业税号
	ProductId                string                                `json:"product_id,omitempty"`                  // 发票产品ID
	BuyerName                string                                `json:"buyer_name,omitempty"`                  // 购买方名称
	BuyerTaxNo               string                                `json:"buyer_tax_no,omitempty"`                // 购买方税号
	BuyerTaxNoType           string                                `json:"buyer_tax_no_type,omitempty"`           // 购买方证件类型
	BuyerAddress             string                                `json:"buyer_address,omitempty"`               // 购买方地址
	BuyerTel                 string                                `json:"buyer_tel,omitempty"`                   // 购买方电话
	BuyerBankName            string                                `json:"buyer_bank_name,omitempty"`             // 购买方开户行
	BuyerBankAccount         string                                `json:"buyer_bank_account,omitempty"`          // 购买方银行账号
	SellerName               string                                `json:"seller_name,omitempty"`                 // 销售方名称
	SellerTaxNo              string                                `json:"seller_tax_no,omitempty"`               // 销售方税号
	InvoiceType              string                                `json:"invoice_type,omitempty"`                // 发票类型
	InvoiceNo                string                                `json:"invoice_no,omitempty"`                  // 发票号码
	InvoiceDate              string                                `json:"invoice_date,omitempty"`                // 开票时间
	InvoiceKind              string                                `json:"invoice_kind,omitempty"`                // 发票票种
	InvoiceRedReason         string                                `json:"invoice_red_reason,omitempty"`          // 红冲原因
	RelatedBlueInvoiceNo     string                                `json:"related_blue_invoice_no,omitempty"`     // 关联蓝票发票号码
	RedConfirmationNo        string                                `json:"red_confirmation_no,omitempty"`         // 红字确认单号
	RedConfirmationUuid      string                                `json:"red_confirmation_uuid,omitempty"`       // 红字确认单UUID
	InvoiceStatus            string                                `json:"invoice_status,omitempty"`              // 开票状态
	InvoiceFailCode          string                                `json:"invoice_fail_code,omitempty"`           // 开票异常错误码
	InvoiceFailDesc          string                                `json:"invoice_fail_desc,omitempty"`           // 开票异常错误说明
	ImgFileUrl               string                                `json:"img_file_url,omitempty"`                // 发票文件地址-图片
	PdfFileUrl               string                                `json:"pdf_file_url,omitempty"`                // 发票文件地址-PDF
	InvoiceAmountWithoutTax  string                                `json:"invoice_amount_without_tax,omitempty"`  // 发票不含税金额
	InvoiceTaxAmount         string                                `json:"invoice_tax_amount,omitempty"`          // 发票税额
	InvoiceAmount            string                                `json:"invoice_amount,omitempty"`              // 发票总金额
	Remark                   string                                `json:"remark,omitempty"`                      // 发票备注
	InvoiceItemList          []IndustryInvoiceItemInfoQuery        `json:"invoice_item_list,omitempty"`           // 商品明细列表
	RealPropertyBusinessList []IndustryInvoiceRealPropertyBusiness `json:"real_property_business_list,omitempty"` // 不动产信息列表
}

// IndustryInvoiceTradeInfo 交易信息
type IndustryInvoiceTradeInfo struct {
	ChannelType string `json:"channel_type,omitempty"` // 交易渠道
	TradeNo     string `json:"trade_no,omitempty"`     // 订单交易号
}

// IndustryInvoiceItemInfoQuery 查询发票商品明细信息
type IndustryInvoiceItemInfoQuery struct {
	SerialNo            int    `json:"serial_no,omitempty"`              // 明细行序号
	ItemCode            string `json:"item_code,omitempty"`              // 商品配置编码
	TaxCode             string `json:"tax_code,omitempty"`               // 税收分类编码
	ItemCategoryName    string `json:"item_category_name,omitempty"`     // 税收分类编码简称
	ItemName            string `json:"item_name,omitempty"`              // 商品名称
	ItemSpec            string `json:"item_spec,omitempty"`              // 规格型号
	ItemUnit            string `json:"item_unit,omitempty"`              // 商品单位
	ItemNum             string `json:"item_num,omitempty"`               // 商品数量
	ItemAmount          string `json:"item_amount,omitempty"`            // 开票金额
	ItemTaxRate         string `json:"item_tax_rate,omitempty"`          // 商品税率
	ItemTaxAmount       string `json:"item_tax_amount,omitempty"`        // 税额
	InvoiceLineProperty string `json:"invoice_line_property,omitempty"`  // 发票行性质
	FavouredPolicyFlag  string `json:"favoured_policy_flag,omitempty"`   // 优惠政策标识
	RelatedBlueSerialNo int    `json:"related_blue_serial_no,omitempty"` // 关联蓝票明细行序号
}

// IndustryInvoiceRealPropertyBusiness 不动产信息
type IndustryInvoiceRealPropertyBusiness struct {
	SerialNo             int      `json:"serial_no,omitempty"`              // 明细行号
	RealPropertyCode     string   `json:"real_property_code,omitempty"`     // 不动产自定义编码
	RealPropertyProvince string   `json:"real_property_province,omitempty"` // 不动产地址省级行政区
	RealPropertyCity     string   `json:"real_property_city,omitempty"`     // 不动产地址所属市级行政区
	RealPropertyAddress  string   `json:"real_property_address,omitempty"`  // 不动产地址
	CrossCityFlag        string   `json:"cross_city_flag,omitempty"`        // 跨地（市）标志
	RealPropertyCertNo   string   `json:"real_property_cert_no,omitempty"`  // 不动产权证号
	RealPropertyArea     string   `json:"real_property_area,omitempty"`     // 不动产使用面积
	PlateNoList          []string `json:"plate_no_list,omitempty"`          // 车牌号列表
	StartTime            string   `json:"start_time,omitempty"`             // 使用开始时间
	EndTime              string   `json:"end_time,omitempty"`               // 使用结束时间
}

// ---------------------------------------------------------------------------
// alipay.commerce.ec.industryinvoice.company.query
// 查询企业开票信息
// 文档: https://opendocs.alipay.com/mini/a52ef6fc_alipay.commerce.ec.industryinvoice.company.query?scene=common&pathHash=83faa4b3
// ---------------------------------------------------------------------------

// IndustryInvoiceCompanyQueryRsp 查询企业信息响应
type IndustryInvoiceCompanyQueryRsp struct {
	Response     *IndustryInvoiceCompanyQuery `json:"alipay_commerce_ec_industryinvoice_company_query_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

// IndustryInvoiceCompanyQuery 查询企业信息响应数据
type IndustryInvoiceCompanyQuery struct {
	ErrorResponse
	TaxNo                  string                       `json:"tax_no,omitempty"`                    // 企业税号
	CompanyName            string                       `json:"company_name,omitempty"`              // 企业名称
	InvoiceClerk           *IndustryInvoiceClerk        `json:"invoice_clerk,omitempty"`             // 企业开票员信息
	CompanyProductInfoList []IndustryCompanyProductInfo `json:"company_product_info_list,omitempty"` // 企业已开通产品信息列表
}

// IndustryInvoiceClerk 企业开票员信息
type IndustryInvoiceClerk struct {
	ClerkCertNo string `json:"clerk_cert_no,omitempty"` // 开票员证件号码
	ClerkName   string `json:"clerk_name,omitempty"`    // 开票员姓名
	ClerkNo     string `json:"clerk_no,omitempty"`      // 开票员身份标识
}

// IndustryCompanyProductInfo 企业已开通产品信息
type IndustryCompanyProductInfo struct {
	ProductId            string                         `json:"product_id,omitempty"`             // 产品ID
	ProductName          string                         `json:"product_name,omitempty"`           // 产品名称
	CompanyProductConfig []IndustryCompanyProductConfig `json:"company_product_config,omitempty"` // 企业产品配置信息
}

// IndustryCompanyProductConfig 企业产品配置信息
type IndustryCompanyProductConfig struct {
	ConfigKey   string `json:"config_key,omitempty"`   // 配置键
	ConfigValue string `json:"config_value,omitempty"` // 配置值
}

// ---------------------------------------------------------------------------
// alipay.commerce.ec.industryinvoice.invoiceapply.retry
// 重试发票申请
// 文档: https://opendocs.alipay.com/mini/99125d7c_alipay.commerce.ec.industryinvoice.invoiceapply.retry?scene=common&pathHash=d5ce1397
// ---------------------------------------------------------------------------

// IndustryInvoiceApplyRetryRsp 重试发票申请响应
type IndustryInvoiceApplyRetryRsp struct {
	Response     *IndustryInvoiceApplyRetry `json:"alipay_commerce_ec_industryinvoice_invoiceapply_retry_response"`
	AlipayCertSn string                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                     `json:"-"`
	Sign         string                     `json:"sign"`
}

// IndustryInvoiceApplyRetry 重试发票申请响应数据
type IndustryInvoiceApplyRetry struct {
	ErrorResponse
}
