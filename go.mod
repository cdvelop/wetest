module github.com/cdvelop/wetest

go 1.20

require (
	github.com/cdvelop/model v0.0.120
	github.com/cdvelop/object v0.0.76
	github.com/cdvelop/strings v0.0.9
)

require github.com/cdvelop/structs v0.0.1 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/object => ../object
