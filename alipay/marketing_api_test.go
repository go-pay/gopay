package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestOpenAppQrcodeCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("url_param", "page/component/component-pages/view/view").
		Set("query_param", "x=1").
		Set("describe", "二维码描述")

	// 发起请求
	aliRsp, err := client.OpenAppQrcodeCreate(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}
