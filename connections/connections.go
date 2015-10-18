package connections

import (
	"bufio"
	"fmt"
	"net"
	"sync"

	"github.com/curiouscain/langbase/data"
	"github.com/curiouscain/langbase/fault"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func HandleLiveConnection(conn net.Conn, wg sync.WaitGroup, collection *mgo.Collection) {
	defer wg.Done()
	fmt.Println("Connection accepted!")

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		fault.Handle(err)

		sentences := data.GetSentences(msg)

		var sentencePairs [][]data.WordPair

		for _, sentence := range sentences {
			sentencePairs = append(sentencePairs, data.GetPairs(sentence))
		}

		for _, sentencePair := range sentencePairs {
			for _, pair := range sentencePair {
				_, err := collection.Upsert(bson.M{pair.Head: pair.Tail}, bson.M{"$inc": bson.M{"weight": 0.1}})
				fault.Handle(err)

				conn.Write([]byte("Done!"))
			}
		}
	}
}
