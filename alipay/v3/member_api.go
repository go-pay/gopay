package alipay

import (
	"context"
	"encoding/json"
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
	authorization, err := a.authorization(MethodPost, v3SystemOauthToken, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3SystemOauthToken, authorization)
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
	uri := v3UserCertifyOpenQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
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
	authorization, err := a.authorization(MethodPost, v3UserCertifyOpenInitialize, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3UserCertifyOpenInitialize, authorization)
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
	err = bm.CheckEmptyError("avatar", "city", "nick_name", "province")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3UserInfoShare, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3UserInfoShare, authorization)
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
	uri := v3UserAuthRelationshipQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
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
	authorization, err := a.authorization(MethodPost, v3UserDelOauthDetailQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3UserDelOauthDetailQuery, authorization)
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
