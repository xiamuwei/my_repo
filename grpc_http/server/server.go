package main

import (
    "context"
    _ "fmt"
    "log"
    "net"

    pb "gprc_helloworld/pb"
    "google.golang.org/grpc"
)

// 定义 gRPC 服务
type server struct {
    pb.UnimplementedGreeterServer
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Printf("Received: %v", req.GetName())
    return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
    // 启动监听器
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 创建 gRPC 服务器
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})

    // 启动 gRPC 服务器
    log.Printf("Server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}