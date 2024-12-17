package apple

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error

	iss = "57246542-96fe-1a63-e053-0824d011072a"
	bid = "com.example.testbundleid2021"
	kid = "2X9R4HXF34"
)

func TestMain(m *testing.M) {
	xlog.SetLevel(xlog.DebugLevel)
	// 初始化客户端
	// iss：issuer ID
	// bid：bundle ID
	// kid：private key ID
	// privateKey：私钥文件读取后的字符串内容
	// isProd：是否是正式环境
	client, err = NewClient(iss, bid, kid, "privateKey", false)
	if err != nil {
		xlog.Error(err)
		return
	}

	os.Exit(m.Run())
}
