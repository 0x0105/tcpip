package main

import (
	"net"

	"github.com/whuheyiwu/tcpip/server/handlers"
)

func main() {
	listener, err := net.Listen("tcp4", "0.0.0.0:12345")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go handlers.HandleConnection(conn)
	}
}
