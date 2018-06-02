package mongodbpersistence

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestNewMongodbPersistence(t *testing.T) {
	t.Log("Actual Mongodb persistence, please disable on build process")
	{
		session, err := mgo.Dial("mongodb://root:jibilal@localhost")

		if err != nil {
			t.Errorf("Failed, %s", err.Error())
		}

		defer session.Close()

		p := NewMongodbPersistence(session, "testdb")

		data := struct {
			ID          bson.ObjectId `bson:"_id"`
			ReferenceId string        `bson:"ref_id"`
			Items       []struct {
				Name     string
				Quantity int
			}
		}{
			ID:          bson.NewObjectId(),
			ReferenceId: "123444",
			Items: []struct {
				Name     string
				Quantity int
			}{
				{
					Name:     "XL First Gigz",
					Quantity: 2,
				},
				{
					Name:     "XL First 8 Gigz",
					Quantity: 2,
				},
			},
		}

		err = p.Store("ordertest", data)
		if err != nil {
			t.Errorf("Test %s", err.Error())
		}

		var result []*struct {
			ID          bson.ObjectId `bson:"_id"`
			ReferenceId string        `bson:"ref_id"`
			Items       interface{}
		}
		err = p.Fetch(nil, "ordertest", &result)

		t.Log("Ze Result", result)
	}
}
