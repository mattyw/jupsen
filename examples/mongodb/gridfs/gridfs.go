// This file really isn't an example of how to use mgo.
// It exists to try to cause/ analyse problems
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"gopkg.in/mgo.v2"
)

var (
	url      = flag.String("url", "localhost:27017", "url of the mongoserver")
	count    = flag.Int("count", 6000, "number of list items")
	majority = flag.Bool("majority", true, `use WMode="majority"`)
)

type File struct {
	filename string
	gridfs   *mgo.GridFS
}

func (f *File) Value() (int, error) {
	file, err := f.gridfs.Open(f.filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	value, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(value))
}

func (f *File) Inc() error {
	value, err := f.Value()
	if err != nil {
		return err
	}
	file, err := f.gridfs.Create(f.filename)
	_, err = fmt.Fprint(file, value+1)
	if err != nil {
		return err
	}
	return file.Close()
}

type Stats struct {
	Expected []int
	Pass     int
	Fail     int
}

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

func collection(session *mgo.Session) *mgo.Collection {
	return session.DB("jupsen").C("list")
}

func WriteList(session *mgo.Session, filename string, count int) Stats {
	stats := Stats{}

	s := session.Copy()
	defer s.Close()
	for i := 0; i < count; i++ {
		F := File{filename, s.DB("").GridFS("fs")}
		err := F.Inc()
		if err != nil {
			stats.Fail += 1
			log.Printf("failed update: %v\n", err)

			s.Refresh()
			continue
		}
		stats.Pass += 1
		stats.Expected = append(stats.Expected, i)
		if i%100 == 0 {
			fmt.Println(i)
		}
	}
	return stats
}

func Compare(session *mgo.Session, filename string, stats Stats) {
	s := session.Copy()
	defer s.Close()

	F := File{filename, s.DB("").GridFS("fs")}
	value, err := F.Value()
	if err != nil {
		log.Fatalf("failed to read results: %v\n", err)
	}
	fmt.Printf("Pass:%d\n", stats.Pass)
	fmt.Printf("Fail:%d\n", stats.Fail)
	fmt.Printf("Expect Value: %d\n", len(stats.Expected))
	fmt.Printf("Actual Value: %d\n", value)
}

func main() {
	flag.Parse()
	fmt.Printf("Connecting to %s and appending %d times\n", *url, *count)
	session, err := mgo.Dial(*url)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer session.Close()

	if *majority {
		session.EnsureSafe(&mgo.Safe{WMode: "majority", W: 1, FSync: true})
	} else {
		session.EnsureSafe(&mgo.Safe{W: 1, FSync: true})
	}

	primary, err := IsPrimary(session)
	if err != nil {
		log.Fatalf("failed to find primary: %v\n", err)
	}
	fmt.Printf("primary is on ip %v\n", primary)

	var filename = "myfile"
	file, err := session.DB("").GridFS("fs").Create(filename)
	if err != nil {
		log.Fatalf("failed to get file for setting start point: %v\n", err)
	}
	_, err = file.Write([]byte("0"))
	if err != nil {
		log.Fatalf("failed to write start point: %v\n", err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalf("failed to close file at start point: %v\n", err)
	}

	stats := WriteList(session, filename, *count)
	Compare(session, filename, stats)
}
