package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type Telnet struct {
	port     string
	ip       string
	server   string
	deadline int
	Writer   *bufio.Writer
	Reader   *bufio.Reader
	Conn     net.Conn
}

//NewTelnet returns new Telnet client
func NewTelnet(port, ip string, deadline int) *Telnet {
	return &Telnet{
		port:     port,
		ip:       ip,
		deadline: deadline,
		server:   fmt.Sprintf("%s:%s", ip, port),
		Writer:   new(bufio.Writer),
	}
}

func (t *Telnet) Dial() (err error) {
	log.Printf("starting Dial connection to %s\n", t.server)
	t.Conn, err = net.Dial("tcp", t.server)
	if err != nil {
		return err
	}

	t.Reader = bufio.NewReader(t.Conn)

	err = t.Conn.SetDeadline(time.Now().Add(time.Duration(t.deadline) * time.Second))
	if err != nil {
		return err
	}
	return nil
}
