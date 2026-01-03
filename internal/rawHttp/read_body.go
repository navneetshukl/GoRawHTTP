package rawHttp

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// isStructPtr returns true if v is a struct or a pointer to a struct.
func isStructPtr(v any) bool {
	if v == nil {
		return false
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Struct {
		return true
	}
	if t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Struct {
		return true
	}
	return false
}

// DecodeBodyStruct decodes the request body into a struct (or pointer to struct).
// The target must be a pointer to allow json.Unmarshal to modify it.
// Returns an error if decoding fails or input is invalid.
func (ctx *Context) DecodeBodyStruct(target any) error {
	if !isStructPtr(target) {
		return ctx.writeError(401, "target must be a struct or pointer to struct")
	}

	if err := json.Unmarshal(ctx.Body, target); err != nil {
		return ctx.writeError(401, "invalid JSON payload")
	}

	return nil
}

// DecodeBodyDynamic decodes the request body into a dynamic type (map, slice, etc.).
// Rejects actual struct types even if passed through interface{}.
// Useful for generic JSON handling when schema is unknown.
func (ctx *Context) DecodeBodyDynamic(target any) error {
	if target == nil {
		return ctx.writeError(401, "target cannot be nil")
	}

	// Reject if the underlying concrete type is a struct
	if isStructPtr(reflect.ValueOf(target).Elem().Interface()) {
		return ctx.writeError(401, "target cannot be a struct type")
	}

	if err := json.Unmarshal(ctx.Body, target); err != nil {
		return ctx.writeError(401, "invalid JSON payload")
	}

	return nil
}

// Helper to write error response and return an error (optional extension)
func (ctx *Context) writeError(code int, message string) error {
	ctx.writeResponse(code, message)
	return fmt.Errorf("%d: %s", code, message) // optional: for logging/calling code
}
