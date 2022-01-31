package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/proto"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
}

// 服务端流模式
func (s server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		err := res.Send(&proto.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// 客户端流模式
func (s server) PutStream(putStream proto.Greeter_PutStreamServer) error {
	for {
		if a, err := putStream.Recv(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

// 双向流模式
func (s server) AllStream(allStream proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStream.Recv()
			fmt.Println("收到客户端信息", data.Data)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {

			allStream.Send(&proto.StreamResData{
				Data: "我是服务端",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
