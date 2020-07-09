package wechat

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xlog"
	"github.com/iGoogle-ink/gotil/xrsa"
)

var (
	client *Client
	appId  = "wxdaa2ab9ef87b5497"
	mchId  = "1368139502"
	apiKey = "GFDS8j98rewnmgl45wHTt980jg543wmg"
)

func TestMain(m *testing.M) {

	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client = NewClient(appId, mchId, apiKey, false)

	// 设置国家，不设置默认就是 China
	client.SetCountry(China)

	//err := client.AddCertFilePath(nil, nil, nil)
	//if err != nil {
	//	panic(err)
	//}

	os.Exit(m.Run())
}

func TestClient_UnifiedOrder(t *testing.T) {
	number := gotil.GetRandomString(32)
	xlog.Info("out_trade_no:", number)
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "H5支付")
	bm.Set("out_trade_no", number)
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "http://www.gopay.ink")
	bm.Set("trade_type", TradeType_H5)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", SignType_MD5)

	sceneInfo := make(map[string]map[string]string)
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "http://www.gopay.ink"
	h5Info["wap_name"] = "H5测试支付"
	sceneInfo["h5_info"] = h5Info
	bm.Set("scene_info", sceneInfo)

	//bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	// 正式
	//sign := gopay.GetWeChatParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	// 沙箱
	//sign, _ := gopay.GetWeChatSanBoxParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	//body.Set("sign", sign)

	// 请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(bm)
	if err != nil {
		xlog.Errorf("client.UnifiedOrder(%+v),error:%+v", bm, err)
		return
	}
	xlog.Info("wxRsp:", *wxRsp)
	//xlog.Info("wxRsp.MwebUrl:", wxRsp.MwebUrl)

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	// 获取小程序支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := GetMiniPaySign(appId, wxRsp.NonceStr, pac, SignType_MD5, timeStamp, apiKey)
	//xlog.Info("paySign:", paySign)

	// 获取H5支付需要的paySign
	pac := "prepay_id=" + wxRsp.PrepayId
	paySign := GetH5PaySign(appId, wxRsp.NonceStr, pac, SignType_MD5, timeStamp, apiKey)
	xlog.Debug("paySign:", paySign)

	// 获取小程序需要的paySign
	//paySign := GetAppPaySign(appId,"partnerid", wxRsp.NonceStr, wxRsp.PrepayId, SignType_MD5, timeStamp, apiKey)
	//xlog.Info("paySign:", paySign)
}

func TestClient_QueryOrder(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", SignType_MD5)

	// 请求订单查询，成功后得到结果
	wxRsp, resBm, err := client.QueryOrder(bm)
	if err != nil {
		xlog.Errorf("client.QueryOrder(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
	xlog.Debug("resBm：", resBm)
}

func TestClient_CloseOrder(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", SignType_MD5)

	// 请求关闭订单，成功后得到结果
	wxRsp, err := client.CloseOrder(bm)
	if err != nil {
		xlog.Errorf("client.CloseOrder(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
}

func TestClient_Micropay(t *testing.T) {
	number := gotil.GetRandomString(32)
	xlog.Info("out_trade_no:", number)
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "扫用户付款码支付")
	bm.Set("out_trade_no", number)
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("auth_code", "134622817080551492")
	bm.Set("sign_type", SignType_MD5)

	// 请求支付，成功后得到结果
	wxRsp, err := client.Micropay(bm)
	if err != nil {
		xlog.Errorf("client.Micropay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("Response:", *wxRsp)
	ok, err := VerifySign(apiKey, SignType_MD5, wxRsp)
	if err != nil {
		xlog.Error(err)
	}
	xlog.Debug("同步验签结果：", ok) // 沙箱环境验签失败请用正式环境测
}

func TestClient_AuthCodeToOpenId(t *testing.T) {
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("auth_code", "134753997737645794")

	wxRsp, err := client.AuthCodeToOpenId(bm)
	if err != nil {
		xlog.Errorf("client.AuthCodeToOpenId(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("Response:", *wxRsp)
}

func TestClient_Refund(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "QRcTBTbJLoDrWSW9FtpSFlgWhft2QbaY")
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", SignType_MD5)
	s := gotil.GetRandomString(64)
	xlog.Info("out_refund_no:", s)
	bm.Set("out_refund_no", s)
	bm.Set("total_fee", 101)
	bm.Set("refund_fee", 101)
	bm.Set("notify_url", "https://www.gopay.ink")

	// 请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	//    certFilePath：cert证书路径
	//    keyFilePath：Key证书路径
	//    pkcs12FilePath：p12证书路径
	wxRsp, resBm, err := client.Refund(bm, nil, nil, nil)
	if err != nil {
		xlog.Errorf("client.Refund(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
	xlog.Debug("resBm：", resBm)
}

func TestClient_QueryRefund(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
	//bm.Set("out_refund_no", "vk4264I1UQ3Hm3E4AKsavK8npylGSgQA092f9ckUxp8A2gXmnsLEdsupURVTcaC7")
	//bm.Set("transaction_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
	//bm.Set("refund_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", SignType_MD5)

	// 请求申请退款
	wxRsp, resBm, err := client.QueryRefund(bm)
	if err != nil {
		xlog.Errorf("client.QueryRefund(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
	xlog.Debug("resBm：", resBm)
}

func TestClient_Reverse(t *testing.T) {
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("out_trade_no", "6aDCor1nUcAihrV5JBlI09tLvXbUp02B")
	bm.Set("sign_type", SignType_MD5)

	// 请求撤销订单，成功后得到结果，沙箱环境下，证书路径参数可传nil
	wxRsp, err := client.Reverse(bm, nil, nil, nil)
	if err != nil {
		xlog.Errorf("client.Reverse(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("Response:", wxRsp)
}

func TestClient_Transfer(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("partner_trade_no", gotil.GetRandomString(32))
	bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")
	bm.Set("check_name", "FORCE_CHECK") // NO_CHECK：不校验真实姓名 , FORCE_CHECK：强校验真实姓名
	bm.Set("re_user_name", "付明明")       // 收款用户真实姓名。 如果check_name设置为FORCE_CHECK，则必填用户真实姓名
	bm.Set("amount", 100)               // 企业付款金额，单位为分
	bm.Set("desc", "点点点")               // 企业付款备注，必填。注意：备注中的敏感词会被转成字符*
	bm.Set("spbill_create_ip", "127.0.0.1")

	// 企业向微信用户个人付款（不支持沙箱环境）
	//    body：参数Body
	//    certFilePath：cert证书路径
	//    keyFilePath：Key证书路径
	//    pkcs12FilePath：p12证书路径
	wxRsp, err := client.Transfer(bm, "cert/apiclient_cert.pem", "cert/apiclient_key.pem", "cert/apiclient_cert.p12")
	if err != nil {
		xlog.Errorf("client.Transfer(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
}

func TestClient_GetTransferInfo(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("partner_trade_no", gotil.GetRandomString(32))

	// 查询企业付款
	//    body：参数Body
	//    certFilePath：cert证书路径
	//    keyFilePath：Key证书路径
	//    pkcs12FilePath：p12证书路径
	wxRsp, err := client.GetTransferInfo(bm, "cert/apiclient_cert.pem", "cert/apiclient_key.pem", "cert/apiclient_cert.p12")
	if err != nil {
		xlog.Errorf("client.GetTransferInfo(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
}

func TestClient_DownloadBill(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", SignType_MD5)
	bm.Set("bill_date", "20190722")
	bm.Set("bill_type", "ALL")

	// 请求下载对账单，成功后得到结果（string类型字符串）
	wxRsp, err := client.DownloadBill(bm)
	if err != nil {
		xlog.Errorf("client.DownloadBill(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_DownloadFundFlow(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bm.Set("bill_date", "20190122")
	bm.Set("account_type", "Basic")

	// 请求下载资金账单，成功后得到结果，沙箱环境下，证书路径参数可传nil
	wxRsp, err := client.DownloadFundFlow(bm, "cert/apiclient_cert.pem", "cert/apiclient_key.pem", "cert/apiclient_cert.p12")
	if err != nil {
		xlog.Errorf("client.DownloadFundFlow(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_BatchQueryComment(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", SignType_HMAC_SHA256)
	bm.Set("begin_time", "20200520000000")
	bm.Set("end_time", "20200622174000")
	bm.Set("offset", "0")

	// 请求拉取订单评价数据，成功后得到结果，沙箱环境下，证书路径参数可传nil
	wxRsp, err := client.BatchQueryComment(bm, "cert/apiclient_cert.pem", "cert/apiclient_key.pem", "cert/apiclient_cert.p12")
	if err != nil {
		xlog.Errorf("client.BatchQueryComment(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_EntrustPublic(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("plan_id", "12535")
	bm.Set("contract_code", "100000")
	bm.Set("request_serial", "1000")
	bm.Set("contract_display_account", "微信代扣")
	bm.Set("notify_url", "https://www.igoogle.ink")
	bm.Set("version", "1.0")
	bm.Set("timestamp", time.Now().Unix())

	// 公众号纯签约
	wxRsp, err := client.EntrustPublic(bm)
	if err != nil {
		xlog.Errorf("client.EntrustPublic(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_EntrustAppPre(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("plan_id", "12535")
	bm.Set("contract_code", "100000")
	bm.Set("request_serial", "1000")
	bm.Set("contract_display_account", "微信代扣")
	bm.Set("notify_url", "https://www.igoogle.ink")
	bm.Set("version", "1.0")
	bm.Set("timestamp", time.Now().Unix())

	// APP纯签约
	wxRsp, err := client.EntrustAppPre(bm)
	if err != nil {
		xlog.Errorf("client.EntrustAppPre(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_EntrustH5(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("plan_id", "12535")
	bm.Set("contract_code", "100000")
	bm.Set("request_serial", "1000")
	bm.Set("contract_display_account", "微信代扣")
	bm.Set("notify_url", "https://www.igoogle.ink")
	bm.Set("version", "1.0")
	bm.Set("timestamp", time.Now().Unix())
	bm.Set("clientip", "127.0.0.1")

	// H5纯签约
	wxRsp, err := client.EntrustH5(bm)
	if err != nil {
		xlog.Errorf("client.EntrustH5(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_EntrustPaying(t *testing.T) {
	number := gotil.GetRandomString(32)
	xlog.Info("out_trade_no:", number)
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("contract_mchid", mchId)
	bm.Set("contract_appid", appId)
	bm.Set("out_trade_no", number)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "测试签约")
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("trade_type", TradeType_App)
	bm.Set("plan_id", "12535")
	bm.Set("contract_code", "100000")
	bm.Set("request_serial", "1000")
	bm.Set("contract_display_account", "微信代扣")
	bm.Set("notify_url", "https://www.igoogle.ink")
	bm.Set("contract_notify_url", "https://www.igoogle.ink")

	//bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	// 支付中签约
	wxRsp, err := client.EntrustPaying(bm)
	if err != nil {
		xlog.Errorf("client.EntrustPaying(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_PayBank(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("partner_trade_no", mchId)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("bank_code", "1001") // 招商银行，https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4&index=5
	bm.Set("amount", 1)

	encryptBank, err := xrsa.RsaEncryptData("621400000000567", "publicKey.pem")
	if err != nil {
		xlog.Error(err)
		return
	}
	encryptName, err := xrsa.RsaEncryptData("Jerry", "publicKey.pem")
	if err != nil {
		xlog.Error(err)
		return
	}
	bm.Set("enc_bank_no", encryptBank)
	bm.Set("enc_true_name", encryptName)

	// 企业付款到银行卡API
	wxRsp, err := client.PayBank(bm, "certFilePath", "keyFilePath", "pkcs12FilePath")
	if err != nil {
		xlog.Errorf("client.EntrustPaying(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

// =======================

func TestDecryptOpenDataToStruct(t *testing.T) {
	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="
	phone := new(UserPhone)
	//解密开放数据
	//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	//    iv:加密算法的初始向量
	//    sessionKey:会话密钥
	//    beanPtr:需要解析到的结构体指针
	err := DecryptOpenDataToStruct(data, iv, session, phone)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("PhoneNumber:", phone.PhoneNumber)
	xlog.Debug("PurePhoneNumber:", phone.PurePhoneNumber)
	xlog.Debug("CountryCode:", phone.CountryCode)
	xlog.Debug("Watermark:", phone.Watermark)

	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	iv2 := "r7BXXKkLb8qrSNn05n0qiA=="

	//微信小程序 用户信息
	userInfo := new(AppletUserInfo)

	err = DecryptOpenDataToStruct(encryptedData, iv2, sessionKey, userInfo)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("NickName:", userInfo.NickName)
	xlog.Debug("AvatarUrl:", userInfo.AvatarUrl)
	xlog.Debug("Country:", userInfo.Country)
	xlog.Debug("Province:", userInfo.Province)
	xlog.Debug("City:", userInfo.City)
	xlog.Debug("Gender:", userInfo.Gender)
	xlog.Debug("OpenId:", userInfo.OpenId)
	xlog.Debug("UnionId:", userInfo.UnionId)
	xlog.Debug("Watermark:", userInfo.Watermark)
}

func TestDecryptOpenDataToBodyMap(t *testing.T) {
	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="

	//解密开放数据
	//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	//    iv:加密算法的初始向量
	//    sessionKey:会话密钥
	bm, err := DecryptOpenDataToBodyMap(data, iv, session)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("WeChatUserPhone:", bm)
}

func TestDecryptRefundNotifyReqInfo(t *testing.T) {
	key := "ziR0QKsTUfMOuochC9RfCdmfHECorQAP"
	data := "YYwp8C48th0wnQzTqeI+41pflB26v+smFj9z6h9RPBgxTyZyxc+4YNEz7QEgZNWj/6rIb2MfyWMZmCc41CfjKSssoSZPXxOhUayb6KvNSZ1p6frOX1PDWzhyruXK7ouNND+gDsG4yZ0XXzsL4/pYNwLLba/71QrnkJ/BHcByk4EXnglju5DLup9pJQSnTxjomI9Rxu57m9jg5lLQFxMWXyeASZJNvof0ulnHlWJswS4OxKOkmW7VEyKyLGV6npoOm03Qsx2wkRxLsSa9gPpg4hdaReeUqh1FMbm7aWjyrVYT/MEZWg98p4GomEIYvz34XfDncTezX4bf/ZiSLXt79aE1/YTZrYfymXeCrGjlbe0rg/T2ezJHAC870u2vsVbY1/KcE2A443N+DEnAziXlBQ1AeWq3Rqk/O6/TMM0lomzgctAOiAMg+bh5+Gu1ubA9O3E+vehULydD5qx2o6i3+qA9ORbH415NyRrQdeFq5vmCiRikp5xYptWiGZA0tkoaLKMPQ4ndE5gWHqiBbGPfULZWokI+QjjhhBmwgbd6J0VqpRorwOuzC/BHdkP72DCdNcm7IDUpggnzBIy0+seWIkcHEryKjge3YDHpJeQCqrAH0CgxXHDt1xtbQbST1VqFyuhPhUjDXMXrknrGPN/oE1t0rLRq+78cI+k8xe5E6seeUXQsEe8r3358mpcDYSmXWSXVZxK6er9EF98APqHwcndyEJD2YyCh/mMVhERuX+7kjlRXSiNUWa/Cv/XAKFQuvUYA5ea2eYWtPRHa4DpyuF1SNsaqVKfgqKXZrJHfAgslVpSVqUpX4zkKszHF4kwMZO3M7J1P94Mxa7Tm9mTOJePOoHPXeEB+m9rX6pSfoi3mJDQ5inJ+Vc4gOkg/Wd/lqiy6TTyP/dHDN6/v+AuJx5AXBo/2NDD3fWhHjkqEKIuARr2ClZt9ZRQO4HkXdZo7CN06sGCHk48Tg8PmxnxKcMZm7Aoquv5yMIM2gWSWIRJhwJ8cUpafIHc+GesDlbF6Zbt+/KXkafJAQq2RklEN+WvZ/zFz113EPgWPjp16TwBoziq96MMekvWKY/vdhjol8VFtGH9F61Oy1Xwf6DJtPw=="
	refundNotify, err := DecryptRefundNotifyReqInfo(data, key)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("refundNotify:", *refundNotify)
}

func TestGetAppletAccessToken(t *testing.T) {
	token, err := GetAppletAccessToken("wxdaa2ab9ef87b5497", "AppSecret")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("token:", token)
}

func TestCode2Session(t *testing.T) {
	session, err := Code2Session("wx2e92b2ff5ed4db71", "AppSecret", "081XxRPj1e8Krp0uGUQj1s0MPj1XxRP5")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("Openid:", session.Openid)
}
