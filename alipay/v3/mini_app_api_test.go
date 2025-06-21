package alipay

import (
	"io"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
	"github.com/go-pay/xlog"
)

func TestOpenMiniVersionAuditApply(t *testing.T) {
	// 请求参数
	logo, err := os.Open("../../logo.png")
	if err != nil {
		xlog.Errorf("os.Open(%s),error:%+v", "../../logo.png", err)
		return
	}
	xlog.Warnf("fileName: %s", logo.Name())
	allBs, err := io.ReadAll(logo)
	if err != nil {
		xlog.Errorf("io.ReadAll(%s),error:%+v", logo.Name(), err)
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set(HeaderAppAuthToken, "202506BB364aecd0262f46afa8e472a46a9ffX29")
	bm.Set("app_version", "0.0.1")                                                         // 商家小程序版本号，格式: x.y.z
	bm.Set("version_desc", "该版本新增了XXX功能，修复了XXX的bug，优化了用户体验。")                              // 小程序版本描述
	bm.Set("app_slogan", "我是简介")                                                           // 小程序简介
	bm.Set("service_phone", "13110101010")                                                 // 小程序客服电话
	bm.Set("region_type", "CHINA")                                                         // 区域类型
	bm.Set("mini_category_ids", "XS1001_XS2001_XS3002;XS1011_XS2089;XS1002_XS2008_XS3024") // 新小程序前台类目
	bm.Set("app_name", "杭州支小宝潮流女装店")                                                       // 小程序名称
	bm.Set("app_desc", "一家专注做潮流女装的店铺，带给消费者价格便宜、质量上乘的衣服")                                   // 小程序描述
	bm.Set("license_no", "licenseNo")                                                      // 营业执照证件号
	bm.Set("license_name", "营业执照名称")                                                       // 营业执照名称
	bm.Set("license_valid_date", "2099-12-31")                                             // 营业执照有效期
	bm.Set("audit_rule", "BASE_PROMOTE")                                                   // 审核类型
	bm.Set("auto_online", "true")                                                          // 是否自动上架

	// 小程序logo图片
	logoFile := &gopay.File{
		Name:    "logo.png",
		Content: allBs,
	}
	bm.SetFormFile("app_logo", logoFile)
	// 门头照
	shopFrontFile := &gopay.File{
		Name:    "logo.png",
		Content: allBs,
	}
	bm.SetFormFile("out_door_pic", shopFrontFile)
	// 第一张营业执照
	licenseFile1 := &gopay.File{
		Name:    "logo.png",
		Content: allBs,
	}
	bm.SetFormFile("first_license_pic", licenseFile1)
	// 第一张特殊资质图片
	licenseFile11 := &gopay.File{
		Name:    "logo.png",
		Content: allBs,
	}
	bm.SetFormFile("first_special_license_pic", licenseFile11)

	aliRsp, err := client.OpenMiniVersionAuditApply(ctx, bm)
	if err != nil {
		xlog.Errorf("client.OpenMiniVersionAuditApply(),error:%+v", err)
		return
	}
	xlog.Debug("aliRsp:", js.MarshalString(aliRsp))
}
