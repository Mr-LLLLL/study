package data

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

//test data
var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func setup() {
	ThreadDeleteAll()
	SessionDeleteAll()
	UserDeleteAll()
}

func Test_setup(t *testing.T) {
	setup()
}

func Test_createUUID(t *testing.T) {
	fmt.Println(createUUID())
	fmt.Println(createUUID())
	fmt.Println(createUUID())
}

func emptyRedis(conn redis.Conn) {
	_, err := conn.Do("flushdb")
	if err != nil {
		log.Printf("empty redis err: %v", err)
	}
}

func checkAndPrint(i interface{}, err error) {
	if err != nil {
		log.Println(err)
	} else if i != nil {
		fmt.Println(i)
	}
}

type RedisSession struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt int64
}

// redist testing
func Test_Redis(t *testing.T) {
	conn := Pool.Get()
	defer conn.Close()

	emptyRedis(conn)

	checkAndPrint(conn.Do("set", "key1", "hello"))
	checkAndPrint(conn.Do("set", "number", 123))
	checkAndPrint(conn.Do("incr", "number"))
	checkAndPrint(conn.Do("hmset", "myHash", "name", "zhaogang", "gender", "male"))
	checkAndPrint(conn.Do("lpush", "myList", "1", "2"))
	checkAndPrint(conn.Do("rpush", "myList", "3", "4"))
	checkAndPrint(conn.Do("sadd", "mySet", "1", "2", "3"))
	checkAndPrint(conn.Do("zadd", "mySort", "3", "three", "2", "two", "1", "one", "1.1", "twoPointTwo"))

	checkAndPrint(redis.String(conn.Do("get", "key1")))
	checkAndPrint(redis.Int(conn.Do("get", "number")))
	checkAndPrint(redis.StringMap(conn.Do("hgetall", "myHash")))
	checkAndPrint(redis.Strings(conn.Do("lrange", "myList", "0", "-1")))
	checkAndPrint(redis.Strings(conn.Do("smembers", "mySet")))
	checkAndPrint(redis.StringMap(conn.Do("zrange", "mySort", "0", "-1", "withscores")))
	checkAndPrint(redis.Strings(conn.Do("zrange", "mySort", "0", "-1")))

	fmt.Println()
	checkAndPrint(conn.Do("hmset", "session", "id", 1, "uuid", 123, "email", "@qq", "created_at", time.Now().Unix()))
	checkAndPrint(redis.StringMap(conn.Do("hgetall", "session")))
	values, err := redis.Values(conn.Do("hgetall", "session"))
	checkAndPrint(values, err)

	var s RedisSession
	checkAndPrint(nil, redis.ScanStruct(values, &s))
	fmt.Println(s)
	fmt.Println()

	// redis implement struct to redis two mothod	(recommend json method)
	s = RedisSession{
		Id:        1,
		Uuid:      "123",
		Email:     "@qq",
		UserId:    123,
		CreatedAt: time.Now().Unix(),
	}

	data, _ := json.Marshal(s)
	_, err = conn.Do("hset", "sessions", fmt.Sprintf("session{%d}", s.Id), data)
	if err != nil {
		log.Println(err)
	}
	rebytes, err := redis.Bytes(conn.Do("hget", "sessions", fmt.Sprintf("session{%d}", s.Id)))
	if err != nil {
		log.Println(err)
	}
	var s1 RedisSession
	json.Unmarshal(rebytes, &s1)
	fmt.Printf("%#v\n", s1)
	fmt.Println()

	s2 := RedisSession{
		Id:        1,
		Uuid:      "111",
		Email:     "@qq.com",
		UserId:    2,
		CreatedAt: time.Now().Unix(),
	}
	_, err = conn.Do("hmset", redis.Args{"session2"}.AddFlat(s2)...)
	if err != nil {
		log.Println(err)
	}
	value, _ := redis.Values(conn.Do("hgetall", "session2"))
	var s3 RedisSession
	redis.ScanStruct(value, &s3)
	fmt.Println("hash type:")
	fmt.Printf("%#v\n", s3)
	fmt.Println()

	// redis store complexity struct
	sessions := make([]RedisSession, 0, 200)
	for i := 0; i < 200; i++ {
		temp := RedisSession{}
		temp.Id = i
		sessions = append(sessions, temp)
	}
	datas, _ := json.Marshal(sessions)
	_, err = conn.Do("hset", "sessions", fmt.Sprintf("session{%d}", 0), datas)
	if err != nil {
		log.Println(err)
	}
	rebytes, err = redis.Bytes(conn.Do("hget", "sessions", fmt.Sprintf("session{%d}", 0)))
	if err != nil {
		log.Println(err)
	}
	var sessions1 []RedisSession
	json.Unmarshal(rebytes, &sessions1)
	fmt.Printf("%#v\n", sessions1[1])
	fmt.Printf("%#v\n", sessions1[3])
}

func Test_Testing(t *testing.T) {
	var m map[string]int
	m1 := make(map[string]int)
	fmt.Printf("%p, %p", &m, &m1)
}
