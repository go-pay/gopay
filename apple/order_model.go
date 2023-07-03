package apple

type LookUpOrderIdRsp struct {
	StatusCodeErr
	Status             int                 `json:"status,omitempty"` // 0-valid，1-invalid
	SignedTransactions []SignedTransaction `json:"signedTransactions,omitempty"`
}
