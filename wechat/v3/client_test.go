package wechat

import (
	"os"
	"testing"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xlog"
)

var (
	client    *ClientV3
	err       error
	Appid     = ""
	MchId     = ""
	ApiV3Key  = ""
	SerialNo  = ""
	PKContent = ``
)

func TestMain(m *testing.M) {
	// NewClientV3 初始化微信客户端 V3
	//	appid：appid
	//	mchid：商户ID
	// 	serialNo 商户证书的证书序列号
	//	pkContent：私钥 apiclient_key.pem 读取后的内容
	client, err = NewClientV3(Appid, MchId, SerialNo, []byte(PKContent))
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	os.Exit(m.Run())
}

func TestGetPlatformCerts(t *testing.T) {
	certs, err := client.GetPlatformCerts(ApiV3Key)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("certs.StatusCode", certs.StatusCode)
	xlog.Debug("certs.Headers", certs.Headers)

	for _, v := range certs.Certs {
		xlog.Debug("cert:", v)
	}
}

func TestClientV3_DecryptCerts(t *testing.T) {
	ciphertext := "XboO6jC96vrZVQRZRCGPb5i93fPyrKTY1bf9jVi1eRVRLmG97LlZYPxSTJJWMREi/VwoiDZroT+gUt03k1yCXMN3x8YnPXAjg7KDrA6E/vXFrX9ZqmVtd/fhNFYcydVdw+EPkl6HVPH8cIWQrG9xVEL535MybANzO15FhGWeHDGbr1r7fYsldvxQwLAlAVcTfJ8M04UmnZtQu9rfTbAa7kq73g2mK9LIV94TE2F8+//ZWCIaYHmwBi5J+Ftg5LpTJ6UmID+c1eFPn6YqrX0cW/EtHulIvolZKHyqApehoH/Y+gprKTltp2CejX3PJWYIIRl5yzSCxBcxFzdMBTUEvOJH52OrkjnsSiAldu9zotAuRkLA0RvqW2gp1V/Gid0SNNOP9f5QsnnvrINS36AagV+YdxJMh3LRuF/tmTqic2+cxQLRDMlhbCroK3hwDl5b4Vt563NKkuVYmwPgp83Y2oeqbS77DCTONsQkvCQ1YwxKLqn1apkD/8bomeCnfNDh/ZJmvEaGjnUxSmAR6EiHMSjxVgksWvvRnTEclTN2/S6bZ0gJH3CzC8OB0VI41A3u9VJQcb21MyWpDFNxadLgo5GyYr16dR/Db7+mp+/Tp2Ev4p3bjdhqUxOpPm6YAmEhwsk4xeBbDWKz04vfQKYvkobIUsqnddnco1H+NgoYmbmRMdQ/mRnSXnWOzhigGpriyCrgBJ5NGhhHr7NMgs4DpIqF5OefJGRn+hq0gbKqNCGnZydM5RhcWVUTtF/K/iC+BffX3Fo+FJNrmSbSDjgN57LfFSaqNNpdUmPS55/17AZEQKvhs6nN5j65eU5iu7opMslYv+/26zlbFoLXjduHd9kRh60EO6gH7nXfa6Zo8LckwAgaO6nmiaw0vG1HStjCAmePC6LDY1bGn/DLCvte834d8NVJ7uSFz4U80OoVm9kAb9sWpPXXqpDCQKYDQCoAlHrPT788KCVTUdup9M1H7U1d5WOqTyZIyjsnI4JO7uddSsOS3o+WZYzFfnOZjO1sr2++ipwbGFHtxrUTMWrBsjP+XgDwoGYIsANmbaRK4KFaRVX0cjofQ0Qforz2Wtrr6ydIVA80/bwWBgxSN4OTPyf1nCXMoVS6klvvhhqZjb0lWM2mymHtV8kcrQTyTyirD7Ufe3gCPXlx5s4BravzEcklfO4maQYpBcfaDgfFolWUnuD6Fph/hLkgN9tH3VpfWd3wtLbps7zXPFd+DA1SARRsrf1imTJAaJNnX4nWutneqq6+Y3EdW3nJogbZfC6AZtvRnGa9gJCKnP79S57ql0TCtbRC7cwReWukeYw8PCxhRAYQstbN0bH7fq7K5PyEW9CSBeU88pgaLAlbt+IWtjiC2JxVg22LxlyugSzHXqQGmDWwp06Q9GimNXdfOzx+IQM9W8XuH5G5CTVeicgw1zQCBsHUfbtj1G+TtpfZgsbvaABXzOixK4eicEDdlH9xFQLgPT34wv04Ab0QxeHjkbdhf/UOUEElkF56LKE989YGEpTEBwxsdqKU0ADrty2nbrPKbRInt/RJfp4NNsCXgOiEEEZfnVTuAOGvepJMm4o1lw/LV2zq0ptJN9zisYLA1s39LN/PcaB+sC2VJ4yC0TMIwe1Gpw/nqJvRZAb4QaXrZAmKNi8Ro2x1sKou+TpFsXa2jV54d/xEwPAERpbxaUz95UOu0sdRaPdRTWV6vA8aV/Q6rU8rzrh2XonAMHUbfBVo+gdlKOdPqpxWgr4ETzcVV5HJ4JXhxzwWtfBCFxvws0UCUbjV25RORkY5Q2MpWSH3YenCjPYBjMKC4FzR2AljmfDZQz841EoeOGKl2fYGGxLiUvwc2wFCKjInLQFSJBRm3dhaJ3xX/M3cPnM8gSFMbqzHUg=="
	add := "certificate"
	nonce := "d393d2a609f1"
	client.DecryptCerts(ciphertext, nonce, add, ApiV3Key)
}

func TestV3Jsapi(t *testing.T) {
	tradeNo := gotil.GetRandomString(32)
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("description", "测试Jsapi支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", "https://www.gopay.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", "asdas")
		})

	wxRsp, err := client.V3TransactionJsapi(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp:", wxRsp)
}

func TestV3Native(t *testing.T) {
	tradeNo := gotil.GetRandomString(32)
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("description", "测试Native支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", "https://api2.fangyiyun.com/api/v1/wechat/callback").
		//Set("notify_url", "https://www.gopay.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3TransactionNative(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp.StatusCode:", wxRsp.StatusCode)
	xlog.Debug("wxRsp.Headers:", wxRsp.Headers)
	xlog.Debug("wxRsp.Response:", wxRsp.Response)
}

func TestV3QueryOrder(t *testing.T) {
	//wxRsp, err := client.V3TransactionQueryOrder(TransactionId, "42000008462020122402449153433")
	wxRsp, err := client.V3TransactionQueryOrder(OutTradeNo, "22LW55HDd8tuxgZgFM445kI52BZVk847")
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.StatusCode == 200 {
		xlog.Debugf("wxRsp:%#v", wxRsp.Response)
	}
	xlog.Debugf("wxRsp:%s", wxRsp.Error)
}
