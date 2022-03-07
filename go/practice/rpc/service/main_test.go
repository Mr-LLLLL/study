package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address, fileType string) (*HelloServiceClient, error) {
	var (
		client *rpc.Client
		err    error
	)
	if fileType == "json" {
		conn, err := net.Dial(network, address)
		if err != nil {
			log.Fatal("net.Dial:", err)
		}
		client = rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	} else {
		client, err = rpc.Dial(network, address)
		if err != nil {
			return nil, err
		}
	}

	return &HelloServiceClient{Client: client}, err
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func (p *HelloServiceClient) AsyncHello(request string, reply *string) error {
	helloCall := p.Client.Go(HelloServiceName+".Hello", request, reply, nil)
	fmt.Println(*reply)

	// do something

	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		return err
	}

	return nil
}

func TestHelloService_Hello(t *testing.T) {
	client, err := DialHelloService("tcp", "localhost:1234", "json")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.AsyncHello("world", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
