package paypal

type ProductCreateRep struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *Product       `json:"response,omitempty"`
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	ImageURL    string `json:"image_url"`
	HomeURL     string `json:"home_url"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	Links       []struct {
		Href   string `json:"href"`
		Rel    string `json:"rel"`
		Method string `json:"method"`
	} `json:"links"`
}

type ProductsListRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *ProductsList  `json:"response,omitempty"`
}

type ProductsList struct {
	TotalItems int              `json:"total_items"`
	TotalPages int              `json:"total_pages"`
	Items      []*ProductDetail `json:"products"`
	Links      []*Link          `json:"links,omitempty"`
}

type ProductDetailsRsp struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
	Response      *ProductDetail `json:"response,omitempty"`
}

type ProductDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type,omitempty"`
	Category    string `json:"category,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	HomeURL     string `json:"home_url,omitempty"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time,omitempty"`
	Links       []struct {
		Href   string `json:"href"`
		Rel    string `json:"rel"`
		Method string `json:"method"`
	} `json:"links"`
}
