/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

package main

import (
	"L2/dev10/telnet"
	"flag"
)

var (
	timeout int
	port    string
	host    string
)

//init ...
func init() {
	flag.IntVar(&timeout, "timeout", 15, "timeout for connect")
	flag.StringVar(&port, "port", "", "port for connecting to")
	flag.StringVar(&host, "host", "", "binded ip for connecting to")
}

type Telnet struct {
	client telnet.Telnet
}

func newTelnet(client telnet.Telnet) *Telnet {
	return &Telnet{client: client}
}

func main() {
	flag.Parse()
	myTelnetClient := telnet.NewTelnetClient(host, port, timeout)
	newTelnet(myTelnetClient)

}
