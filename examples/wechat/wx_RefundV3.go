package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/go-pay/xlog"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/util"
)

/*
Create 退款申请

# 应用场景
当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，微信支付将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。

注意：
1、交易时间超过一年的订单无法提交退款
2、微信支付退款支持单笔交易分多次退款，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。申请退款总金额不能超过订单金额。 一笔退款失败后重新提交，请不要更换退款单号，请使用原商户退款单号
3、请求频率限制：150qps，即每秒钟正常的申请退款请求次数不超过150次
    错误或无效请求频率限制：6qps，即每秒钟异常或错误的退款申请请求不超过6次
4、每个支付订单的部分退款次数不能超过50次
5、如果同一个用户有多笔退款，建议分不同批次进行退款，避免并发退款导致退款失败
6、申请退款接口的返回仅代表业务的受理情况，具体退款是否成功，需要通过退款查询接口获取结果


# 错误码
|名称|描述|原因|解决方案|
|-|-|-|-|
|SYSTEM_ERROR|接口返回错误|系统超时等|请不要更换商户退款单号，请使用相同参数再次调用API。|
|USER_ACCOUNT_ABNORMAL|退款请求失败|用户帐号注销|此状态代表退款申请失败，商户可自行处理退款。|
|NOT_ENOUGH|余额不足|商户可用退款余额不足|此状态代表退款申请失败，商户可根据具体的错误提示做相应的处理。|
|PARAM_ERROR|参数错误|请求参数未按指引进行填写|请求参数错误，请重新检查再调用退款申请|
|MCH_NOT_EXISTS|MCHID不存在|参数中缺少MCHID|请检查MCHID是否正确|
|RESOURCE_NOT_EXISTS|订单号不存在|缺少有效的订单号|请检查你的订单号是否正确且是否已支付，未支付的订单不能发起退款|
|SIGN_ERROR|签名错误|参数签名结果不正确|请检查签名参数和方法是否都符合签名算法要求|
|FREQUENCY_LIMITED|频率限制|2个月之前的订单申请退款有频率限制|该笔退款未受理，请降低频率后重试|
|INVALID_REQUEST|请求参数符合参数格式，但不符合业务规则|不符合业务规则|此状态代表退款申请失败，商户可根据具体的错误提示做相应的处理。|
|NO_AUTH|没有退款权限|没有此单的退款权限|此状态代表退款申请失败，请检查是否有退这笔订单的权限|
*/

type res struct {
	Code    string
	Message string
}

// RefundV3  退款V3版本
func RefundV3() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// NewClientV3 初始化微信客户端 V3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户API证书的证书序列号
	// apiV3Key：APIv3Key，商户平台获取
	// privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	client, err := wechat.NewClientV3("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", "")
	if err != nil {
		return
	}
	//生成订单流水号
	s := util.RandomString(64)
	orderNo := fmt.Sprintf("CX-%s", s)

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	// 选填 商户订单号（支付后返回的，一般是以42000开头）
	bm.Set("transaction_id", "4200001650202212082601288339").
		Set("sign_type", "MD5").
		// 必填 退款订单号（程序员定义的）
		Set("out_refund_no", orderNo).
		// 选填 退款描述
		Set("reason", "这是一退款操作").
		Set("notify_url", "https://www.fmm.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			// 退款金额:单位是分
			bm.Set("refund", "1"). //实际退款金额
				Set("total", "1"). // 折扣前总金额（不是实际退款数）
				Set("currency", "CNY")
		})
	//请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	refund, err := client.V3Refund(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 将非正常退款异常记录
	// 返回：404 > {"code":"RESOURCE_NOT_EXISTS","message":"订单不存在"}
	if refund.Code == 404 || refund.Code == 400 || refund.Code == 403 {
		// 这里时对非正常退款的一些处理message，我们将code统一使用自定义的，然后把message抛出去
		var temp res
		err = json.Unmarshal([]byte(refund.Error), &temp)
		if err != nil {
			xlog.Error("json序列化失败")
		}
		xlog.Infof(fmt.Sprintf("code:%d,message:%s", 50000, temp.Message))
		return
	}

	// 这里可以进行业务处理
	// 比如插入支付表，回写订单表等
	return
}
