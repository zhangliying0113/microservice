package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// 例题：golang 实现 RPC 程序，实现求矩形面积和周长

// 传参结构体
type Params struct {
	Width, Height int
}

// 矩形结构体
type Rect struct {
	
}

// 服务端求矩形面积函数
func (r *Rect)Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

// 服务端求矩形周长函数
func (r *Rect)Perimeter(p Params, ret *int) error {
	*ret = 2 * (p.Width + p.Height)
	return nil
}

func main()  {
	// 1. 实例化 Rect 矩形结构体对象
	rect := new(Rect)
	// 为 rect 对象注册 rpc 服务
	rpc.Register(rect)
	// 2. 将服务绑定到 http 协议上
	rpc.HandleHTTP()
	// 3. 监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
