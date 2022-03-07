package main

import (
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"
)

func main() {
	go gorun()

	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}

func test() {
}

func gorun() {
	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
		log.Fatal(err)
	}
}
