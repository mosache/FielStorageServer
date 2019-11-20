package redis

import "github.com/gomodule/redigo/redis"

import "time"

var (
	pool *redis.Pool
)

func init() {
	pool = newRedisPool()
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			//1.打开连接
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")

			return err
		},
	}
}

//RedisPool RedisPool	return pool
func RedisPool() *redis.Pool {
	return pool
}
