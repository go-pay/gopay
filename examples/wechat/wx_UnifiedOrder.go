package wechat

import (
	"fmt"
	"github.com/iGoogle-ink/gopay"
)

func main() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	number := gopay.GetRandomString(32)
	fmt.Println("out_trade_no:", number)
	//初始化参数Map
	body := make(gopay.BodyMap)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("body", "H5测试支付")
	body.Set("out_trade_no", number)
	body.Set("total_fee", 1)
	body.Set("spbill_create_ip", "127.0.0.1")
	body.Set("notify_url", "http://www.gopay.ink")
	body.Set("trade_type", gopay.TradeType_H5)
	body.Set("device_info", "WEB")
	body.Set("sign_type", gopay.SignType_MD5)

	sceneInfo := make(map[string]map[string]string)
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "http://www.gopay.ink"
	h5Info["wap_name"] = "H5测试支付"
	sceneInfo["h5_info"] = h5Info
	body.Set("scene_info", sceneInfo)

	//body.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp:", *wxRsp)

	//timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	////获取小程序支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := gopay.GetMiniPaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, gopay.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	//获取H5支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := gopay.GetH5PaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, gopay.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	//获取小程序需要的paySign
	//paySign := gopay.GetAppPaySign("wxdaa2ab9ef87b5497","", wxRsp.NonceStr, wxRsp.PrepayId, gopay.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)
}
