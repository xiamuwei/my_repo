package main

import (
    "context"
    "log"
    "os"
    "time"

    pb "gprc_helloworld/pb"
    "google.golang.org/grpc"
)

func main() {
    // 连接到 gRPC 服务器
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // 创建 Greeter 客户端
    c := pb.NewGreeterClient(conn)

    // 准备请求
    name := "world"
    if len(os.Args) > 1 {
        name = os.Args[1]
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // 调用 SayHello 方法
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMessage())
}