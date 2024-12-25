package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

func main() {
	conn, err := jsonrpc.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)

	err3 := conn.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：", ret)
}
