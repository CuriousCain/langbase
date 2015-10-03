package connections

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/curiouscain/langbase/fault"
)

func StartListening(port string) net.Listener {
	ln, err := net.Listen("tcp", port)
	fault.Handle(err)

	return ln
}

func Accept(ln net.Listener) net.Conn {
	conn, err := ln.Accept()
	fault.Handle(err)

	return conn
}

func HandleLiveConnection(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		fault.Handle(err)

		fmt.Println(msg)
		newMsg := strings.ToUpper(msg)
		conn.Write([]byte(newMsg + "\n"))
	}
}
