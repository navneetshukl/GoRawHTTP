package rawHttp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func (r *Router) Run() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		go r.handleConnection(conn)
	}
}

// handleConnection handle every request data
func (r *Router) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	parts := strings.Split(strings.TrimSpace(line), " ")
	if len(parts) < 3 {
		fmt.Println("Invalid request line")
		return
	}

	method := parts[0]
	path := parts[1]
	protocol := parts[2]
	headers := map[string]string{}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		if line == "\r\n" {
			break
		}
		header := strings.Split(line, ":")
		headers[header[0]] = strings.ReplaceAll(header[1], " ", "")
	}
	var bodySize int = 0
	bodyLength, ok := headers["Content-Length"]
	if ok {
		bodySize, err = strconv.Atoi(strings.Split(bodyLength, "\r\n")[0])
		if err != nil {
			fmt.Println("Error in converting string to int ", err)
			return
		}
	}
	body := make([]byte, bodySize)
	totalSize := 0

	for totalSize < bodySize {
		n, err := reader.Read(body[totalSize:])
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		totalSize += n
	}

	ctx := &Context{
		Conn:    conn,
		Method:  method,
		Path:    path,
		Proto:   protocol,
		Headers: headers,
		Body:    body,
		Status:  200,
		RespHeader: map[string]string{
			"Content-Type": "text/plain",
		},
		CurrentHandler: 0,
	}

	r.executeHandler(ctx)

}

// executeHandler will execute the required route and middleware
func (r *Router) executeHandler(ctx *Context) {

	for _, route := range r.routes {
		if route.method == ctx.Method && route.path == ctx.Path {

			// appending the required global middleware and handler middleware to execute
			allhandlers := make([]Handler, 0, len(r.middleware)+len(route.handler))
			allhandlers = append(allhandlers, r.middleware...)
			allhandlers = append(allhandlers, route.handler...)
			ctx.Handlers = allhandlers
			log.Println("Total Handler is ", len(ctx.Handlers))
			ctx.CurrentHandler = -1
			ctx.Next()
			return
		}

	}
	ctx.String(401, "HANDLER ERROR : handler does not exist")

}
