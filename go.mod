module github.com/cdvelop/wetest

go 1.20

require github.com/cdvelop/model v0.0.77

require (
	github.com/cdvelop/object v0.0.42
	github.com/cdvelop/strings v0.0.8
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/logclient => ../logclient

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/strings => ../strings
