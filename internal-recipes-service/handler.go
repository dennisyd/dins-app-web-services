package main

import (
	"context"
	"fmt"

	pb "github.com/dins-app/web-services/proto/internal-recipes-service"
)

// Handler is our service handler
type Handler struct{}

// Create creates a recipes in the db
func (h *Handler) Create(ctx context.Context, r *pb.Recipe) (*pb.Response, error) {
	res := pb.Response{}

	// create recipe in db
	err := DB.Insert(r)
	if err != nil {
		res.Errors = append(res.Errors, fmt.Sprintf("unable to create new recipe: %s", err.Error()))
		return &res, nil
	}

	// set res.Recipe to created recipe and return res
	res.Recipe = r
	return &res, nil
}

// Get gets all recipes from the db
func (h *Handler) Get(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res := pb.Response{}
	recipes := []*pb.Recipe{}

	// get all recipes from db
	err := DB.Find(nil).All(&recipes)
	if err != nil {
		res.Errors = append(res.Errors, fmt.Sprintf("unable to find recipes: %s", err.Error()))
		return &res, nil
	}

	// set res.Recipes to found recipes and return res
	res.Recipes = recipes
	return &res, nil
}
