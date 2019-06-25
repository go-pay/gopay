package gopay

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/alipay"
	"testing"
)

func TestAliPayClient_AliPayTradeWapPay(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		//SetReturnUrl("https://www.igoogle.ink").
		//SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify").
		SetNotifyUrl("https://gytmp.iguiyu.com/park-api/pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "手机网站测试支付")
	body.Set("out_trade_no", "GYWX201901301040355706100423")
	body.Set("quit_url", "https://www.igoogle.ink")
	body.Set("total_amount", "100.00")
	body.Set("product_code", "QUICK_WAP_WAY")
	//手机网站支付请求
	payUrl, err := client.AliPayTradeWapPay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payUrl:", payUrl)
}

func TestAliPayClient_AliPayTradeAppPay(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://gytmp.iguiyu.com/park-api/pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "测试APP支付")
	body.Set("out_trade_no", "GYWX201901301040355706100411")
	body.Set("total_amount", "1.00")
	body.Set("product_code", "QUICK_MSECURITY_PAY")
	//手机APP支付参数请求
	payParam, err := client.AliPayTradeAppPay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payParam:", payParam)
}

func TestAliPayClient_AliPayTradePagePay(t *testing.T) {
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://gytmp.iguiyu.com/park-api/pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "网站测试支付")
	body.Set("out_trade_no", "GYWX201901301040355706100433")
	body.Set("quit_url", "https://www.igoogle.ink")
	body.Set("total_amount", "88.88")
	body.Set("product_code", "FAST_INSTANT_TRADE_PAY")

	//电脑网站支付请求
	payUrl, err := client.AliPayTradePagePay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payUrl:", payUrl)
}

func TestAliPayParams(t *testing.T) {

	bodyMap := make(BodyMap)
	bodyMap.Set("subject", "测试支付")
	bodyMap.Set("out_trade_no", "GYWX201901301040355706100401")
	bodyMap.Set("quit_url", "https://www.igoogle.ink")
	bodyMap.Set("total_amount", "1.00")
	bodyMap.Set("product_code", "QUICK_WAP_WAY")

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

func TestFormatAliPayPublicKey(t *testing.T) {
	aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"

	pKey := FormatAliPayPublicKey(aliPayPublicKey)
	fmt.Println(pKey)
}

type List struct {
	BillList []FundBillListInfo `json:"bill_list"`
}

func TestJsonToString(t *testing.T) {

	list := new(List)
	infos := make([]FundBillListInfo, 0)

	infos = append(infos, FundBillListInfo{FundChannel: "iguiyu", Amount: "1.0.0"})
	infos = append(infos, FundBillListInfo{FundChannel: "Jerry", Amount: "2.0.2"})

	list.BillList = infos

	bs, err := json.Marshal(list)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("string:", string(bs))
	alipay.AckNotification()
}
