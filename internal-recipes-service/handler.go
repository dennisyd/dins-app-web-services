package main

import (
	"context"

	pb "github.com/dins-app/web-services/proto/internal-recipes-service"
)

type handler struct{}

func (h *Handler) Create(ctx context.Context, r *pb.Recipe) (res *pb.Response, err error) {
	res.Created = true
	res.Recipe = r
	return res, err
}

func (h *Handler) Get(ctx context.Context, req *pb.Request) (res *pb.Response, err error) {
	recipes := []*pb.Recipe{
		&pb.Recipe{
			Name:        "test",
			Description: "hello this is a test",
		},
	}
	res.Recipes = recipes
	return res, err
}
