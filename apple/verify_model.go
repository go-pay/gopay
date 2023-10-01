package apple

// VerifyRequest 校验请求体
// https://developer.apple.com/documentation/appstorereceipts/requestbody
type VerifyRequest struct {
	// Receipt app解析出的票据信息
	Receipt string `json:"receipt-data"`

	// Password App的秘钥
	Password string `json:"password,omitempty"`

	// ExcludeOldTranscations Set this value to true for the response to include only the latest renewal transaction for any subscriptions. Use this field only for app receipts that contain auto-renewable subscriptions.
	ExcludeOldTranscations bool `json:"exclude-old-transactions"`
}

// VerifyResponse 校验响应体
// https://developer.apple.com/documentation/appstorereceipts/responsebody
type VerifyResponse struct {
	// Environment is which the receipt was generated. Possible values: Sandbox, Production
	Environment string `json:"environment"`

	// IsRetryable is an indicator that an error occurred during the request. A value of 1 indicates a temporary issue; retry validation for this receipt at a later time. A value of 0 indicates an unresolvable issue; do not retry validation for this receipt. Only applicable to status codes 21100-21199.
	IsRetryable bool `json:"is-retryable"`

	// LatestReceipt The latest Base64 encoded app receipt. Only returned for receipts that contain auto-renewable subscriptions
	LatestReceipt string `json:"latest_receipt,omitempty"`

	// LatestReceiptInfo is an array that contains all in-app purchase transactions. This excludes transactions for consumable products that have been marked as finished by your app. Only returned for receipts that contain auto-renewable subscriptions.
	LatestReceiptInfo []*LatestReceiptInfo `json:"latest_receipt_info,omitempty"`

	// PendingRenewalInfo ,in the JSON file, an array where each element contains the pending renewal information for each auto-renewable subscription identified by the product_id. Only returned for app receipts that contain auto-renewable subscriptions.
	PendingRenewalInfo []*PendingRenewalInfo `json:"pending_renewal_info,omitempty"`

	// Receipt is a JSON representation of the receipt that was sent for verification.
	Receipt *Receipt `json:"receipt,omitempty"`

	// Status either 0 if the receipt is valid, or a status code if there is an error. The status code reflects the status of the app receipt as a whole. See status for possible status codes and descriptions.
	// =0时就表示校验成功
	Status int `json:"status"`
}

// LatestReceiptInfo
// https://developer.apple.com/documentation/appstorereceipts/responsebody/latest_receipt_info
type LatestReceiptInfo struct {
	// The time Apple customer support canceled a transaction, in a date-time format similar to the ISO 8601. This field is only present for refunded transactions.
	CancellationDate string `json:"cancellation_date"`

	// The time Apple customer support canceled a transaction, or the time an auto-renewable subscription plan was upgraded, in UNIX epoch time format, in milliseconds. This field is only present for refunded transactions. Use this time format for processing dates.
	// https://developer.apple.com/documentation/appstorereceipts/cancellation_date_ms
	CancellationDateTimestamp string `json:"cancellation_date_ms"`

	// The time Apple customer support canceled a transaction, in the Pacific Time zone. This field is only present for refunded transactions.
	CancellationDatePST string `json:"cancellation_date_pst"`

	// The reason for a refunded transaction. When a customer cancels a transaction, the App Store gives them a refund and provides a value for this key. A value of “1” indicates that the customer canceled their transaction due to an actual or perceived issue within your app. A value of “0” indicates that the transaction was canceled for another reason; for example, if the customer made the purchase accidentally.
	// Possible values: 1, 0
	CancellationReason string `json:"cancellation_reason"`

	// The time a subscription expires or when it will renew, in a date-time format similar to the ISO 8601.
	ExpiresDate string `json:"expires_date"`

	// The time a subscription expires or when it will renew, in UNIX epoch time format, in milliseconds. Use this time format for processing dates.
	// https://developer.apple.com/documentation/appstorereceipts/expires_date_ms
	ExpiresDateTimestamp string `json:"expires_date_ms"`

	// The time a subscription expires or when it will renew, in the Pacific Time zone.
	ExpiresDatePST string `json:"expires_date_pst"`

	// A value that indicates whether the user is the purchaser of the product, or is a family member with access to the product through Family Sharing.
	// https://developer.apple.com/documentation/appstorereceipts/in_app_ownership_type
	InAppOwnershipType string `json:"in_app_ownership_type"`

	// An indicator of whether an auto-renewable subscription is in the introductory price period.
	// Possible values: true, false
	IsInIntroOfferPeriod string `json:"is_in_intro_offer_period"`

	// An indicator of whether a subscription is in the free trial period.
	// https://developer.apple.com/documentation/appstorereceipts/is_trial_period
	IsTrialPeriod string `json:"is_trial_period"`

	// An indicator that a subscription has been canceled due to an upgrade. This field is only present for upgrade transactions.
	// Value: true
	IsUpgraded string `json:"is_upgraded"`

	// The reference name of a subscription offer that you configured in App Store Connect. This field is present when a customer redeemed a subscription offer code. For more information about offer codes
	// https://help.apple.com/app-store-connect/#/dev6a098e4b1
	// https://developer.apple.com/documentation/storekit/original_api_for_in-app_purchase/subscriptions_and_offers/implementing_offer_codes_in_your_app
	OfferCodeRefName string `json:"offer_code_ref_name"`

	// The time of the original app purchase, in a date-time format similar to ISO 8601.
	OriginalPurchaseDate string `json:"original_purchase_date"`

	// The time of the original app purchase, in UNIX epoch time format, in milliseconds. Use this time format for processing dates. For an auto-renewable subscription, this value indicates the date of the subscription’s initial purchase. The original purchase date applies to all product types and remains the same in all transactions for the same product ID. This value corresponds to the original transaction’s transactionDate property in StoreKit.
	OriginalPurchaseDateTimestamp string `json:"original_purchase_date_ms"`

	// The time of the original app purchase, in the Pacific Time zone.
	OriginalPurchaseDatePST string `json:"original_purchase_date_pst"`

	// The transaction identifier of the original purchase.
	// https://developer.apple.com/documentation/appstorereceipts/original_transaction_id
	OriginalTransactionId string `json:"original_transaction_id"`

	// The unique identifier of the product purchased. You provide this value when creating the product in App Store Connect, and it corresponds to the productIdentifier property of the SKPayment object stored in the transaction’s payment property.
	ProductId string `json:"product_id"`

	// The identifier of the subscription offer redeemed by the user.
	// https://developer.apple.com/documentation/appstorereceipts/promotional_offer_id
	PromotionalOfferId string `json:"promotional_offer_id"`

	// The time the App Store charged the user’s account for a purchased or restored product, or the time the App Store charged the user’s account for a subscription purchase or renewal after a lapse, in a date-time format similar to ISO 8601.
	PurchaseDate string `json:"purchase_date"`

	// For consumable, non-consumable, and non-renewing subscription products, the time the App Store charged the user’s account for a purchased or restored product, in the UNIX epoch time format, in milliseconds. For auto-renewable subscriptions, the time the App Store charged the user’s account for a subscription purchase or renewal after a lapse, in the UNIX epoch time format, in milliseconds. Use this time format for processing dates.
	PurchaseDateTimestamp string `json:"purchase_date_ms"`

	// The time the App Store charged the user’s account for a purchased or restored product, or the time the App Store charged the user’s account for a subscription purchase or renewal after a lapse, in the Pacific Time zone.
	PurchaseDatePST string `json:"purchase_date_pst"`

	// The number of consumable products purchased. This value corresponds to the quantity property of the SKPayment object stored in the transaction’s payment property. The value is usually “1” unless modified with a mutable payment. The maximum value is 10.
	Quantity string `json:"quantity"`

	// The identifier of the subscription group to which the subscription belongs. The value for this field is identical to the subscriptionGroupIdentifier property in SKProduct.
	// https://developer.apple.com/documentation/storekit/skproduct/2981047-subscriptiongroupidentifier
	SubscriptionGroupIdentifier string `json:"subscription_group_identifier"`

	// A unique identifier for purchase events across devices, including subscription-renewal events. This value is the primary key for identifying subscription purchases.
	WebOrderLineItemId string `json:"web_order_line_item_id"`

	// A unique identifier for a transaction such as a purchase, restore, or renewal
	TransactionId string `json:"transaction_id"`

	// https://developer.apple.com/documentation/appstorereceipts/app_account_token
	AppAccountToken string `json:"app_account_token"`
}

// PendingRenewalInfo
// https://developer.apple.com/documentation/appstorereceipts/responsebody/pending_renewal_info
type PendingRenewalInfo struct {
	// The value for this key corresponds to the productIdentifier property of the product that the customer’s subscription renews.
	AutoRenewProductId string `json:"auto_renew_product_id"`

	// The current renewal status for the auto-renewable subscription.
	// https://developer.apple.com/documentation/appstorereceipts/auto_renew_status
	AutoRenewStatus string `json:"auto_renew_status"`

	// The reason a subscription expired. This field is only present for a receipt that contains an expired auto-renewable subscription.
	// https://developer.apple.com/documentation/appstorereceipts/expiration_intent
	ExpirationIntent string `json:"expiration_intent"`

	// The time at which the grace period for subscription renewals expires, in a date-time format similar to the ISO 8601.
	GracePeriodExpiresDate string `json:"grace_period_expires_date"`

	// The time at which the grace period for subscription renewals expires, in UNIX epoch time format, in milliseconds. This key is only present for apps that have Billing Grace Period enabled and when the user experiences a billing error at the time of renewal. Use this time format for processing dates.
	GracePeriodExpiresDateTimestamp string `json:"grace_period_expires_date_ms"`

	// The time at which the grace period for subscription renewals expires, in the Pacific Time zone.
	GracePeriodExpiresDatePST string `json:"grace_period_expires_date_pst"`

	// A flag that indicates Apple is attempting to renew an expired subscription automatically. This field is only present if an auto-renewable subscription is in the billing retry state.
	// https://developer.apple.com/documentation/appstorereceipts/is_in_billing_retry_period
	IsInBillingRetryPeriod string `json:"is_in_billing_retry_period"`

	// The reference name of a subscription offer that you configured in App Store Connect. This field is present when a customer redeemed a subscription offer code
	// https://developer.apple.com/documentation/appstorereceipts/offer_code_ref_name
	OfferCodeRefName string `json:"offer_code_ref_name"`

	// The transaction identifier of the original purchase.
	OriginalTransactionId string `json:"original_transaction_id"`

	// The price consent status for a subscription price increase. This field is only present if the customer was notified of the price increase. The default value is "0" and changes to "1" if the customer consents.
	// Possible values: 1, 0
	PriceConsentStatus string `json:"price_consent_status"`

	// The unique identifier of the product purchased. You provide this value when creating the product in App Store Connect, and it corresponds to the productIdentifier property of the SKPayment object stored in the transaction's payment property.
	// https://developer.apple.com/documentation/storekit/skpayment
	ProductId string `json:"product_id"`

	// The identifier of the promotional offer for an auto-renewable subscription that the user redeemed. You provide this value in the Promotional Offer Identifier field when you create the promotional offer in App Store Connect.
	// https://developer.apple.com/documentation/appstorereceipts/promotional_offer_id
	Promotionalofferid string `json:"promotional_offer_id"`
}

// Receipt is the decoded version of the encoded receipt data sent with the request to the App Store
// https://developer.apple.com/documentation/appstorereceipts/responsebody/receipt
type Receipt struct {
	// See app_item_id.
	AdamId int64 `json:"adam_id"`

	// Generated by App Store Connect and used by the App Store to uniquely identify the app purchased. Apps are assigned this identifier only in production. Treat this value as a 64-bit long integer.
	AppItemId int64 `json:"app_item_id"`

	// The app’s version number. The app's version number corresponds to the value of CFBundleVersion (in iOS) or CFBundleShortVersionString (in macOS) in the Info.plist. In production, this value is the current version of the app on the device based on the receipt_creation_date_ms. In the sandbox, the value is always "1.0".
	ApplicationVersion string `json:"application_version"`

	// The bundle identifier for the app to which the receipt belongs. You provide this string on App Store Connect. This corresponds to the value of CFBundleIdentifier in the Info.plist file of the app.
	BundleId string `json:"bundle_id"`

	// A unique identifier for the app download transaction.
	DownloadId int64 `json:"download_id"`

	// The time the receipt expires for apps purchased through the Volume Purchase Program, in a date-time format similar to the ISO 8601.
	ExpirationDate string `json:"expiration_date"`

	// The time the receipt expires for apps purchased through the Volume Purchase Program, in UNIX epoch time format, in milliseconds. If this key is not present for apps purchased through the Volume Purchase Program, the receipt does not expire. Use this time format for processing dates.
	ExpirationDateTimestamp string `json:"expiration_date_ms"`

	// The time the receipt expires for apps purchased through the Volume Purchase Program, in the Pacific Time zone.
	ExpirationDatePST string `json:"expiration_date_pst"`

	// An array that contains the in-app purchase receipt fields for all in-app purchase transactions.
	InApp []*InApp `json:"in_app,omitempty"`

	// The version of the app that the user originally purchased. This value does not change, and corresponds to the value of CFBundleVersion (in iOS) or CFBundleShortVersionString (in macOS) in the Info.plist file of the original purchase. In the sandbox environment, the value is always "1.0".
	OriginalApplicationVersion string `json:"original_application_version"`

	// The time of the original app purchase, in a date-time format similar to ISO 8601.
	OriginalPurchaseDate string `json:"original_purchase_date"`

	// The time of the original app purchase, in UNIX epoch time format, in milliseconds. Use this time format for processing dates.
	OriginalPurchaseDateTimestamp string `json:"original_purchase_date_ms"`

	// The time of the original app purchase, in the Pacific Time zone.
	OriginalPurchaseDatePST string `json:"original_purchase_date_pst"`

	// The time the user ordered the app available for pre-order, in a date-time format similar to ISO 8601.
	PreorderDate string `json:"preorder_date"`

	// The time the user ordered the app available for pre-order, in UNIX epoch time format, in milliseconds. This field is only present if the user pre-orders the app. Use this time format for processing dates.
	PreorderDateTimestamp string `json:"preorder_date_ms"`

	// The time the user ordered the app available for pre-order, in the Pacific Time zone.
	PreorderDatePST string `json:"preorder_date_pst"`

	// The time the App Store generated the receipt, in a date-time format similar to ISO 8601.
	ReceiptCreationDate string `json:"receipt_creation_date"`

	// The time the App Store generated the receipt, in UNIX epoch time format, in milliseconds. Use this time format for processing dates. This value does not change.
	ReceiptCreationDateTimestamp string `json:"receipt_creation_date_ms"`

	// The time the App Store generated the receipt, in the Pacific Time zone.
	ReceiptCreationDatePST string `json:"receipt_creation_date_pst"`

	// The type of receipt generated. The value corresponds to the environment in which the app or VPP purchase was made.
	//  Possible values: Production, ProductionVPP, ProductionSandbox, ProductionVPPSandbox
	ReceiptType string `json:"receipt_type"`

	// The time the request to the verifyReceipt endpoint was processed and the response was generated, in a date-time format similar to ISO 8601.
	RequestDate string `json:"request_date"`

	// The time the request to the verifyReceipt endpoint was processed and the response was generated, in UNIX epoch time format, in milliseconds. Use this time format for processing dates.
	RequestDateTimestamp string `json:"request_date_ms"`

	// The time the request to the verifyReceipt endpoint was processed and the response was generated, in the Pacific Time zone.
	RequestDatePST string `json:"request_date_pst"`

	// An arbitrary number that identifies a revision of your app. In the sandbox, this key's value is 0.
	VersionExternalIdentifier int64 `json:"version_external_identifier"`
}

// InApp is the in-app purchase receipt fields for all in-app purchase transactions.
// https://developer.apple.com/documentation/appstorereceipts/responsebody/receipt/in_app
type InApp struct {
	// The time the App Store refunded a transaction or revoked it from family sharing, in a date-time format similar to the ISO 8601. This field is present only for refunded or revoked transactions.
	CancellationDate string `json:"cancellation_date"`

	// The time the App Store refunded a transaction or revoked it from family sharing, in UNIX epoch time format, in milliseconds. This field is present only for refunded or revoked transactions. Use this time format for processing dates. The time the App Store refunded a transaction or revoked it from family sharing, in UNIX epoch time format, in milliseconds. This field is present only for refunded or revoked transactions. Use this time format for processing dates.
	// https://developer.apple.com/documentation/appstorereceipts/cancellation_date_ms
	CancellationDateTimestamp string `json:"cancellation_date_ms"`

	// The time Apple customer support canceled a transaction, in the Pacific Time zone. This field is only present for refunded transactions.
	CancellationDatePST string `json:"cancellation_date_pst"`

	// The reason for a refunded transaction. When a customer cancels a transaction, the App Store gives them a refund and provides a value for this key. A value of “1” indicates that the customer canceled their transaction due to an actual or perceived issue within your app. A value of “0” indicates that the transaction was canceled for another reason; for example, if the customer made the purchase accidentally.
	// Possible values: 1, 0
	CancellationReason string `json:"cancellation_reason"`

	// The time a subscription expires or when it will renew, in a date-time format similar to the ISO 8601.
	ExpiresDate string `json:"expires_date"`

	// The time a subscription expires or when it will renew, in UNIX epoch time format, in milliseconds. Use this time format for processing dates.
	// https://developer.apple.com/documentation/appstorereceipts/expires_date_ms
	ExpiresDateTimestamp string `json:"expires_date_ms"`

	// The time a subscription expires or when it will renew, in the Pacific Time zone.
	ExpiresDatePST string `json:"expires_date_pst"`

	// An indicator of whether an auto-renewable subscription is in the introductory price period.
	// https://developer.apple.com/documentation/appstorereceipts/is_in_intro_offer_period
	IsInIntroOfferPeriod string `json:"is_in_intro_offer_period"`

	// An indication of whether a subscription is in the free trial period.
	// https://developer.apple.com/documentation/appstorereceipts/is_trial_period
	IsTrialPeriod string `json:"is_trial_period"`

	// The time of the original in-app purchase, in a date-time format similar to ISO 8601.
	OriginalPurchaseDate string `json:"original_purchase_date"`

	// The time of the original in-app purchase, in UNIX epoch time format, in milliseconds. For an auto-renewable subscription, this value indicates the date of the subscription's initial purchase. The original purchase date applies to all product types and remains the same in all transactions for the same product ID. This value corresponds to the original transaction’s transactionDate property in StoreKit. Use this time format for processing dates.
	OriginalPurchaseDateTimestamp string `json:"original_purchase_date_ms"`

	// The time of the original in-app purchase, in the Pacific Time zone.
	OriginalPurchaseDatePST string `json:"original_purchase_date_pst"`

	// The transaction identifier of the original purchase.
	// https://developer.apple.com/documentation/appstorereceipts/original_transaction_id
	OriginalTransactionId string `json:"original_transaction_id"`

	// The unique identifier of the product purchased. You provide this value when creating the product in App Store Connect, and it corresponds to the productIdentifier property of the SKPayment object stored in the transaction's payment property.
	ProductId string `json:"product_id"`

	// The identifier of the subscription offer redeemed by the user.
	// https://developer.apple.com/documentation/appstorereceipts/promotional_offer_id
	PromotionalOfferId string `json:"promotional_offer_id"`

	// The time the App Store charged the user's account for a purchased or restored product, or the time the App Store charged the user’s account for a subscription purchase or renewal after a lapse, in a date-time format similar to ISO 8601.
	PurchaseDate string `json:"purchase_date"`

	// For consumable, non-consumable, and non-renewing subscription products, the time the App Store charged the user's account for a purchased or restored product, in the UNIX epoch time format, in milliseconds. For auto-renewable subscriptions, the time the App Store charged the user’s account for a subscription purchase or renewal after a lapse, in the UNIX epoch time format, in milliseconds. Use this time format for processing dates.
	PurchaseDateTimestamp string `json:"purchase_date_ms"`

	// The time the App Store charged the user's account for a purchased or restored product, or the time the App Store charged the user’s account for a subscription purchase or renewal after a lapse, in the Pacific Time zone.
	PurchaseDatePST string `json:"purchase_date_pst"`

	// The number of consumable products purchased. This value corresponds to the quantity property of the SKPayment object stored in the transaction's payment property. The value is usually “1” unless modified with a mutable payment. The maximum value is 10.
	Quantity string `json:"quantity"`

	// A unique identifier for a transaction such as a purchase, restore, or renewal. See transaction_id for more information.
	TransactionId string `json:"transaction_id"`

	// A unique identifier for purchase events across devices, including subscription-renewal events. This value is the primary key for identifying subscription purchases.
	WebOrderLineItemId string `json:"web_order_line_item_id"`
}
