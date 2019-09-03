module github.com/iGoogle-ink/gopay

go 1.12

require (
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/parnurzeal/gorequest v0.2.15
	github.com/pkg/errors v0.8.1 // indirect
	github.com/tjfoc/gmsm v1.0.1
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190710143415-6ec70d6a5542
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190711191110-9a621aea19f8
)
