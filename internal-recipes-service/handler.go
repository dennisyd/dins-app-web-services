package main

import (
	"context"

	pb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

type handler struct{}

func (h *handler) Create(ctx context.Context, r *pb.Recipe, res *pb.Response) error {
	res.Created = true
	res.Recipe = r
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.Request, res *pb.Response) error {
	recipes := []*pb.Recipe{
		&pb.Recipe{
			Name:        "test",
			Description: "hello this is a test",
		},
	}
	res.Recipes = recipes
	return nil
}
