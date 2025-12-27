package rawHttp

import (
	"encoding/json"
	"fmt"
)

func (ctx *Context) SetStatus(status int) {
	ctx.Status = status
}

func (ctx *Context) SetBody(body []byte) {
	ctx.Body = body
}

func (ctx *Context) AddHeader(key, value string) {
	ctx.RespHeader[key] = value
}

func (ctx *Context) writeResponse(status int, body string) {
	statusText := "OK"
	if status == 404 {
		statusText = "Not Found"
	}

	response := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\n"+
			"Content-Type: text/plain\r\n"+
			"Content-Length: %d\r\n"+
			"Connection: close\r\n"+
			"\r\n%s",
		status,
		statusText,
		len(body),
		body,
	)
	ctx.Conn.Write([]byte(response))
}

func (ctx *Context) String(status int, data string) {
	ctx.writeResponse(status, data)
}

func (ctx *Context) JSON(status int, data interface{}) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		ctx.writeResponse(500, "Internal Server Error")
		return
	}
	ctx.writeResponse(status, string(jsonBytes))
}
