package testnet

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
)

// NewMemConn returns client and server side connections. The arguments
// are only used to fullfill the net.Addr
func Dial(network, address string) (clientRW *Conn, serverRW *Conn) {
	fromServer, toClient := io.Pipe()
	fromClient, toServer := io.Pipe()

	clientRW = &Conn{
		Reader:      fromServer,
		WriteCloser: toServer,
		remote:      NewAddr(network, address),
	}

	serverRW = &Conn{
		Reader:      fromClient,
		WriteCloser: toClient,
		remote:      RandAddr(network),
	}
	return
}

type Conn struct {
	io.Reader
	io.WriteCloser

	remote *Addr
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.remote
}

func NewAddr(network, address string) *Addr {
	return &Addr{
		network: network,
		address: address,
	}
}

// RandAddr resturns a new Addr with a random IP and port 1024-2^16
func RandAddr(network string) *Addr {
	return &Addr{
		network: network,
		// port range 1024-2^16
		address: fmt.Sprintf("%s:%v", RandIPv4(), 1024+rand.Int31n(64512)),
	}
}

// Addr implements net.Addr
type Addr struct {
	network string
	address string
}

func (a *Addr) Network() string { return a.network }
func (a *Addr) String() string  { return a.address }

func RandIPv4() string {
	buf := make([]byte, 4)
	ip := rand.Uint32()
	binary.LittleEndian.PutUint32(buf, ip)
	return net.IP(buf).String()
}
