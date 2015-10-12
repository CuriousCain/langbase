package main

import (
	"fmt"

	"github.com/curiouscain/langbase/connections"
)

func main() {
	fmt.Println("Server starting...")

	ln := connections.StartListening(":8000")

	db := connections.GetConnection("localhost", "test")
	collection := db.C("words")

	for {
		conn := connections.Accept(ln)

		go connections.HandleLiveConnection(conn, collection)
	}
}
