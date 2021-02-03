package xhttp

type RequestType string

const (
	POST                       = "POST"
	GET                        = "GET"
	TypeJSON       RequestType = "json"
	TypeXML        RequestType = "xml"
	TypeUrlencoded RequestType = "urlencoded"
	TypeForm       RequestType = "form"
	TypeFormData   RequestType = "form-data"
)

var types = map[RequestType]string{
	TypeJSON:       "application/json",
	TypeXML:        "application/xml",
	TypeForm:       "application/x-www-form-urlencoded",
	TypeFormData:   "application/x-www-form-urlencoded",
	TypeUrlencoded: "application/x-www-form-urlencoded",
}
