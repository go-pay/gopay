package douyin

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/go-pay/xlog"
)

// mockNotifyRequest 构造一个模拟的回调 http.Request，用于本地演示 ParseNotify 流程
// 真实环境中，直接把回调 http.Handler 里的 *http.Request 传给 ParseNotify 即可
func mockNotifyRequest(body string, timestamp, nonce, signature, serial string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/notify", io.NopCloser(bytes.NewBufferString(body)))
	req.Header.Set(HeaderTimestamp, timestamp)
	req.Header.Set(HeaderNonce, nonce)
	req.Header.Set(HeaderSignature, signature)
	req.Header.Set(HeaderSerial, serial)
	return req
}

// TestParseNotify 演示回调解析 + 验签 + 按 event_type 分发解密
func TestParseNotify(t *testing.T) {
	// 以下参数请从实际接收到的回调 http.Request 中获取
	body := ""
	timestamp := ""
	nonce := ""
	signature := ""
	serial := ""

	req := mockNotifyRequest(body, timestamp, nonce, signature, serial)

	notifyReq, err := ParseNotify(req)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("event_type: %s", notifyReq.EventType)

	// 1) 验签（按 header Douyinpay-Serial 匹配已注册的平台证书公钥）
	if err = notifyReq.VerifySignByPKMap(client.PlatformCertMap()); err != nil {
		xlog.Error("verify sign err:", err)
		return
	}

	// 2) 根据 event_type 解密对应载荷
	switch notifyReq.EventType {
	case EventTransactionSuccess:
		pay, err := notifyReq.DecryptPayCipherText(ApiKey)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Debugf("pay: %+v", pay)
	case EventRefundSuccess:
		r, err := notifyReq.DecryptRefundCipherText(ApiKey)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Debugf("refund: %+v", r)
	case EventAsyncSplitFinish:
		s, err := notifyReq.DecryptProfitResultCipherText(ApiKey)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Debugf("profit result: %+v", s)
	case EventSplitSuccess:
		d, err := notifyReq.DecryptProfitDynamicCipherText(ApiKey)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Debugf("profit dynamic: %+v", d)
	case EventTransferSuccess:
		tr, err := notifyReq.DecryptTransferCipherText(ApiKey)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Debugf("transfer: %+v", tr)
	}
}

// TestParseNotifyToBodyMap 演示回调解析到 gopay.BodyMap
func TestParseNotifyToBodyMap(t *testing.T) {
	req := mockNotifyRequest("", "", "", "", "")

	bm, err := ParseNotifyToBodyMap(req)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("bm: %+v", bm)
}
