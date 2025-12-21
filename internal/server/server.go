package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func Listen() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error in starting the server ", err)
		return
	}
	fmt.Println("Starting the Server on port 8080")
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Error in accepting the connection ", err)
			return
		}
		defer conn.Close()
		reader := bufio.NewReader(conn)

		for {
			n, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error in Reading ", err)
				return
			}
			fmt.Println("Read Length is ", n)
		}

	}
}
