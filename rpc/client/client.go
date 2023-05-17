package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ArithReq struct {
	A int
	B int
}

type ArithRes struct {
	Pro int // *
	Quo int // /
	Rem int // %
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	req := ArithReq{9, 2}
	var res ArithRes
	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
