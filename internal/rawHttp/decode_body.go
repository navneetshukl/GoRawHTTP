package rawHttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func isValidDecodeTarget(v interface{}) bool {
	if v == nil {
		return false
	}
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Pointer {
		return false
	}

	elem := t.Elem().Kind()
	return elem == reflect.Struct ||
		elem == reflect.Map ||
		elem == reflect.Interface
}

// DecodeBody function will decode the body into struct or interface
func (ctx *Context) DecodeBody(target interface{}) error {
	if !isValidDecodeTarget(target) {
		return fmt.Errorf("DecodeBody expects pointer to struct, map, or interface")
	}

	if len(ctx.Body) == 0 {
		return fmt.Errorf("request body is empty")
	}

	decoder := json.NewDecoder(bytes.NewReader(ctx.Body))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("invalid JSON payload: %w", err)
	}

	// prevent trailing garbage
	if decoder.More() {
		return fmt.Errorf("unexpected data after JSON body")
	}

	return nil
}
