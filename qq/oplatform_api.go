package qq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
)

// 获取开放平台，access_token 返回结构体
type AccessToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// 获取开放平台，access_token 返回结构体
type OpenIdInfo struct {
	ClientId         string `json:"client_id,omitempty"`         // 用户ClientID
	OpenId           string `json:"openid,omitempty"`            // 用户OpenID
	UnionId          string `json:"unionid,omitempty"`           // 用户UnionID
	Error            int    `json:"error,omitempty"`             // 错误代码
	ErrorDescription string `json:"error_description,omitempty"` // 错误描述
}

// QQ开放平台用户信息
type UserInfo struct {
	Ret      int    `json:"sex,omitempty"`      // 返回码
	Msg      string `json:"msg,omitempty"`      // 如果ret<0，会有相应的错误信息提示，返回数据全部用UTF-8编码。
	Nickname string `json:"nickname,omitempty"` // 用户在QQ空间的昵称。

	Figureurl    string `json:"figureurl,omitempty"`      // 大小为30×30像素的QQ空间头像URL。
	Figureurl1   string `json:"figureurl_1,omitempty"`    // 大小为50×50像素的QQ空间头像URL。
	Figureurl2   string `json:"figureurl_2,omitempty"`    // 大小为100×100像素的QQ空间头像URL。
	FigureurlQq1 string `json:"figureurl_qq_1,omitempty"` // 大小为40×40像素的QQ头像URL。
	FigureurlQq2 string `json:"figureurl_qq_2,omitempty"` // 大小为100×100像素的QQ头像URL。需要注意，不是所有的用户都拥有QQ的100x100的头像，但40x40像素则是一定会有。

	Gender          string `json:"gender,omitempty"`             // 性别。 如果获取不到则默认返回"男"
	IsYellowVip     string `json:"is_yellow_vip,omitempty"`      // 标识用户是否为黄钻用户（0：不是；1：是）。
	Vip             string `json:"vip,omitempty"`                // 标识用户是否为黄钻用户（0：不是；1：是）
	YellowVipLevel  string `json:"yellow_vip_level,omitempty"`   // 黄钻等级
	Level           string `json:"level,omitempty"`              // 黄钻等级
	IsYellowYearVip string `json:"is_yellow_year_vip,omitempty"` // 标识是否为年费黄钻用户（0：不是； 1：是）
}

/*{
"ret":0,
"msg":"",
"nickname":"Peter",
"figureurl":"http://qzapp.qlogo.cn/qzapp/111111/942FEA70050EEAFBD4DCE2C1FC775E56/30",
"figureurl_1":"http://qzapp.qlogo.cn/qzapp/111111/942FEA70050EEAFBD4DCE2C1FC775E56/50",
"figureurl_2":"http://qzapp.qlogo.cn/qzapp/111111/942FEA70050EEAFBD4DCE2C1FC775E56/100",
"figureurl_qq_1":"http://q.qlogo.cn/qqapp/100312990/DE1931D5330620DBD07FB4A5422917B6/40",
"figureurl_qq_2":"http://q.qlogo.cn/qqapp/100312990/DE1931D5330620DBD07FB4A5422917B6/100",
"gender":"男",
"is_yellow_vip":"1",
"vip":"1",
"yellow_vip_level":"7",
"level":"7",
"is_yellow_year_vip":"1"
}*/

// 参数 是否必须 含义
// grant_type 必须 授权类型，在本步骤中，此值为“authorization_code”。
// client_id 必须 申请QQ登录成功后，分配给网站的appid。
// client_secret 必须 申请QQ登录成功后，分配给网站的appkey。
// code 必须 上一步返回的authorization code。
// 如果用户成功登录并授权，则会跳转到指定的回调地址，并在URL中带上Authorization Code。
// 例如，回调地址为www.qq.com/my.php，则跳转到：
// http://www.qq.com/my.php?code=520DD95263C1CFEA087******
// 注意此code会在10分钟内过期。
// redirect_uri 必须 与上面一步中传入的redirect_uri保持一致。
// fmt 可选 因历史原因，默认是x-www-form-urlencoded格式，如果填写json，则返回json格式
// 文档：https://wiki.connect.qq.com/%E4%BD%BF%E7%94%A8authorization_code%E8%8E%B7%E5%8F%96access_token#Step2.EF.BC.9A.E9.80.9A.E8.BF.87AuthorizationCode.E8.8E.B7.E5.8F.96AccessToken
func GetAccessToken(ctx context.Context, appId, appSecret, code, redirectUri string) (accessToken *AccessToken, err error) {
	accessToken = new(AccessToken)
	url := "https://graph.qq.com/oauth2.0/token?client_id=" + appId + "&client_secret=" + appSecret + "&code=" + code + "&redirect_uri=" + redirectUri + "&fmt=json" + "&grant_type=authorization_code"

	_, err = xhttp.NewClient().Req().Get(url).EndStruct(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

// GetOpenId QQ开放平台：使用Access Token来获取用户的OpenID
// accessToken：接口调用凭据
// openId：用户的OpenID
// oauthConsumerKey：AppID
// lang:默认为 zh_CN ，可选填 zh_CN 简体，zh_TW 繁体，en 英语
// 文档：https://wiki.open.qq.com/wiki/website/%E5%BC%80%E5%8F%91%E6%94%BB%E7%95%A5_Server-side#Step2.EF.BC.9A.E8.8E.B7.E5.8F.96Authorization_Code
func GetOpenId(ctx context.Context, accessToken string, lang ...string) (openid *OpenIdInfo, err error) {
	openid = new(OpenIdInfo)
	url := "https://graph.qq.com/oauth2.0/me?access_token=" + accessToken + "&unionid=1"
	if len(lang) == 1 {
		url += "&lang=" + lang[0]
	}
	_, bs, err := xhttp.NewClient().Req().Get(url).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	bs = bytes.ReplaceAll(bs, []byte("callback("), []byte(""))
	bs = bytes.ReplaceAll(bs, []byte(");"), []byte(""))
	err = json.Unmarshal(bs, openid)
	if err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, bs)
	}

	return
}

// GetUserInfo QQ开放平台：获取用户个人信息
// accessToken：接口调用凭据
// openId：用户的OpenID
// oauthConsumerKey：AppID
// lang:默认为 zh_CN ，可选填 zh_CN 简体，zh_TW 繁体，en 英语
// 文档：https://wiki.open.qq.com/wiki/website/get_user_info
func GetUserInfo(ctx context.Context, accessToken, openId string, oauthConsumerKey string, lang ...string) (userInfo *UserInfo, err error) {
	userInfo = new(UserInfo)
	url := "https://graph.qq.com/user/get_user_info?access_token=" + accessToken + "&openid=" + openId + "&oauth_consumer_key=" + oauthConsumerKey
	if len(lang) == 1 {
		url += "&lang=" + lang[0]
	}
	_, err = xhttp.NewClient().Req().Get(url).EndStruct(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
