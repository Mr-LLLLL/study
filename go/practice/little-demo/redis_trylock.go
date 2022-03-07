package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type lock struct {
	c chan struct{}
}

var counter int

func redis_try_incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}

func newLock() lock {
	var l lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l lock) lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}

	return lockResult
}

func (l lock) unLock() {
	l.c <- struct{}{}
}

func trylock() {
	var l = newLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.lock() {
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.unLock()
		}()
	}
	wg.Wait()
}
