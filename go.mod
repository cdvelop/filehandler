module github.com/cdvelop/filehandler

go 1.20

require (
	github.com/cdvelop/model v0.0.67
	github.com/cdvelop/object v0.0.25
)

require github.com/cdvelop/strings v0.0.3 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/object => ../object
