package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Arith struct {
}

type ArithReq struct {
	A int
	B int
}

type ArithRes struct {
	Pro int // *
	Quo int // /
	Rem int // %
}

func (this *Arith) Multiply(req ArithReq, res *ArithRes) error {
	res.Pro = req.A * req.B
	return nil
}

func (this *Arith) Divide(req ArithReq, res *ArithRes) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	// 注册服务
	rpc.Register(new(Arith))
	// 服务绑定 http 协议
	rpc.HandleHTTP()
	// 监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
