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

type EntityService struct {
  cache map[string]interface{}
}

func NewEntityService() *EntityService {
  return &EntityService{
    cache: make(map[string]interface{}),
  }
}

func (self *EntityService) FindById(tab *mgo.Collection, id string, result interface{}) error {
  if val, ok := self.cache[id]; ok {
    result = val
    log.Printf("cache %+v", val)
    return nil
  }

  err := tab.Find(bson.M{"id": id}).One(result)
  self.cache[id] = result
  log.Printf("store %+v", self.cache[id])

  return err
}

func (self *EntityService) FindByDuration(tab *mgo.Collection, start string, end string, size int, results []interface{}) error {
  from := ConvertTimestamp(start)
  to := ConvertTimestamp(end)

  return tab.Find(bson.M{"timestamp": bson.M{"$gte": from, "lt": to}}).Sort("-timestamp").Limit(size).All(results)
}

func (self *EntityService) All(tab *mgo.Collection, size int, results []interface{}) error {
  return tab.Find(bson.M{}).Sort("-timestamp").Limit(size).All(results)
}
