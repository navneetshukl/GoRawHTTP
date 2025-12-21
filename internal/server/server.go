package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Listen() {
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	arr := strings.Split(line, " ")

	for idx := range arr {
		fmt.Println(idx, " ", arr[idx])
	}
	headers := map[string]interface{}{}

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
	fmt.Println("Headers is ", headers)

	bodyLength := headers["Content-Length"]
	bodySize, err := strconv.Atoi(strings.Split(bodyLength.(string), "\r\n")[0])
	if err != nil {
		fmt.Println("Error in converting string to int ", err)
		return
	}
	fmt.Println("Body size is ", bodySize)
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

	fmt.Println("Body is", string(body))

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: 2\r\n" +
		"Content-Type: text/plain\r\n" +
		"Connection: close\r\n" +
		"\r\n" +
		"OK"

	conn.Write([]byte(response))
}
