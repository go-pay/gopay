package gopay

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/alipay"
	"log"
	"testing"
)

func TestPay(t *testing.T) {
	//网页&移动应用
	var AlipayAppId = "2016091200494382"
	var AlipayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	//私钥通用
	var AlipayPrivateKey = "MIIEowIBAAKCAQEAxorj0SisU93LxawVWybZzDYfq8bxzw9VIOKDSW1M5ombl/MvZTT9IuBSGvqUxXFSWqY05e+CBlquFrOHTzKttmzeZJEKHpGAkW+4TsX6N4ZCHfz3Sz5MBAzBqQvscUD0FTZ+VlQbQ1HjQiKLbHaNY+fFVn3q2XL6dNtWwFhad592M1lCbzEjZ7yqFcxJ0jIDBh1lWwemHu52iW3YkvLD5lE8IXRiJedh++gPFXCFAkX5HLcc0/jicRTarQ9gRNpHC6B87T2SDjKmgxqcws4CVdYJbztU1KLrsbMALFMOp24x8xsLgR88XeraQyUWe6V3lt1OYEOd9XLHWRx33bRHSwIDAQABAoIBAHjrGfjG1r11Nae8OH19WeRfikZqMdcztVsD2YWcxdsaL+MJPvJapVjaWecIehcN/2QqGcl4Zy5Lh/9Xc68uZFHYWFHTa+BWKYFqE0wWk1/Bqv7slAgFdvJ4enHkSypmrsFEoQkezEPh2ZDrzRJP2ajg/XTB14h72EHXXCxlIyP6q9ldlHyYSc+KOdC3WYOg1FoFXsliZHVKGZUxo4jck+xwdTGRAIKYpdLjpw7DqAWS6N25cx4XK9GuBYoV7AkIM0kpdjDDXAciG2aws6ef4kumWuW/JSbdrFWGLdiAN8GVpBx6+9WTeDKerS9KyDLNf6rsz9Hm3LNWOYCMLlrEFiECgYEA/XB6pZGrbmSebn2lZO+WFs3LYCXoCUy4ePouLYZO7lKsHNeTYRxCc6cbEmap0hpuMCYVPJkqK+nL/CDwBad1QChN5rPVFl2rLtLu0owvoAuTVjWjYNPgDfWb3spXh0p+lZ9oi53kZd4iDGe/jQJzAcpUa3Yj2me6XFFTD+8FNCkCgYEAyIxrkCo2oqGg1aJ7xc9+aBcpsrVg/uppt++dTgiMGSVZ3XuixufZV36lzJOdmCBoFGKwmLgKkStJSOm3SHUqdEKQBbHI95aG3HgnAMRXOtkn1exExfpAmCnAAnAx8RONorOTjrMrW0BO0/NII7O7NkLg/ocahr/lXEylsC8dLlMCgYBDiwiEu6/Oee5nUAEWR2veo/YBp9iRMeswAqzv4Q2EInBQN3vFs7xaCj0CyG2V2wlmt5+NSNyeW27LwRN2zkxHTvaD94VgspH+pqSTZF0E8FDR9vWVxqG91qk11QNCwS2/Pn6kRu4p3+t/Ft9L+00fOwcIpLGlcWOPWvUiF/dxEQKBgQCNEEhwpWC80FejLaFGKIdPjEtmSrKpXBV0rfTF+LkizuUBJ3/9zQNRyeGxnnuRj+nlvO1e3sWACySHRu4G53MvR8qqVr13ecfuuA0nOvPojuq4THKrlzVsUqGelXBrlEdiFFJMY7axfvBzoYIyqq+aoTxFjJ6Z/czFOZyp6tnpxQKBgClvDZ9pUc+WH28fDnDPd36bC6HmBq/fkxo92RJey1aRFSoCtKNW5Eaqem8iDD+WAVYak2Vg7xUHkhwIEyVVfHIxZBXc0X1w3jNFjG1/Fyul4hLjqCH2QI8gOjHXAcDZe+MJa8b33ZTiiilUu5A0N8+Xz8qpMQ84cXODHJcPMPYb"
	client := alipay.New(AlipayAppId, AlipayPublicKey, AlipayPrivateKey, false)

	p := alipay.AliPayTradeWapPay{}
	p.NotifyURL = "https://wx1.iguiyu.com/v1/parking/alipay/recharge/notify"
	p.Subject = "充值"
	p.OutTradeNo = GetRandomString(32)
	p.TotalAmount = "0.01"
	p.ProductCode = "充值"
	var payURL, err = client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("支付URL：", payURL)
}

func TestAliPayClient_AliPayTradeWapPay(t *testing.T) {
	aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"
	privateKey := "MIIEowIBAAKCAQEAxcVdmw2Rie/nPjylx/VOIKJa4bofSZhdikpb38vVQvVOjgiV6xfhNTpv5dA5+Sdx6O2lF+KKngkC0sdbY3jWTWgk0fx5rTB1M6wzqZxJdchU8HcGm/dIIq5aBnet4RjsYU++3q0cQbZSUMzznMgf1Lz7gHFZS3IIYUMAZloYpcklqxJNXq1ZREthW96FbZyn1rbycdXjlpEbqIHDJuVqUL74eh+53xL9Tj8GCOoRBA0iqlPLVL4JpL7e4E6defftaypqB2aC0fJO8trNrbWCYxsIyeFjZeRKqVzYMFc+RFfDyIK3yC3hjzMOKoxWXkYw4uml/Qrkwt7NVvCzvR9NeQIDAQABAoIBABbBqAywWf/KOAyEQ/snMc81f0mb9f+s5Y6FEd9FgAuNWHWlbUK447QRPlDuTc1qiYPo3GdMLPEUTlvcjpp6jAYqJpp297VC7yl79hHdJuLDo2pr97m4kXdUIo299acCDCkCWQ8cUjUJep1Lh/iRWoBLIpFb+Y9h1q8CW6hrU4y3zZYRaww3rjj7zl19vaEXF5uf/mmsoIToa2bl3vVW4GK1utkhF8iiK5VvBgiWEmTCVpNCCOKbv5Ma8xDrip0CqzCghqIYVTOwqS65FOoGrwzq/4Ivli2/tjnfm6BZcgbxUf6CoKj4ism8P3r5io9yrj/9RXkihjuPw6DmTWVYkLUCgYEA8ECyT9oi2b04c9RzE5s2nzO87TqI/zL0yhSd7Xlmn4JfNg49YI0TXfyupJ7BQSyuu/DAGdiLPWSoHolDhtDBLSaA9l1x26iGcs/ihu8H5jIz0De1NuM4Ef3vLD6LF9khCpzqnyWyCcOFh384Tu3q7kqmncF4Mqj/zJtglfed15sCgYEA0rvZc5ECUZuT0MyRU4Mj2IrRwc6aqO2tnxqEevZRnfmnTWiGfT8Op2CvHw9hKnQdWJUYqWrudG63HVx8/8Lvyx8XZlXVlKjU92rlaek6dO1Pw/z4DTuxE6WPoup2WfA1kjzrmhdTNzacrj6pJt/38HPRdhLnWVMkHTE6bBEkgnsCgYEA0zneS4xaRZsyDcxEHIHDBTdErEFhfxU62IxFySqKCkViFjFwzvlZhLGKjhsxh26UdZIWIuMakDB2CtrdrqLMpDiM+41udBP3mOuimsV+6WlL2o2P2iDtBAyBAiI+wgnZHe6V7LQEksb/GADG7cYJXdXuJRaa6ddhhm84/MDGWm8CgYAzISrQdGWIoWPK7GdySMZAuuXLzTIPPKO8j7WHFA6XcsRZ7rt61frbN4Ul1xhvMX8RSBOUv4Ids+Mv94nIkGaX9PI7fSX2DMSnR0NkYBcz2YRZ/B2/MDV1m9zu3U5b4gFNewR6/Z/OLKz1RfTKntrMd31h1ZJWROrlPlV7dOlT0wKBgEKY2m8grGRWFOmaMlDuN/fICtcv2KCDmC+ogs/QlGjsIld8YXkOrZiDr3u3V4YYlqJqKSmiJm2pDffDC7xtAiJ3o2dU9MH5RoLkefJSnDC3npKUY7iu7oBpsJJ0vFFWTkH8KHnnfWhElBSU9vPXakxIHLXLltF27hn3Ry4N826k"
	client := NewAliPayClient("2018112762320936", aliPayPublicKey, privateKey, true)
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetReturnUrl("https://api.iguiyu.com").
		SetNotifyUrl("https://api.iguiyu.com")

	body := make(BodyMap)
	body.Set("subject", "测试支付")
	body.Set("out_trade_no", "GYWX201901301040355706100401")
	body.Set("total_amount", "1.00")
	body.Set("product_code", "QUICK_WAP_WAY")

	err := client.AliPayTradeWapPay(body)
	if err != nil {
		log.Println("err2:", err)
	}

}

func TestAliPayParams(t *testing.T) {
	//body := new(aliPayPublicBody)
	//body.AppId = "2018112762320936"
	//body.Method = "alipay.trade.wap.pay"
	//body.Format = "JSON"
	//body.ReturnUrl = "https://www"
	//body.Charset = "utf-8"
	//body.SignType = "RSA2"
	//body.Sign = "asddgfnkjsdklasjdklfl"
	//body.Timestamp = "2006-01-02 15:04:05"
	//body.Version = "1.0"
	//body.NotifyUrl = "https://www"

	bodyMap := make(BodyMap)
	bodyMap.Set("subject", "测试支付")
	bodyMap.Set("out_trade_no", "GYWX201901301040355706100401")
	bodyMap.Set("total_amount", "1.00")
	bodyMap.Set("product_code", "QUICK_WAP_WAY")

	//fmt.Println("bodyMap:", bodyMap)

	//body.BizContent = bodyMap

	//fmt.Println(body.BizContent)

	bytes, err := json.Marshal(bodyMap)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("result:", string(bytes))
}

func TestFormatPrivateKey(t *testing.T) {
	privateKey := "MIIEowIBAAKCAQEAxorj0SisU93LxawVWybZzDYfq8bxzw9VIOKDSW1M5ombl/MvZTT9IuBSGvqUxXFSWqY05e+CBlquFrOHTzKttmzeZJEKHpGAkW+4TsX6N4ZCHfz3Sz5MBAzBqQvscUD0FTZ+VlQbQ1HjQiKLbHaNY+fFVn3q2XL6dNtWwFhad592M1lCbzEjZ7yqFcxJ0jIDBh1lWwemHu52iW3YkvLD5lE8IXRiJedhgPFXCFAkX5HLcc0/jicRTarQ9gRNpHC6B87T2SDjKmgxqcws4CVdYJbztU1KLrsbMALFMOp24x8xsLgR88XeraQyUWe6V3lt1OYEOd9XLHWRx33bRHSwIDAQABAoIBAHjrGfjG1r11Nae8OH19WeRfikZqMdcztVsD2YWcxdsaL+MJPvJapVjaWecIehcN/2QqGcl4Zy5Lh/9Xc68uZFHYWFHTa+BWKYFqE0wWk1/Bqv7slAgFdvJ4enHkSypmrsFEoQkezEPh2ZDrzRJP2ajg/XTB14h72EHXXCxlIyP6q9ldlHyYSc+KOdC3WYOg1FoFXsliZHVKGZUxo4jck+xwdTGRAIKYpdLjpw7DqAWS6N25cx4XK9GuBYoV7AkIM0kpdjDDXAciG2aws6ef4kumWuW/JSbdrFWGLdiAN8GVpBx6+9WTeDKerS9KyDLNf6rsz9Hm3LNWOYCMLlrEFiECgYEA/XB6pZGrbmSebn2lZO+WFs3LYCXoCUy4ePouLYZO7lKsHNeTYRxCc6cbEmap0hpuMCYVPJkqK+nL/CDwBad1QChN5rPVFl2rLtLu0owvoAuTVjWjYNPgDfWb3spXh0p+lZ9oi53kZd4iDGe/jQJzAcpUa3Yj2me6XFFTD+8FNCkCgYEAyIxrkCo2oqGg1aJ7xc9+aBcpsrVg/upptdTgiMGSVZ3XuixufZV36lzJOdmCBoFGKwmLgKkStJSOm3SHUqdEKQBbHI95aG3HgnAMRXOtkn1exExfpAmCnAAnAx8RONorOTjrMrW0BO0/NII7O7NkLg/ocahr/lXEylsC8dLlMCgYBDiwiEu6/Oee5nUAEWR2veo/YBp9iRMeswAqzv4Q2EInBQN3vFs7xaCj0CyG2V2wlmt5+NSNyeW27LwRN2zkxHTvaD94VgspH+pqSTZF0E8FDR9vWVxqG91qk11QNCwS2/Pn6kRu4p3+t/Ft9L+00fOwcIpLGlcWOPWvUiF/dxEQKBgQCNEEhwpWC80FejLaFGKIdPjEtmSrKpXBV0rfTF+LkizuUBJ3/9zQNRyeGxnnuRj+nlvO1e3sWACySHRu4G53MvR8qqVr13ecfuuA0nOvPojuq4THKrlzVsUqGelXBrlEdiFFJMY7axfvBzoYIyqq+aoTxFjJ6Z/czFOZyp6tnpxQKBgClvDZ9pUc+WH28fDnDPd36bC6HmBq/fkxo92RJey1aRFSoCtKNW5Eaqem8iDD+WAVYak2Vg7xUHkhwIEyVVfHIxZBXc0X1w3jNFjG1/Fyul4hLjqCH2QI8gOjHXAcDZe+MJa8b33ZTiiilUu5A0N8+Xz8qpMQ84cXODHJcPMPYb"

	pKey := FormatPrivateKey(privateKey)
	fmt.Println(pKey)

}

func TestFormatPublickKey(t *testing.T) {
	aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"

	pKey := FormatPublickKey(aliPayPublicKey)
	fmt.Println(pKey)
}
