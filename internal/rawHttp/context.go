package rawHttp

import "net"

type Context struct {
	Conn net.Conn

	// request
	Method  string
	Path    string
	Proto   string
	Headers map[string]string
	Body    []byte

	// response
	Status     int
	RespBody   []byte
	RespHeader map[string]string
}

func (ctx *Context) GetMethod() string {
	if ctx == nil {
		return "No Method Present"
	} else {
		return ctx.Method
	}
}

func (ctx *Context) GetPath() string {
	if ctx == nil {
		return "No Path Present"
	} else {
		return ctx.Path
	}
}

func(ctx *Context)GetHeader(key string)(string){
	if _,ok:=ctx.Headers[key];!ok{
		return "Header Not Present"
	}
	return ctx.Headers[key]
}

func(ctx *Context)GetAllHeaders()map[string]string{
	return ctx.Headers
}

func(ctx *Context)SetStatus(status int){
	ctx.Status=status
}

func(ctx *Context)SetBody(body []byte){
	ctx.Body=body
}

func(ctx *Context)AddHeader(key,value string){
	ctx.RespHeader[key]=value
}
