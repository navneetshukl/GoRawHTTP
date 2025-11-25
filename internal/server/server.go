package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func Listen() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}
	log.Println("Server Started on Port 8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	var allData []byte

	for {
		buf := make([]byte, 4096)
		n, err := reader.Read(buf)

		if n > 0 {
			allData = append(allData, buf[:n]...)
		}
		if n<4096{
			break
		}

		if err != nil {
			if err != io.EOF {
				log.Println("Read error:", err)
			}
			break
		}
	}

	fmt.Println("----- RAW REQUEST BEGIN -----")
	fmt.Println(string(allData))
	fmt.Println("----- RAW REQUEST END -----")

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: 5\r\n" +
		"\r\n" +
		"Hello"

	conn.Write([]byte(response))
}
