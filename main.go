package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
	"github.com/zibilal/essynctest/persistence/mongodbpersistence"
)

func main() {
	fmt.Println("Starts...")

	sessiondb, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{
			"localhost:28017",
		},
		Timeout: 5 * time.Second,
		Database: "terracotta",
		Username: "terracotta",
		Password: "terracotta",
		Direct: true,
	})

	if err != nil {
		panic(err)
	}

	defer sessiondb.Close()

	mpo := mongodbpersistence.NewMongodbPersistence(sessiondb, "terracotta")
	err = mpo.Store("events", struct {
		Name string
		Address string
	}{
		"Test", "Jl. Kemang Raya No. 15",
	})

	if err != nil {
		fmt.Errorf("Failed saving test data: %s\n", err.Error())
	}

	fmt.Println("Data is saved")

	fmt.Println("Ends...")
}
