package main

import (
	"context"
	"log"

	"go-demo/grpc/proto/love"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接服务
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	// 很关键
	defer conn.Close()

	// 初始化客户端
	c := love.NewLoveClient(conn)

	// 发起请求
	response, err := c.Confession(context.Background(), &love.Request{Name: "qq"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Res)
}
