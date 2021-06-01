package commands

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// connect cmd
var _ cmder = (*connectCmd)(nil)

type connectCmd struct {
	serverAddress string
	peerName      string
	selfName      string
	*baseCmd
}

func newConnectCmd() *connectCmd {
	c := &connectCmd{}

	c.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "connect",
		Short: "connect to peer client.",
		Run: func(cmd *cobra.Command, args []string) {
			handleMessage(c.serverAddress)
		},
	})

	c.cmd.Flags().StringVarP(&c.serverAddress, "serverAddress", "", "127.0.0.1:12345", "")
	c.cmd.Flags().StringVarP(&c.peerName, "peerName", "", "", "")
	return c
}

var reader io.ReadCloser

func init() {
	reader = os.Stdin
}

func handleMessage(serverAddress string) {
	input := make(chan string, 1)
	conn, err := net.Dial("tcp4", serverAddress)
	if err != nil {
		fmt.Println("Dial server error: ", err.Error())
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	ctx, cancel := context.WithCancel(context.Background())
	go readLines(input, cancel)
	go sendMessage(ctx, conn, input)
	receiveMessage(ctx, conn)
	fmt.Println("Bye, have a nice day")
}

func readLines(in chan string, cancel context.CancelFunc) {
	reader := bufio.NewReader(reader)
	for {
		s, err := reader.ReadString('\n')
		// check for errors
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF in read string", err)
			} else {
				fmt.Println("Error in read string", err)
			}
			close(in)
			cancel()
			return
		}
		in <- strings.TrimSpace(s)
	}
}

func sendMessage(ctx context.Context, conn net.Conn, input chan string) {
	for {
		select {
		case in := <-input:
			_, err := conn.Write([]byte(in))
			if err != nil {
				fmt.Println("Write to server error: ", err.Error())
			}
		case <-ctx.Done():
			return
		}
	}
}

func receiveMessage(ctx context.Context, conn net.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			var buf [128]byte
			n, err := conn.Read(buf[:])
			if err != nil {
				if err == io.EOF {
					fmt.Println("Read from tcp server EOF")
					return
				} else {
					fmt.Println("Read from tcp server with unexpected error:", err)
				}
				return
			}
			ret := []byte{}
			for _, b := range buf[:n] {
				if b == 0x00 {
					break
				}
				ret = append(ret, b)
			}
			if string(ret) == "EXIT" {
				return
			}
			if len(ret) > 0 {
				fmt.Printf("Recived from server, data: %s\n", string(ret))
			}
		}
	}
}
