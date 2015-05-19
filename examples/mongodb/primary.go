package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	url   = flag.String("url", "localhost:27017", "url of the mongoserver")
	watch = flag.Duration("watch", 0, "The interval at which to query the cluster for the primary")
)

type IsPrimaryResults struct {
	PrimaryAddress string `bson:"primary"`
}

func IsPrimary(session *mgo.Session) (string, error) {
	results := &IsPrimaryResults{}
	err := session.Run("isMaster", results)
	if err != nil {
		return "", err
	}

	return results.PrimaryAddress, nil
}

func main() {
	flag.Parse()
	session, err := mgo.Dial(*url)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer session.Close()

	for {
		primary, err := IsPrimary(session)
		if err != nil {
			log.Fatalf("failed to find primary: %v\n", err)
		}
		if *watch == 0 {
			fmt.Printf("%v\n", primary)
			break
		}
		log.Println(primary)
		time.Sleep(*watch)
	}
}
