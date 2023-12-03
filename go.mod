module github.com/cdvelop/wetest

go 1.20

require github.com/cdvelop/model v0.0.75

require (
	github.com/cdvelop/object v0.0.39
	github.com/cdvelop/strings v0.0.7
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timeserver => ../timeserver

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/logserver => ../logserver
