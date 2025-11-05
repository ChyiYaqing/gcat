package main

import (
	"context"
	"net"
	"os"
	"testing"

	v1 "github.com/chyiyaqing/gcat/pkg/proto/demo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
)

func init() {
	if os.Getenv("VTCODEC") != "" {
		encoding.RegisterCodec(VTCodecV1{})
	}
}

// 1. Client marshalls request
// 2. Client sends request
// 3. Server receives request
// 4. Server unmarshalls request
// 5. Server executes handler
// 6. Server marshalls response
// 7. Server sends response
// 8. Client receives response
// 9. Client unmarshalls response

func BenchmarkUnary(b *testing.B) {
	// Error handling elided
	// Setup server
	lis, err := net.Listen("tcp", "localhost:0") // 使用随机端口
	if err != nil {
		b.Fatal(err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	v1.RegisterMyServiceServer(grpcServer, &BasicImpl{})

	// 等待服务器启动
	go func() {
		grpcServer.Serve(lis)
	}()
	defer grpcServer.Stop()

	// Setup client
	cc, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		b.Fatal(err)
	}
	defer cc.Close()

	client := v1.NewMyServiceClient(cc)
	ctx := context.Background()

	req := &v1.CreateUserRequest{
		Id:            1234,
		Name:          "my-name",
		Age:           42,
		PaidPlan:      true,
		CreatedAt:     567153,
		Status:        "PREMIUM",
		Subscriptions: []string{"sub1", "sub2"},
		Email:         "name@gmail.com",
	}

	// 预热: 建立连接
	if _, err := client.CreateUser(ctx, req); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	// benchmark
	for i := 0; i < b.N; i++ {
		_, err = client.CreateUser(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// 测试标准 protobuf
func BenchmarkMarshalProto(b *testing.B) {
	req := &v1.CreateUserRequest{
		Id:            1234,
		Name:          "my-name",
		Age:           42,
		PaidPlan:      true,
		CreatedAt:     567153,
		Status:        "PREMIUM",
		Subscriptions: []string{"sub1", "sub2"},
		Email:         "name@gmail.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(req)
		if err != nil {
			b.Fatal(err)
		}
	}

}

// 测试 vtproto
func BenchmarkMarshalVT(b *testing.B) {
	req := &v1.CreateUserRequest{
		Id:            1234,
		Name:          "my-name",
		Age:           42,
		PaidPlan:      true,
		CreatedAt:     567153,
		Status:        "PREMIUM",
		Subscriptions: []string{"sub1", "sub2"},
		Email:         "name@gmail.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := req.MarshalVT()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	req := &v1.CreateUserRequest{
		Id:            1234,
		Name:          "my-name",
		Age:           42,
		PaidPlan:      true,
		CreatedAt:     567153,
		Status:        "PREMIUM",
		Subscriptions: []string{"sub1", "sub2"},
		Email:         "name@gmail.com",
	}
	codec := VTCodecV1{}
	data, _ := codec.Marshal(req)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := &v1.CreateUserRequest{}
		err := codec.Unmarshal(data, dst)
		if err != nil {
			b.Fatal(err)
		}
	}
}
