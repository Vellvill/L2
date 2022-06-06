package main

import (
	"fmt"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {
	srv, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = srv.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		sigConn, err := srv.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func(sigConn net.Conn) {
			for {
				input := make([]byte, 0)

				num, err := sigConn.Read(input)
				if err != nil || num == 0 {
					fmt.Println(err)
				}

				num, err = sigConn.Write(input)
				if err != nil || num == 0 {
					fmt.Println(err)
				}
			}
		}(sigConn)
	}
}
