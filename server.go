package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)

	fmt.Println("Received:", string(buffer[:n]))

	conn.Write([]byte("Chunk sent"))
}

func main() {
	listener, _ := net.Listen("tcp", ":8081")
	defer listener.Close()

	fmt.Println("Peer running on port 8081")

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}
