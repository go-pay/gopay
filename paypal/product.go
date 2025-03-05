package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建产品（Create product）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/catalog-products/v1/#products_create
func (c *Client) ProductCreate(ctx context.Context, bm gopay.BodyMap) (ppRsp *ProductCreateRep, err error) {
	if err = bm.CheckEmptyError("name", "type"); err != nil {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, productCreate)
	if err != nil {
		return nil, err
	}
	ppRsp = &ProductCreateRep{Code: Success}
	ppRsp.Response = new(Product)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 产品列表（List products）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/catalog-products/v1/#products_list
func (c *Client) ProductList(ctx context.Context, bm gopay.BodyMap) (ppRsp *ProductsListRsp, err error) {
	uri := productList + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &ProductsListRsp{Code: Success}
	ppRsp.Response = new(ProductsList)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 产品详情（Show product details）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/catalog-products/v1/#products_get
func (c *Client) ProductDetails(ctx context.Context, productId string, bm gopay.BodyMap) (ppRsp *ProductDetailsRsp, err error) {
	if productId == gopay.NULL {
		return nil, errors.New("product_id is empty")
	}
	uri := fmt.Sprintf(productDetail, productId) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &ProductDetailsRsp{Code: Success}
	ppRsp.Response = new(ProductDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}

// 更新产品（Update product）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/catalog-products/v1/#products_patch
func (c *Client) ProductUpdate(ctx context.Context, productId string, patchs []*Patch) (ppRsp *EmptyRsp, err error) {
	if productId == gopay.NULL {
		return nil, errors.New("product_id is empty")
	}
	url := fmt.Sprintf(productUpdate, productId)
	res, bs, err := c.doPayPalPatch(ctx, patchs, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &EmptyRsp{Code: Success}
	if res.StatusCode != http.StatusNoContent {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
