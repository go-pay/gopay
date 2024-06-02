package apple

import (
	"context"

	"github.com/go-pay/gopay/pkg/xhttp"
)

const (
	// is the URL when testing your app in the sandbox and while your application is in review
	UrlSandbox = "https://sandbox.itunes.apple.com/verifyReceipt"
	// is the URL when your app is live in the App Store
	UrlProd = "https://buy.itunes.apple.com/verifyReceipt"
)

// VerifyReceipt 请求APP Store 校验支付请求,实际测试时发现这个文档介绍的返回信息只有那个status==0表示成功可以用，其他的返回信息跟文档对不上
// url：取 UrlProd 或 UrlSandbox
// pwd：苹果APP秘钥，https://help.apple.com/app-store-connect/#/devf341c0f01
// 文档：https://developer.apple.com/documentation/appstorereceipts/verifyreceipt
func VerifyReceipt(ctx context.Context, url, pwd, receipt string) (rsp *VerifyResponse, err error) {
	req := &VerifyRequest{Receipt: receipt, Password: pwd}
	rsp = new(VerifyResponse)
	_, err = xhttp.NewClient().Req(xhttp.TypeJSON).Post(url).SendStruct(req).EndStruct(ctx, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
