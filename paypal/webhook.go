package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// CreateWebhook 创建Webhook
func (c *Client) CreateWebhook(ctx context.Context, bm gopay.BodyMap) (ppRsp *CreateWebhookRsp, err error) {
	if err = bm.CheckEmptyError("url", "event_types"); nil != err {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, createWebhook)
	if nil != err {
		return nil, err
	}
	ppRsp = &CreateWebhookRsp{Code: Success}
	ppRsp.Response = new(Webhook)
	if err = json.Unmarshal(bs, &ppRsp.Response); nil != err {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// ListWebhook 查询Webhook列表
func (c *Client) ListWebhook(ctx context.Context) (ppRsp *ListWebhookRsp, err error) {
	res, bs, err := c.doPayPalGet(ctx, listWebhook)
	if nil != err {
		return nil, err
	}
	ppRsp = &ListWebhookRsp{Code: Success}
	if err = json.Unmarshal(bs, &ppRsp.Response); nil != err {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// ShowWebhookDetail 查询Webhook消息消息
func (c *Client) ShowWebhookDetail(ctx context.Context, webhookId string) (ppRsp *WebhookDetailRsp, err error) {
	url := fmt.Sprintf(showWebhookDetail, webhookId)
	res, bs, err := c.doPayPalGet(ctx, url)
	if nil != err {
		return nil, err
	}
	ppRsp = &WebhookDetailRsp{Code: Success}
	if err = json.Unmarshal(bs, &ppRsp.Response); nil != err {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// UpdateWebhook 更新Webhook消息
func (c *Client) UpdateWebhook(ctx context.Context, webhookId string, patchs []*Patch) (ppRsp *WebhookDetailRsp, err error) {
	url := fmt.Sprintf(updateWebhook, webhookId)
	res, bs, err := c.doPayPalPatch(ctx, patchs, url)
	if nil != err {
		return nil, err
	}
	ppRsp = &WebhookDetailRsp{Code: Success}
	if err = json.Unmarshal(bs, &ppRsp.Response); nil != err {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// DeleteWebhook 删除Webhook消息
func (c *Client) DeleteWebhook(ctx context.Context, webhookId string) (ppRsp *WebhookDetailRsp, err error) {
	url := fmt.Sprintf(deleteWebhook, webhookId)
	res, bs, err := c.doPayPalDelete(ctx, url)
	if nil != err {
		return nil, err
	}
	ppRsp = &WebhookDetailRsp{Code: Success}
	if res.StatusCode != http.StatusNoContent {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
