package main

import (
	"context"
	"fmt"

	v1 "github.com/chyiyaqing/gcat/pkg/proto/demo/v1"
)

type BasicImpl struct {
	v1.UnimplementedMyServiceServer
}

func (BasicImpl) CreateUser(_ context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	// Dumb echo
	return &v1.CreateUserResponse{
		Id:            req.Id,
		Name:          req.Name,
		Age:           req.Age,
		PaidPlan:      req.PaidPlan,
		CreatedAt:     req.CreatedAt,
		Status:        req.Status,
		Subscriptions: req.Subscriptions,
		Email:         req.Email,
	}, nil
}

type VTCodecV1 struct{}

type vtprotoMessage interface {
	MarshalVT() ([]byte, error)
	UnmarshalVT([]byte) error
	// Name() string
}

func (VTCodecV1) Marshal(v interface{}) ([]byte, error) {
	vt, ok := v.(vtprotoMessage)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T", v)
	}
	return vt.MarshalVT()
}

func (VTCodecV1) Unmarshal(data []byte, v interface{}) error {
	vt, ok := v.(vtprotoMessage)
	if !ok {
		return fmt.Errorf("failed to unmarshal; message is %T", v)
	}
	return vt.UnmarshalVT(data)
}

func (VTCodecV1) Name() string {
	return "vtproto-v1"
}
