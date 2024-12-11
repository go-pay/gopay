package bytedance

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error
	Token  = ""
	MchId  = ""
	Salt   = ""
)

func TestMain(m *testing.M) {
	// NewClient 初始化字节客户端
	// 	token：Token 令牌
	//	mchid：商户号
	//	salt：SALT
	client, err = NewClient(Token, MchId, Salt)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	os.Exit(m.Run())
}
