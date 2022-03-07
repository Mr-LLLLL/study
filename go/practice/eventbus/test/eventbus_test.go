package main

import (
	"strconv"
	"testing"
	"time"
)

func TestPublish(t *testing.T) {
	InitEventBus()
	InitEvent()
	Wait()

	for i := 0; i < 20; i++ {
		Publish(strconv.Itoa(i))
	}

	Quit()
	time.Sleep(10 * time.Second)
}
