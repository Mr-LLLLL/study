package data

import (
	"crypto/sha1"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	uuid "github.com/nu7hatch/gouuid"
)

var (
	Db          *sql.DB
	Pool        *redis.Pool
	redisServer = flag.String("redisServer", ":6379", "assign redis server address and port")
)

func init() {
	initPG()
	initRedis()
	initLogFile()
}

var logger *log.Logger
func initLogFile() {
	file, err := os.OpenFile("data.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open data log file", err)
	}
	logger = log.New(file, "DATA", log.Ldate|log.Ltime|log.Lshortfile)
}

func checkErr(err error) {
	if err != nil {
		logger.Println(err)
	}
}

func initPG() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp password=gwp dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func initRedis() {
	if strings.HasSuffix(os.Args[0], ".test") {
		fmt.Println("in testing mode")
	} else {
		flag.Parse()
	}
	Pool = newPool(*redisServer)
	close()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         3,
		MaxActive:       0,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

// the function is testing
func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() string {
	// use third package
	u1, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	return u1.String()
	/*
		 * original
		u := new([16]byte)
		_, err := rand.Read(u[:])
		if err != nil {
			log.Fatalln("Cannot generate UUID", err)
		}

		// 0x40 is reserved variant from RFC 4122
		u[8] = (u[8] | 0x40) & 0x7F
		// Set the four most significant bits (bits 12 throught 15) of the
		// time_hi_and_version field to the 4-bit version number.
		u[6] = (u[6] & 0xF) | 0x40
		uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
		return
	*/
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
