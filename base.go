package util

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Entity struct {
	ID        string    `json:"id" bson:"id"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
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

func (service *Service) FindByTimestamp(tab *mgo.Collection, start string, end string, size int) (interface{}, error) {
	from := ConvertTimestamp(start)
	to := ConvertTimestamp(end)
	var results []interface{}
	err := tab.Find(bson.M{"timestamp": bson.M{"$gte": from, "$lt": to}}).Sort("-timestamp").Limit(size).All(&results)
	return results, err
}

func (service *Service) All(tab *mgo.Collection, size int) (interface{}, error) {
	var results []interface{}
	err := tab.Find(bson.M{}).Sort("-timestamp").Limit(size).All(&results)
	return results, err
}
