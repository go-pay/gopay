package apple

import (
	"fmt"

	"github.com/go-pay/gopay/pkg/jwt"
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

// TransactionsItem
// Doc: https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
type TransactionsItem struct {
	jwt.StandardClaims
	TransactionId               string `json:"transactionId"`
	OriginalTransactionId       string `json:"originalTransactionId"`
	WebOrderLineItemId          string `json:"webOrderLineItemId"`
	BundleId                    string `json:"bundleId"`
	ProductId                   string `json:"productId"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	PurchaseDate                int64  `json:"purchaseDate"`
	OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`
	ExpiresDate                 int64  `json:"expiresDate"`
	Quantity                    int    `json:"quantity"`
	Type                        string `json:"type"`
	InAppOwnershipType          string `json:"inAppOwnershipType"`
	SignedDate                  int64  `json:"signedDate"`
	OfferType                   int    `json:"offerType"`
	Environment                 string `json:"environment"`
}

func (s *SignedTransaction) DecodeSignedTransaction() (ti *TransactionsItem, err error) {
	if *s == "" {
		return nil, fmt.Errorf("signedTransactions is empty")
	}
	ti = &TransactionsItem{}
	_, err = ExtractClaims(string(*s), ti)
	if err != nil {
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
	_, err = ExtractClaims(t.SignedTransactionInfo, ti)
	if err != nil {
		return nil, err
	}
	return ti, nil
}
