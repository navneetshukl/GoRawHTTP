package server

import (
	"bufio"
	"fmt"
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

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}

		fmt.Println("Received from client : ",data)
		conn.Write([]byte(data))
	}
}
