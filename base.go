package util

import (
  "log"
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

func (service *Service) Remove(tab *mgo.Collection, id string, doc interface{}) error {
  delete(service.cache, id)
  return tab.Remove(bson.M{"id": id})
}

func (service *Service) RemoveAll(tab *mgo.Collection) error {
  service.cache = make(map[string]interface{})
  _, err := tab.RemoveAll(nil)
  return err
}

func (service *Service) Find(tab *mgo.Collection, id string, result interface{}) error {
  if result, ok := service.cache[id]; ok {
    log.Printf("cache %+v", result)
    return nil
  }

  err := tab.Find(bson.M{"id": id}).One(result)
  if err == nil {
    service.cache[id] = result
    log.Printf("store %+v", service.cache[id])
  }
  return err
}

func (service *Service) FindByDuration(tab *mgo.Collection, start string, end string, size int, results []interface{}) error {
  from := ConvertTimestamp(start)
  to := ConvertTimestamp(end)
  return tab.Find(bson.M{"timestamp": bson.M{"$gte": from, "lt": to}}).Sort("-timestamp").Limit(size).All(results)
}

func (service *Service) All(tab *mgo.Collection, size int, results []interface{}) error {
  return tab.Find(bson.M{}).Sort("-timestamp").Limit(size).All(results)
}
