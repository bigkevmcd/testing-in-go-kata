package integration

import (
	"os"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
)

//TEST OMIT
func TestRedis(t *testing.T) {
	pool, cleanup := setupRedisPool(t)
	defer cleanup()
	t.Run("Set initial value", func(t *testing.T) {
		c := pool.Get() //OMIT
		defer c.Close() //OMIT
		res, err := c.Do("SET", "testing", "value")
		if err != nil {
			t.Fatal(err)
		}

		if res != "OK" {
			t.Fatalf("unexpected result: got '%s', expected OK", res)
		}
	})
	t.Run("Handle no initial value", func(t *testing.T) {
		c := pool.Get() //OMIT
		defer c.Close() //OMIT
		res, err := c.Do("GET", "testing")
		if err != nil {
			t.Fatal(err)
		}

		if res != nil {
			t.Fatalf("unexpected result: got '%s', expected nil", res) // HL
		}
	})
}

//ENDTEST OMIT

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
