package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	dbUri := "neo4j://localhost" // scheme://host(:port) (default port is 7687)
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth("neo4j", "asdfjkl;", ""))
	if err != nil {
		panic(err)
	}
	// Starting with 5.0, you can control the execution of most driver APIs
	// To keep things simple, we create here a never-cancelling context
	// Read https://pkg.go.dev/context to learn more about contexts
	ctx := context.Background()
	// Handle driver lifetime based on your application lifetime requirements.
	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per
	// application

	defer driver.Close(ctx) // Make sure to handle errors during deferred calls
	item, err := insertItem(ctx, driver)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", item)
}

func insertItem(ctx context.Context, driver neo4j.DriverWithContext) (*Item, error) {
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- struct{}{}
		i := i
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()
			start := 10000 * i
			end := 10000 * (i + 1)

			for i := start; i < end; i++ {
				_, err := neo4j.ExecuteQuery(ctx, driver,
					"CREATE (n:Item { id: $id, name: $name, text: $text })",
					map[string]any{
						"id":   i,
						"name": strconv.Itoa(i),
						"text": "hellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohello",
					}, neo4j.EagerResultTransformer)
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
	// _, err := neo4j.ExecuteQuery(ctx, driver,
	// 	"CREATE (n:Item { id: $id, name: $name })",
	// 	map[string]any{
	// 		"id":   1,
	// 		"name": "Item 1",
	// 	}, neo4j.EagerResultTransformer)
	// if err != nil {
	// 	return nil, err
	// }
	// itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
	// if err != nil {
	// 	return nil, fmt.Errorf("could not find node n")
	// }
	// id, err := neo4j.GetProperty[int64](itemNode, "id")
	// if err != nil {
	// 	return nil, err
	// }
	// name, err := neo4j.GetProperty[string](itemNode, "name")
	// if err != nil {
	// 	return nil, err
	// }
	// return &Item{Id: id, Name: name}, nil
	return nil, nil
}

type Item struct {
	Id   int64
	Name string
}

func (i *Item) String() string {
	return fmt.Sprintf("Item (id: %d, name: %q)", i.Id, i.Name)
}
