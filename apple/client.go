package apple

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// UrlProd is the URL when testing your app in the sandbox and while your application is in review
	UrlSandbox = "https://sandbox.itunes.apple.com/verifyReceipt"

	// UrlSandbox is the URL when your app is live in the App Store
	UrlProd = "https://buy.itunes.apple.com/verifyReceipt"
)

// Client 苹果应用内支付Client
type Client struct {
	// url 苹果的校验URL：UrlProd，UrlSandbox
	url string

	// pwd 密码，苹果APP秘钥，https://help.apple.com/app-store-connect/#/devf341c0f01
	pwd string
}

// NewClient 根据校验url和app秘钥来创建一个客户端
// 	url:取值为 UrlProd，UrlSandbox
// 	pwd: 苹果APP秘钥，https://help.apple.com/app-store-connect/#/devf341c0f01
// 	https://help.apple.com/app-store-connect/#/devf341c0f01
func NewClient(url, pwd string) Client {
	return Client{
		url: url,
		pwd: pwd,
	}
}

// Verify 请求APP Store 校验支付请求,实际测试时发现这个文档介绍的返回信息只有那个status==0表示成功可以用，其他的返回信息跟文档对不上
// 	文档：https://developer.apple.com/documentation/appstorereceipts/verifyreceipt
func (c Client) Verify(ctx context.Context, receipt string) (VerifyResponse, error) {
	var vr VerifyResponse

	data, err := json.Marshal(VerifyRequest{Receipt: receipt, Password: c.pwd})
	if err != nil {
		return vr, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewReader(data))
	if err != nil {
		return vr, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return vr, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return vr, fmt.Errorf("response status failed,status:%d body:%s", res.StatusCode, body)
	}

	// show body content
	// body, _ := io.ReadAll(res.Body)
	// fmt.Printf("verify body:%s", string(body))

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&vr)
	if err != nil {
		return vr, err
	}

	return vr, nil
}
