package main

import (
	"diploma_search/biz/server/rpc"
	"diploma_search/biz/server/service"
	"flag"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var port int
	//命令行接收prot
	flag.IntVar(&port, "port", 9110, "help message for port")
	flag.Parse()

	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	rpc.RegisterPersonServer(s, &service.Person{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	log.Println("grpc sever success: 0.0.0.0:" + strconv.Itoa(port))

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
