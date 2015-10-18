package connections

import (
	"github.com/curiouscain/langbase/fault"
	"gopkg.in/mgo.v2"
)

func GetConnection(host string) *mgo.Session {
	session, err := mgo.Dial(host)
	fault.Handle(err)

	return session
}
