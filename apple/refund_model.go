package apple

type RefundHistoryRsp struct {
	StatusCodeErr
	HasMore            bool                `json:"hasMore"`
	Revision           string              `json:"revision"`
	SignedTransactions []SignedTransaction `json:"signedTransactions"`
}
