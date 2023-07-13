package paypal

import (
	"context"
	"encoding/base64"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

var (
	client   *Client
	ctx      = context.Background()
	err      error
	Clientid = "AQH4MMv6pFY9HChKHZYmphzJrtbMM-9Vh1ETzoe5W-koYBuk-xSnmjY_dMBCwmmBOczNHtmenCRSgbU4"
	Secret   = "EMV4AUd59OWVQWpAcVvUgtDvwIn50pcAZmdNE-5zR6iBPHijqGmbELcY9ruvK7_FGIlHcGQNR3_rICbZ"
)

func TestMain(m *testing.M) {
	client, err = NewClient(Clientid, Secret, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	xlog.Debugf("Appid: %s", client.Appid)
	xlog.Debugf("AccessToken: %s", client.AccessToken)
	xlog.Debugf("ExpiresIn: %d", client.ExpiresIn)
	os.Exit(m.Run())
}

func TestBasicAuth(t *testing.T) {
	uname := "jerry"
	passwd := "12346"
	auth := base64.StdEncoding.EncodeToString([]byte(uname + ":" + passwd))
	xlog.Debugf("Basic %s", auth)
}
