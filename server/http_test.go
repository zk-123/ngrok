package server

import (
	"fmt"
	"net"
	"testing"
)

func TestServerListen(t *testing.T) {
	if listen, err := net.Listen("tcp", ":889"); err != nil {
		panic(err)
	} else {
		for {
			c, _ := listen.Accept()
			go func(conn net.Conn) {
				defer c.Close()

				rc := make([]byte, 1024*1)
				for {
					_, err := conn.Read(rc)
					if err != nil {
						fmt.Printf("%v\n", err)
						break
					}
					fmt.Printf("message:%v\n", string(rc[:len(rc)]))
				}
				c.Write([]byte("hello world\n\r"))
			}(c)
		}
	}
}

func TestClientConnect(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:889")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("hello world"))
}
