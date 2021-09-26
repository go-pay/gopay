package paypal

type AccessToken struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Appid       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

type EmptyRsp struct {
	Code  int    `json:"-"`
	Error string `json:"-"`
}

type CreateOrderRsp struct {
	Code     int          `json:"-"`
	Error    string       `json:"-"`
	Response *OrderDetail `json:"response,omitempty"`
}

type OrderDetailRsp struct {
	Code     int          `json:"-"`
	Error    string       `json:"-"`
	Response *OrderDetail `json:"response,omitempty"`
}

// ==================================分割==================================

type Patch struct {
	Op    string      `json:"op"` // The possible values are: add、remove、replace、move、copy、test
	Path  string      `json:"path,omitempty"`
	Value interface{} `json:"value"` // The value to apply. The remove operation does not require a value.
	From  string      `json:"from,omitempty"`
}

type OrderDetail struct {
	CreateTime    string          `json:"create_time"`
	UpdateTime    string          `json:"update_time"`
	Id            string          `json:"id"`
	PaymentSource *PaymentSource  `json:"payment_source,omitempty"`
	Intent        string          `json:"intent"`
	Payer         *Payer          `json:"payer,omitempty"`
	PurchaseUnits []*PurchaseUnit `json:"purchase_units,omitempty"`
	Status        string          `json:"status"` // CREATED、SAVED、APPROVED、VOIDED、COMPLETED、PAYER_ACTION_REQUIRED
	Links         []*Link         `json:"links,omitempty"`
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
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"` // Possible values: GET,POST,PUT,DELETE,HEAD,CONNECT,OPTIONS,PATCH
}
