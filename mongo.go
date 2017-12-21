package util

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

func NewMongoSession(host string, port int) *mgo.Session {
	CheckStr(host, "host")

	url := fmt.Sprintf("mongodb://%s:%d", host, port)
	// Connect to our local mongo
	session, err := mgo.Dial(url)
	// Check if connection error, is mongo running?
	CheckErr(err)

	session.SetMode(mgo.Monotonic, true)
	return session
}

func NewMgoSession(url string) *mgo.Session {
	CheckStr(url, "url")
	// Connect to our local mongo
	session, err := mgo.Dial(url)
	// Check if connection error, is mongo running?
	CheckErr(err)

	session.SetMode(mgo.Monotonic, true)
	return session
}
