package alipay

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 换取授权访问令牌 alipay.system.oauth.token
// StatusCode = 200 is success
func (a *ClientV3) SystemOauthToken(ctx context.Context, bm gopay.BodyMap) (aliRsp *SystemOauthTokenRsp, err error) {
	err = bm.CheckEmptyError("grant_type")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3SystemOauthToken, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3SystemOauthToken, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &SystemOauthTokenRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 身份认证记录查询 alipay.user.certify.open.query
// StatusCode = 200 is success
func (a *ClientV3) UserCertifyOpenQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserCertifyOpenQueryRsp, err error) {
	err = bm.CheckEmptyError("certify_id")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3UserCertifyOpenQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &UserCertifyOpenQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 身份认证初始化服务 alipay.user.certify.open.initialize
// StatusCode = 200 is success
func (a *ClientV3) UserCertifyOpenInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserCertifyOpenInitializeRsp, err error) {
	err = bm.CheckEmptyError("outer_order_no", "biz_code", "identity_param")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3UserCertifyOpenInitialize, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3UserCertifyOpenInitialize, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &UserCertifyOpenInitializeRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 支付宝会员授权信息查询接口 alipay.user.info.share
// StatusCode = 200 is success
func (a *ClientV3) UserInfoShare(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserInfoShareRsp, err error) {
	err = bm.CheckEmptyError("auth_token")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	uri := v3UserInfoShare + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodPost, uri, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &UserInfoShareRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 用户授权关系查询 alipay.open.auth.userauth.relationship.query
// StatusCode = 200 is success
func (a *ClientV3) UserAuthRelationshipQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserAuthRelationshipQueryRsp, err error) {
	err = bm.CheckEmptyError("scopes")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3UserAuthRelationshipQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &UserAuthRelationshipQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 查询解除授权明细 alipay.user.deloauth.detail.query
// StatusCode = 200 is success
func (a *ClientV3) UserDelOauthDetailQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *UserDelOauthDetailQueryRsp, err error) {
	err = bm.CheckEmptyError("date", "limit", "offset")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3UserDelOauthDetailQuery, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3UserDelOauthDetailQuery, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &UserDelOauthDetailQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// MobilePhoneNumberDecryption 解密手机号
func (a *ClientV3) MobilePhoneNumberDecryption(response string) (*MobilePhoneNumberDecryptionResp, error) {
	if a.aesKey == "" {
		return nil, fmt.Errorf("aes key is empty")
	}
	var data, err = base64.StdEncoding.DecodeString(a.aesKey)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(data)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, 16)
	for i := 0; i < 16; i++ {
		iv[i] = 0
	}

	encryptedBytes, err := base64.StdEncoding.DecodeString(response)
	if err != nil {
		return nil, err
	}

	if len(encryptedBytes)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encryptedBytes, encryptedBytes)

	decryptedBytes, err := a.pkcs5Unpad(encryptedBytes)
	if err != nil {
		return nil, err
	}

	var resp = new(MobilePhoneNumberDecryptionResp)
	if err := json.Unmarshal(decryptedBytes, resp); err != nil {
		return nil, err
	}

	if resp.Code != "10000" {
		return resp, fmt.Errorf("%s:%s", resp.Msg, resp.SubMsg)
	}

	return resp, nil
}

func (a *ClientV3) pkcs5Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpadding error")
	}

	return src[:(length - unpadding)], nil
}
