package apple

import (
	"fmt"
)

type SignedTransaction string

// Doc: HistoryResponse https://developer.apple.com/documentation/appstoreserverapi/historyresponse
type TransactionHistoryRsp struct {
	StatusCodeErr
	AppAppleId         int                 `json:"appAppleId"`
	BundleId           string              `json:"bundleId"`
	Environment        string              `json:"environment"`
	HasMore            bool                `json:"hasMore"`
	Revision           string              `json:"revision"`
	SignedTransactions []SignedTransaction `json:"signedTransactions"`
}

func (s *SignedTransaction) DecodeSignedTransaction() (ti *TransactionsItem, err error) {
	if *s == "" {
		return nil, fmt.Errorf("signedTransactions is empty")
	}
	ti = &TransactionsItem{}
	if err = ExtractClaims(string(*s), ti); err != nil {
		return nil, err
	}
	return ti, nil
}

// Doc: https://developer.apple.com/documentation/appstoreserverapi/transactioninforesponse
type TransactionInfoRsp struct {
	StatusCodeErr
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

func (t *TransactionInfoRsp) DecodeSignedTransaction() (ti *TransactionsItem, err error) {
	if t.SignedTransactionInfo == "" {
		return nil, fmt.Errorf("signedTransactionInfo is empty")
	}
	ti = &TransactionsItem{}
	if err = ExtractClaims(t.SignedTransactionInfo, ti); err != nil {
		return nil, err
	}
	return ti, nil
}
