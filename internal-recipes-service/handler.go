package main

import (
	"context"
	"fmt"

	pb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

type handler struct{}

// CreateRecipe creates a recipe
func (h *handler) CreateRecipe(ctx context.Context, recipe *pb.Recipe) (*pb.Response, error) {
	fmt.Println("Received Create Req", recipe)
	return &pb.Response{}, nil
}

// GetRecipes gets all recipes
func (h *handler) GetRecipes(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	fmt.Println("Received Get All Req")
	return &pb.Response{}, nil
}
