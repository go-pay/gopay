package wechat

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/go-pay/crypto/xrsa"
	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestClient_Transfer(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("partner_trade_no", util.RandomString(32)).
		Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8").
		Set("check_name", "FORCE_CHECK"). // NO_CHECK：不校验真实姓名 , FORCE_CHECK：强校验真实姓名
		Set("re_user_name", "付明明").       // 收款用户真实姓名。 如果check_name设置为FORCE_CHECK，则必填用户真实姓名
		Set("amount", 30).                // 企业付款金额，单位为分
		Set("desc", "测试转账").              // 企业付款备注，必填。注意：备注中的敏感词会被转成字符*
		Set("spbill_create_ip", "127.0.0.1")

	// 企业向微信用户个人付款（不支持沙箱环境）
	//    body：参数Body
	wxRsp, err := client.Transfer(ctx, bm)
	if err != nil {
		xlog.Errorf("client.Transfer(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
}

func Test_ProfitSharing(t *testing.T) {
	type Receiver struct {
		Type        string `json:"type"`
		Account     string `json:"account"`
		Amount      int    `json:"amount"`
		Description string `json:"description"`
	}

	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("transaction_id", "4208450740201411110007820472").
		Set("out_order_no", "P20150806125346")

	var rs []*Receiver
	item := &Receiver{
		Type:        "MERCHANT_ID",
		Account:     "190001001",
		Amount:      100,
		Description: "分到商户",
	}
	rs = append(rs, item)
	item2 := &Receiver{
		Type:        "PERSONAL_OPENID",
		Account:     "86693952",
		Amount:      888,
		Description: "分到个人",
	}
	rs = append(rs, item2)
	bs, _ := json.Marshal(rs)

	bm.Set("receivers", string(bs))

	wxRsp, err := client.ProfitSharing(ctx, bm)
	if err != nil {
		xlog.Errorf("client.ProfitSharingAddReceiver(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp:", wxRsp)
}

func Test_ProfitSharingAddReceiver(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32))

	receiver := make(gopay.BodyMap)
	receiver.Set("type", "MERCHANT_ID").
		Set("account", "190001001").
		Set("name", "商户全称").
		Set("relation_type", "STORE_OWNER")

	bm.Set("receiver", receiver.JsonBody())

	wxRsp, err := client.ProfitSharingAddReceiver(ctx, bm)
	if err != nil {
		xlog.Errorf("client.ProfitSharingAddReceiver(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp:", wxRsp)
}

func Test_ProfitSharingRemoveReceiver(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32))

	receiver := make(gopay.BodyMap)
	receiver.Set("type", "MERCHANT_ID").
		Set("account", "190001001")

	bm.Set("receiver", receiver.JsonBody())

	wxRsp, err := client.ProfitSharingRemoveReceiver(ctx, bm)
	if err != nil {
		xlog.Errorf("client.ProfitSharingRemoveReceiver(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp:", wxRsp)
}

func TestClient_GetRSAPublicKey(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("sign_type", SignType_MD5)
	publicKey, err := client.GetRSAPublicKey(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("publicKey:%#v", publicKey)
}

func TestClient_PayBank(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("partner_trade_no", mchId).
		Set("nonce_str", util.RandomString(32)).
		Set("bank_code", "1001"). // 招商银行，https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4&index=5
		Set("amount", 1)

	// publicKey 通过 client.GetRSAPublicKey() 获取
	// 加密 银行账号，需要转 base64，微信解密使用的是 sha1
	encryptBank, err := xrsa.RsaEncryptOAEP(sha1.New(), xrsa.PKCS1, "publicKey content", []byte("621400000000567"), nil)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 加密 银行收款人，需要转 base64，微信解密使用的是 sha1
	encryptName, err := xrsa.RsaEncryptOAEP(sha1.New(), xrsa.PKCS1, "publicKey content", []byte("Jerry"), nil)
	if err != nil {
		xlog.Error(err)
		return
	}
	bm.Set("enc_bank_no", base64.StdEncoding.EncodeToString(encryptBank)).
		Set("enc_true_name", base64.StdEncoding.EncodeToString(encryptName))

	// 企业付款到银行卡API
	wxRsp, err := client.PayBank(ctx, bm)
	if err != nil {
		xlog.Errorf("client.EntrustPaying(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}
