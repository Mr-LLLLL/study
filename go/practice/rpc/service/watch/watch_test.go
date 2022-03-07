package watch

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"testing"
	"time"
)

func TestKVStoreService_Watch(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	doClientWork(client)
}

func TestTest(t *testing.T) {
	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
		}()
	}()
	<-done
	fmt.Println("Done")
}

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()

	for i := 0; i < 10; i++ {
		err := client.Call(
			"KVStoreService.Set", [2]string{"abc" + strconv.Itoa(i), "abc-value"},
			new(struct{}),
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(time.Second * 3)
}
