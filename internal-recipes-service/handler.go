package main

import (
	"context"

	pb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

type handler struct{}

func (h *handler) CreateRecipe(ctx context.Context, recipe *pb.Recipe) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (h *handler) GetRecipes(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
