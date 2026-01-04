package rawHttp

import (
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
func TestDecodeBody_StructSuccess(t *testing.T) {
	ctx := &Context{
		Body: []byte(`{"name":"navneet","age":20}`),
	}

	var u User
	err := ctx.DecodeBody(&u)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if u.Name != "navneet" || u.Age != 20 {
		t.Fatalf("unexpected struct value: %+v", u)
	}
}
func TestDecodeBody_MapSuccess(t *testing.T) {
	ctx := &Context{
		Body: []byte(`{"name":"navneet","age":20}`),
	}

	var m map[string]interface{}
	err := ctx.DecodeBody(&m)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if m["name"] != "navneet" {
		t.Fatalf("expected name navneet, got %v", m["name"])
	}

	if _, ok := m["age"].(float64); !ok {
		t.Fatalf("expected age to be float64, got %T", m["age"])
	}
}
func TestDecodeBody_InterfaceSuccess(t *testing.T) {
	ctx := &Context{
		Body: []byte(`{"name":"navneet","age":20}`),
	}

	var data interface{}
	err := ctx.DecodeBody(&data)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	m, ok := data.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map[string]interface{}, got %T", data)
	}

	if m["name"] != "navneet" {
		t.Fatalf("unexpected value: %v", m["name"])
	}
}
func TestDecodeBody_NonPointerTarget(t *testing.T) {
	ctx := &Context{
		Body: []byte(`{"name":"navneet"}`),
	}

	var u User
	err := ctx.DecodeBody(u)

	if err == nil {
		t.Fatalf("expected error for non-pointer target")
	}
}
func TestDecodeBody_NilTarget(t *testing.T) {
	ctx := &Context{
		Body: []byte(`{"name":"navneet"}`),
	}

	err := ctx.DecodeBody(nil)

	if err == nil {
		t.Fatalf("expected error for nil target")
	}
}
func TestDecodeBody_EmptyBody(t *testing.T) {
	ctx := &Context{
		Body: []byte{},
	}

	var u User
	err := ctx.DecodeBody(&u)

	if err == nil {
		t.Fatalf("expected error for empty body")
	}
}
func TestDecodeBody_InvalidJSON(t *testing.T) {
	ctx := &Context{
		Body: []byte(`{"name":`),
	}

	var u User
	err := ctx.DecodeBody(&u)

	if err == nil {
		t.Fatalf("expected error for invalid JSON")
	}
}
func TestDecodeBody_UnknownField(t *testing.T) {
	ctx := &Context{
		Body: []byte(`{"name":"navneet","age":20,"email":"x@test.com"}`),
	}

	var u User
	err := ctx.DecodeBody(&u)

	if err == nil {
		t.Fatalf("expected error for unknown field")
	}
}
