package rawHttp

func (ctx *Context) GetMethod() string {
	if ctx == nil {
		return "No Method Present"
	} else {
		return ctx.Method
	}
}

func (ctx *Context) GetPath() string {
	if ctx == nil || ctx.Path == "" {
		return "No Path Present"
	} else {
		return ctx.Path
	}
}

func (ctx *Context) GetHeader(key string) string {
	if _, ok := ctx.Headers[key]; !ok {
		return "Header Not Present"
	}
	return ctx.Headers[key]
}

func (ctx *Context) GetAllHeaders() map[string]string {
	return ctx.Headers
}
