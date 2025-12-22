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
