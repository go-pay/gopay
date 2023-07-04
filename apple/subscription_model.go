package apple

import "fmt"

type AllSubscriptionStatusesRsp struct {
	StatusCodeErr
	AppAppleId  int                                `json:"appAppleId"`
	BundleId    string                             `json:"bundleId"`
	Environment string                             `json:"environment"`
	Data        []*SubscriptionGroupIdentifierItem `json:"data"`
}

type SubscriptionGroupIdentifierItem struct {
	SubscriptionGroupIdentifier string                  `json:"subscriptionGroupIdentifier"`
	LastTransactions            []*LastTransactionsItem `json:"lastTransactions"`
}

type LastTransactionsItem struct {
	OriginalTransactionId string `json:"originalTransactionId"`
	Status                int    `json:"status"`
	SignedRenewalInfo     string `json:"signedRenewalInfo"`
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

func (d *LastTransactionsItem) DecodeRenewalInfo() (ri *RenewalInfo, err error) {
	if d.SignedRenewalInfo == "" {
		return nil, fmt.Errorf("SignedRenewalInfo is empty")
	}
	ri = &RenewalInfo{}
	if err = ExtractClaims(d.SignedRenewalInfo, ri); err != nil {
		return nil, err
	}
	return ri, nil
}

func (d *LastTransactionsItem) DecodeTransactionInfo() (ti *TransactionInfo, err error) {
	if d.SignedTransactionInfo == "" {
		return nil, fmt.Errorf("signedTransactionInfo is empty")
	}
	ti = &TransactionInfo{}
	if err = ExtractClaims(d.SignedTransactionInfo, ti); err != nil {
		return nil, err
	}
	return ti, nil
}
