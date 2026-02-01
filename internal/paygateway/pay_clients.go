package paygateway

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xhttp"
	wechatv3 "github.com/go-pay/gopay/wechat/v3"
)

type ClientManager struct {
	store *MerchantStore
	xhc   *xhttp.Client
	hc    *http.Client

	wechatV3 sync.Map // key tenant:merchant -> *wechatv3.ClientV3
	alipay   sync.Map // key tenant:merchant -> *alipay.Client
}

func NewClientManager(cfg *Config, store *MerchantStore) (*ClientManager, error) {
	transport, err := newSecureHTTPTransport(cfg.TLS.CAFile)
	if err != nil {
		return nil, err
	}
	hc := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	xhc := xhttp.NewClient()
	xhc.SetHttpTransport(transport)
	xhc.SetTimeout(30 * time.Second)

	return &ClientManager{
		store: store,
		xhc:   xhc,
		hc:    hc,
	}, nil
}

func (m *ClientManager) HTTPClient() *http.Client { return m.hc }

func (m *ClientManager) InvalidateKeys(keys []string) {
	for _, key := range keys {
		m.wechatV3.Delete(key)
		m.alipay.Delete(key)
	}
}

func (m *ClientManager) WechatV3(tenantID, merchantID string) (*wechatv3.ClientV3, *WechatV3Config, error) {
	mc, ok := m.store.Get(tenantID, merchantID)
	if !ok || mc.WechatV3 == nil {
		return nil, nil, fmt.Errorf("wechatV3 config not found for %s/%s", tenantID, merchantID)
	}
	key := merchantKey(tenantID, merchantID)
	if v, ok := m.wechatV3.Load(key); ok {
		return v.(*wechatv3.ClientV3), mc.WechatV3, nil
	}
	c, err := wechatv3.NewClientV3(mc.WechatV3.MchID, mc.WechatV3.SerialNo, mc.WechatV3.ApiV3Key, mc.WechatV3.PrivateKey)
	if err != nil {
		return nil, nil, err
	}
	c.SetHttpClient(m.xhc)
	if mc.WechatV3.WechatPayPublicKey != "" && mc.WechatV3.WechatPayPublicKeyID != "" {
		if err := c.AutoVerifySignByPublicKey([]byte(mc.WechatV3.WechatPayPublicKey), mc.WechatV3.WechatPayPublicKeyID); err != nil {
			return nil, nil, err
		}
	} else {
		if err := c.AutoVerifySign(); err != nil {
			return nil, nil, err
		}
	}
	actual, _ := m.wechatV3.LoadOrStore(key, c)
	return actual.(*wechatv3.ClientV3), mc.WechatV3, nil
}

func (m *ClientManager) Alipay(tenantID, merchantID string) (*alipay.Client, *AlipayConfig, error) {
	mc, ok := m.store.Get(tenantID, merchantID)
	if !ok || mc.Alipay == nil {
		return nil, nil, fmt.Errorf("alipay config not found for %s/%s", tenantID, merchantID)
	}
	key := merchantKey(tenantID, merchantID)
	if v, ok := m.alipay.Load(key); ok {
		return v.(*alipay.Client), mc.Alipay, nil
	}
	c, err := alipay.NewClient(mc.Alipay.AppID, mc.Alipay.PrivateKey, mc.Alipay.IsProd)
	if err != nil {
		return nil, nil, err
	}
	c.SetHttpClient(m.xhc)
	actual, _ := m.alipay.LoadOrStore(key, c)
	return actual.(*alipay.Client), mc.Alipay, nil
}
