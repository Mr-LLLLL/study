package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"server/spider"
	"time"

	"context"

	"google.golang.org/grpc"
)

type server struct{}

const (
	Address string = "localhost:12345"
	Method  string = "tcp"
)

func (s *server) GetMessage(ctx context.Context, req *spider.MessageReq) (*spider.MessageRes, error) {
	data := make([]string, len(req.Messages))

	for i, v := range req.GetMessages() {
		data[i] = v
	}

	res := &spider.MessageRes{
		Messages: data,
	}

	fmt.Println(res.GetMessages())
	fmt.Println("GetMessage Finished")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("slepping")
	}

	// return res, nil
}

func (s *server) GetClientStream(stream spider.GoSpider_GetClientStreamServer) error {
	data := make([]string, 0)
	for {
		req, err := stream.Recv()
		fmt.Println(req.GetMessages())

		if err == io.EOF {
			fmt.Println(req.GetMessages())
			fmt.Println("GetClientStream Finished")
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()

			return stream.SendAndClose(&spider.MessageRes{
				Messages: data,
			})
		}
		if err != nil {
			return err
		}
		data = append(data, req.GetMessages()...)
	}
}

func (s *server) GetServerStream(req *spider.MessageReq, stream spider.GoSpider_GetServerStreamServer) error {
	for _, v := range req.GetMessages() {
		stream.Send(&spider.MessageRes{
			Messages: []string{v},
		})
		fmt.Println(v)
	}
	fmt.Println("GetServerStream Finished")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	return nil
}

func (s *server) GetTwoStream(stream spider.GoSpider_GetTwoStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("GetTwoStream Finished")
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			return nil
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(req.GetMessages())
		err = stream.Send(&spider.MessageRes{
			Messages: req.GetMessages(),
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	listener, err := net.Listen(Method, Address)
	if err != nil {
		return
	}
	s := grpc.NewServer()
	spider.RegisterGoSpiderServer(s, &server{})

	// reflection.Register(s)
	err = s.Serve(listener)
	if err != nil {
		return
	}
}
