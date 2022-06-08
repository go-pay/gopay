package paypal

type AccessToken struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Appid       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

type ErrorResponse struct {
	Name    string        `json:"name,omitempty"`
	Message string        `json:"message,omitempty"`
	DebugId string        `json:"debug_id,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
	Links   []Link        `json:"links,omitempty"`
}

type ErrorDetail struct {
	Issue       string `json:"issue,omitempty"`
	Field       string `json:"field,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
}

type EmptyRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
}

type CreateOrderRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *OrderDetail   `json:"response,omitempty"`
}

type OrderDetailRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *OrderDetail   `json:"response,omitempty"`
}

type OrderAuthorizeRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *OrderDetail   `json:"response,omitempty"`
}

type OrderCaptureRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *OrderDetail   `json:"response,omitempty"`
}

type OrderConfirmRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *OrderDetail   `json:"response,omitempty"`
}

type PaymentAuthorizeDetailRsp struct {
	Code          int                     `json:"-"`
	Error         string                  `json:"-"`
	ErrorResponse *ErrorResponse          `json:"-"`
	Response      *PaymentAuthorizeDetail `json:"response,omitempty"`
}

type PaymentReauthorizeRsp struct {
	Code          int                     `json:"-"`
	Error         string                  `json:"-"`
	ErrorResponse *ErrorResponse          `json:"-"`
	Response      *PaymentAuthorizeDetail `json:"response,omitempty"`
}

type PaymentAuthorizeCaptureRsp struct {
	Code          int                      `json:"-"`
	Error         string                   `json:"-"`
	ErrorResponse *ErrorResponse           `json:"-"`
	Response      *PaymentAuthorizeCapture `json:"response,omitempty"`
}

type PaymentCaptureDetailRsp struct {
	Code          int                      `json:"-"`
	Error         string                   `json:"-"`
	ErrorResponse *ErrorResponse           `json:"-"`
	Response      *PaymentAuthorizeCapture `json:"response,omitempty"`
}

type PaymentCaptureRefundRsp struct {
	Code          int                   `json:"-"`
	Error         string                `json:"-"`
	ErrorResponse *ErrorResponse        `json:"-"`
	Response      *PaymentCaptureRefund `json:"response,omitempty"`
}

type PaymentRefundDetailRsp struct {
	Code          int                   `json:"-"`
	Error         string                `json:"-"`
	ErrorResponse *ErrorResponse        `json:"-"`
	Response      *PaymentCaptureRefund `json:"response,omitempty"`
}

type CreateBatchPayoutRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *BatchPayout   `json:"response,omitempty"`
}

type PayoutBatchDetailRsp struct {
	Code          int                `json:"-"`
	Error         string             `json:"-"`
	ErrorResponse *ErrorResponse     `json:"-"`
	Response      *PayoutBatchDetail `json:"response,omitempty"`
}

type PayoutItemDetailRsp struct {
	Code          int               `json:"-"`
	Error         string            `json:"-"`
	ErrorResponse *ErrorResponse    `json:"-"`
	Response      *PayoutItemDetail `json:"response,omitempty"`
}

type CancelUnclaimedPayoutItemRsp struct {
	Code          int               `json:"-"`
	Error         string            `json:"-"`
	ErrorResponse *ErrorResponse    `json:"-"`
	Response      *PayoutItemDetail `json:"response,omitempty"`
}

// ==================================分割==================================

type Patch struct {
	Op    string      `json:"op"` // The possible values are: add、remove、replace、move、copy、test
	Path  string      `json:"path,omitempty"`
	Value interface{} `json:"value"` // The value to apply. The remove operation does not require a value.
	From  string      `json:"from,omitempty"`
}

type OrderDetail struct {
	Id                    string          `json:"id,omitempty"`
	Status                string          `json:"status,omitempty"` // CREATED、SAVED、APPROVED、VOIDED、COMPLETED、PAYER_ACTION_REQUIRED
	PaymentSource         *PaymentSource  `json:"payment_source,omitempty"`
	Intent                string          `json:"intent,omitempty"`
	ProcessingInstruction string          `json:"processing_instruction,omitempty"`
	Payer                 *Payer          `json:"payer,omitempty"`
	PurchaseUnits         []*PurchaseUnit `json:"purchase_units,omitempty"`
	Links                 []*Link         `json:"links,omitempty"`
	CreateTime            string          `json:"create_time,omitempty"`
	UpdateTime            string          `json:"update_time,omitempty"`
}

type PaymentSource struct {
	Bancontact *Bancontact `json:"bancontact,omitempty"`
	Blik       *Blik       `json:"blik,omitempty"`
	Card       *Card       `json:"card,omitempty"`
	Eps        *Eps        `json:"eps,omitempty"`
	Giropay    *Giropay    `json:"giropay,omitempty"`
	Ideal      *Ideal      `json:"ideal,omitempty"`
	Mybank     *Mybank     `json:"mybank,omitempty"`
	P24        *P24        `json:"p24,omitempty"`
	Sofort     *Sofort     `json:"sofort,omitempty"`
	Trustly    *Trustly    `json:"trustly,omitempty"`
}

type Bancontact struct {
	Bic            string `json:"bic,omitempty"`
	CardLastDigits string `json:"card_last_digits,omitempty"`
	CountryCode    string `json:"country_code,omitempty"`
	IbanLastChars  string `json:"iban_last_chars,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Blik struct {
	CountryCode string `json:"country_code,omitempty"`
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
}
type Eps struct {
	Bic         string `json:"bic,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Giropay struct {
	Bic         string `json:"bic,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Ideal struct {
	Bic           string `json:"bic,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	IbanLastChars string `json:"iban_last_chars,omitempty"`
	Name          string `json:"name,omitempty"`
}

type Mybank struct {
	Bic           string `json:"bic,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	IbanLastChars string `json:"iban_last_chars,omitempty"`
	Name          string `json:"name,omitempty"`
}

type P24 struct {
	CountryCode       string `json:"country_code,omitempty"`
	Email             string `json:"email,omitempty"`
	MethodDescription string `json:"method_description,omitempty"`
	MethodId          string `json:"method_id,omitempty"`
	Name              string `json:"name,omitempty"`
	PaymentDescriptor string `json:"payment_descriptor,omitempty"`
}

type Sofort struct {
	Bic           string `json:"bic,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	IbanLastChars string `json:"iban_last_chars,omitempty"`
	Name          string `json:"name,omitempty"`
}

type Trustly struct {
	Bic           string `json:"bic,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	IbanLastChars string `json:"iban_last_chars,omitempty"`
	Name          string `json:"name,omitempty"`
}

type Card struct {
	Name                 string                `json:"name"`
	BillingAddress       *Address              `json:"billing_address,omitempty"`
	LastDigits           string                `json:"last_digits"`
	Brand                string                `json:"brand"`
	Type                 string                `json:"type"` // The payment card type：CREDIT、DEBIT、PREPAID、UNKNOWN
	AuthenticationResult *AuthenticationResult `json:"authentication_result,omitempty"`
}

type Address struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	AdminArea1   string `json:"admin_area_1"`
	AdminArea2   string `json:"admin_area_2"`
	PostalCode   string `json:"postal_code"`
	CountryCode  string `json:"country_code"`
}

type AuthenticationResult struct {
	LiabilityShift string        `json:"liability_shift"`
	ThreeDSecure   *ThreeDSecure `json:"three_d_secure,omitempty"`
}

type ThreeDSecure struct {
	AuthenticationStatus string `json:"authentication_status"`
	EnrollmentStatus     string `json:"enrollment_status"`
}

type Payer struct {
	Name         *Name    `json:"name"`
	Phone        *Phone   `json:"phone,omitempty"`
	BirthDate    string   `json:"birth_date"`
	TaxInfo      *TaxInfo `json:"tax_info,omitempty"`
	EmailAddress string   `json:"email_address"`
	PayerId      string   `json:"payer_id"`
	Address      *Address `json:"address"`
}

type TaxInfo struct {
	TaxId     string `json:"tax_id"`
	TaxIdType string `json:"tax_id_type"`
}

type PurchaseUnit struct {
	Id                 string              `json:"id,omitempty"`
	ReferenceId        string              `json:"reference_id,omitempty"`
	Amount             *Amount             `json:"amount,omitempty"`
	Payee              *Payee              `json:"payee,omitempty"`
	PaymentInstruction *PaymentInstruction `json:"payment_instruction,omitempty"`
	Description        string              `json:"description,omitempty"`
	CustomId           string              `json:"custom_id,omitempty"`
	InvoiceId          string              `json:"invoice_id,omitempty"`
	SoftDescriptor     string              `json:"soft_descriptor,omitempty"`
	Items              []*Item             `json:"items,omitempty"`
	Shipping           *Shipping           `json:"shipping,omitempty"`
	Payments           *Payments           `json:"payments,omitempty"`
}

type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type Payee struct {
	EmailAddress string `json:"email_address,omitempty"`
	MerchantId   string `json:"merchant_id,omitempty"`
}

type PaymentInstruction struct {
	PlatformFees       []*PlatformFee `json:"platform_fees,omitempty"`
	DisbursementMode   string         `json:"disbursement_mode,omitempty"`
	PayeePricingTierId string         `json:"payee_pricing_tier_id,omitempty"`
}

type PlatformFee struct {
	Amount *Amount `json:"amount,omitempty"`
	Payee  *Payee  `json:"payee,omitempty"`
}

type Item struct {
	Name        string  `json:"name,omitempty"`
	UnitAmount  *Amount `json:"unit_amount,omitempty"`
	Tax         *Amount `json:"tax,omitempty"`
	Quantity    string  `json:"quantity,omitempty"`
	Description string  `json:"description,omitempty"`
	Sku         string  `json:"sku,omitempty"`
	Category    string  `json:"category,omitempty"`
}

type Shipping struct {
	Name    *Name    `json:"name,omitempty"`
	Type    string   `json:"type,omitempty"` // SHIPPING、PICKUP_IN_PERSON
	Address *Address `json:"address,omitempty"`
}

type Name struct {
	Prefix     string `json:"prefix,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	Suffix     string `json:"suffix,omitempty"`
	FullName   string `json:"full_name,omitempty"`
}

type Phone struct {
	PhoneType   string       `json:"phone_type"`
	PhoneNumber *PhoneNumber `json:"phone_number"`
}

type PhoneNumber struct {
	NationalNumber string `json:"national_number"`
}

type Payments struct {
	Authorizations []*Authorization `json:"authorizations,omitempty"`
	Captures       []*Capture       `json:"captures,omitempty"`
	Refunds        []*Refund        `json:"refunds,omitempty"`
}

type Authorization struct {
	ProcessorResponse *Processor `json:"processor_response,omitempty"`
}

type Processor struct {
	AvsCode           string `json:"avs_code"`
	CvvCode           string `json:"cvv_code"`
	ResponseCode      string `json:"response_code"`
	PaymentAdviceCode string `json:"payment_advice_code"`
}

type StatusDetails struct {
	Reason string `json:"reason"`
}

type Capture struct {
	Id                        string                     `json:"id,omitempty"`
	Status                    string                     `json:"status,omitempty"`
	StatusDetails             *StatusDetails             `json:"status_details,omitempty"`
	Amount                    *Amount                    `json:"amount,omitempty"`
	FinalCapture              bool                       `json:"final_capture,omitempty"`
	DisbursementMode          string                     `json:"disbursement_mode,omitempty"`
	SellerProtection          *SellerProtection          `json:"seller_protection,omitempty"`
	SellerReceivableBreakdown *SellerReceivableBreakdown `json:"seller_receivable_breakdown,omitempty"`
	Links                     []*Link                    `json:"links,omitempty"`
	CreateTime                string                     `json:"create_time,omitempty"`
	UpdateTime                string                     `json:"update_time,omitempty"`
}

type Refund struct {
	Status        string         `json:"status,omitempty"`
	StatusDetails *StatusDetails `json:"status_details,omitempty"`
}

type Link struct {
	Href   string `json:"href,omitempty"`
	Rel    string `json:"rel,omitempty"`
	Method string `json:"method,omitempty"` // Possible values: GET,POST,PUT,DELETE,HEAD,CONNECT,OPTIONS,PATCH
}

type PaymentAuthorizeDetail struct {
	Id               string            `json:"id,omitempty"`
	Status           string            `json:"status,omitempty"` // CREATED、CAPTURED、DENIED、EXPIRED、PARTIALLY_CAPTURED、PARTIALLY_CREATED、VOIDED、PENDING
	StatusDetails    *StatusDetails    `json:"status_details,omitempty"`
	Amount           *Amount           `json:"amount,omitempty"`
	InvoiceId        string            `json:"invoice_id,omitempty"`
	CustomId         string            `json:"custom_id,omitempty"`
	SellerProtection *SellerProtection `json:"seller_protection,omitempty"`
	Links            []*Link           `json:"links,omitempty"`
	ExpirationTime   string            `json:"expiration_time,omitempty"`
	CreateTime       string            `json:"create_time,omitempty"`
	UpdateTime       string            `json:"update_time,omitempty"`
}

type SellerProtection struct {
	Status            string   `json:"status,omitempty"` // ELIGIBLE、PARTIALLY_ELIGIBLE、NOT_ELIGIBLE
	DisputeCategories []string `json:"dispute_categories,omitempty"`
}

type PaymentAuthorizeCapture struct {
	Id                        string                     `json:"id,omitempty"`
	Status                    string                     `json:"status,omitempty"` // COMPLETED、DECLINED、PARTIALLY_REFUNDED、PENDING、REFUNDED、FAILED
	StatusDetails             *StatusDetails             `json:"status_details,omitempty"`
	Amount                    *Amount                    `json:"amount,omitempty"`
	InvoiceId                 string                     `json:"invoice_id,omitempty"`
	CustomId                  string                     `json:"custom_id,omitempty"`
	FinalCapture              bool                       `json:"final_capture,omitempty"`
	SellerProtection          *SellerProtection          `json:"seller_protection,omitempty"`
	SellerReceivableBreakdown *SellerReceivableBreakdown `json:"seller_receivable_breakdown,omitempty"`
	DisbursementMode          string                     `json:"disbursement_mode,omitempty"`
	Links                     []*Link                    `json:"links,omitempty"`
	ProcessorResponse         *Processor                 `json:"processor_response,omitempty"`
	CreateTime                string                     `json:"create_time,omitempty"`
	UpdateTime                string                     `json:"update_time,omitempty"`
}

type SellerReceivableBreakdown struct {
	GrossAmount                   *Amount        `json:"gross_amount"`
	PaypalFee                     *Amount        `json:"paypal_fee,omitempty"`
	PaypalFeeInReceivableCurrency *Amount        `json:"paypal_fee_in_receivable_currency,omitempty"`
	NetAmount                     *Amount        `json:"net_amount,omitempty"`
	ReceivableAmount              *Amount        `json:"receivable_amount,omitempty"`
	ExchangeRate                  *ExchangeRate  `json:"exchange_rate,omitempty"`
	PlatformFees                  []*PlatformFee `json:"platform_fees,omitempty"`
}

type ExchangeRate struct {
	SourceCurrency string `json:"source_currency,omitempty"`
	TargetCurrency string `json:"target_currency,omitempty"`
	Value          string `json:"value,omitempty"`
}

type PaymentCaptureRefund struct {
	Id                     string                  `json:"id,omitempty"`
	Status                 string                  `json:"status,omitempty"` // CANCELLED、PENDING、COMPLETED
	StatusDetails          *StatusDetails          `json:"status_details,omitempty"`
	Amount                 *Amount                 `json:"amount,omitempty"`
	InvoiceId              string                  `json:"invoice_id,omitempty"`
	NoteToPayer            string                  `json:"note_to_payer,omitempty"`
	SellerPayableBreakdown *SellerPayableBreakdown `json:"seller_payable_breakdown,omitempty"`
	Links                  []*Link                 `json:"links,omitempty"`
	CreateTime             string                  `json:"create_time,omitempty"`
	UpdateTime             string                  `json:"update_time,omitempty"`
}

type SellerPayableBreakdown struct {
	GrossAmount                   *Amount               `json:"gross_amount"`
	PaypalFee                     *Amount               `json:"paypal_fee,omitempty"`
	PaypalFeeInReceivableCurrency *Amount               `json:"paypal_fee_in_receivable_currency,omitempty"`
	NetAmount                     *Amount               `json:"net_amount,omitempty"`
	NetAmountInReceivableCurrency *Amount               `json:"net_amount_in_receivable_currency,omitempty"`
	PlatformFees                  []*PlatformFee        `json:"platform_fees,omitempty"`
	NetAmountBreakdown            []*NetAmountBreakdown `json:"net_amount_breakdown,omitempty"`
	TotalRefundedAmount           *Amount               `json:"total_refunded_amount,omitempty"`
}

type NetAmountBreakdown struct {
	PayableAmount   *Amount       `json:"payable_amount,omitempty"`
	ConvertedAmount *Amount       `json:"converted_amount,omitempty"`
	ExchangeRate    *ExchangeRate `json:"exchange_rate,omitempty"`
}

// =============== V1 API Payout ==================================

type V1Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type Errors struct {
	Name            string `json:"name"`
	Message         string `json:"message"`
	InformationLink string `json:"information_link"`
}

type PayoutCurrencyConversion struct {
	ExchangeRate string    `json:"exchange_rate"`
	FromAmount   *V1Amount `json:"from_amount"`
	ToAmount     *V1Amount `json:"to_amount"`
}

type PayoutItem struct {
	RecipientType string    `json:"recipient_type"`
	Amount        *V1Amount `json:"amount"`
	Note          string    `json:"note"`
	Receiver      string    `json:"receiver"`
	SenderItemId  string    `json:"sender_item_id"`
}

type SenderBatchHeader struct {
	SenderBatchId string `json:"sender_batch_id"`
	EmailSubject  string `json:"email_subject"`
	EmailMessage  string `json:"email_message,omitempty"`
}

type BatchHeader struct {
	PayoutBatchId     string             `json:"payout_batch_id"`
	BatchStatus       string             `json:"batch_status"` // DENIED、PENDING、PROCESSING、SUCCESS、CANCELED
	TimeCreated       string             `json:"time_created,omitempty"`
	TimeCompleted     string             `json:"time_completed,omitempty"`
	SenderBatchHeader *SenderBatchHeader `json:"sender_batch_header"`
	Amount            *V1Amount          `json:"amount,omitempty"`
	Fees              *V1Amount          `json:"fees,omitempty"`
}

type BatchPayout struct {
	BatchHeader *BatchHeader `json:"batch_header"`
	Links       []*Link      `json:"links,omitempty"`
}

type PayoutItemDetail struct {
	ActivityId         string                    `json:"activity_id,omitempty"`
	CurrencyConversion *PayoutCurrencyConversion `json:"currency_conversion,omitempty"`
	Errors             *Errors                   `json:"errors,omitempty"`
	Links              []*Link                   `json:"links,omitempty"`
	PayoutBatchId      string                    `json:"payout_batch_id"`
	PayoutItem         *PayoutItem               `json:"payout_item"`
	PayoutItemFee      *V1Amount                 `json:"payout_item_fee"`
	PayoutItemId       string                    `json:"payout_item_id"`
	SenderBatchId      string                    `json:"sender_batch_id,omitempty"`
	TimeProcessed      string                    `json:"time_processed"`
	TransactionId      string                    `json:"transaction_id"`
	TransactionStatus  string                    `json:"transaction_status"` // SUCCESS、FAILED、PENDING、UNCLAIMED、RETURNED、ONHOLD、BLOCKED、REFUNDED、REVERSED
}

type PayoutBatchDetail struct {
	BatchHeader *BatchHeader        `json:"batch_header"`
	Items       []*PayoutItemDetail `json:"items"`
	Links       []*Link             `json:"links"`
	TotalItems  int64               `json:"total_items,omitempty"`
	TotalPage   int64               `json:"total_page,omitempty"`
}

// Subscription Model

type Frequency struct {
	IntervalUnit  string `json:"interval_unit"`
	IntervalCount int    `json:"interval_count"`
}

type PricingScheme struct {
	FixedPrice *FixedPrice `json:"fixed_price"`
}

type FixedPrice struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type BillingCycles struct {
	Frequency     *Frequency     `json:"frequency"`
	TenureType    string         `json:"tenure_type"`
	Sequence      int            `json:"sequence"`
	TotalCycles   int            `json:"total_cycles"`
	PricingScheme *PricingScheme `json:"pricing_scheme"`
}

type Plans struct {
	ProductId          string              `json:"product_id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	BillingCycles      []*BillingCycles    `json:"billing_cycles"`
	PaymentDefinitions *PaymentPreferences `json:"payment_preferences"`
}

type PaymentPreferences struct {
	AutoBillOutstanding     bool   `json:"auto_bill_outstanding"`
	SetupFeeFailureAction   string `json:"setup_fee_failure_action"`
	PaymentFailureThreshold int    `json:"payment_failure_threshold"`
}

type CreateBillingRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *BillingDetail `json:"response,omitempty"`
}

type BillingDetail struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
	UsageType   string  `json:"usage_type"`
	CreateTime  string  `json:"create_time"`
	Links       []*Link `json:"links"`
}
