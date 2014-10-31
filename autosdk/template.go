package autosdk

var (
	StructTmpl = `
		package domain

		/* {{.Desc}} */
		type {{.Name}} struct { {{range .Props}} 
			/* {{.Desc}} */
			{{GetHump .Name}} {{$parseType:=parseType .Type .Level}}{{$parseType}} 'json:"{{.Name}}{{jsonName $parseType}}{{if isFloat64 .Type}},string{{end}}"' 
		{{end}} }
	`
	RequestTmpl = `
		package request
		import (
			"net/url"
		)
		{{$CustomerName:=GetRequestResponseStructName .Name "Request"}}
		type {{$CustomerName}} struct{
			values url.Values
		}
		
		func (r *{{$CustomerName}})GetApiMethodName() string{
			return "{{.Name}}"
		}
		func (r *{{$CustomerName}})Set(key,value string){
			if r.values == nil {
				r.values = url.Values{}
			}
			r.values.Set(key, value)
		}
		func (r *{{$CustomerName}})GetValues() url.Values{
			return r.values
		}

		{{range .Params}}
		/* {{.Desc}} */
		func (r *{{$CustomerName}}) Set{{GetHump .Name}}(value string) {
			r.Set("{{.Name}}", value)
		}
		{{end}}
	`
	ResponseTmpl = `
		{{$package:="github.com/ohohco/taobao/api"}}
		package response
		import (
			"{{$package}}"
			{{range .Params}}{{if hasObject .Level}}."{{$package}}/domain"{{end}}{{end}}
		)

		{{$CustomerName:=GetRequestResponseStructName .Name "Response"}}
		{{$Name:=.Name}}
		type {{GetRequestResponseStructName .Name "Response"}} struct { 
			*api.ErrorResponse 'json:"error_response,omitempty"'
			{{GetRequestResponseStructName .Name "Response"}}Result 'json:"{{GetRequestResponseStructJsonName $Name}}"'
		}
		type {{GetRequestResponseStructName .Name "Response"}}Result struct{
			{{range .Params}}/* {{.Desc}} */
			{{GetHump .Name}} {{parseType .Type .Level}} 'json:"{{.Name}}"'
			{{end}} 
		}
	`

	AllInOneTmpl = `
		package api
		import "net/url"
		{{range .Structs}}
			/* {{.Desc}} */
			type {{.Name}} struct { {{range .Props}} 
				/* {{.Desc}} */
				{{GetHump .Name}} {{$parseType:=parseType .Type .Level}}{{$parseType}} 'json:"{{.Name}}{{jsonName $parseType}}{{if isFloat64 .Type}},string{{end}}"' 
			{{end}} }
		{{end}}
		
		{{range .Apis}}
			{{$CustomerName:=GetRequestResponseStructName .Name "Request"}}
			/*{{.Name}}*/
			type {{$CustomerName}} struct{
				values url.Values
			}
			
			func (r *{{$CustomerName}})GetApiMethodName() string{
				return "{{.Name}}"
			}
			func (r *{{$CustomerName}})Set(key,value string){
				if r.values == nil {
					r.values = url.Values{}
				}
				r.values.Set(key, value)
			}
			func (r *{{$CustomerName}})GetValues() url.Values{
				return r.values
			}

			{{range .Request.Params}}
			/* {{.Desc}} */
			func (r *{{$CustomerName}}) Set{{GetHump .Name}}(value string) {
				r.Set("{{.Name}}", value)
			}
			{{end}}

			{{$ResponseCustomerName:=GetRequestResponseStructName .Name "Response"}}
			type {{$ResponseCustomerName}} struct { 
				*ErrorResponse 'json:"error_response,omitempty"'
				{{GetRequestResponseStructName .Name "Response"}}Result 'json:"{{GetRequestResponseStructJsonName .Name}}"'
			}
			type {{$ResponseCustomerName}}Result struct{
				{{range .Response.Params}}/* {{.Desc}} */
				{{GetHump .Name}} {{parseType .Type .Level}} 'json:"{{.Name}}"'
				{{end}} 
			}
		{{end}}


	`
)
