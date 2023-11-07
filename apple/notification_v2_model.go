package apple

import (
	"fmt"

	"github.com/go-pay/gopay/pkg/jwt"
)

const (
	// 通知类型常量
	// https://developer.apple.com/documentation/appstoreservernotifications/notificationtype
	NotificationTypeV2ConsumptionRequest     = "CONSUMPTION_REQUEST"
	NotificationTypeV2DidChangeRenewalPref   = "DID_CHANGE_RENEWAL_PREF"
	NotificationTypeV2DidChangeRenewalStatus = "DID_CHANGE_RENEWAL_STATUS"
	NotificationTypeV2DidFailToRenew         = "DID_FAIL_TO_RENEW"
	NotificationTypeV2DidRenew               = "DID_RENEW"
	NotificationTypeV2Expired                = "EXPIRED"
	NotificationTypeV2GracePeriodExpired     = "GRACE_PERIOD_EXPIRED"
	NotificationTypeV2OfferRedeemed          = "OFFER_REDEEMED"
	NotificationTypeV2PriceIncrease          = "PRICE_INCREASE"
	NotificationTypeV2Refund                 = "REFUND"
	NotificationTypeV2RefundDeclined         = "REFUND_DECLINED"
	NotificationTypeV2RenewalExtended        = "RENEWAL_EXTENDED"
	NotificationTypeV2Revoke                 = "REVOKE"
	NotificationTypeV2Subscribed             = "SUBSCRIBED"

	// 子类型常量
	// https://developer.apple.com/documentation/appstoreservernotifications/subtype
	SubTypeV2InitialBuy        = "INITIAL_BUY"
	SubTypeV2Resubscribe       = "RESUBSCRIBE"
	SubTypeV2Downgrade         = "DOWNGRADE"
	SubTypeV2Upgrade           = "UPGRADE"
	SubTypeV2AutoRenewEnabled  = "AUTO_RENEW_ENABLED"
	SubTypeV2AutoRenewDisabled = "AUTO_RENEW_DISABLED"
	SubTypeV2Voluntary         = "VOLUNTARY"
	SubTypeV2BillingRetry      = "BILLING_RETRY"
	SubTypeV2PriceIncrease     = "PRICE_INCREASE"
	SubTypeV2GracePeriod       = "GRACE_PERIOD"
	SubTypeV2BillingRecovery   = "BILLING_RECOVERY"
	SubTypeV2Pending           = "PENDING"
	SubTypeV2Accepted          = "ACCEPTED"
)

// https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2
type NotificationV2Req struct {
	SignedPayload string `json:"signedPayload"`
}

// https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2decodedpayload
type NotificationV2Payload struct {
	jwt.StandardClaims
	NotificationType string `json:"notificationType"`
	Subtype          string `json:"subtype"`
	NotificationUUID string `json:"notificationUUID"`
	Version          string `json:"version"`
	Data             *Data  `json:"data"`
}

func (d *NotificationV2Payload) DecodeRenewalInfo() (ri *RenewalInfo, err error) {
	if d.Data == nil {
		return nil, fmt.Errorf("data is nil")
	}
	if d.Data.SignedRenewalInfo == "" {
		return nil, fmt.Errorf("data.signedRenewalInfo is empty")
	}
	ri = &RenewalInfo{}
	if err = ExtractClaims(d.Data.SignedRenewalInfo, ri); err != nil {
		return nil, err
	}
	return
}

func (d *NotificationV2Payload) DecodeTransactionInfo() (ti *TransactionInfo, err error) {
	if d.Data == nil {
		return nil, fmt.Errorf("data is nil")
	}
	if d.Data.SignedTransactionInfo == "" {
		return nil, fmt.Errorf("data.signedTransactionInfo is empty")
	}
	ti = &TransactionInfo{}
	if err = ExtractClaims(d.Data.SignedTransactionInfo, ti); err != nil {
		return nil, err
	}
	return
}

// https://developer.apple.com/documentation/appstoreservernotifications/data
type Data struct {
	AppAppleID            int    `json:"appAppleId"`
	BundleID              string `json:"bundleId"`
	BundleVersion         string `json:"bundleVersion"`
	Environment           string `json:"environment"`
	SignedRenewalInfo     string `json:"signedRenewalInfo"`
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

// RenewalInfo https://developer.apple.com/documentation/appstoreservernotifications/jwsrenewalinfodecodedpayload
type RenewalInfo struct {
	jwt.StandardClaims
	AutoRenewProductId          string `json:"autoRenewProductId"`
	AutoRenewStatus             int64  `json:"autoRenewStatus"`
	Environment                 string `json:"environment"`
	ExpirationIntent            int64  `json:"expirationIntent"`
	GracePeriodExpiresDate      int64  `json:"gracePeriodExpiresDate"`
	IsInBillingRetryPeriod      bool   `json:"isInBillingRetryPeriod"`
	OfferIdentifier             string `json:"offerIdentifier"`
	OfferType                   int64  `json:"offerType"` // 1:An introductory offer. 2:A promotional offer. 3:An offer with a subscription offer code.
	OriginalTransactionId       string `json:"originalTransactionId"`
	PriceIncreaseStatus         int64  `json:"priceIncreaseStatus"` // 0: The customer hasn’t responded to the subscription price increase. 1:The customer consented to the subscription price increase.
	ProductId                   string `json:"productId"`
	RecentSubscriptionStartDate int64  `json:"recentSubscriptionStartDate"`
	RenewalDate                 int64  `json:"renewalDate,omitempty"` // The UNIX time, in milliseconds, that the most recent auto-renewable subscription purchase expires.
	SignedDate                  int64  `json:"signedDate"`
}

// TransactionInfo https://developer.apple.com/documentation/appstoreservernotifications/jwstransactiondecodedpayload
type TransactionInfo struct {
	jwt.StandardClaims
	AppAccountToken             string `json:"appAccountToken"`
	BundleId                    string `json:"bundleId"`
	Currency                    string `json:"currency"`
	Environment                 string `json:"environment"`
	ExpiresDate                 int64  `json:"expiresDate"`
	InAppOwnershipType          string `json:"inAppOwnershipType"` // FAMILY_SHARED  PURCHASED
	IsUpgraded                  bool   `json:"isUpgraded"`
	OfferDiscountType           string `json:"offerDiscountType"`
	OfferIdentifier             string `json:"offerIdentifier"`
	OfferType                   int64  `json:"offerType"` // 1:An introductory offer. 2:A promotional offer. 3:An offer with a subscription offer code.
	OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`
	OriginalTransactionId       string `json:"originalTransactionId"`
	Price                       int64  `json:"price"`
	ProductId                   string `json:"productId"`
	PurchaseDate                int64  `json:"purchaseDate"`
	Quantity                    int64  `json:"quantity"`
	RevocationDate              int64  `json:"revocationDate"`
	RevocationReason            int    `json:"revocationReason"`
	SignedDate                  int64  `json:"signedDate"` // Auto-Renewable Subscription: An auto-renewable subscription.  Non-Consumable: A non-consumable in-app purchase.  Consumable: A consumable in-app purchase.  Non-Renewing Subscription: A non-renewing subcription.
	Storefront                  string `json:"storefront"`
	StorefrontId                string `json:"storefrontId"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	TransactionId               string `json:"transactionId"`
	TransactionReason           string `json:"transactionReason"`
	Type                        string `json:"type"`
	WebOrderLineItemId          string `json:"webOrderLineItemId"`
}

type NotificationHistoryRsp struct {
	StatusCodeErr
	HasMore             bool                `json:"hasMore"`
	PaginationToken     string              `json:"paginationToken"`
	NotificationHistory []*NotificationItem `json:"notificationHistory"`
}

type NotificationItem struct {
	SendAttempts  []*SendAttemptItem `json:"sendAttempts"`
	SignedPayload string             `json:"signedPayload"`
}

type SendAttemptItem struct {
	AttemptDate       int64  `json:"attemptDate"`
	SendAttemptResult string `json:"sendAttemptResult"`
}
