package alipay

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"testing"
)

func TestAntMerchantShopModify(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)

	aliRsp, err := client.AntMerchantShopModify(bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopModify(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantShopCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.SetBodyMap("business_address", func(bm gopay.BodyMap) {
		bm.Set("city_code", "37100")
		bm.Set("district_code", "371002")
		bm.Set("address", "万塘路18号黄龙时代广场B座")
		bm.Set("province_code", "310000")
	})
	bm.Set("shop_category", "B0001")
	bm.Set("store_id", "NO0001")
	bm.Set("shop_type", "01")
	bm.Set("ip_role_id", "2088301155943087")
	bm.Set("shop_name", "肯德基中关村店")

	aliRsp, err := client.AntMerchantShopCreate(bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopCreate(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantShopConsult(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.SetBodyMap("business_address", func(bm gopay.BodyMap) {
		bm.Set("city_code", "37100")
		bm.Set("district_code", "371002")
		bm.Set("address", "万塘路18号黄龙时代广场B座")
		bm.Set("province_code", "310000")
	})
	bm.Set("shop_category", "B0001")
	bm.Set("store_id", "NO0001")
	bm.Set("shop_type", "01")
	bm.Set("ip_role_id", "2088301155943087")
	bm.Set("shop_name", "肯德基中关村店")

	aliRsp, err := client.AntMerchantShopConsult(bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopConsult(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
