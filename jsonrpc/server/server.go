package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Rect struct {
}

type Req struct {
	Width  int
	Height int
}

func (this *Rect) Area(req Req, res *int) error {
	*res = req.Width * req.Height
	return nil
}

func main() {
	// 注册服务
	rpc.Register(new(Rect))
	// 监听端口
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	//
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)

		i++
		fmt.Println(i)
	}
}
