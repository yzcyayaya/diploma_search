package client

import (
	"diploma_search/biz/server/rpc"
	"log"

	"google.golang.org/grpc"
)

var (
	Conn *grpc.ClientConn
	C    rpc.PersonClient
)

func init() {
	// 连接grpc服务器
	var err error
	Conn, err = grpc.Dial("127.0.0.1:9110", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}


	// 初始化Greeter服务客户端
	C = rpc.NewPersonClient(Conn)
}
