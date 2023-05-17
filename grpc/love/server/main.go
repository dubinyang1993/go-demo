package main

import (
	"context"
	"log"
	"net"

	"dubinyang1993/go-demo/grpc/proto/love"

	"google.golang.org/grpc"
)

// 定义服务
type Love struct {
	love.UnimplementedLoveServer // 必须有
}

// 实现接口
func (l *Love) Confession(ctx context.Context, req *love.Request) (*love.Response, error) {
	res := love.Response{}
	res.Res = "I love you " + req.Name
	return &res, nil
}

func main() {
	// 监听8888端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	// 实例化grpc server
	s := grpc.NewServer()

	// 注册Love服务
	love.RegisterLoveServer(s, &Love{})

	log.Println("Listen on 127.0.0.1:8888...")
	s.Serve(listen)
}
