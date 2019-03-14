package main

import (
	"context"
	"log"

	pb "github.com/dins-app/web-services/proto/internal-recipes-service"
)

type Handler struct{}

func (h *Handler) Create(ctx context.Context, r *pb.Recipe) (*pb.Response, error) {
	res := pb.Response{}
	res.Created = true
	res.Recipe = r
	log.Print("Received request")
	return &res, nil
}

func (h *Handler) Get(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res := pb.Response{}
	recipes := []*pb.Recipe{
		&pb.Recipe{
			Name:        "test",
			Description: "hello this is a test",
		},
	}
	res.Recipes = recipes
	log.Print("Received request")
	return &res, nil
}
