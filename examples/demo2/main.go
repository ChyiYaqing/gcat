package main

import (
	"context"

	demov1 "github.com/chyiyaqing/gcat/pkg/proto/demo/v1"
)

type BasicImpl struct {
	demov1.UnimplementedMyServiceServer
}

func (BasicImpl) CreateUser(_ context.Context, req *demov1.CreateUserRequest) (*demov1.CreateUserResponse, error) {
	// Dumb echo
	return &demov1.CreateUserResponse{
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
