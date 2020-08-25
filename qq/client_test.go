package qq

import (
	"os"
	"testing"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xlog"
)

var (
	client *Client
	mchId  = "1368139502"
	apiKey = "GFDS8j98rewnmgl45wHTt980jg543abc"
)

func TestMain(m *testing.M) {

	// 初始化QQ客户端
	//    mchId：商户ID
	//    apiKey：API秘钥值
	client = NewClient(mchId, apiKey)

	//err := client.AddCertFilePath(nil, nil, nil)
	//if err != nil {
	//	panic(err)
	//}
	os.Exit(m.Run())
}

func TestClient_MicroPay(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))

	qqRsp, err := client.MicroPay(bm)
	if err != nil {
		xlog.Errorf("client.Micropay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("qqRsp:", *qqRsp)
}

func TestNotifyResponse_ToXmlString(t *testing.T) {
	n := new(NotifyResponse)
	n.ReturnCode = "SUCCESS"
	xlog.Info(n.ToXmlString())

	n.ReturnCode = "FAIL"
	n.ReturnMsg = "abc"
	xlog.Info(n.ToXmlString())

}
