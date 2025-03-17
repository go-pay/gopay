// Package paypal
// Copyright 2025 Giga Inc. All Rights Reserved
// Author: Stuart <stuartjing@sina.com>
// Date: 2025/3/17
/*
   DESCRIPTION: xx
   使用文档：xx
*/
package paypal

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"testing"
)

func TestClient_InvoiceCreate(t *testing.T) {

	type args struct {
		header string
		body   gopay.BodyMap
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "001-发票创建-with header",
			args: args{header: "return=representation", body: gopay.BodyMap{
				"detail": map[string]any{
					"invoice_number": "INV-202503141441011",
					"reference":      "",
					"invoice_date":   "2025-03-17",
					"currency_code":  "USD",
					"note":           "发票收件人须知。也会显示在发票通知邮件中。",
					"term":           "",
					"memo":           "",
				},
				"invoicer": map[string]any{
					"name": map[string]any{
						"given_name": "san",
						"surname":    "zhang",
					},
					"address": map[string]any{
						"address_line_1": "1234 Main Street",
						"admin_area_1":   "CA",
						"admin_area_2":   "Anytown",
						"postal_code":    "98765",
						"country_code":   "US",
					},
					"email_address": "xxxx@business.example.com",
					"logo_url":      "xxxx",
				},
				"items": []map[string]any{map[string]any{
					"name": "product vip member 1",
					"unit_amount": map[string]any{
						"currency_code": "USD",
						"value":         "10.00",
					},
					"quantity": "1",
				}},
				"primary_recipients": []map[string]any{{
					"billing_info": map[string]any{
						"name": map[string]any{
							"given_name": "san",
							"surname":    "zhang",
						},
						"address": map[string]any{
							"address_line_1": "1234 Main Street",
							"admin_area_1":   "CA",
							"admin_area_2":   "Anytown",
							"postal_code":    "98765",
							"country_code":   "US",
						},
						"email_address":         "vvvv@personal.example.com",
						"additional_info_value": "add-info",
					},
					"shipping_info": map[string]any{
						"name": map[string]any{
							"given_name": "san",
							"surname":    "zhang",
						},
						"address": map[string]any{
							"address_line_1": "1234 Main Street",
							"address_line_2": "",
							"address_line_3": "",
							"admin_area_1":   "CA",
							"admin_area_2":   "Anytown",
							"admin_area_3":   "",
							"admin_area_4":   "",
							"postal_code":    "98765",
							"country_code":   "US",
						},
					},
				}},
			},
			},
		},
		{name: "002-发票创建-without header",
			args: args{header: "", body: gopay.BodyMap{
				"detail": map[string]any{
					"invoice_number": "INV-202503141441020",
					"reference":      "",
					"invoice_date":   "2025-03-17",
					"currency_code":  "USD",
					"note":           "发票收件人须知。也会显示在发票通知邮件中。",
					"term":           "",
					"memo":           "",
				},
				"invoicer": map[string]any{
					"name": map[string]any{
						"given_name": "san",
						"surname":    "zhang",
					},
					"address": map[string]any{
						"address_line_1": "1234 Main Street",
						"admin_area_1":   "CA",
						"admin_area_2":   "Anytown",
						"postal_code":    "98765",
						"country_code":   "US",
					},
					"email_address": "xxxx@business.example.com",
					"logo_url":      "vvv",
				},
				"items": []map[string]any{map[string]any{
					"name": "product vip member 1",
					"unit_amount": map[string]any{
						"currency_code": "USD",
						"value":         "10.00",
					},
					"quantity": "1",
				}},
				"primary_recipients": []map[string]any{{
					"billing_info": map[string]any{
						"name": map[string]any{
							"given_name": "san",
							"surname":    "zhang",
						},
						"address": map[string]any{
							"address_line_1": "1234 Main Street",
							"admin_area_1":   "CA",
							"admin_area_2":   "Anytown",
							"postal_code":    "98765",
							"country_code":   "US",
						},
						"email_address":         "vvvv@personal.example.com",
						"additional_info_value": "add-info",
					},
					"shipping_info": map[string]any{
						"name": map[string]any{
							"given_name": "san",
							"surname":    "zhang",
						},
						"address": map[string]any{
							"address_line_1": "1234 Main Street",
							"address_line_2": "",
							"address_line_3": "",
							"admin_area_1":   "CA",
							"admin_area_2":   "Anytown",
							"admin_area_3":   "",
							"admin_area_4":   "",
							"postal_code":    "98765",
							"country_code":   "US",
						},
					},
				}},
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newCtx := context.Background()
			if len(tt.args.header) > 0 {
				newCtx = context.WithValue(newCtx, PreferHeaderKey, tt.args.header)
			}
			got, err := client.InvoiceCreate(newCtx, tt.args.body)
			fmt.Println(got, err)
		})
	}
}
