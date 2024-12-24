package wechat

import "time"

type InvoiceCardTemplateCreateRsp struct {
	Code     int                        `json:"-"`
	SignInfo *SignInfo                  `json:"-"`
	Response *InvoiceCardTemplateCreate `json:"response,omitempty"`
	Error    string                     `json:"-"`
}

type InvoiceMerchantDevConfigRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *InvoiceMerchantDevConfig `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

type InvoiceMerchantDevConfigQueryRsp struct {
	Code     int                            `json:"-"`
	SignInfo *SignInfo                      `json:"-"`
	Response *InvoiceMerchantDevConfigQuery `json:"response,omitempty"`
	Error    string                         `json:"-"`
}

type InvoiceQueryRsp struct {
	Code     int           `json:"-"`
	SignInfo *SignInfo     `json:"-"`
	Response *InvoiceQuery `json:"response,omitempty"`
	Error    string        `json:"-"`
}

type InvoiceUserTitleUrlRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *InvoiceUserTitleUrl `json:"response,omitempty"`
	Error    string               `json:"-"`
}

type InvoiceUserTitleRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *InvoiceUserTitle `json:"response,omitempty"`
	Error    string            `json:"-"`
}

type InvoiceMerchantBaseInfoRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *InvoiceMerchantBaseInfo `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

type InvoiceMerchantTaxCodesRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *InvoiceMerchantTaxCodes `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

type InvoiceFileUrlRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *InvoiceFileUrl `json:"response,omitempty"`
	Error    string          `json:"-"`
}

type InvoiceUploadFileRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *InvoiceUploadFile `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// =========================================================分割=========================================================

type InvoiceCardTemplateCreate struct {
	CardAppid string `json:"card_appid"`
	CardId    string `json:"card_id"`
}

type InvoiceMerchantDevConfig struct {
	CallbackUrl    string `json:"callback_url"`
	ShowFapiaoCell bool   `json:"show_fapiao_cell"`
}

type InvoiceMerchantDevConfigQuery struct {
	CallbackUrl    string `json:"callback_url"`
	ShowFapiaoCell bool   `json:"show_fapiao_cell"`
}

type InvoiceQuery struct {
	TotalCount        int `json:"total_count"`
	FapiaoInformation []struct {
		FapiaoId   string `json:"fapiao_id"`
		Status     string `json:"status"`
		BlueFapiao struct {
			FapiaoCode   string    `json:"fapiao_code"`
			FapiaoNumber string    `json:"fapiao_number"`
			CheckCode    string    `json:"check_code"`
			Password     string    `json:"password"`
			FapiaoTime   time.Time `json:"fapiao_time"`
		} `json:"blue_fapiao"`
		RedFapiao struct {
			FapiaoCode   string    `json:"fapiao_code"`
			FapiaoNumber string    `json:"fapiao_number"`
			CheckCode    string    `json:"check_code"`
			Password     string    `json:"password"`
			FapiaoTime   time.Time `json:"fapiao_time"`
		} `json:"red_fapiao"`
		CardInformation struct {
			CardAppid  string `json:"card_appid"`
			CardOpenid string `json:"card_openid"`
			CardId     string `json:"card_id"`
			CardCode   string `json:"card_code"`
			CardStatus string `json:"card_status"`
		} `json:"card_information"`
		TotalAmount       int `json:"total_amount"`
		TaxAmount         int `json:"tax_amount"`
		Amount            int `json:"amount"`
		SellerInformation struct {
			Name        string `json:"name"`
			TaxpayerId  string `json:"taxpayer_id"`
			Address     string `json:"address"`
			Telephone   string `json:"telephone"`
			BankName    string `json:"bank_name"`
			BankAccount string `json:"bank_account"`
		} `json:"seller_information"`
		BuyerInformation struct {
			Type        string `json:"type"`
			Name        string `json:"name"`
			TaxpayerId  string `json:"taxpayer_id"`
			Address     string `json:"address"`
			Telephone   string `json:"telephone"`
			BankName    string `json:"bank_name"`
			BankAccount string `json:"bank_account"`
			Phone       string `json:"phone"`
			Email       string `json:"email"`
		} `json:"buyer_information"`
		ExtraInformation struct {
			Payee    string `json:"payee"`
			Reviewer string `json:"reviewer"`
			Drawer   string `json:"drawer"`
		} `json:"extra_information"`
		Items []struct {
			TaxCode       string `json:"tax_code"`
			GoodsName     string `json:"goods_name"`
			Specification string `json:"specification"`
			Unit          string `json:"unit"`
			Quantity      int    `json:"quantity"`
			UnitPrice     int64  `json:"unit_price"`
			Amount        int    `json:"amount"`
			TaxAmount     int    `json:"tax_amount"`
			TotalAmount   int    `json:"total_amount"`
			TaxRate       int    `json:"tax_rate"`
			TaxPreferMark string `json:"tax_prefer_mark"`
			Discount      bool   `json:"discount"`
		} `json:"items"`
		Remark string `json:"remark"`
	} `json:"fapiao_information"`
}

type InvoiceUserTitleUrl struct {
	MiniprogramAppid    string `json:"miniprogram_appid"`
	MiniprogramPath     string `json:"miniprogram_path"`
	MiniprogramUserName string `json:"miniprogram_user_name"`
}

type InvoiceUserTitle struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	TaxpayerId  string `json:"taxpayer_id"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	BankName    string `json:"bank_name"`
	BankAccount string `json:"bank_account"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

type InvoiceMerchantBaseInfo struct {
	SellerInformation struct {
		Name        string `json:"name"`
		TaxpayerId  string `json:"taxpayer_id"`
		Address     string `json:"address"`
		Telephone   string `json:"telephone"`
		BankName    string `json:"bank_name"`
		BankAccount string `json:"bank_account"`
	} `json:"seller_information"`
	ExtraInformation struct {
		Payee    string `json:"payee"`
		Reviewer string `json:"reviewer"`
		Drawer   string `json:"drawer"`
	} `json:"extra_information"`
}

type InvoiceMerchantTaxCodes struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Links  struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
		Self string `json:"self"`
	} `json:"links"`
	Data []struct {
		GoodsName     string `json:"goods_name"`
		GoodsId       int    `json:"goods_id"`
		GoodsCategory string `json:"goods_category"`
		TaxCode       string `json:"tax_code"`
		TaxRate       int    `json:"tax_rate"`
		TaxPreferMark string `json:"tax_prefer_mark"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type InvoiceFileUrl struct {
	FapiaoDownloadInfoList []struct {
		FapiaoId    string `json:"fapiao_id"`
		DownloadUrl string `json:"download_url"`
		Status      string `json:"status"`
	} `json:"fapiao_download_info_list"`
}

type InvoiceUploadFile struct {
	FapiaoMediaId string `json:"fapiao_media_id"`
}
