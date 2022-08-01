package main

import (
	"fmt"
	"net"
)

const mtu = 1 << 16

func generate(n int) []byte {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = byte(i)
	}
	return res
}

func main() {
	conn, err := net.Dial("udp", ":8001")
	fmt.Println(err)

	for i := 0; i < mtu; i++ {
		_, err := conn.Write(generate(mtu))
		if err != nil {
			fmt.Println(i, err)
			return
		}
	}
}
