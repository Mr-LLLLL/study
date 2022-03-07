package main

import (
	"client/spider"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	Address string = "localhost:12345"
)

var client spider.GoSpiderClient

func init() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client = spider.NewGoSpiderClient(conn)
}

func main() {
	go getMessage()
	time.Sleep(time.Second * 5)
	getClientStream()
	getServerStream()
	getTwoStream()
}

func getMessage() {
	fmt.Println("getMessage")

	req := spider.MessageReq{
		Messages: []string{"hello", "world", "nihao"},
	}
	res, err := client.GetMessage(context.Background(), &req)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res)
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func getClientStream() {
	fmt.Println("getClientStream")

	msg := []string{"hello", "world", "nihao"}

	stream, err := client.GetClientStream(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range msg {
		err := stream.Send(&spider.MessageReq{
			Messages: []string{v},
		})
		if err != nil {
			log.Println(err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res.GetMessages())
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func getServerStream() {
	fmt.Println("getServerStream")

	req := spider.MessageReq{
		Messages: []string{"hello", "world", "nihao"},
	}
	stream, err := client.GetServerStream(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 0
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Messages, i)
		i++
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func getTwoStream() {
	fmt.Println("getTwoStream")

	msg := []string{"hello", "world", "nihao"}

	stream, err := client.GetTwoStream(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range msg {
		err := stream.Send(&spider.MessageReq{
			Messages: []string{v},
		})
		if err != nil {
			log.Println(err)
			break
		}

		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		fmt.Println(res)
	}
	stream.CloseSend()

	fmt.Println()
	fmt.Println()
	fmt.Println()
}
