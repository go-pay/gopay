package qq

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	mchId  = "1368139502"
	apiKey = "GFDS8j98rewnmgl45wHTt980jg543abc"
)

func TestMain(m *testing.M) {

	// 初始化QQ客户端
	//    mchId：商户ID
	//    apiKey：API秘钥值
	client = NewClient(mchId, apiKey)

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOn

	//err := client.AddCertFilePath(nil, nil, nil)
	//if err != nil {
	//	panic(err)
	//}
	os.Exit(m.Run())
}

func TestClient_MicroPay(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32))

	qqRsp, err := client.MicroPay(ctx, bm)
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

func TestClient_DownloadRedListFile(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("date", 20160803)
	file, err := client.DownloadRedListFile(ctx, bm)
	if err != nil {
		xlog.Errorf("client.DownloadRedListFile(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("qqRsp:", file)
}
