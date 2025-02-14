package integration

import (
	"os"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
)

func TestRedis(t *testing.T) {
	pool, cleanup := setupRedisPool(t)
	defer cleanup()

	c := pool.Get()
	defer c.Close()

	res, err := c.Do("SET", "testing", "value")
	if err != nil {
		t.Fatal(err)
	}

	if res != "OK" {
		t.Fatal("unexpected result: got %s, expected OK", res)
	}
}

func setupRedisPool(t *testing.T) (*redis.Pool, func()) {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		t.Fatal("no REDIS_URL configured")
	}

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redisurl.ConnectToURL(redisURL)
			if err != nil {
				t.Fatal(err)
			}
			return conn, nil
		},
		MaxIdle:   1,
		MaxActive: 50,
	}

	return pool, func() {
		c := pool.Get()
		c.Do("FLUSHALL")
		pool.Close()
	}
}
