module github.com/cdvelop/wetest

go 1.20

require github.com/cdvelop/model v0.0.72

require (
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/timetools v0.0.21 // indirect
)

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/logserver v0.0.6
	github.com/cdvelop/object v0.0.35
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/timeserver v0.0.20
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timeserver => ../timeserver

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/logserver => ../logserver
