package xhttp

import "encoding/json"

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"

	ResTypeJSON = "json"
	ResTypeXML  = "xml"

	TypeJSON              = "json"
	TypeXML               = "xml"
	TypeFormData          = "form-data"
	TypeMultipartFormData = "multipart-form-data"
)

var (
	_ReqContentTypeMap = map[string]string{
		TypeJSON:              "application/json",
		TypeXML:               "application/xml",
		TypeFormData:          "application/x-www-form-urlencoded",
		TypeMultipartFormData: "multipart/form-data",
	}

	_ResTypeMap = map[string]string{
		ResTypeJSON: "application/json",
		ResTypeXML:  "application/xml",
	}
)

func ConvertToString(v any) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
