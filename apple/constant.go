package apple

const (
	hostUrl        = "https://api.storekit.itunes.apple.com"
	sandBoxHostUrl = "https://api.storekit-sandbox.itunes.apple.com"

	// Get Transaction History
	getTransactionHistory = "/inApps/v1/history/%s" // transactionId

	// Get Transaction Info
	getTransactionInfo = "/inApps/v1/transactions/%s" // transactionId

	// Get All Subscription Statuses
	getAllSubscriptionStatuses = "/inApps/v1/subscriptions/%s" // transactionId

	// Send Consumption Information
	sendConsumptionInformation = "/inApps/v1/transactions/consumption/%s" // transactionId

	// Look Up Order ID
	lookUpOrderID = "/inApps/v1/lookup/%s" // orderId

	// Get Subscription Status
	getRefundHistory = "/inApps/v2/refund/lookup/%s" // transactionId

	// Get Notification History
	getNotificationHistory = "/inApps/v1/notifications/history"
)
