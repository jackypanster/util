package util

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Entity struct {
	ID        string    `json:"id" bson:"id"`
	Date      string    `json:"date" bson:"date"`
	Name      string    `json:"name" bson:"name"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type Operator interface {
	All(interface{}) error
	Update(interface{}, interface{}) error
	Remove(string) error
	Empty() error
	Find(string, interface{}) error
	FindByDate(string, interface{}) error
	FindByName(string, interface{}) error
	FindByTimestamp(string, string, interface{}) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service *Service) Insert(tab *mgo.Collection, id string, doc interface{}) error {
	return tab.Insert(doc)
}

func (service *Service) Remove(tab *mgo.Collection, id string) error {
	return tab.Remove(bson.M{"id": id})
}

func (service *Service) RemoveAll(tab *mgo.Collection) error {
	_, err := tab.RemoveAll(nil)
	return err
}

func (service *Service) Find(tab *mgo.Collection, id string, result interface{}) error {
	return tab.Find(bson.M{"id": id}).One(result)
}

func (service *Service) Update(tab *mgo.Collection, selector interface{}, update interface{}) error {
	return tab.Update(selector, update)
}

func (service *Service) FindByDate(tab *mgo.Collection, date string, size int, results interface{}) error {
	return tab.Find(bson.M{"date": date}).Sort("-timestamp").Limit(size).All(results)
}

func (service *Service) FindByName(tab *mgo.Collection, name string, size int, results interface{}) error {
	return tab.Find(bson.M{"name": name}).Sort("-timestamp").Limit(size).All(results)
}

func (service *Service) FindByTimestamp(tab *mgo.Collection, start string, end string, size int, results interface{}) error {
	from := ConvertTimestamp(start)
	to := ConvertTimestamp(end)
	return tab.Find(bson.M{"timestamp": bson.M{"$gte": from, "$lt": to}}).Sort("-timestamp").Limit(size).All(results)
}

func (service *Service) All(tab *mgo.Collection, size int, results interface{}) error {
	return tab.Find(bson.M{}).Sort("-timestamp").Limit(size).All(results)
}
