package apple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// rootPEM is from `openssl x509 -inform der -in AppleRootCA-G3.cer -out apple_root.pem`
const rootPEM = `
-----BEGIN CERTIFICATE-----
MIICQzCCAcmgAwIBAgIILcX8iNLFS5UwCgYIKoZIzj0EAwMwZzEbMBkGA1UEAwwS
QXBwbGUgUm9vdCBDQSAtIEczMSYwJAYDVQQLDB1BcHBsZSBDZXJ0aWZpY2F0aW9u
IEF1dGhvcml0eTETMBEGA1UECgwKQXBwbGUgSW5jLjELMAkGA1UEBhMCVVMwHhcN
MTQwNDMwMTgxOTA2WhcNMzkwNDMwMTgxOTA2WjBnMRswGQYDVQQDDBJBcHBsZSBS
b290IENBIC0gRzMxJjAkBgNVBAsMHUFwcGxlIENlcnRpZmljYXRpb24gQXV0aG9y
aXR5MRMwEQYDVQQKDApBcHBsZSBJbmMuMQswCQYDVQQGEwJVUzB2MBAGByqGSM49
AgEGBSuBBAAiA2IABJjpLz1AcqTtkyJygRMc3RCV8cWjTnHcFBbZDuWmBSp3ZHtf
TjjTuxxEtX/1H7YyYl3J6YRbTzBPEVoA/VhYDKX1DyxNB0cTddqXl5dvMVztK517
IDvYuVTZXpmkOlEKMaNCMEAwHQYDVR0OBBYEFLuw3qFYM4iapIqZ3r6966/ayySr
MA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgEGMAoGCCqGSM49BAMDA2gA
MGUCMQCD6cHEFl4aXTQY2e3v9GwOAEZLuN+yRhHFD/3meoyhpmvOwgPUnPWTxnS4
at+qIxUCMG1mihDK1A3UT82NQz60imOlM27jbdoXt2QfyFMm+YhidDkLF1vLUagM
6BgD56KyKA==
-----END CERTIFICATE-----
`

// DecodeSignedPayload 解析SignedPayload数据
func DecodeSignedPayload(signedPayload string) (payload *NotificationV2Payload, err error) {
	if signedPayload == "" {
		return nil, fmt.Errorf("signedPayload is empty")
	}
	payload = &NotificationV2Payload{}
	if err = ExtractClaims(signedPayload, payload); err != nil {
		return nil, err
	}
	return
}

// GetNotificationHistory Get Notification History
// rsp.NotificationHistory[x].SignedPayload use apple.DecodeSignedPayload() to decode
// Doc: https://developer.apple.com/documentation/appstoreserverapi/get_notification_history
func (c *Client) GetNotificationHistory(ctx context.Context, paginationToken string, bm gopay.BodyMap) (rsp *NotificationHistoryRsp, err error) {
	path := getNotificationHistory
	// Note: Omit this parameter the first time you call this endpoint.
	if paginationToken != "" {
		path += "?paginationToken=" + paginationToken
	}

	res, bs, err := c.doRequestPost(ctx, path, bm)
	if err != nil {
		return nil, err
	}
	rsp = &NotificationHistoryRsp{}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode == http.StatusOK {
		return rsp, nil
	}
	if err = statusCodeErrCheck(rsp.StatusCodeErr); err != nil {
		return rsp, err
	}
	return rsp, nil
}
