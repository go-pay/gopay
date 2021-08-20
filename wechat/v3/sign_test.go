package wechat

import (
	"testing"
	"time"

	"github.com/go-pay/gopay/pkg/util"

	"github.com/go-pay/gopay/pkg/xlog"
)

func TestPaySignOfJSAPIp(t *testing.T) {
	jsapi, err := client.PaySignOfJSAPI("appid", "prepayid", &SignInfo{
		HeaderTimestamp: util.Int642String(time.Now().Unix()),
		HeaderNonce:     util.GetRandomString(32),
	})
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("jsapi:%#v", jsapi)
}

func TestPaySignOfApp(t *testing.T) {
	app, err := client.PaySignOfApp("appid", "prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("app:%#v", app)
}

func TestPaySignOfApplet(t *testing.T) {
	applet, err := client.PaySignOfApplet("appid", "prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("applet:%#v", applet)
}
