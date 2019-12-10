package wechat

import (
	"fmt"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
)

func QueryRefund() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	//初始化参数结构体
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
	//body.Set("out_refund_no", "vk4264I1UQ3Hm3E4AKsavK8npylGSgQA092f9ckUxp8A2gXmnsLEdsupURVTcaC7")
	//body.Set("transaction_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
	//body.Set("refund_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("sign_type", wechat.SignType_MD5)

	//请求申请退款
	wxRsp, err := client.QueryRefund(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp：", *wxRsp)

}
