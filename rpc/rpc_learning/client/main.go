package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	A, B int
}

// 返回给客户端的结果
type Response struct {
	ChengJi int
	Shang   int
	YuShu   int
}

func main() {
	conn, err := jsonrpc.Dial("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	params := Params{9, 2}
	var res Response
	err2 := conn.Call("Operation.Multi", params, &res)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%d * %d = %d\n", params.A, params.B, res.ChengJi)
	err3 := conn.Call("Operation.Division", params, &res)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Printf("%d / %d 商 %d, 余 %d ", params.A, params.B, res.Shang, res.YuShu)
}
