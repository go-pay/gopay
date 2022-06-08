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
	Id            string          `json:"id,omitempty"`
	PaymentSource *PaymentSource  `json:"payment_source,omitempty"`
	Intent        string          `json:"intent,omitempty"`
	Payer         *Payer          `json:"payer,omitempty"`
	PurchaseUnits []*PurchaseUnit `json:"purchase_units,omitempty"`
	Status        string          `json:"status,omitempty"` // CREATED、SAVED、APPROVED、VOIDED、COMPLETED、PAYER_ACTION_REQUIRED
	Links         []*Link         `json:"links,omitempty"`
	CreateTime    string          `json:"create_time,omitempty"`
	UpdateTime    string          `json:"update_time,omitempty"`
}

type PaymentSource struct {
	Card *Card `json:"card,omitempty"`
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
	EmailAddress string `json:"email_address"`
	PayerId      string `json:"payer_id"`
}

type PurchaseUnit struct {
	ReferenceId        string              `json:"reference_id,omitempty"`
	Amount             *Amount             `json:"amount,omitempty"`
	Payee              *Payee              `json:"payee,omitempty"`
	PaymentInstruction *PaymentInstruction `json:"payment_instruction,omitempty"`
	Description        string              `json:"description,omitempty"`
	CustomId           string              `json:"custom_id,omitempty"`
	InvoiceId          string              `json:"invoice_id,omitempty"`
	Id                 string              `json:"id,omitempty"`
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
	FullName string `json:"full_name,omitempty"`
}

type Payments struct {
	Authorizations []*Authorization `json:"authorizations"`
	Captures       []*Capture       `json:"captures"`
	Refunds        []*Refund        `json:"refunds"`
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
	Status        string         `json:"status,omitempty"`
	StatusDetails *StatusDetails `json:"status_details,omitempty"`
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
	Status            string             `json:"status,omitempty"` // ELIGIBLE、PARTIALLY_ELIGIBLE、NOT_ELIGIBLE
	DisputeCategories []*DisputeCategory `json:"dispute_categories,omitempty"`
}

type DisputeCategory struct {
	DisputeCategory string `json:"dispute_category,omitempty"`
}

type PaymentAuthorizeCapture struct {
	Id                        string                     `json:"id,omitempty"`
	Status                    string                     `json:"status,omitempty"` // COMPLETED、DECLINED、PARTIALLY_REFUNDED、PENDING、REFUNDED、FAILED
	StatusDetails             *StatusDetails             `json:"status_details,omitempty"`
	Amount                    *Amount                    `json:"amount,omitempty"`
	InvoiceId                 string                     `json:"invoice_id,omitempty"`
	CustomId                  string                     `json:"custom_id,omitempty"`
	SellerProtection          *SellerProtection          `json:"seller_protection,omitempty"`
	FinalCapture              bool                       `json:"final_capture,omitempty"`
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
