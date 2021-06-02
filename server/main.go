package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp4", "0.0.0.0:12345")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	cm := NewConnMngr()
	msg := make(chan Msg)
	defer close(msg)
	go func() {
		for {
			conn, _ := listener.Accept()
			cm.Add(conn.RemoteAddr().String(), conn)
			go receiveMsg(conn, msg)
		}
	}()
	defer cm.Close()
	for {
		select {
		case m := <-msg:
			all := cm.All()
			for _, c := range all {
				_, err = (c).Write([]byte(m.Message))
				if err == io.EOF {
					cm.Remove(m.From)
					_ = c.Close()
					break
				}
				fmt.Printf("Send msg %s to %s.\n", m.Message, m.From)
			}
		}
	}

}

func receiveMsg(conn net.Conn, msg chan Msg) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from: " + remoteAddr)

	buf := make([]byte, 1024)
	for {
		reqLen, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Disconnected from ", remoteAddr)
			} else {
				fmt.Println("Unexpected error:", err.Error())
			}
			break
		}
		str := string(buf[:reqLen])
		fmt.Printf("len: %d, recv: %s\n", reqLen, str)
		msg <- Msg{
			Message: str,
			From:    remoteAddr,
			To:      "",
		}
	}
}
