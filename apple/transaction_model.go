package apple

type TransactionHistoryRsp struct {
	AppAppleId         int    `json:"appAppleId"`
	BundleId           string `json:"bundleId"`
	Environment        string `json:"environment"`
	HasMore            bool   `json:"hasMore"`
	Revision           string `json:"revision"`
	SignedTransactions string `json:"signedTransactions"`
}
