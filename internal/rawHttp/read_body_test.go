package rawHttp

import (
	"encoding/json"
	"reflect"
	"testing"
)

// mockContext embeds Context by value and captures writeResponse calls
type mockContext struct {
	Context
	statusCode int
	body       []byte
	written    bool
}

func (m *mockContext) writeResponse(code int, message string) {
	m.statusCode = code
	m.body = []byte(message)
	m.written = true
}

func TestDecodeBodyStruct_Success(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	payload := User{Name: "Alice", Age: 30}
	bodyBytes, _ := json.Marshal(payload)

	ctx := &mockContext{
		Context: Context{
			Body:       bodyBytes,
			Headers:    make(map[string]string),
			RespHeader: make(map[string]string),
		},
	}

	var result User
	ctx.DecodeBodyStruct(&result)

	if ctx.written {
		t.Fatal("Success case should not write any response")
	}

	if !reflect.DeepEqual(result, payload) {
		t.Fatalf("Decoded struct mismatch: got %+v, want %+v", result, payload)
	}
}

func TestDecodeBodyStruct_InvalidJSON(t *testing.T) {
	ctx := &mockContext{
		Context: Context{
			Body:       []byte(`{invalid json`),
			Headers:    make(map[string]string),
			RespHeader: make(map[string]string),
		},
	}

	var target struct{ Name string }
	ctx.DecodeBodyStruct(&target)

	if !ctx.written {
		t.Fatal("Invalid JSON should trigger response")
	}
	if ctx.statusCode != 401 {
		t.Fatalf("Expected status 401, got %d", ctx.statusCode)
	}
	if string(ctx.body) != "Invalid struct to decode" {
		t.Fatalf("Expected 'Invalid struct to decode', got %q", string(ctx.body))
	}
}

func TestDecodeBodyStruct_NotAStruct(t *testing.T) {
	ctx := &mockContext{
		Context: Context{
			Body:       []byte(`{}`),
			Headers:    make(map[string]string),
			RespHeader: make(map[string]string),
		},
	}

	var target string // not a struct
	ctx.DecodeBodyStruct(&target)

	if !ctx.written {
		t.Fatal("Non-struct target should trigger response")
	}
	if ctx.statusCode != 401 {
		t.Fatalf("Expected 401, got %d", ctx.statusCode)
	}
	if string(ctx.body) != "Not a struct" {
		t.Fatalf("Expected 'Not a struct', got %q", string(ctx.body))
	}
}

func TestDecodeBodyStruct_PlainStructWithoutPointer(t *testing.T) {
	ctx := &mockContext{
		Context: Context{
			Body:       []byte(`{"name":"Bob"}`),
			Headers:    make(map[string]string),
			RespHeader: make(map[string]string),
		},
	}

	type Person struct {
		Name string `json:"name"`
	}
	var target Person
	ctx.DecodeBodyStruct(target) // pass by value

	// json.Unmarshal requires a pointer, so this should fail
	if !ctx.written {
		t.Fatal("Passing struct by value should fail and write response")
	}
	if ctx.statusCode != 401 {
		t.Fatalf("Expected 401, got %d", ctx.statusCode)
	}
	if string(ctx.body) != "Invalid struct to decode" {
		t.Fatalf("Expected 'Invalid struct to decode', got %q", string(ctx.body))
	}
}

func TestDecodeBodyInterface_Success(t *testing.T) {
	bodyBytes, _ := json.Marshal(map[string]interface{}{
		"name": "Charlie",
		"age":  25,
	})

	ctx := &mockContext{
		Context: Context{
			Body:       bodyBytes,
			Headers:    make(map[string]string),
			RespHeader: make(map[string]string),
		},
	}

	var result interface{}
	ctx.DecodeBodyDynamic(&result)

	if ctx.written {
		t.Fatal("Success case should not write response")
	}

	expected := map[string]interface{}{
		"name": "Charlie",
		"age":  float64(25),
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Decoded interface mismatch: got %+v, want %+v", result, expected)
	}
}

func TestDecodeBodyInterface_InvalidJSON(t *testing.T) {
	ctx := &mockContext{
		Context: Context{
			Body:       []byte(`{bad json`),
			Headers:    make(map[string]string),
			RespHeader: make(map[string]string),
		},
	}

	var target interface{}
	ctx.DecodeBodyDynamic(&target)

	if !ctx.written {
		t.Fatal("Invalid JSON should trigger response")
	}
	if ctx.statusCode != 401 {
		t.Fatalf("Expected 401, got %d", ctx.statusCode)
	}
	if string(ctx.body) != "Invalid struct to decode" {
		t.Fatalf("Expected 'Invalid struct to decode', got %q", string(ctx.body))
	}
}

func TestDecodeBodyInterface_RejectsStruct(t *testing.T) {
	ctx := &mockContext{
		Context: Context{
			Body:       []byte(`{}`),
			Headers:    make(map[string]string),
			RespHeader: make(map[string]string),
		},
	}

	type Dummy struct{}
	var bad interface{} = &Dummy{} // pointer to struct

	ctx.DecodeBodyDynamic(bad)

	if !ctx.written {
		t.Fatal("Passing struct (even via interface) should be rejected")
	}
	if ctx.statusCode != 401 {
		t.Fatalf("Expected 401, got %d", ctx.statusCode)
	}
	if string(ctx.body) != "Not a interface" {
		t.Fatalf("Expected 'Not a interface', got %q", string(ctx.body))
	}
}

func TestDecodeBodyInterface_AcceptsNonStructTypes(t *testing.T) {
	testCases := []interface{}{
		map[string]interface{}{},
		[]interface{}{},
		"string",
		123,
		true,
		nil,
	}

	validJSON := []byte(`null`) // safe for all

	for _, tc := range testCases {
		ctx := &mockContext{
			Context: Context{
				Body:       validJSON,
				Headers:    make(map[string]string),
				RespHeader: make(map[string]string),
			},
		}

		// Create pointer to the value
		v := reflect.New(reflect.TypeOf(tc))
		ptr := v.Interface()

		ctx.DecodeBodyDynamic(ptr)

		if ctx.written {
			t.Fatalf("Should accept type %T, but wrote response: %s", tc, ctx.body)
		}
	}
}