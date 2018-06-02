package mongodbpersistence

import (
	"gopkg.in/mgo.v2"
)

type MongodbPersistence struct {
	dbSession *mgo.Session
	dbName    string
}

func NewMongodbPersistence(dbSession *mgo.Session, dbName string) *MongodbPersistence {
	p := new(MongodbPersistence)
	p.dbSession = dbSession
	p.dbName = dbName
	return p
}

func (m *MongodbPersistence) Store(name string, data interface{}) error {

	sCopy := m.dbSession.Copy()

	defer sCopy.Close()

	return sCopy.DB(m.dbName).C(name).Insert(data)
}

func (m *MongodbPersistence) Fetch(query interface{}, name string, output interface{}) error {

	sCopy := m.dbSession.Copy()

	defer sCopy.Close()

	return sCopy.DB(m.dbName).C(name).Find(query).All(output)
}
