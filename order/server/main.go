package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/echo-music/go-k8s-test/order/proto" // 替换为你的模块路径
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *server) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	log.Printf("Received: %v", in.GetName())
	hostName, _ := os.Hostname()
	return &pb.CreateOrderReply{Message: "Hello " + in.GetName() + ",host:" + hostName}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
