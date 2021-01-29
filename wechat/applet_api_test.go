package wechat

import (
	"testing"

	"github.com/iGoogle-ink/gotil/xlog"
)

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
