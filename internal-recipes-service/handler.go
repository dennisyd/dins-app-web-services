package main

import (
	"context"
	"fmt"

	pb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

type handler struct{}

func (h *handler) CreateRecipe(ctx context.Context, in *pb.Recipe, out *pb.Response) error {
	fmt.Println("Recipe created")
	return nil
}

func (h *handler) GetRecipes(ctx context.Context, in *pb.GetRequest, out *pb.Response) error {
	fmt.Println("Get All Recipes")
	return nil
}
