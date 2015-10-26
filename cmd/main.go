package main

import (
	"fmt"
	"github.com/curiouscain/langbase/connections"
	"sync"
)

func main() {
	fmt.Println("Server starting...")

	ln := connections.StartListening(":8000")
	connections.StartListeningForWebSocket(":8001")

	session := connections.GetConnection("localhost")
	defer session.Close()

	collection := session.DB("test").C("words")

	var wg sync.WaitGroup

	for {
		conn := connections.Accept(ln)

		wg.Add(1)
		go connections.HandleLiveConnection(conn, wg, collection)
	}

	wg.Wait()
}
