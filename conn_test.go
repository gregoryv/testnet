package testnet

import (
	"fmt"
	"math/rand"
)

func Example_dial() {
	rand.Seed(0) // to get same IP every time in example

	// dial any server, arguments are there to mimic net.Dial
	client, server := Dial("tcp", "example.com:1234")
	defer client.Close()
	defer server.Close()

	// server echoes message from client once
	go func() {
		buf := make([]byte, 1024)
		n, _ := server.Read(buf)
		_, _ = server.Write(buf[:n])
	}()

	// client sends data
	fmt.Println(server.RemoteAddr(), "->", client.RemoteAddr())
	client.Write([]byte("hello"))

	// print response
	got := make([]byte, 10)
	n, _ := client.Read(got)
	fmt.Print("server responded ", string(got[:n]))
	
	// output:
	// 245.95.248.241:28690 -> example.com:1234
	// server responded hello
}
