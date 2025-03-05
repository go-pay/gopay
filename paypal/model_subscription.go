package paypal

type SubscriptionCreateRsp struct {
	Code          int                 `json:"-"`
	Error         string              `json:"-"`
	ErrorResponse *ErrorResponse      `json:"-"`
	Response      *SubscriptionDetail `json:"response,omitempty"`
}

type SubscriptionDetail struct {
	ID               string          `json:"id"`
	Status           string          `json:"status"`
	StatusUpdateTime string          `json:"status_update_time"`
	PlanID           string          `json:"plan_id"`
	PlanOverridden   bool            `json:"plan_overridden"`
	StartTime        string          `json:"start_time"`
	Quantity         string          `json:"quantity"`
	ShippingAmount   *CommonAmount   `json:"shipping_amount"`
	Subscriber       *Subscriber     `json:"subscriber"`
	BillingInfo      *BillingInfoNew `json:"billing_info,omitempty"`
	CreateTime       string          `json:"create_time"`
	UpdateTime       string          `json:"update_time,omitempty"`
	Links            []*Link         `json:"links,omitempty"`
}

type Subscriber struct {
	Name            *Name            `json:"name"`
	EmailAddress    string           `json:"email_address"`
	PayerId         string           `json:"payer_id"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
	Name    *Name    `json:"name"`
	Address *Address `json:"address"`
}

type BillingInfoNew struct {
	OutstandingBalance  *CommonAmount     `json:"outstanding_balance"`
	CycleExecutions     []*CycleExecution `json:"cycle_executions"`
	LastPayment         *LastPayment      `json:"last_payment"`
	NextBillingTime     string            `json:"next_billing_time"`
	FailedPaymentsCount int               `json:"failed_payments_count"`
}

type CycleExecution struct {
	TenureType      string `json:"tenure_type"`
	Sequence        int    `json:"sequence"`
	CyclesCompleted int    `json:"cycles_completed"`
	CyclesRemaining int    `json:"cycles_remaining"`
	TotalCycles     int    `json:"total_cycles"`
}

type LastPayment struct {
	Amount *CommonAmount `json:"amount"`
	Time   string        `json:"time"`
}

type SubscriptionDetailRsp struct {
	Code          int                 `json:"-"`
	Error         string              `json:"-"`
	ErrorResponse *ErrorResponse      `json:"-"`
	Response      *SubscriptionDetail `json:"response,omitempty"`
}

type SubscriptionReviseRsp struct {
	Code          int                 `json:"-"`
	Error         string              `json:"-"`
	ErrorResponse *ErrorResponse      `json:"-"`
	Response      *ReviseSubscription `json:"response,omitempty"`
}

type ReviseSubscription struct {
	PlanId          string       `json:"plan_id"`
	Quantity        string       `json:"quantity"`
	PlanOverridden  bool         `json:"plan_overridden"`
	ShippingAmount  CommonAmount `json:"shipping_amount"`
	ShippingAddress struct {
		Type string `json:"type"`
		Name struct {
			FullName string `json:"full_name"`
		} `json:"name"`
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			AdminArea2   string `json:"admin_area_2"`
			AdminArea1   string `json:"admin_area_1"`
			PostalCode   string `json:"postal_code"`
			CountryCode  string `json:"country_code"`
		} `json:"address"`
	} `json:"shipping_address"`
	Plan struct {
		BillingCycles []struct {
			Sequence      int `json:"sequence"`
			TotalCycles   int `json:"total_cycles"`
			PricingScheme struct {
				Version      int    `json:"version"`
				PricingModel string `json:"pricing_model"`
				Tiers        []struct {
					StartingQuantity string `json:"starting_quantity"`
					EndingQuantity   string `json:"ending_quantity"`
					Amount           struct {
						CurrencyCode string `json:"currency_code"`
						Value        string `json:"value"`
					} `json:"amount"`
				} `json:"tiers"`
				FixedPrice FixedPrice `json:"fixed_price"`
				CreateTime string     `json:"create_time"`
				UpdateTime string     `json:"update_time"`
			} `json:"pricing_scheme"`
		} `json:"billing_cycles"`
		PaymentPreferences struct {
			AutoBillOutstanding     bool         `json:"auto_bill_outstanding"`
			SetupFeeFailureAction   string       `json:"setup_fee_failure_action"`
			PaymentFailureThreshold int          `json:"payment_failure_threshold"`
			SetupFee                CommonAmount `json:"setup_fee"`
		} `json:"payment_preferences"`
		Taxes struct {
			Inclusive  bool   `json:"inclusive"`
			Percentage string `json:"percentage"`
		} `json:"taxes"`
	} `json:"plan"`
	Links []struct {
		Href   string `json:"href"`
		Rel    string `json:"rel"`
		Method string `json:"method"`
	} `json:"links"`
}

type SubscriptionTransactionListRsp struct {
	Code          int                      `json:"-"`
	Error         string                   `json:"-"`
	ErrorResponse *ErrorResponse           `json:"-"`
	Response      *SubscriptionTransaction `json:"response,omitempty"`
}

type SubscriptionTransaction struct {
	Transactions []*Transaction `json:"transactions"`
	Links        []*Link        `json:"links,omitempty"`
}
type Transaction struct {
	ID                  string               `json:"id"`
	Status              string               `json:"status"`
	PayerEmail          string               `json:"payer_email"`
	PayerName           *Name                `json:"payer_name"`
	AmountWithBreakdown *AmountWithBreakdown `json:"amount_with_breakdown"`
	Time                string               `json:"time"`
}

type AmountWithBreakdown struct {
	GrossAmount *CommonAmount `json:"gross_amount"`
	FeeAmount   *CommonAmount `json:"fee_amount"`
	NetAmount   *CommonAmount `json:"net_amount"`
}
