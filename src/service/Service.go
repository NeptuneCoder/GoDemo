package main

import (
	"net"
	"time"
	"github.com/yanghai23/GoLib/aterr"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	aterr.CheckErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	aterr.CheckErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime)) // don't care about return value
	              // we're finished with this client
}

