package main

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
)

func main() {
	dialInfo := &mgo.DialInfo{
		Addrs:   []string{"127.0.0.1:27017"},
		Timeout: 10 * time.Second,
		Direct:  true,
	}
	fmt.Println("Attempting to connect to MongoDB...")
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer session.Close()

	err = session.Ping()
	if err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")
}
