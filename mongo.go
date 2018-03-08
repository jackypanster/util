package util

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

func NewMongoSession(host string, port int) *mgo.Session {
	CheckStr(host, "host")

	url := fmt.Sprintf("mongodb://%s:%d", host, port)
	// Connect to our local mongo
	session, err := mgo.DialWithTimeout(url, time.Minute*3)
	// Check if connection error, is mongo running?
	CheckErrf(err, "fail to dial mongo")

	session.SetMode(mgo.Nearest, true)
	return session
}

func NewMgoSession(url string) *mgo.Session {
	CheckStr(url, "url")
	// Connect to our local mongo
	session, err := mgo.DialWithTimeout(url, time.Minute*3)
	// Check if connection error, is mongo running?
	CheckErrf(err, "fail to dial mongo")

	session.SetMode(mgo.Nearest, true)
	return session
}
