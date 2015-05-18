package main

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

var url = flag.String("url", "localhost:27017", "url of the mongoserver")

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

	primary, err := IsPrimary(session)
	if err != nil {
		log.Fatalf("failed to find primary: %v\n", err)
	}
	fmt.Printf("%v", primary)
}
