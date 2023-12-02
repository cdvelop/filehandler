module github.com/cdvelop/filehandler

go 1.20

require (
	github.com/cdvelop/model v0.0.74
	github.com/cdvelop/object v0.0.38
)

require github.com/cdvelop/timetools v0.0.23 // indirect

require (
	github.com/cdvelop/input v0.0.57
	github.com/cdvelop/maps v0.0.7
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/unixid v0.0.23
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/maps => ../maps

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/strings => ../strings
