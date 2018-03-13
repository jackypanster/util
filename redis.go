package util

import (
	"fmt"
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
			CheckErr(err)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
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
			CheckErr(err)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return pool
}
