package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"git.dustess.com/mk-base/util/snowflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Init(); err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go test(1, &wg)
	wg.Add(1)
	go test(2, &wg)
	wg.Wait()
}

func test(x int, wg *sync.WaitGroup) {
	defer wg.Done()

	co := client.Database("test").Collection("foobar")

	models := make([]mongo.WriteModel, 0, 5)
	for i := 113; i < 10000; i++ {
		upDoc := bson.M{
			"$inc": bson.M{
				"version": 1,
				"num":     x,
			},
			"$set": bson.M{
				"updatedTime": time.Now().Format("2006-01-02 15:04:05"),
			},
		}
		temp := mongo.NewUpdateOneModel().SetFilter(bson.M{"id": i, "version": 0}).SetUpdate(upDoc)
		models = append(models, temp)
	}

	opts := options.BulkWrite().SetOrdered(false)
	result, err := co.BulkWrite(context.Background(), models, opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.ModifiedCount)
}

func test1() {
	co := client.Database("test").Collection("foobar")

	t := time.Now().Format("2006-01-02 15:04:05")
	arr := make([]Order, 10)
	for i := range arr {
		arr[i].ID = snowflake.BaseNumber()
		arr[i].CreateTime = t
		arr[i].UpdateTime = t
	}
	co.InsertMany(context.Background(), nil, nil)
}

// BaseModel BaseModel
type BaseModel struct {
	ID         string `bson:"_id"`        // id
	CreateTime string `bson:"createTime"` // 创建时间
	UpdateTime string `bson:"updateTime"` // 更新时间
}

type Order struct {
	BaseModel `bson:"inline"`
	UserId    string `bson:"userId"`
	Filed     string `bson:"Filed"`
}
