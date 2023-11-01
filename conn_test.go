package testnet

import (
	"fmt"
	"math/rand"
	"testing"
)

func Example_dial() {
	rand.Seed(0) // to get same IP every time in example

	// dial any server, arguments are there to mimic net.Dial
	conn, srvconn := Dial("tcp", "example.com:1234")
	defer conn.Close()
	defer srvconn.Close()

	// server echoes message from client once
	go func() {
		buf := make([]byte, 1024)
		n, _ := srvconn.Read(buf)
		_, _ = srvconn.Write(buf[:n])
	}()

	// client sends data
	fmt.Println(srvconn.RemoteAddr(), "->", conn.RemoteAddr())
	conn.Write([]byte("hello"))

	// print response
	got := make([]byte, 10)
	n, _ := conn.Read(got)
	fmt.Print("server responded ", string(got[:n]))

	// output:
	// 245.95.248.241:28690 -> example.com:1234
	// server responded hello
}

func TestClose(t *testing.T) {
	// closing one side should affect the other
	conn, srvconn := Dial("tcp", "example.com:1234")
	srvconn.Close()
	if _, err := conn.Write([]byte("...")); err == nil {
		t.Error("client still open even if server side closed")
	}

	conn, srvconn = Dial("tcp", "example.com:1234")
	conn.Close()
	if _, err := srvconn.Write([]byte("...")); err == nil {
		t.Error("server still open even if client side closed")
	}
}

func TestAddr_Network(t *testing.T) {
	var a Addr
	if v := a.Network(); v != "" {
		t.Error("Network", v)
	}
}
