package alipay

import (
	"io"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestAntMerchantShopModify(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)

	aliRsp, err := client.AntMerchantShopModify(ctx, bm)
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

	aliRsp, err := client.AntMerchantShopCreate(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
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

	aliRsp, err := client.AntMerchantShopConsult(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopConsult(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantOrderQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("order_id", "2017112200502000000004754299")

	aliRsp, err := client.AntMerchantOrderQuery(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantOrderQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantShopQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("shop_id", "2018011900502000000005124744")
	bm.Set("store_id", "NO0001")
	bm.Set("ip_role_id", "2088301155943087")

	aliRsp, err := client.AntMerchantShopQuery(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantShopClose(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("shop_id", "2018011900502000000005124744")
	bm.Set("store_id", "NO0001")
	bm.Set("ip_role_id", "2088301155943087")

	aliRsp, err := client.AntMerchantShopClose(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopClose(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantExpandIndirectImageUpload(t *testing.T) {
	// 请求参数
	logo, err := os.Open("../logo.png")
	if err != nil {
		xlog.Errorf("os.Open(%s),error:%+v", "../logo.png", err)
		return
	}
	xlog.Warnf("fileName: %s", logo.Name())
	allBs, err := io.ReadAll(logo)
	if err != nil {
		xlog.Errorf("io.ReadAll(%s),error:%+v", logo.Name(), err)
		return
	}
	f := &gopay.File{
		Name:    "logo.png",
		Content: allBs,
	}
	bm := make(gopay.BodyMap)
	bm.Set("image_type", "png")
	bm.SetFormFile("image_content", f)
	aliRsp, err := client.AntMerchantExpandIndirectImageUpload(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantExpandIndirectImageUpload(),error:%+v", err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
