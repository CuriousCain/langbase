package connections

import (
	"github.com/curiouscain/langbase/fault"
	"gopkg.in/mgo.v2"
)

func GetConnection(host string, database string) *mgo.Database {
	session, err := mgo.Dial(host)
	fault.Handle(err)

	defer session.Close()

	db := session.DB(database)

	return db
}
