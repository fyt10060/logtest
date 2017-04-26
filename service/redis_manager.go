// redis_manager
package service

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

const (
	keyIsDoingOperation = "isDoingOpertaion"
	keyShouldBeNum      = "shouldBeNum"
	keyElementList      = "elementlist"
)

var (
	MaxPoolSize = 20
	redisPool   chan redis.Conn
)

// redis pool related
func putRedis(conn redis.Conn) {
	if redisPool == nil {
		redisPool = make(chan redis.Conn, MaxPoolSize)
	}
	if len(redisPool) >= MaxPoolSize {
		conn.Close()
		return
	}
	fmt.Printf("redis conn pool size: %d\n", len(redisPool))
	redisPool <- conn
}

func InitRedis(address string) redis.Conn {
	if len(redisPool) == 0 {
		redisPool = make(chan redis.Conn, MaxPoolSize)
		go func() {
			for i := 0; i < MaxPoolSize/2; i++ {
				c, err := redis.Dial("tcp", address, redis.DialPassword("123456redis"))

				if err != nil {
					panic(err)
				}
				putRedis(c)
			}
		}()
	}
	return <-redisPool
}

func getRedisConn() redis.Conn {
	return InitRedis("123.56.137.103:6379")
}

func CheckDoingOperation() bool {
	return getBoolResult(keyIsDoingOperation, false)
}

func CheckShouldBeNum() bool {
	return getBoolResult(keyShouldBeNum, true)
}

func getBoolResult(withKey string, nilValue bool) bool {
	conn := getRedisConn()
	defer putRedis(conn)
	is, err := redis.Bool(conn.Do("get", withKey))
	if err != nil {
		return nilValue
	} else {
		return is
	}
}

func AddToElementList(value string) {
	conn := getRedisConn()
	defer putRedis(conn)
	conn.Do("lpush", keyElementList, value)
}

func GetAllNumberInList() []string {
	conn := getRedisConn()
	defer putRedis(conn)
	lens, err := redis.Int(conn.Do("llen", keyElementList))
	if err != nil {
		return nil
	}
	var list []string
	for i := 0; i < lens; i++ {
		value, _ := redis.String(conn.Do("lpop", keyElementList))
		list = append(list, value)
	}

	return reverse(list)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func setToRedisWithKey(key string, value interface{}) {
	conn := getRedisConn()
	defer putRedis(conn)
	conn.Do("set", key, value)
}

func SetDoingOperation(set bool) {
	setToRedisWithKey(keyIsDoingOperation, set)
}

func SetShouldBeNumber(set bool) {
	setToRedisWithKey(keyShouldBeNum, set)
}
