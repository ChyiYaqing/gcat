package main

import (
	"context"
	"net"
	"testing"

	demov1 "github.com/chyiyaqing/gcat/pkg/proto/demo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
	lis, _ := net.Listen("tcp", "localhost:1234")
	defer lis.Close()
	grpcServer := grpc.NewServer()
	demov1.RegisterMyServiceServer(grpcServer, &BasicImpl{})
	go func() {
		grpcServer.Serve(lis)
	}()

	// Setup client
	cc, _ := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := demov1.NewMyServiceClient(cc)

	// benchmark
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.CreateUser(ctx, &demov1.CreateUserRequest{
			Id:            1234,
			Name:          "my-name",
			Age:           42,
			PaidPlan:      true,
			CreatedAt:     567153,
			Status:        "PREMIUM",
			Subscriptions: []string{"sub1", "sub2"},
			Email:         "name@gmail.com",
		})
	}
}
