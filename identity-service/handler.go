package main

import (
	"context"
	"fmt"

	pb "github.com/dins-app/web-services/proto/identity-service"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

// Handler is our service handler
type Handler struct{}

// Create creates a user in the db
func (h *Handler) Create(ctx context.Context, u *pb.User) (*pb.Response, error) {
	res := pb.Response{}

	// hash password
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return &res, nil
	}
	u.Password = string(password)

	// create user in db
	err = DB.Insert(u)
	if err != nil {
		u.Password = ""
		res.Errors = append(res.Errors, fmt.Sprintf("unable to create new user: %s", err.Error()))
		return &res, nil
	}

	// set new user's password to "" so it isn't exposed
	u.Password = ""

	// set res.User to new user and return res
	res.User = u
	return &res, nil
}

// Get gets a user from the db by request.query
func (h *Handler) Get(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	res := pb.Response{}
	user := pb.User{}

	// get user from db by r.Query
	err := DB.Find(r.Query).One(&user)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return &res, nil
	}

	// set res.User to found user and return res
	res.User = &user
	return &res, nil
}

// Auth authorizes a user by email and password
func (h *Handler) Auth(ctx context.Context, u *pb.User) (*pb.Response, error) {
	res := pb.Response{}
	user := pb.User{}

	// get user from db by email
	err := DB.Find(bson.M{"email": u.Email}).One(&user)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return &res, nil
	}

	// compare found user password hash to hash of given user
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return &res, nil
	}

	// set new user's password to "" so it isn't exposed
	u.Password = ""

	// set res.User to new user and return res
	res.User = u
	return &res, nil
}
