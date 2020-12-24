package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 客户机传参
type Params struct {
	Width, Height int
}

func main() {
	// 1. 连接远程 rpc 服务
	coon, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	// 2. 调用方法
	// 求矩形面积
	ret := 0
	err2 := coon.Call("Rect.Area", Params{5, 10}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积:", ret)
	// 求矩形周长
	err3 := coon.Call("Rect.Perimeter", Params{5,10}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长:", ret)
}
