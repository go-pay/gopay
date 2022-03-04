package apple

import "github.com/golang-jwt/jwt"

type (
	// NotificationTypeV2 通知类型
	NotificationTypeV2 string

	// SubtypeV2 子类型
	SubtypeV2 string

	// NotificationV2SignedPayload
	// https://developer.apple.com/documentation/appstoreservernotifications/signedpayload
	NotificationV2SignedPayload struct {
		SignedPayload string `json:"signedPayload"`
	}

	// SubscriptionNotificationV2DecodedPayload
	// https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2decodedpayload
	SubscriptionNotificationV2DecodedPayload struct {
		NotificationType    NotificationTypeV2             `json:"notificationType"`
		Subtype             SubtypeV2                      `json:"subtype"`
		NotificationUUID    string                         `json:"notificationUUID"`
		NotificationVersion string                         `json:"notificationVersion"`
		Data                SubscriptionNotificationV2Data `json:"data"`
	}

	// SubscriptionNotificationV2Data
	// https://developer.apple.com/documentation/appstoreservernotifications/data
	SubscriptionNotificationV2Data struct {
		AppAppleID            int    `json:"appAppleId"`
		BundleID              string `json:"bundleId"`
		BundleVersion         string `json:"bundleVersion"`
		Environment           string `json:"environment"`
		SignedRenewalInfo     string `json:"signedRenewalInfo"`
		SignedTransactionInfo string `json:"signedTransactionInfo"`
	}

	// SubscriptionNotificationV2JWSDecodedHeader jws 解析头部
	SubscriptionNotificationV2JWSDecodedHeader struct {
		Alg string   `json:"alg"`
		Kid string   `json:"kid"`
		X5c []string `json:"x5c"`
	}

	// SignedRenewalInfo https://developer.apple.com/documentation/appstoreservernotifications/jwsrenewalinfodecodedpayload
	SignedRenewalInfo struct {
		AutoRenewProductId     string `json:"autoRenewProductId"`
		AutoRenewStatus        int64  `json:"autoRenewStatus"`
		ExpirationIntent       int64  `json:"expirationIntent"`
		GracePeriodExpiresDate int64  `json:"gracePeriodExpiresDate"`
		IsInBillingRetryPeriod bool   `json:"isInBillingRetryPeriod"`
		OfferIdentifier        string `json:"offerIdentifier"`
		OfferType              int64  `json:"offerType"` // 1:An introductory offer. 2:A promotional offer. 3:An offer with a subscription offer code.
		OriginalTransactionId  string `json:"originalTransactionId"`
		PriceIncreaseStatus    int64  `json:"priceIncreaseStatus"` // 0: The customer hasn’t responded to the subscription price increase. 1:The customer consented to the subscription price increase.
		ProductId              string `json:"productId"`
		SignedDate             int64  `json:"signedDate"`
	}

	// SignedTransactionInfo https://developer.apple.com/documentation/appstoreservernotifications/jwstransactiondecodedpayload
	SignedTransactionInfo struct {
		AppAccountToken             string `json:"appAccountToken"`
		BundleId                    string `json:"bundleId"`
		ExpiresDate                 int64  `json:"expiresDate"`
		InAppOwnershipType          string `json:"inAppOwnershipType"` // FAMILY_SHARED  PURCHASED
		IsUpgraded                  bool   `json:"isUpgraded"`
		OfferIdentifier             string `json:"offerIdentifier"`
		OfferType                   int64  `json:"offerType"` // 1:An introductory offer. 2:A promotional offer. 3:An offer with a subscription offer code.
		OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`
		OriginalTransactionId       string `json:"originalTransactionId"`
		ProductId                   string `json:"productId"`
		PurchaseDate                int64  `json:"purchaseDate"`
		Quantity                    int64  `json:"quantity"`
		RevocationDate              int64  `json:"revocationDate"`
		RevocationReason            string `json:"revocationReason"`
		SignedDate                  int64  `json:"signedDate"` // Auto-Renewable Subscription: An auto-renewable subscription.  Non-Consumable: A non-consumable in-app purchase.  Consumable: A consumable in-app purchase.  Non-Renewing Subscription: A non-renewing subcription.
		SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
		TransactionId               string `json:"transactionId"`
		Type                        string `json:"type"`
		WebOrderLineItemId          string `json:"webOrderLineItemId"`
	}

	JWSNotificationV2Payload struct {
		jwt.StandardClaims
		SubscriptionNotificationV2DecodedPayload
		RenewalInfo     *JWSSignedRenewalInfoPayload     `json:"renewalInfo"`
		TransactionInfo *JWSSignedTransactionInfoPayload `json:"transactionInfo"`
	}

	JWSSignedRenewalInfoPayload struct {
		jwt.StandardClaims
		SignedRenewalInfo
	}

	JWSSignedTransactionInfoPayload struct {
		jwt.StandardClaims
		SignedTransactionInfo
	}
)

// 通知类型常量
// https://developer.apple.com/documentation/appstoreservernotifications/notificationtype
const (
	NotificationTypeV2ConsumptionRequest     NotificationTypeV2 = "CONSUMPTION_REQUEST"
	NotificationTypeV2DidChangeRenewalPref   NotificationTypeV2 = "DID_CHANGE_RENEWAL_PREF"
	NotificationTypeV2DidChangeRenewalStatus NotificationTypeV2 = "DID_CHANGE_RENEWAL_STATUS"
	NotificationTypeV2DidFailToRenew         NotificationTypeV2 = "DID_FAIL_TO_RENEW"
	NotificationTypeV2DidRenew               NotificationTypeV2 = "DID_RENEW"
	NotificationTypeV2Expired                NotificationTypeV2 = "EXPIRED"
	NotificationTypeV2GracePeriodExpired     NotificationTypeV2 = "GRACE_PERIOD_EXPIRED"
	NotificationTypeV2OfferRedeemed          NotificationTypeV2 = "OFFER_REDEEMED"
	NotificationTypeV2PriceIncrease          NotificationTypeV2 = "PRICE_INCREASE"
	NotificationTypeV2Refund                 NotificationTypeV2 = "REFUND"
	NotificationTypeV2RefundDeclined         NotificationTypeV2 = "REFUND_DECLINED"
	NotificationTypeV2RenewalExtended        NotificationTypeV2 = "RENEWAL_EXTENDED"
	NotificationTypeV2Revoke                 NotificationTypeV2 = "REVOKE"
	NotificationTypeV2Subscribed             NotificationTypeV2 = "SUBSCRIBED"
)

// 子类型常量
// https://developer.apple.com/documentation/appstoreservernotifications/subtype
const (
	SubTypeV2InitialBuy        SubtypeV2 = "INITIAL_BUY"
	SubTypeV2Resubscribe       SubtypeV2 = "RESUBSCRIBE"
	SubTypeV2Downgrade         SubtypeV2 = "DOWNGRADE"
	SubTypeV2Upgrade           SubtypeV2 = "UPGRADE"
	SubTypeV2AutoRenewEnabled  SubtypeV2 = "AUTO_RENEW_ENABLED"
	SubTypeV2AutoRenewDisabled SubtypeV2 = "AUTO_RENEW_DISABLED"
	SubTypeV2Voluntary         SubtypeV2 = "VOLUNTARY"
	SubTypeV2BillingRetry      SubtypeV2 = "BILLING_RETRY"
	SubTypeV2PriceIncrease     SubtypeV2 = "PRICE_INCREASE"
	SubTypeV2GracePeriod       SubtypeV2 = "GRACE_PERIOD"
	SubTypeV2BillingRecovery   SubtypeV2 = "BILLING_RECOVERY"
	SubTypeV2Pending           SubtypeV2 = "PENDING"
	SubTypeV2Accepted          SubtypeV2 = "ACCEPTED"
)
