package main_test

import (
	"fmt"
	"net"
	"testing"
)

func handleConnection(conn net.Conn) {
	fmt.Println(conn.RemoteAddr())
	conn.Close()
}

func accept(ln net.Listener) {
	for {
		fmt.Println("err")
		conn, err := ln.Accept()
		fmt.Println(err)
		if err != nil {
			go handleConnection(conn)
		}
	}
}

func TestToto(t *testing.T) {
	ln, err := net.Listen("tcp", ":8080")
	t.Log(err)
	go accept(ln)

	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn.Write([]byte("hello"))
	t.Log(err)

	ln.Close()
}
