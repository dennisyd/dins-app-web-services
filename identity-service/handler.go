package main

import (
	"context"

	pb "github.com/dins-app/web-services/proto/identity-service"
)

// Handler is our service handler
type Handler struct{}

// Create creates a user in the db
func (h *Handler) Create(ctx context.Context, u *pb.User) (*pb.Response, error) {
	res := pb.Response{}
	return &res, nil
}

// Get gets a user from the db by request.query
func (h *Handler) Get(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	res := pb.Response{}
	return &res, nil
}

// Auth authorizes a user by email and password
func (h *Handler) Auth(ctx context.Context, u *pb.User) (*pb.Response, error) {
	res := pb.Response{}
	return &res, nil
}
