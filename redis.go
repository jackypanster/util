package util

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisService struct {
	pool *redis.Pool
	list string
}

func NewRedisService(pool *redis.Pool, list string) *RedisService {
	CheckStr(list, "list")
	CheckCondition(pool == nil, "redis pool should not be empty")
	return &RedisService{
		pool: pool,
		list: list,
	}
}

func (self *RedisService) Rpush(data interface{}) error {
	CheckCondition(data == nil, "data should not be nil")

	c := self.pool.Get()
	defer c.Close()

	_, err := c.Do("RPUSH", self.list, data)
	if err != nil {
		return err
	}
	return nil
}

func (self *RedisService) Lpop() (string, error) {
	c := self.pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("LPOP", self.list))
	if err != nil {
		if err == redis.ErrNil {
			return "", nil
		}
		//Errorf(Map{"error": err}, "unable to lpop of %s", self.list)
		return "", err
	}
	return reply, nil
}

func (self *RedisService) Size() (int, error) {
	c := self.pool.Get()
	defer c.Close()

	size, err := redis.Int(c.Do("LLEN", self.list))
	if err != nil {
		return -1, err
	}
	return size, nil
}

func NewRedisPool(host string, port int) *redis.Pool {
	CheckStr(host, "host")
	hostPort := fmt.Sprintf("%s:%d", host, port)

	pool := &redis.Pool{
		MaxActive:   65536,
		MaxIdle:     256,
		IdleTimeout: 128 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", hostPort)
			if err != nil {
				log.Printf("fail to dial %s, error %s", hostPort, err.Error())
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				log.Printf("fail to ping %s, error %s", host, err.Error())
			}
			return err
		},
	}

	return pool
}

func GetRedisPool(url string) *redis.Pool {
	CheckStr(url, "url")
	pool := &redis.Pool{
		MaxActive:   65536,
		MaxIdle:     256,
		IdleTimeout: 128 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url)
			if err != nil {
				log.Printf("fail to dial %s, error %s", url, err.Error())
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				log.Printf("fail to ping %s, error %s", url, err.Error())
			}
			return err
		},
	}

	return pool
}

func GetPool(url string, password string, db int) *redis.Pool {
	CheckStr(url, "url")
	pool := &redis.Pool{
		MaxActive:   65536,
		MaxIdle:     256,
		IdleTimeout: 128 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url)
			if err != nil {
				log.Panicf("fail to dial %s, error %s", url, err.Error())
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				log.Panicf("fail to auth %s, error %s", password, err.Error())
				return nil, err
			}
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				log.Panicf("fail to select db %d, error %s", db, err.Error())
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				log.Panicf("fail to ping %s, error %s", url, err.Error())
			}
			return err
		},
	}

	return pool
}
