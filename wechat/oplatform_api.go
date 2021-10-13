/*
	微信开放平台
	移动应用文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/Resource_Center_Homepage.html
	网站应用文档：https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
	第三方平台文档：https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Third_party_platform_appid.html
*/

package wechat

import "github.com/cedarwu/gopay/pkg/xhttp"

// GetOauth2AccessToken 微信第三方登录，code 换取 access_token
//	appId：应用唯一标识，在微信开放平台提交应用审核通过后获得
//	appSecret：应用密钥AppSecret，在微信开放平台提交应用审核通过后获得
//	code：App用户换取access_token的code
//	文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html
func GetOauth2AccessToken(appId, appSecret, code string) (accessToken *Oauth2AccessToken, err error) {
	accessToken = new(Oauth2AccessToken)
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appId + "&secret=" + appSecret + "&code=" + code + "&grant_type=authorization_code"

	_, errs := xhttp.NewClient().Get(url).EndStruct(accessToken)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return accessToken, nil
}

// RefreshOauth2AccessToken 刷新微信第三方登录后，获取到的 access_token
//	appId：应用唯一标识，在微信开放平台提交应用审核通过后获得
//	refreshToken：填写通过获取 access_token 获取到的 refresh_token 参数
//	文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html
func RefreshOauth2AccessToken(appId, refreshToken string) (accessToken *Oauth2AccessToken, err error) {
	accessToken = new(Oauth2AccessToken)
	url := "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=" + appId + "&grant_type=refresh_token&refresh_token=" + refreshToken

	_, errs := xhttp.NewClient().Get(url).EndStruct(accessToken)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return accessToken, nil
}

// CheckOauth2AccessToken 检验授权凭证（access_token）是否有效
//	accessToken：调用接口凭证
//	openid：普通用户标识，对该公众帐号唯一，获取 access_token 是获取的
//	文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func CheckOauth2AccessToken(accessToken, openid string) (result *CheckAccessTokenRsp, err error) {
	result = new(CheckAccessTokenRsp)
	url := "https://api.weixin.qq.com/sns/auth?access_token=" + accessToken + "&openid=" + openid

	_, errs := xhttp.NewClient().Get(url).EndStruct(result)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return result, nil
}

// GetOauth2UserInfo 微信开放平台：获取用户个人信息
//	accessToken：接口调用凭据
//	openId：用户的OpenID
//	lang:默认为 zh_CN ，可选填 zh_CN 简体，zh_TW 繁体，en 英语
//	文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func GetOauth2UserInfo(accessToken, openId string, lang ...string) (userInfo *Oauth2UserInfo, err error) {
	userInfo = new(Oauth2UserInfo)
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openId
	if len(lang) == 1 {
		url += "&lang=" + lang[0]
	}
	_, errs := xhttp.NewClient().Get(url).EndStruct(userInfo)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return userInfo, nil
}

// Deprecated
// 请替换 GetOauth2AccessToken 使用
func GetAppLoginAccessToken(appId, appSecret, code string) (accessToken *Oauth2AccessToken, err error) {
	accessToken = new(Oauth2AccessToken)
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appId + "&secret=" + appSecret + "&code=" + code + "&grant_type=authorization_code"

	_, errs := xhttp.NewClient().Get(url).EndStruct(accessToken)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return accessToken, nil
}

// Deprecated
// 请替换 RefreshOauth2AccessToken 使用
func RefreshAppLoginAccessToken(appId, refreshToken string) (accessToken *RefreshAppLoginAccessTokenRsp, err error) {
	accessToken = new(RefreshAppLoginAccessTokenRsp)
	url := "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=" + appId + "&grant_type=refresh_token&refresh_token=" + refreshToken

	_, errs := xhttp.NewClient().Get(url).EndStruct(accessToken)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return accessToken, nil
}

// Deprecated
// 推荐使用 GetOauth2UserInfo
func GetUserInfoOpen(accessToken, openId string, lang ...string) (userInfo *Oauth2UserInfo, err error) {
	userInfo = new(Oauth2UserInfo)
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openId
	if len(lang) == 1 {
		url += "&lang=" + lang[0]
	}
	_, errs := xhttp.NewClient().Get(url).EndStruct(userInfo)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return userInfo, nil
}
