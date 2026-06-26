// Package paypal
// Copyright 2025 Giga Inc. All Rights Reserved
// Author: Stuart <stuartjing@sina.com>
// Date: 2025/3/4
/*
   DESCRIPTION: xx
   使用文档：xx
*/
package paypal

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
	"testing"
)

func TestCreateProduct(t *testing.T) {

	bm := make(gopay.BodyMap)
	// can be AUTHORIZE
	bm.Set("name", "Video Gen Service").
		Set("description", "Video generate service").
		Set("type", "DIGITAL").
		Set("category", "PICTURE_VIDEO_PRODUCTION")
	//Set("image_url", "https://example.com/streaming.jpg").
	//Set("home_url", "https://example.com/home")
	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.ProductCreate(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestListProduct(t *testing.T) {
	ppRsp, err := client.ProductList(ctx, nil)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
	for _, v := range ppRsp.Response.Items {
		xlog.Debugf("ppRsp.Response.Item: %+v", v)
	}
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestProductDetail(t *testing.T) {
	ppRsp, err := client.ProductDetails(ctx, "PROD-10J947659N0823244", nil)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestUpdateProduct(t *testing.T) {
	var ps []*Patch
	item := &Patch{
		Op:    "replace",
		Path:  "/description", // reference_id is yourself set when create order
		Value: "Video or model generate service",
	}

	ps = append(ps, item)

	ppRsp, err := client.ProductUpdate(ctx, "PROD-10J947659N0823244", ps)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
}
