/*
	微信公众号
	文档：https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html
*/

package wechat

import "github.com/cedarwu/gopay/pkg/xhttp"

// GetPublicUserInfo 获取用户基本信息（微信公众号）
//	accessToken：接口调用凭据
//	openId：用户的OpenID
//	lang:默认为 zh_CN ，可选填 zh_CN 简体，zh_TW 繁体，en 英语
//	获取公众号用户基本信息文档：https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140839
func GetPublicUserInfo(accessToken, openId string, lang ...string) (userInfo *PublicUserInfo, err error) {
	userInfo = new(PublicUserInfo)
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openId
	if len(lang) == 1 {
		url += "&lang=" + lang[0]
	}
	_, errs := xhttp.NewClient().Get(url).EndStruct(userInfo)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return userInfo, nil
}

// GetPublicUserInfoBatch 批量获取用户基本信息（微信公众号）
//	accessToken：接口调用凭据
//	注意：开发者可通过该接口来批量获取用户基本信息。最多支持一次拉取100条。
//	文档：https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140839
func GetPublicUserInfoBatch(accessToken string, users *PublicOpenids) (userInfos *PublicUserInfoBatch, err error) {
	userInfos = new(PublicUserInfoBatch)
	url := "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=" + accessToken
	_, errs := xhttp.NewClient().Post(url).Type(xhttp.TypeJSON).SendStruct(users).EndStruct(userInfos)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return userInfos, nil
}

// Deprecated
// 推荐使用 GetPublicUserInfo
func GetUserInfo(accessToken, openId string, lang ...string) (userInfo *PublicUserInfo, err error) {
	userInfo = new(PublicUserInfo)
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openId + "&lang=zh_CN"
	if len(lang) > 0 {
		url = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openId + "&lang=" + lang[0]
	}
	_, errs := xhttp.NewClient().Get(url).EndStruct(userInfo)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return userInfo, nil
}
