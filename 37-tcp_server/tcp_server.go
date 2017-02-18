package main

import (
	"fmt"
	"net"
	"os"
)

type Config struct {
	Listenaddr string
	Net        string
}

func main() {
	var cfg = Config{"0.0.0.0:3333", "tcp"}

	l, err := net.Listen(cfg.Net, cfg.Listenaddr)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on: " + cfg.Listenaddr)

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go HandleRequest(conn)
	}
}

func HandleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}

	conn.Write([]byte("Message received."))

	conn.Close()
}
