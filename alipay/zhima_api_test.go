package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

// 芝麻企业信用信用评估初始化测试
func TestZhimaCreditEpSceneRatingInitialize(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("credit_category", "ZMSCCO_5_1_1")
	bm.Set("product_code", "w1010100100000000001")
	bm.Set("out_order_no", "201805301527674106562F0000954216")
	bm.Set("user_id", "2088302248028263")

	aliRsp, err := client.ZhimaCreditEpSceneRatingInitialize(bm)
	if err != nil {
		xlog.Errorf("client.ZhimaCreditEpSceneRatingInitialize(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
