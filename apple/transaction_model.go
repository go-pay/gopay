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
	AppAccountToken             string `json:"appAccountToken"`
	BundleId                    string `json:"bundleId"`
	Currency                    string `json:"currency"`
	Environment                 string `json:"environment"`
	ExpiresDate                 int64  `json:"expiresDate"`
	InAppOwnershipType          string `json:"inAppOwnershipType"`
	IsUpgraded                  bool   `json:"isUpgraded"`
	OfferDiscountType           string `json:"offerDiscountType"`
	OfferIdentifier             string `json:"offerIdentifier"`
	OfferType                   int    `json:"offerType"`
	OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`
	OriginalTransactionId       string `json:"originalTransactionId"`
	Price                       int64  `json:"price"`
	ProductId                   string `json:"productId"`
	PurchaseDate                int64  `json:"purchaseDate"`
	Quantity                    int    `json:"quantity"`
	RevocationDate              int64  `json:"revocationDate"`
	RevocationReason            string `json:"revocationReason"`
	SignedDate                  int64  `json:"signedDate"`
	Storefront                  string `json:"storefront"`
	StorefrontId                string `json:"storefrontId"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	TransactionId               string `json:"transactionId"`
	TransactionReason           string `json:"transactionReason"`
	Type                        string `json:"type"`
	WebOrderLineItemId          string `json:"webOrderLineItemId"`
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
