package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/prometheus/common/log"
	"sisyphus/common/setting"
	"time"
)

var DefaultConn *Conn

func SetUp() {
	DefaultConn = NewConn(setting.DefaultRedisSetting)
}

type Conn struct {
	RedisConn *redis.Pool
}

func NewConn(conf *setting.Redis) *Conn {
	r := &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxActive,
		IdleTimeout: conf.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.Host)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			if conf.Password != "" {
				if _, err := c.Do("AUTH", conf.Password); err != nil {
					log.Error(err)
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			log.Error(err)
			return err
		},
	}
	return &Conn{RedisConn: r}
}

func (c *Conn) Set(key string, data interface{}, time int) error {
	conn := c.RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	log.Info("set cache: ", key, " expire: ", time)
	return nil
}

func (c *Conn) Exists(key string) (bool, error) {
	conn := c.RedisConn.Get()
	defer conn.Close()

	exist, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	log.Info("exists cache: ", key, "v: ", exist)
	return exist, nil
}

func (c *Conn) Get(key string) ([]byte, error) {
	conn := c.RedisConn.Get()
	defer conn.Close()

	resp, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	log.Info("get cache: ", key)
	return resp, nil
}

func (c *Conn) Delete(key string) (bool, error) {
	conn := c.RedisConn.Get()
	defer conn.Close()

	log.Info("delete cache: ", key)
	return redis.Bool(conn.Do("DEL", key))
}

func (c *Conn) BatchDelete(key string) error {
	conn := c.RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err = c.Delete(key)
		if err != nil {
			return err
		}
		log.Info("batch delete cache: ", key)
	}

	return nil
}
