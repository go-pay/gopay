//==================================
//  * Name：Jerry
//  * DateTime：2019/5/6 13:16
//  * Desc：
//==================================
package gopay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strings"
)

//JSAPI支付，支付参数后，再次计算出小程序用的paySign
func GetMiniPaySign(appId, nonceStr, prepayId, signType, timeStamp, secretKey string) (paySign string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("appId=")
	buffer.WriteString(appId)

	buffer.WriteString("&nonceStr=")
	buffer.WriteString(nonceStr)

	buffer.WriteString("&package=")
	buffer.WriteString(prepayId)

	buffer.WriteString("&signType=")
	buffer.WriteString(signType)

	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timeStamp)

	buffer.WriteString("&key=")
	buffer.WriteString(secretKey)

	signStr := buffer.String()

	var hashSign []byte
	if signType == SignType_MD5 {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := hmac.New(sha256.New, []byte(secretKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	paySign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//获取微信用户的OpenId、SessionKey、UnionId
func GetWeChatUserId(appId, secretKey, wxCode string) (userRsp *WeChatUserIdRsp, err error) {
	userRsp = new(WeChatUserIdRsp)

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", appId, secretKey, wxCode)

	agent := gorequest.New()
	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	agent.TLSClientConfig(tlsCfg)
	_, _, errs := agent.Get(url).EndStruct(userRsp)
	if len(errs) > 0 {
		return nil, errs[0]
	} else {
		return userRsp, nil
	}
}
