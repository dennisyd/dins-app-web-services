package main

import (
	"context"
	"log"

	pb "github.com/dins-app/web-services/proto/identity-service"
	"golang.org/x/crypto/bcrypt"
)

// Handler is our service handler
type Handler struct{}

// Create creates a user in the db
func (h *Handler) Create(ctx context.Context, u *pb.User) (*pb.Response, error) {
	log.Print("Received request")
	res := pb.Response{}
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return &res, nil
	}
	u.Password = string(password)
	errors := DB.Create(&u).GetErrors()
	for _, err := range errors {
		res.Errors = append(res.Errors, err.Error())
	}
	if len(errors) == 0 {
		u.Password = ""
		res.User = u
	}
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
