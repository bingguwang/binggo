package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-study/client/service"
	pb "grpc-study/server/proto"
	"grpc-study/server/utils"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

/**
这里是TLS双向认证
	也就是客户端有了公钥和服务名后并不能随心所欲的调用服务，服务端对客户端也要进行筛选
*/
func main() {
	flag.Parse()
	fmt.Println(utils.ToJsonString(addr))

	// 双向TLS认证
	opts := utils.GetBothSideTlsClientOpts()

	conn, err := grpc.Dial(
		*addr,
		opts...,
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewScoreServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service.CallStreamBidirectional(ctx, client)

}
