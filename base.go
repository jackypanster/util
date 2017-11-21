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
	All() (interface{}, error)
	Find(string) (interface{}, error)
	Remove(string) error
	Empty() error
	FindByDate(string) (interface{}, error)
	FindByName(string) (interface{}, error)
	FindByTimestamp(string, string) (interface{}, error)
}

type Service struct {
	cache map[string]interface{}
}

func NewService() *Service {
	return &Service{
		cache: make(map[string]interface{}),
	}
}

func (service *Service) Insert(tab *mgo.Collection, id string, doc interface{}) (interface{}, error) {
	err := tab.Insert(doc)
	if err == nil {
		service.cache[id] = doc
	}
	return service.cache[id], err
}

func (service *Service) Remove(tab *mgo.Collection, id string) error {
	err := tab.Remove(bson.M{"id": id})
	if err == nil {
		delete(service.cache, id)
	}
	return err
}

func (service *Service) RemoveAll(tab *mgo.Collection) error {
	_, err := tab.RemoveAll(nil)
	if err == nil {
		service.cache = make(map[string]interface{})
	}
	return err
}

func (service *Service) Find(tab *mgo.Collection, id string) (interface{}, error) {
	if _, ok := service.cache[id]; ok {
		return service.cache[id], nil
	}

	var result interface{}
	err := tab.Find(bson.M{"id": id}).One(&result)
	if err == nil {
		service.cache[id] = &result
	}
	return service.cache[id], err
}

func (service *Service) Update(tab *mgo.Collection, id string, selector interface{}, update interface{}, doc interface{}) (interface{}, error) {
	err := tab.Update(selector, update)
	if err == nil {
		service.cache[id] = doc
	}
	return service.cache[id], err
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
