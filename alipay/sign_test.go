package alipay

import (
	"os"
	"testing"

	"github.com/go-pay/crypto/xrsa"
	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestSyncVerifySign(t *testing.T) {
	signData := `{"code":"10000","msg":"Success","buyer_logon_id":"854***@qq.com","buyer_pay_amount":"0.01","buyer_user_id":"2088102363632794","fund_bill_list":[{"amount":"0.01","fund_channel":"PCREDIT"}],"gmt_payment":"2019-08-29 20:14:05","invoice_amount":"0.01","out_trade_no":"GZ201901301040361012","point_amount":"0.00","receipt_amount":"0.01","total_amount":"0.01","trade_no":"2019082922001432790585537960"}`
	sign := "bk3SzX0CZRI811IJioS2XKQHcgMixUT8mYyGQj+vcOAQas7GIYi6LpykqqSc3m7+yvqoG0TdX/c2JjYnpw/J53JxtC2IC4vsLuIPIgghVo5qafsfSxEJ22w20RZDatI2dYqFVcj8Jp+4aesQ8zMMNw7cX9NLyk7kw3DecYeyQp+zrZMueZPqLh88Z+54G+e6QuSU++0ouqQVd4PkpPqy6YI+8MdMUX4Ve0jOQxMmYH8BC6n5ZsTH/uEaLEtzYVZdSw/xdSQ7K1SH73aEH8XbRYx6rL7RkKksrdvhezX+ThDjQ+fTWjvNFrGcg3fmqXRy2elvoalu+BQmqlkWWjEJYA=="
	alipayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"
	pKey := xrsa.FormatAlipayPublicKey(alipayPublicKey)
	err := verifySign(signData, sign, RSA2, pKey)
	if err != nil {
		xlog.Errorf("verifySign(),error:%+v", err)
	}
}

func TestVerifySign(t *testing.T) {
	// 测试，假数据，无法验签通过
	publicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"

	bm := make(gopay.BodyMap)
	bm.Set("sign", "f19WZ3rko3cVpSG3uEEJF0eb4DuZVLt4/GXnNw9qg8iHUsJLkav0V91R5SSTDhW5lgkn3Xhq7TkFRJiDXdVXMu3XUlsONArp3Iu4tXagYJWt9jbcnc2/l29VYDXPLNcs7BXEWFEaCZLutQY2U82AumEwSc1XBKtsLC4mVX3M3f/ExFQHWklJEBHArYBGe4535uFRlsT2fk6WVuX8CuYZatCrVF1o02xMS5aD29eICPkmin/h87OcTbE1syktyCU1WVgcypagUdGGPTF0SVDFf7FRov7+w7fiCGGGL10tNlK/MLzcewtN2dyGF6RLUX3m+HQ7sNEk2wylRXLNUFig==")
	bm.Set("seller_email", "imonkey@100tal.com")
	bm.Set("sign_type", RSA2)
	bm.Set("total_amount", "0.02")
	bm.Set("buyer_id", "2088812847201551")
	bm.Set("invoice_amount", "0.02")
	bm.Set("fund_bill_list", `[{"amount":"0.02","fundChannel":"PCREDIT"}]`)
	bm.Set("trade_no", "2020010222001401551430614892")
	bm.Set("receipt_amount", "0.02")
	bm.Set("buyer_pay_amount", "0.02")
	bm.Set("notify_time", "2020-01-02 16:18:21")
	bm.Set("subject", "商品")
	bm.Set("auth_app_id", "2015102700040153")
	bm.Set("charset", "utf-8")
	bm.Set("point_amount", "0.00")
	bm.Set("notify_type", "trade_status_sync")
	bm.Set("out_trade_no", "1086209247658383466")
	bm.Set("gmt_payment", "2020-01-02 16:18:21")
	bm.Set("trade_status", "TRADE_SUCCESS")
	bm.Set("version", "1.0")
	bm.Set("buyer_logon_id", "185****2920")
	bm.Set("gmt_create", "2020-01-02 16:18:21")
	bm.Set("app_id", "2015102700040153")
	bm.Set("seller_id", "2088631240818980")
	bm.Set("notify_id", "2020010200222161821001551453140885")

	ok, err := VerifySign(publicKey, bm)
	if err != nil {
		xlog.Errorf("VerifySign(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("OK:", ok)
}

func TestVerifySignWithCert(t *testing.T) {
	// 测试，假数据，无法验签通过
	bm := make(gopay.BodyMap)
	bm.Set("sign", "kPbQIjX+xQc8F0/A6/AocELIjhhZnGbcBN6G4MM/HmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo+lHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l/KL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM=")
	bm.Set("sign_type", RSA2)
	bm.Set("total_amount", "2.00")
	bm.Set("buyer_id", "2088102116773037")
	bm.Set("body", "大乐透2.1")
	bm.Set("trade_no", "2016071921001003030200089909")
	bm.Set("refund_fee", "0.00")
	bm.Set("notify_time", "2016-07-19 14:10:49")
	bm.Set("subject", "大乐透2.1")
	bm.Set("charset", "utf-8")
	bm.Set("notify_type", "trade_status_sync")
	bm.Set("out_trade_no", "0719141034-6418")
	bm.Set("gmt_close", "2016-07-19 14:10:46")
	bm.Set("gmt_payment", "2016-07-19 14:10:47")
	bm.Set("trade_status", "TRADE_SUCCESS")
	bm.Set("version", "1.0")
	bm.Set("gmt_create", "2016-07-19 14:10:44")
	bm.Set("app_id", "2015102700040153")
	bm.Set("seller_id", "2088102119685838")
	bm.Set("notify_id", "4a91b7a78a503640467525113fb7d8bg8e")
	// filePath
	filepath := "/cert/alipayPublicCert.crt"
	ok, err := VerifySignWithCert(filepath, bm)
	if err != nil {
		xlog.Errorf("VerifySignWithCert(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("VerifySignWithCert", "OK:", ok)
	// fileByte
	bts, err := os.ReadFile(filepath)
	if err != nil {
		xlog.Errorf("VerifySignWithCert(%+v),error:%+v", bm, err)
		return
	}
	ok, err = VerifySignWithCert(bts, bm)
	if err != nil {
		xlog.Errorf("VerifySignWithCert(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("VerifySignWithCert", "OK:", ok)
}

func TestGetCertSN(t *testing.T) {
	sn, err := GetCertSN("cert/alipayPublicCert.crt")
	if err != nil {
		xlog.Errorf("GetCertSN(),error:%+v", err)
		return
	}
	xlog.Debug(sn)
	pubKeyPath := "cert/appPublicCert.crt"
	sn, err = GetCertSN(pubKeyPath)
	if err != nil {
		xlog.Errorf("GetCertSN(),error:%+v", err)
		return
	}
	xlog.Debug(sn)
	bts, _ := os.ReadFile(pubKeyPath)
	sn, err = GetCertSN(bts)
	if err != nil {
		xlog.Errorf("GetCertSN(),error:%+v", err)
		return
	}
	xlog.Debug(sn)
	rootCrtPath := "cert/alipayRootCert.crt"
	sn, err = GetRootCertSN(rootCrtPath)
	if err != nil {
		xlog.Errorf("GetCertSN(),error:%+v", err)
		return
	}
	xlog.Debug(sn)
	bts, _ = os.ReadFile(rootCrtPath)
	sn, err = GetRootCertSN(bts)
	if err != nil {
		xlog.Errorf("GetCertSN(),error:%+v", err)
		return
	}
	xlog.Debug(sn)
	// Output:
	// 04afd423ea5bd6f5c5482854ed73278c
	// 4498aaa8ab0c8986c15c41b36186db7d
	// 4498aaa8ab0c8986c15c41b36186db7d
	// 687b59193f3f462dd5336e5abf83c5d8_02941eef3187dddf3d3b83462e1dfcf6
	// 687b59193f3f462dd5336e5abf83c5d8_02941eef3187dddf3d3b83462e1dfcf6
}
