package data

import "github.com/garyburd/redigo/redis"

func hitTheTopRanking(thread Thread) {
	conn := Pool.Get()
	defer conn.Close()
	currPostNum := thread.NumReplies()
	_, err := conn.Do("zadd", "topRanking", -currPostNum, thread.Uuid)
	checkErr(err)
	_, err = conn.Do("zremrangebyrank", "topRanking", 3, -1)
	checkErr(err)
}

func GetHotspotThreads() (threads []Thread) {
	conn := Pool.Get()
	defer conn.Close()
	UUIDs, err := redis.Strings(conn.Do("zrange", "topRanking", 0, -1))
	if err != nil && err != redis.ErrNil {
		checkErr(err)
		return
	}
	for _, v := range UUIDs {
		t, err := ThreadByUUID(v)
		if err != nil {
			checkErr(err)
		}
		threads = append(threads, t)
	}
	return
}
