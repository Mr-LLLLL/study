package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	config "gitlab.cd.anpro/go/common/config"
	request "gitlab.cd.anpro/go/common/request"
)

func init() {
	config.Get().Log.Enable = false
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type Vertex struct {
	Label      string `json:"label"`
	Properties Person `json:"properties"`
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 20)
	for i := 0; i < 20000; i++ {
		ch <- struct{}{}
		wg.Add(1)
		i := i
		go func() {
			defer func() {
				<-ch
				wg.Done()
			}()

			start := i * 500
			end := 500 * (i + 1)
			vers := make([]*Vertex, 0)
			for i := start; i < end; i++ {
				vers = append(vers, &Vertex{
					Label: "person",
					Properties: Person{
						Name: strconv.Itoa(i),
						Age:  i,
						City: "chengdu",
					},
				})
			}

			err := request.PostAPI(context.Background(), "http://localhost:18080/graphs/hugegraph/graph/vertices/batch", vers)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
}
