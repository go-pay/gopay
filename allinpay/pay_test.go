package allinpay

import (
	"testing"
	"time"

	"github.com/go-pay/xlog"

	"github.com/go-pay/gopay"
)

func TestClient_ScanPay(t *testing.T) {
	client, err = NewClient("660121065134QCZ", "00382908",
		"MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC1izXRKdSXuwse3SDP9xDoHg3UnvNOq8RhcU/5479w5pYJiSTCkOrC45+2SsdItwJSxxZOnvRTZl46n/q/IoO5HDWSk45Ox2jY9Adii7GPyBJxl0Bl+skTKkMUovfbc1G4Yy+2mwhRd3JD7MWs42M/d91ENFYyQV9QF3f77TKriZX20o0apthnC88L2AFnBG8CkzK+B2izoMmbkffhC2L088UZQYUHHpC/GjLVcgdHwEzIlRSHwy1rn7PmYJ6s7hiwXd9Q/uqjncCJ9yunvwvxs6y00F2T9hvU7hvjn5xlW6DSY/MkXYLVblNWhYmD0ivCeKtx96SjdBocUXLG9WZHAgMBAAECggEAAfjIMtvAizBCKnsbSKzKHsQEzPusQfGUwKUZhAEhcHOTnIRbiyC6XX1ZnAIrjWbQ6rwopyPGHKyxma1om0HCtlGhZA63UYWMN58Ho1KN1BKB31JVyOiuxgzhGOnrf57fwaAH/BeTWHZywOhLz4ejo8czJVAiwbfjsDkCfMrp1iiQ/lLI+npHYnG9qGAuUqa9C0NSJlTP+tG1m4DBkodJIXPsmTmqTudav1cQKsMQr8zcf8yWtAdB9XhqTwqLAiLR76w2KWwRSxjtjCaX8QTITcjWX2Rpj4iCmuFT0qy0THZRwURX2faXO3GTg7e23wTJ6GdPmYqCbGME+ZSAvm81wQKBgQDa3Vo8qYYRcqraSJPbdFfQUR4DXwsdgdj/39dCSxIAWlYAKdsZiWjfwh3u69ENZxKYjEvJvil78KIahqggNAobARDV+2mSdsnTn0UxtTGxbF+AGXVARDK+GRXu/aN1qpuOBoXXRmTWfKLIsRnBjT2m4u14Rs+0shT17ByGcyB5FwKBgQDUWMgDe2ZV9oYxAo5PbrqXkLJ4Vw3nBPg4ETGHFvDPmrHTTL8ww0ZT4qLb4Bi0145GVvdf4fB4Q6rUfEBKTwmhHLW4HxjIGGmAPmiTfaMLhcS370uZ9q1YddYoMIkbpsfG1Ovlaw3BpsaN6HYIG7aD69s8Rzv1kiswqIPc0etaUQKBgQCxCj8VNPgbIwtreSwJaAokm4lQptAh9UgoatQAbyNHn8tTZIg0Fv/7iFWYT68STV8wgMRJlAaQmC04kdZr/kxyXaFVxoI2lNpb0ckyFWT2JTj3MSC0dLYrKbWhVhCkfPZJo6MeaXXmPnmbKA8yjuLhHU2EbptTin4EFBNa+sO4zwKBgQCi1bILeYVwRiuBScR3hSHxODSjs54lmdn+pLCmNyFTEf1rW18Btb0odpMMqTxI8UNZDTeFf203zCwj5WQnl5R6x9lR6AbI6m98uF7ZO27cVygTJJ2E43TqcPJYv19kpPBtuLlspDtYVNvN1oUskpPTdfoYr0d7eStFlPQJUbB5AQKBgC8pfh6GaYRKB+GfLWe7EIZ86V8r/aTzGn/eFIJyl+b5ocXjb6LaysUz0A+UWe+jgiUYZOZ9fuY2rYhf9eVPQPjEWiwnwCM29HXCqvom0LcYus8vORsgHxgEHV5HM26YzoASECzsEoOAerfMHPlBrRQ2PrR+2bQev+VylhtO8rNI",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCm9OV6zH5DYH/ZnAVYHscEELdCNfNTHGuBv1nYYEY9FrOzE0/4kLl9f7Y9dkWHlc2ocDwbrFSm0Vqz0q2rJPxXUYBCQl5yW3jzuKSXif7q1yOwkFVtJXvuhf5WRy+1X5FOFoMvS7538No0RpnLzmNi3ktmiqmhpcY/1pmt20FHQQIDAQAB",
		true)
	if err != nil {
		t.Fatal(err)
	}

	// 扫码支付
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("trxamt", "1").
		Set("reqsn", "larry01").
		Set("body", "支付测试").
		SetBodyMap("terminfo", func(b gopay.BodyMap) {
			b.Set("devicetype", "10").
				Set("termno", "00000001")
		}).
		Set("authcode", "131104796948096102")
	// 创建订单
	resp, err := client.ScanPay(ctx, bm)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rsp:%+v", resp)

}

func TestClient_Query(t *testing.T) {
	// 查询订单
	resp, err := client.Query(ctx, OrderTypeReqSN, "larry01")
	xlog.Debugf("aliRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_Refund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("trxamt", "1").
		Set("reqsn", "relarry01").
		Set("remark", "支付测试退款").
		Set("oldreqsn", "larry01")

	// 退款
	resp, err := client.Refund(ctx, bm)
	xlog.Debugf("allRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_Cancel(t *testing.T) {
	// 订单退款
	bm := make(gopay.BodyMap)
	bm.Set("trxamt", "1").
		Set("reqsn", "cclarry01").
		Set("remark", "支付测试取消").
		Set("oldreqsn", "larry01")

	// 取消订单
	resp, err := client.Cancel(ctx, bm)
	xlog.Debugf("allRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_Close(t *testing.T) {
	// 订单关闭
	bm := make(gopay.BodyMap)
	bm.Set("oldreqsn", "larry01")
	// 创建订单
	resp, err := client.Close(ctx, bm)
	xlog.Debugf("allRsp:%+v", resp)
	if err != nil {
		xlog.Errorf("%+v", err)
		return
	}
}

func TestClient_NativePay(t *testing.T) {
	client, err = NewClient("660121065134QCZ", "00382908",
		"MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC1izXRKdSXuwse3SDP9xDoHg3UnvNOq8RhcU/5479w5pYJiSTCkOrC45+2SsdItwJSxxZOnvRTZl46n/q/IoO5HDWSk45Ox2jY9Adii7GPyBJxl0Bl+skTKkMUovfbc1G4Yy+2mwhRd3JD7MWs42M/d91ENFYyQV9QF3f77TKriZX20o0apthnC88L2AFnBG8CkzK+B2izoMmbkffhC2L088UZQYUHHpC/GjLVcgdHwEzIlRSHwy1rn7PmYJ6s7hiwXd9Q/uqjncCJ9yunvwvxs6y00F2T9hvU7hvjn5xlW6DSY/MkXYLVblNWhYmD0ivCeKtx96SjdBocUXLG9WZHAgMBAAECggEAAfjIMtvAizBCKnsbSKzKHsQEzPusQfGUwKUZhAEhcHOTnIRbiyC6XX1ZnAIrjWbQ6rwopyPGHKyxma1om0HCtlGhZA63UYWMN58Ho1KN1BKB31JVyOiuxgzhGOnrf57fwaAH/BeTWHZywOhLz4ejo8czJVAiwbfjsDkCfMrp1iiQ/lLI+npHYnG9qGAuUqa9C0NSJlTP+tG1m4DBkodJIXPsmTmqTudav1cQKsMQr8zcf8yWtAdB9XhqTwqLAiLR76w2KWwRSxjtjCaX8QTITcjWX2Rpj4iCmuFT0qy0THZRwURX2faXO3GTg7e23wTJ6GdPmYqCbGME+ZSAvm81wQKBgQDa3Vo8qYYRcqraSJPbdFfQUR4DXwsdgdj/39dCSxIAWlYAKdsZiWjfwh3u69ENZxKYjEvJvil78KIahqggNAobARDV+2mSdsnTn0UxtTGxbF+AGXVARDK+GRXu/aN1qpuOBoXXRmTWfKLIsRnBjT2m4u14Rs+0shT17ByGcyB5FwKBgQDUWMgDe2ZV9oYxAo5PbrqXkLJ4Vw3nBPg4ETGHFvDPmrHTTL8ww0ZT4qLb4Bi0145GVvdf4fB4Q6rUfEBKTwmhHLW4HxjIGGmAPmiTfaMLhcS370uZ9q1YddYoMIkbpsfG1Ovlaw3BpsaN6HYIG7aD69s8Rzv1kiswqIPc0etaUQKBgQCxCj8VNPgbIwtreSwJaAokm4lQptAh9UgoatQAbyNHn8tTZIg0Fv/7iFWYT68STV8wgMRJlAaQmC04kdZr/kxyXaFVxoI2lNpb0ckyFWT2JTj3MSC0dLYrKbWhVhCkfPZJo6MeaXXmPnmbKA8yjuLhHU2EbptTin4EFBNa+sO4zwKBgQCi1bILeYVwRiuBScR3hSHxODSjs54lmdn+pLCmNyFTEf1rW18Btb0odpMMqTxI8UNZDTeFf203zCwj5WQnl5R6x9lR6AbI6m98uF7ZO27cVygTJJ2E43TqcPJYv19kpPBtuLlspDtYVNvN1oUskpPTdfoYr0d7eStFlPQJUbB5AQKBgC8pfh6GaYRKB+GfLWe7EIZ86V8r/aTzGn/eFIJyl+b5ocXjb6LaysUz0A+UWe+jgiUYZOZ9fuY2rYhf9eVPQPjEWiwnwCM29HXCqvom0LcYus8vORsgHxgEHV5HM26YzoASECzsEoOAerfMHPlBrRQ2PrR+2bQev+VylhtO8rNI",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCm9OV6zH5DYH/ZnAVYHscEELdCNfNTHGuBv1nYYEY9FrOzE0/4kLl9f7Y9dkWHlc2ocDwbrFSm0Vqz0q2rJPxXUYBCQl5yW3jzuKSXif7q1yOwkFVtJXvuhf5WRy+1X5FOFoMvS7538No0RpnLzmNi3ktmiqmhpcY/1pmt20FHQQIDAQAB",
		true)
	if err != nil {
		t.Fatal(err)
	}

	expire := time.Now().Add(10 * time.Minute).Format("20060102150405")
	bm := make(gopay.BodyMap)
	bm.Set("reqsn", "nativelarry01").Set("trxamt", "1").Set("body", "支付测试").Set("expiretime", expire)
	resp, err := client.NativePay(ctx, bm)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rsp:%+v", resp)

	// Output:
	// &{RspBase:{RetCode:SUCCESS RetMsg: Sign:*** Cusid:*** Appid:***} ReqSn:nativelarry01 TrxStatus:0000 ErrMsg:生成收款码成功 PayInfo:https://syb.allinpay.com/apiweb/h5unionpay/unionnative?token=** RandomStr:636549522285}
}

func TestClient_NativeClose(t *testing.T) {
	resp, err := client.NativeClose(ctx, OrderTypeReqSN, "nativelarry01")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rsp:%+v", resp)

	// Output:
	// &{RspBase:{RetCode:SUCCESS RetMsg: Sign:**** Cusid:***** Appid:*****} TrxStatus:0000 ErrMsg:交易关闭成功 RandomStr:476373212870}
}

func TestClient_QueryConfirm(t *testing.T) {
	resp, err := client.QueryConfirm(ctx, OrderTypeReqSN, "nativelarry01")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rsp:%+v", resp)
}
