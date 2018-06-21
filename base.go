package util

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Entity struct {
	ID           string    `json:"id" bson:"id"`
	Date         string    `json:"date" bson:"date"`
	Name         string    `json:"name" bson:"name"`
	Milliseconds int64     `json:"milliseconds" bson:"milliseconds"`
	Timestamp    time.Time `json:"timestamp" bson:"timestamp"`
}

type Operator interface {
	All(interface{}) error
	Insert(interface{}) error
	Update(interface{}, interface{}) error
	Remove(string) error
	Empty() error
	Find(string, interface{}) error
	FindByDate(string, interface{}) error
	FindByName(string, interface{}) error
	FindByTimestamp(string, string, interface{}) error
	FindOne(interface{}, interface{}) error
	Search(interface{}, interface{}) error
}

type Service struct {
	session  *mgo.Session
	database string
	table    string
	limit    int
}

func NewService(session *mgo.Session, database string, table string, size int) *Service {
	return &Service{
		session:  session,
		database: database,
		table:    table,
		limit:    size,
	}
}

func (self *Service) Insert(doc interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Insert(doc)
}

func (self *Service) Remove(id string) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Remove(bson.M{"id": id})
}

func (self *Service) Empty() error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	_, err := c.RemoveAll(nil)
	return err
}

func (self *Service) Find(id string, result interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Find(bson.M{"id": id}).One(result)
}

func (self *Service) Update(selector interface{}, update interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Update(selector, update)
}

func (self *Service) FindByDate(date string, results interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Find(bson.M{"date": date}).Sort("-timestamp").Limit(self.limit).All(results)
}

func (self *Service) FindByName(name string, results interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Find(bson.M{"name": name}).Sort("-timestamp").Limit(self.limit).All(results)
}

func (self *Service) FindByTimestamp(start string, end string, results interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	from := ConvertTimestamp(start)
	to := ConvertTimestamp(end)
	return c.Find(bson.M{"timestamp": bson.M{"$gte": from, "$lt": to}}).Sort("-timestamp").Limit(self.limit).All(results)
}

func (self *Service) All(results interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Find(bson.M{}).Sort("-timestamp").Limit(self.limit).All(results)
}

func (self *Service) Search(query interface{}, results interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Find(query).Sort("-timestamp").Limit(self.limit).All(results)
}

func (self *Service) FindOne(query interface{}, result interface{}) error {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	return c.Find(query).One(result)
}

func (self *Service) Exists(query interface{}) bool {
	s := self.session.Copy()
	defer s.Close()
	c := s.DB(self.database).C(self.table)

	n, _ := c.Find(query).Select(bson.M{"_id": 1}).Limit(1).Count()
	return n == 1
}
