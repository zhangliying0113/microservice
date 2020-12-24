package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 自己实现 rpc 程序，服务端接收2个参数，可以做乘法运算，也可以做商和余数的运算

// 运算结构体，用于注册 rpc 服务
type Operation struct {
}

// 声明传参结构体
type Params struct {
	A, B int
}

// 返回给客户端的结果
type Response struct {
	ChengJi int
	Shang   int
	YuShu   int
}

// 乘法实现
func (o *Operation) Multi(p Params, res *Response) error {
	res.ChengJi = p.A * p.B
	return nil
}

// 商和余数实现
func (o *Operation) Division(p Params, res *Response) error {
	if p.B == 0 {
		return errors.New("除数不能为0")
	}
	// 除法
	res.Shang = p.A / p.B
	// 取模
	res.YuShu = p.A % p.B
	return nil
}

// 主函数
func main() {
	// 结构体实例化并注册 rpc 服务
	opt := new(Operation)
	rpc.Register(opt)
	// 监听服务
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}

	// 循环监听服务
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
