package alipay

import (
	"io"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestMerchantItemFileUpload(t *testing.T) {
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
	bm.Set("scene", "SYNC_ORDER") // 素材固定值
	bm.SetFormFile("file_content", f)
	aliRsp, err := client.MerchantItemFileUpload(ctx, bm)
	if err != nil {
		xlog.Errorf("client.MerchantItemFileUpload(),error:%+v", err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
