package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 算数运算请求结构体
type ArithRequest4 struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse4 struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := ArithRequest4{9, 2}
	var res ArithResponse4

	err = conn.Call("Arith.Multiply3", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arith.Divide3", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
