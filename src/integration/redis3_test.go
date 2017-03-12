package integration

import (
	"os"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
)

//TESTSTART OMIT
func TestRedis(t *testing.T) {
	t.Run("Set initial value", setupAndTeardown(func(t *testing.T, pool *redis.Pool) {
		c := pool.Get()
		defer c.Close()

		res, err := c.Do("SET", "testing", "value")
		if err != nil {
			t.Fatal(err)
		}

		if res != "OK" {
			t.Fatal("unexpected result: got %s, expected OK", res)
		}
	}))

	t.Run("Handle no initial value", setupAndTeardown(func(t *testing.T, pool *redis.Pool) {
		c := pool.Get()
		defer c.Close()

		res, err := c.Do("GET", "testing")
		if err != nil {
			t.Fatal(err)
		}

		if res != nil {
			t.Fatal("unexpected result: got %s, expected nil", res)
		}
	}))

}

//TESTEND OMIT

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

//SETUP OMIT
func setupAndTeardown(test func(t *testing.T, p *redis.Pool)) func(t *testing.T) {
	return func(t *testing.T) {
		pool, cleanup := setupRedisPool(t)
		defer cleanup()
		test(t, pool)
	}
}

//ENDSETUP OMIT
