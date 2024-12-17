package main

import (
	"context"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/xlog"
)

func main() {
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCRbpW2f4M7ciRpfcEwnE4QQsoLMN4LiINSgOxhS6cguO2g7GzKWn8Vj5pyE1Px3NJ0ymlBl0sQBnpid1v/CnEjd9erz4YWs+AplljDwvhAzD1S7BY4zSYOi/KZaQnLpSAFhp7B3L1az2uMB07pezzn6aaQr5E3qpusLNzvtiVUX+V21LZh7aga33qTe6ovbDh9b/SNb2l/cDDJ0oWVbi663173RfCo0LRG0caGUl8jRQ2X+Um3yI/S0BRMHxkr97NuPt+oH2JSrwLy1ueJz+OXxiVK7cvxwX94PhqvGD9rL5VQBkcnuM+d5jbzwQDaGu99ofFGA5gr8zIOpM3R1fzfAgMBAAECggEBAJB//Td4mQ8OyYkbj0HafGIByri11FpHSJxIRxYhUizcMhfL8ghZRJ5Ksh2YgLb3PxEWMOEgD1Ab76w4hmrxmBqyr9Mhfky+awNHRGk2Ju5TQLdMpoPHqoGXofO1/yqdpiR3bvSWfHpnT6exd0Hb1ldju8FVAIf793NqnBd9treOlPAsSI2lr+2Nxs56yCIZUIrD4iNbkkQCq9ZDZAxlf5hW/u8QS7ZsCu7KfK9NL+dhYPQzoI/Avtbkle2CI1A2pW9QyIDh8OhM2bqfFbw/uoxSp0QH6fPtgeoq5kEH/QX7WeVcYeroa9nTEpoKLlXaujss8aBzRLGFA5gScGgBqNkCgYEAy9N2Hp57Lb+i/kJTORc0a6TzixV5b7i2tR5QDwLS9c2H9hnytCtz6kq26IA1U7cTRYg6whiQEXpVLcW7bezDrED+phrImbWQ06KyWRuPvirpyDEsM4U56Q0U+zZ18x1wYYS5LNE4uC8lx9bSAZ3iPW1DS5liqABAGNusaIqsjtsCgYEAtqic9Bqm97EiXzdFfAQ2nEHhfP+vaKeeouHBBKVaW8WIqjU220ODWb6QuXJ1gQARhu+SSnO4/hHj8J/TPAbQxxVi8pB8fYoAIxqwQLe471NDQ+7hPTtuxPSnWKRGe/nNbzXWMXAEhGriLesdrUMEwfiAaCB/sJvg/zVYP68vn00CgYBxP2uotYthLtHkDXvqA4+Xo43BoEvZnDq/xTjcLkiCVBEP1vj3zBDag2q/IcT/X3/wqTFkeHtLirna9bse52cMxQv3xHfil3QBcqs/QUYLdhJwrYmOkp5Lc35N2cONMbUoKbMtSI4IIAIQ17XVOiDi2luCnJWpHIKAjUVszGf18wKBgQCgJ9GOuzvBHGBN6lGTfsI/WgiTKEEt2yLeIGG0aCOFKfs9enFB6b1m2A9fevuKg3vau36ipzsCEr+wwQJetH8kwBzFIGj3TiP2o7T82dpehbwJ+Y8muFDUgsukjk168LnvE95d9KERVgJBTtDzlTq7tN8p7azoNpUeUDCzjjCMVQKBgGH9tkNeEZ2Tcfuh8QKEQewZYNc+c3w/E1L/uFOuganmQUeLWbdXpFk0fAyJLEsfJoYeJK7OjmNG7mcwSQLbiqrtlGlT6n22pc+AhaToSZuM1Z7ExkzKIIVm/ijCLploIHW16SWe8qecax1xPK2XElIAhcGCe8sRH8Z+fiv2B9Vp"
	client, err := alipay.NewClient("2021002103640726", privateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl("http://crm.deepic.cn/ali/pay/notify")
	client.SetCertSnByPath("./appPublicCert.crt", "./alipayRootCert.crt", "./alipayPublicCert.crt")

	//请求参数
	bm := gopay.BodyMap{}
	bm.Set("out_order_no", "202104021339585117785701")
	bm.Set("out_request_no", "20210402133958511778570101")
	bm.Set("remark", "测试取消")
	rs, err := client.FundAuthOperationCancel(context.Background(), bm)
	xlog.Infof("rs:%v, err:%+v", rs, err)
}
