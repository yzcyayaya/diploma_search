package client

import (
	"diploma_search/biz/server/rpc"
	"flag"
	"log"

	"google.golang.org/grpc"
)

var (
	Conn    *grpc.ClientConn
	C       rpc.PersonClient
	grpcUrl string
)

func init() {
	// 连接grpc服务器

	//没有flage则从yaml读取meili配置
	flag.StringVar(&grpcUrl, "grpc_url", "0.0.0.0:9110", "grpc_url: grpcurl:9110")
	flag.Parse()
	var err error
	Conn, err = grpc.Dial(grpcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 初始化Greeter服务客户端
	C = rpc.NewPersonClient(Conn)
}
