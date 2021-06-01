package handlers

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from: " + remoteAddr)

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	for {
		// Read the incoming connection into the buffer.
		reqLen, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Disconnected from ", remoteAddr)
				break
			} else {
				fmt.Println("Unexpected error:", err.Error())
				break
			}
		}
		_, _ = conn.Write([]byte(strings.ToUpper(string(buf))))
		fmt.Printf("len: %d, recv: %s\n", reqLen, string(buf[:reqLen]))
	}
}