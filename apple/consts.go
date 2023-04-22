package apple

const (
	hostUrl        = "https://api.storekit.itunes.apple.com"
	sandBoxHostUrl = "https://api.storekit-sandbox.itunes.apple.com"

	// Get Transaction History
	getTransactionHistory = "/inApps/v1/history/%s" // originalTransactionId

	// Get All Subscription Statuses
	getAllSubscriptionStatuses = "/inApps/v1/subscriptions/%s" // originalTransactionId
)
