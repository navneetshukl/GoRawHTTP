package rawHttp

import (
	"context"
	"net"
)

type H map[string]interface{}

type Context struct {
	Conn       net.Conn
	Context    context.Context
	CancelFunc context.CancelFunc

	// request
	Method    string
	Path      string
	Proto     string
	Headers   map[string]string
	Body      []byte
	UrlParams map[string]string

	// response
	Status     int
	RespBody   []byte
	RespHeader map[string]string

	// handling current executing function in case of middleware channing

	CurrentHandler int
	Handlers       []Handler
}

func newContext() *Context {
	return &Context{
		Headers:        make(map[string]string),
		RespBody:       make([]byte, 0),
		RespHeader:     make(map[string]string),
		CurrentHandler: 0,
	}
}

func (ctx *Context) Next() {
	ctx.CurrentHandler = ctx.CurrentHandler + 1
	if ctx.CurrentHandler < len(ctx.Handlers) {
		ctx.Handlers[ctx.CurrentHandler](ctx)
	}
}
