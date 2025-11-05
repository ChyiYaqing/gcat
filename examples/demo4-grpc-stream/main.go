package main

import (
	"crypto/rand"
	"io"

	v1 "github.com/chyiyaqing/gcat/pkg/proto/demo/v1"
	"google.golang.org/grpc"
)

type BasicImpl struct {
	v1.UnimplementedMyServiceServer
}

type EfficientImpl struct {
	v1.UnimplementedMyServiceServer
}

func (BasicImpl) Put(stream grpc.ClientStreamingServer[v1.Chunk, v1.PutResult]) error {
	for {
		chunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&v1.PutResult{})
			}
			return err
		}
		// do something with received chunk
		_ = chunk
	}
}

func (EfficientImpl) Put(stream grpc.ClientStreamingServer[v1.Chunk, v1.PutResult]) error {
	for {
		chunk := v1.ChunkFromVTPool()
		chunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&v1.PutResult{})
			}
			return err
		}
		// do something with received chunk
		_ = chunk
	}
}

func (BasicImpl) Get(req *v1.GetRequest, stream grpc.ServerStreamingServer[v1.Chunk]) error {
	data := make([]byte, 16*1024)
	_, _ = rand.Read(data)
	for i := 0; i < 1024; i++ {
		chunk := &v1.Chunk{Data: data}
		stream.Send(chunk)
	}
	return nil
}
