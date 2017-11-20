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

type EntityService struct {
  cache map[string]*Entity
}

func NewEntityService() *EntityService {
  return &EntityService{
    cache: make(map[string]*Entity),
  }
}

func (self *EntityService) FindById(tab *mgo.Collection, id string) (*Entity, error) {
  if val, ok := self.cache[id]; ok {
    return val, nil
  }

  var entity Entity
  if err := tab.Find(bson.M{"id": id}).One(&entity); err != nil {
    return nil, err
  }
  self.cache[entity.ID] = &entity
  return &entity, nil
}

func (self *EntityService) FindByDuration(tab *mgo.Collection, start string, end string, size int) ([]*Entity, error) {

  var entities []*Entity
  var err error

  from := ConvertTimestamp(start)
  to := ConvertTimestamp(end)

  err = tab.Find(bson.M{"timestamp": bson.M{"$gte": from, "lt": to}}).Sort("-timestamp").Limit(size).All(&entities)
  return entities, err
}

func (self *EntityService) All(tab *mgo.Collection, size int) ([]*Entity, error) {
  var entities []*Entity
  var err error

  err = tab.Find(bson.M{}).Sort("-timestamp").Limit(size).All(&entities)
  return entities, err
}
