package alipay

import (
	"fmt"
	"github.com/iGoogle-ink/gopay"
	"log"
	"net/http"
)

func FormatPrivateKey() {
	privateKey := "MIIEowIBAAKCAQEAxorj0SisU93LxawVWybZzDYfq8bxzw9VIOKDSW1M5ombl/MvZTT9IuBSGvqUxXFSWqY05e+CBlquFrOHTzKttmzeZJEKHpGAkW+4TsX6N4ZCHfz3Sz5MBAzBqQvscUD0FTZ+VlQbQ1HjQiKLbHaNY+fFVn3q2XL6dNtWwFhad592M1lCbzEjZ7yqFcxJ0jIDBh1lWwemHu52iW3YkvLD5lE8IXRiJedhgPFXCFAkX5HLcc0/jicRTarQ9gRNpHC6B87T2SDjKmgxqcws4CVdYJbztU1KLrsbMALFMOp24x8xsLgR88XeraQyUWe6V3lt1OYEOd9XLHWRx33bRHSwIDAQABAoIBAHjrGfjG1r11Nae8OH19WeRfikZqMdcztVsD2YWcxdsaL+MJPvJapVjaWecIehcN/2QqGcl4Zy5Lh/9Xc68uZFHYWFHTa+BWKYFqE0wWk1/Bqv7slAgFdvJ4enHkSypmrsFEoQkezEPh2ZDrzRJP2ajg/XTB14h72EHXXCxlIyP6q9ldlHyYSc+KOdC3WYOg1FoFXsliZHVKGZUxo4jck+xwdTGRAIKYpdLjpw7DqAWS6N25cx4XK9GuBYoV7AkIM0kpdjDDXAciG2aws6ef4kumWuW/JSbdrFWGLdiAN8GVpBx6+9WTeDKerS9KyDLNf6rsz9Hm3LNWOYCMLlrEFiECgYEA/XB6pZGrbmSebn2lZO+WFs3LYCXoCUy4ePouLYZO7lKsHNeTYRxCc6cbEmap0hpuMCYVPJkqK+nL/CDwBad1QChN5rPVFl2rLtLu0owvoAuTVjWjYNPgDfWb3spXh0p+lZ9oi53kZd4iDGe/jQJzAcpUa3Yj2me6XFFTD+8FNCkCgYEAyIxrkCo2oqGg1aJ7xc9+aBcpsrVg/upptdTgiMGSVZ3XuixufZV36lzJOdmCBoFGKwmLgKkStJSOm3SHUqdEKQBbHI95aG3HgnAMRXOtkn1exExfpAmCnAAnAx8RONorOTjrMrW0BO0/NII7O7NkLg/ocahr/lXEylsC8dLlMCgYBDiwiEu6/Oee5nUAEWR2veo/YBp9iRMeswAqzv4Q2EInBQN3vFs7xaCj0CyG2V2wlmt5+NSNyeW27LwRN2zkxHTvaD94VgspH+pqSTZF0E8FDR9vWVxqG91qk11QNCwS2/Pn6kRu4p3+t/Ft9L+00fOwcIpLGlcWOPWvUiF/dxEQKBgQCNEEhwpWC80FejLaFGKIdPjEtmSrKpXBV0rfTF+LkizuUBJ3/9zQNRyeGxnnuRj+nlvO1e3sWACySHRu4G53MvR8qqVr13ecfuuA0nOvPojuq4THKrlzVsUqGelXBrlEdiFFJMY7axfvBzoYIyqq+aoTxFjJ6Z/czFOZyp6tnpxQKBgClvDZ9pUc+WH28fDnDPd36bC6HmBq/fkxo92RJey1aRFSoCtKNW5Eaqem8iDD+WAVYak2Vg7xUHkhwIEyVVfHIxZBXc0X1w3jNFjG1/Fyul4hLjqCH2QI8gOjHXAcDZe+MJa8b33ZTiiilUu5A0N8+Xz8qpMQ84cXODHJcPMPYb"

	pKey := gopay.FormatPrivateKey(privateKey)
	fmt.Println(pKey)
}

func FormatAliPayPublicKey() {
	aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"

	pKey := gopay.FormatAliPayPublicKey(aliPayPublicKey)
	fmt.Println(pKey)
}

func AliPaySystemOauthToken() {
	privateKey := "MIIEpAIBAAKCAQEA0N6CzO4T6NEgWrlW+mrZYMZSNAQ/vBeaCzo1qT+VYNBbklYWxwyUT4JptswSkupdSZGzKVX8rgTBnIbJXjaL9OuUS4KpkOVtqnW9uaPlei7kh2vHkyzbJS85HG8r4XwrrNGY/o4kp/ur51X5dSNPR+Bd5L/0/SUxztsMS8w2onkKk+VNh95yNQx3ZX0XXDhJGmURHFhuWf3bd0XYGoAsLfuEVj95dKOVUodP6RduGJYd/23CBWQfsapA2SxB/X8/ol/7pCgI9UxYkSm88dFZF5BygBv4Alc77ZP2G05VLRtyaEgvzooXzb7N9QEPWsMD4e7sjuuyhUWXakP78BN7tQIDAQABAoIBAGSqeVwPeqUA9ZkELrlueeo8ay4KxjT/Iw06NtalgHcytI+o0j2JAIn5jr29SU5piiCFClZnzOqJ/E5WYAL6QyX9zHFAPj1Jdjtrl887AMSMiTQDbCKwwcYAJoRfyERatioKsrQtCdm0YDuUojAknndD2b0ADzFaldDeneSCPeq7kHVpAylhLIbzGmE7rGWy1gAnJ5ljza/QMQQkVX7MP852IFEwl+8dYeMK4snfKH4SBWbGgLSkfgemEL3mbHaocl2encFHQYjjFSWabUedOfmf0s81+KytsylKLxzZePAKCpbw7cBzf09wXO7YfFPWP3ORSZtEb6UsJAZBD7Kr1ZkCgYEA9DFdgInepyyFgg4w0ljg7HzRT+2U8xGV/rzO6Yc3HerXKWlhUQlAxq+nLW//MZ4RKRilkQmp9876S2jchATXroU+D620q7RmwN80EPhJC3HuAD/vv13rXJpRLhOElL4eMfbllqN/p7Xu0i4P8qrN/WmyICBN4HQxwiAPacOejNMCgYEA2vfoQpcpfDh+pfFvh1kpZLqEsqNdap9Su/dBmtv0oXDST78Azh4WXskoF0g9MQIkUZDjM3Aaa3qWeVI7p8FW8pk1vC1ktD/3x5K19d/eTJYjw4zaxLSDu/8KXkk8+NjpbBJNiUmidFk8y9n15o98AmLmRx5xYeHosH2Y+W6D4FcCgYEA0F7rbYq3/qkEB7V6PZPv4nyj03NlHYENIEEWwrlJ8/J3xXORwieGKcSrgB9IgJtkA1Bvv3Kiob3xYLXLLbBAJ0C1c4WDXN1krQIJmM107xo5v+bO+tn2w58+1HV9NbuBYptO0BsWDBM4BgueOjE78xjlEXrfWT9tUjwvSs3NFlcCgYBakX2ExYdINQnVU1aBEhT379nhFYcXCc0rp+p/xgiaGdMzXtpTZshz7hnh/2Qn+C3Teu0uomibGfpRNf0stuBiuZIrQk9L6sVuy7TtfOoynUsvn/wArnVqdePw+bP5baamp69gYI+MNjjaTE3UTnbJeeSrd+EHzmvzTA/Q47AyQwKBgQC50qZ0jV1WsNfyvv3YyGmlLiyLFEPZ6aDY+gntgeXJ3fCORy90LOk+4FfkaWdIEGH21+dIx7g7oEa9rTt2foqZnCmWWC3vRlDYsdpRYXz4UOKbHsBocWvSCSOmQ8XgwoDve1MhAz7yOlrHIzX816Ld7QHQdI/d/Z4j1LrMyIWP7A=="
	rsp, err := gopay.AliPaySystemOauthToken("2019071565794663", privateKey, "", "06e8961891d647c0ac99bb1cebe7SE69")
	if err != nil {
		fmt.Println("gopay.AliPaySystemOauthToken:", err)
		return
	}
	fmt.Println("rsp:", *rsp)
}

func DecryptAliPayOpenDataToStruct() {
	data := "MkvuiIZsGOC8S038cu/JIpoRKnF+ZFjoIRGf5d/K4+ctYjCtb/eEkwgrdB5TeH/93bxff1Ylb+SE+UGStlpvcg=="
	key := "TDftre9FpItr46e9BVNJcw=="
	rsp := new(gopay.PhoneNumberResponse)
	err := gopay.DecryptAliPayOpenDataToStruct(data, key, rsp)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("rsp.Code:", rsp.Code)
	fmt.Println("rsp.Msg:", rsp.Msg)
	fmt.Println("rsp.SubCode:", rsp.SubCode)
	fmt.Println("rsp.SubMsg:", rsp.SubMsg)
	fmt.Println("rsp.Mobile:", rsp.Mobile)
}

func ParseAliPayNotifyResultAndVerifyAliPayResultSign(req *http.Request) {
	aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"

	//解析请求参数
	notifyReq, err := gopay.ParseAliPayNotifyResult(req)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("notifyReq:", *notifyReq)

	//验签
	ok, err := gopay.VerifyAliPayResultSign(aliPayPublicKey, notifyReq)
	if err != nil {
		fmt.Println("err:", err)
	}
	log.Println("支付宝验签是否通过:", ok)
}
