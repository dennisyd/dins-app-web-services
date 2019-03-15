package main

import (
	"context"
	"log"

	pb "github.com/dins-app/web-services/proto/internal-recipes-service"
)

type Handler struct{}

func (h *Handler) Create(ctx context.Context, r *pb.Recipe) (*pb.Response, error) {
	res := pb.Response{}
	DB.Create(&r)
	res.Recipe = r
	log.Print("Received request")
	return &res, nil
}

func (h *Handler) Get(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res := pb.Response{}
	DB.Find(&res.Recipes)
	log.Print("Received request")
	return &res, nil
}
