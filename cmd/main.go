package main

import (
	"fmt"

	"github.com/curiouscain/langbase/connections"
)

func main() {
	fmt.Println("Server starting...")

	ln := connections.StartListening(":8000")

	for {
		conn := connections.Accept(ln)

		go connections.HandleLiveConnection(conn)
	}
}
