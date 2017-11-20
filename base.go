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

func (service *Service) Insert(tab *mgo.Collection, id string, doc interface{}) error {
	service.cache[id] = doc
	return tab.Insert(doc)
}

func (service *Service) Remove(tab *mgo.Collection, id string) error {
	delete(service.cache, id)
	return tab.Remove(bson.M{"id": id})
}

func (service *Service) RemoveAll(tab *mgo.Collection) error {
	service.cache = make(map[string]interface{})
	_, err := tab.RemoveAll(nil)
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

func (service *Service) Update(tab *mgo.Collection, id string, selector interface{}, update interface{}, doc interface{}) error {
	service.cache[id] = doc
	return tab.Update(selector, update)
}

func (service *Service) FindByTimestamp(tab *mgo.Collection, start string, end string, size int, result interface{}) error {
	from := ConvertTimestamp(start)
	to := ConvertTimestamp(end)
	return tab.Find(bson.M{"timestamp": bson.M{"$gte": from, "$lt": to}}).Sort("-timestamp").Limit(size).All(result)
}

func (service *Service) All(tab *mgo.Collection, size int) (interface{}, error) {
	var results []interface{}
	err := tab.Find(bson.M{}).Sort("-timestamp").Limit(size).All(results)
	return results, err
}
