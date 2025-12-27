package rawHttp

import (
	"encoding/json"
	"reflect"
)

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

// checkStruct check if the provided interface is struct or not
func checkStruct(v interface{}) bool {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Struct {
		return true
	}

	if t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Struct {
		return true
	}

	return false
}

// DecodeBodyStruct decodes the request body to given struct type
func (ctx *Context) DecodeBodyStruct(decodeType interface{}) {
	if !checkStruct(decodeType) {
		ctx.writeResponse(401, "Not a struct")
		return

	}

	err := json.Unmarshal(ctx.Body, &decodeType)
	if err != nil {
		ctx.writeResponse(401, "Invalid struct to decode")
		return
	}

}
