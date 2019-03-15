package main

import (
	"context"
	"log"

	pb "github.com/dins-app/web-services/proto/internal-recipes-service"
)

type Handler struct{}

func (h *Handler) Create(ctx context.Context, r *pb.Recipe) (*pb.Response, error) {
	res := pb.Response{}
	errors := DB.Create(&r).GetErrors()
	for _, err := range errors {
		res.Errors = append(res.Errors, err.Error())
	}
	if len(errors) == 0 {
		res.Recipe = r
	}
	log.Print("Received request")
	return &res, nil
}

func (h *Handler) Get(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res := pb.Response{}
	recipes := []*pb.Recipe{}
	DB.Find(&recipes)
	res.Recipes = recipes
	log.Print("Received request")
	return &res, nil
}
