package connections

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/curiouscain/langbase/data"
	"github.com/curiouscain/langbase/fault"

	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func StartListening(port string) net.Listener {
	ln, err := net.Listen("tcp", port)
	fault.Handle(err)

	return ln
}

func StartListeningForWebSocket(port string, collection *mgo.Collection) {
	http.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		//var in []byte
		var in string

		err := websocket.Message.Receive(ws, &in)
		fault.Handle(err)

		fmt.Println(in)

		handleMessage(in, collection)
	}))

	err := http.ListenAndServe(port, nil)
	fault.Handle(err)
}

func Accept(ln net.Listener) net.Conn {
	conn, err := ln.Accept()
	fault.Handle(err)

	return conn
}

func HandleLiveConnection(conn net.Conn, wg sync.WaitGroup, collection *mgo.Collection) {
	defer wg.Done()
	fmt.Println("Connection accepted!")

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		fault.Handle(err)

		fmt.Println(msg)

		handleMessage(msg, collection)

		conn.Write([]byte("Done!"))
	}
}

func handleMessage(message string, collection *mgo.Collection) {
	sentences := data.GetSentences(message)

	var sentencePairs [][]data.WordPair

	for _, sentence := range sentences {
		sentencePairs = append(sentencePairs, data.GetPairs(sentence))
	}

	for _, sentencePair := range sentencePairs {
		for _, pair := range sentencePair {
			_, err := collection.Upsert(bson.M{pair.Head: pair.Tail}, bson.M{"$inc": bson.M{"weight": 0.1}})
			fault.Handle(err)
		}
	}
}
