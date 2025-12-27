package rawHttp

import (
	"net"
)

type H map[string]interface{}

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

func newContext()*Context{
	return &Context{
		Headers: make(map[string]string),
		RespBody: make([]byte, 0),
		RespHeader: make(map[string]string),
	}
}
