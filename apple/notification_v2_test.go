package apple

import (
	"testing"

	"github.com/go-pay/gopay/pkg/xlog"
)

func TestDecodeSignedPayload(t *testing.T) {
	signedPayload := `eyJhbGciOiJFUzI1NiIsIng1YyI6WyJNSUlOW...mnpo2QrItvA`
	// decode signedPayload
	payload, err := DecodeSignedPayload(signedPayload)
	if err != nil {
		xlog.Error(err)
		return
	}
	/*
		{
		    "notificationType":"DID_RENEW",
		    "subtype":"",
		    "notificationUUID":"7a33cb9d-2503-4104-a1e9-24bb9ea8bb75",
		    "notificationVersion":"",
		    "data":{
		        "appAppleId":0,
		        "bundleId":"bundleId",
		        "bundleVersion":"7",
		        "environment":"Sandbox",
		        "signedRenewalInfo":"xxxxxxxxxxxxxxxx",
		        "signedTransactionInfo":"xxxxxxxxxxxxx"
		    }
		}
	*/
	xlog.Debugf("payload.NotificationType:", payload.NotificationType)
	xlog.Debugf("payload.Subtype:", payload.Subtype)
	xlog.Debugf("payload.NotificationUUID:", payload.NotificationUUID)
	xlog.Debugf("payload.NotificationVersion:", payload.NotificationVersion)

	// decode renewalInfo
	renewalInfo, err := payload.DecodeRenewalInfo()
	if err != nil {
		xlog.Error(err)
		return
	}
	/*
		{
		    "autoRenewProductId":"bundleId.productId",
		    "autoRenewStatus":1,
		    "expirationIntent":0,
		    "gracePeriodExpiresDate":0,
		    "isInBillingRetryPeriod":false,
		    "offerIdentifier":"",
		    "offerType":0,
		    "originalTransactionId":"2000000000842607",
		    "priceIncreaseStatus":0,
		    "productId":"bundleId.productId",
		    "signedDate":1646327704113
		}
	*/
	xlog.Debugf("data.renewalInfo: %+v", renewalInfo)

	// decode transactionInfo
	transactionInfo, err := payload.DecodeTransactionInfo()
	if err != nil {
		xlog.Error(err)
		return
	}
	/*
		{
		    "appAccountToken":"",
		    "bundleId":"bundleId",
		    "expiresDate":1646329907000,
		    "inAppOwnershipType":"PURCHASED",
		    "isUpgraded":false,
		    "offerIdentifier":"",
		    "offerType":0,
		    "originalPurchaseDate":1646046037000,
		    "originalTransactionId":"2000000000842607",
		    "productId":"bundleId.productId",
		    "purchaseDate":1646327747000,
		    "quantity":1,
		    "revocationDate":0,
		    "revocationReason":"",
		    "signedDate":1646327704135,
		    "subscriptionGroupIdentifier":"20929536",
		    "transactionId":"2000000003595767",
		    "type":"Auto-Renewable Subscription",
		    "webOrderLineItemId":"2000000000264942"
		}
	*/
	xlog.Debugf("data.transactionInfo: %+v", transactionInfo)
}
