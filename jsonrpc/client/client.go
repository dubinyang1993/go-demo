package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Req struct {
	Width, Height int
}

func main() {
	conn, err := jsonrpc.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	req := Req{5, 10}
	var res int
	err = conn.Call("Rect.Area", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
