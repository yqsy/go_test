package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

type Config struct {
	Proto      string
	ListenAddr string
}

func main() {
	cfg := Config{"tcp", ":3333"}

	l, err := net.Listen(cfg.Proto, cfg.ListenAddr)
	if err != nil {
		fmt.Println("Listen error: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accept errror: ", err.Error())
			os.Exit(1)
		}

		go func(conn net.Conn) {
			fmt.Println("Remote addr: ", conn.RemoteAddr())
			buf := make([]byte, 1024)
			length, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error reading: ", err.Error())
			}
			s := string(buf[:length])
			fmt.Printf("Read: [%v] %v \n", length, s)

			conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))

			conn.Close()
		}(conn)

	}

}
