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
	CheckNil(pool, "redis pool should not be empty")
	return &RedisService{
		pool: pool,
		list: list,
	}
}

func (self *RedisService) Set(key string, v interface{}, ttl int) error {
	CheckStr(key, "key")
	CheckNil(v, "v")
	CheckCondition(ttl <= 0, "ttl should be positive")

	c := self.pool.Get()
	defer c.Close()

	val, err := Encode(v)
	if err != nil {
		return err
	}
	_, err = c.Do("SET", key, val, "EX", ttl)
	if err != nil {
		return err
	}
	return nil
}

func (self *RedisService) Get(key string, result interface{}) error {
	CheckStr(key, "key")
	CheckNil(result, "result")

	c := self.pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("GET", key))
	if err != nil {
		return err
	}

	err = Decode(reply, result)
	if err != nil {
		return err
	}

	return nil
}

func (self *RedisService) Rpush(v interface{}) error {
	CheckNil(v, "arg should not be nil")

	c := self.pool.Get()
	defer c.Close()

	_, err := c.Do("RPUSH", self.list, v)
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
