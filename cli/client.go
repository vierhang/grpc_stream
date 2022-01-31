package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/proto"
	"sync"
	"time"
)

func GetStream() {
	dial, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	c := proto.NewGreeterClient(dial)
	stream, err := c.GetStream(context.Background(), &proto.StreamReqData{Data: "client stream test"})
	if err != nil {
		return
	}
	for {
		a, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}
}

func PutStream() {
	dial, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	c := proto.NewGreeterClient(dial)
	stream, err := c.PutStream(context.Background())
	i := 0
	for {
		i++
		err := stream.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("weihang_test %d", i),
		})

		time.Sleep(time.Second)
		if err != nil {
			return
		}
		if i > 10 {
			break
		}
	}

}

func AllStream() {
	dial, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	c := proto.NewGreeterClient(dial)
	allStream, err := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStream.Recv()
			fmt.Println("收到服务端消息", data.Data)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			err := allStream.Send(&proto.StreamReqData{
				Data: "我是客户端",
			})
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
func main() {
	//GetStream()
	//PutStream()
	AllStream()
}
