// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: internal-recipes-service.proto

/*
Package lasagna_api_recipes is a generated protocol buffer package.

It is generated from these files:
	internal-recipes-service.proto

It has these top-level messages:
	Recipe
	Request
	Response
*/
package lasagna_api_recipes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Recipes service

type RecipesService interface {
	Create(ctx context.Context, in *Recipe, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type recipesService struct {
	c    client.Client
	name string
}

func NewRecipesService(name string, c client.Client) RecipesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "lasagna.api.recipes"
	}
	return &recipesService{
		c:    c,
		name: name,
	}
}

func (c *recipesService) Create(ctx context.Context, in *Recipe, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Recipes.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesService) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Recipes.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Recipes service

type RecipesHandler interface {
	Create(context.Context, *Recipe, *Response) error
	Get(context.Context, *Request, *Response) error
}

func RegisterRecipesHandler(s server.Server, hdlr RecipesHandler, opts ...server.HandlerOption) error {
	type recipes interface {
		Create(ctx context.Context, in *Recipe, out *Response) error
		Get(ctx context.Context, in *Request, out *Response) error
	}
	type Recipes struct {
		recipes
	}
	h := &recipesHandler{hdlr}
	return s.Handle(s.NewHandler(&Recipes{h}, opts...))
}

type recipesHandler struct {
	RecipesHandler
}

func (h *recipesHandler) Create(ctx context.Context, in *Recipe, out *Response) error {
	return h.RecipesHandler.Create(ctx, in, out)
}

func (h *recipesHandler) Get(ctx context.Context, in *Request, out *Response) error {
	return h.RecipesHandler.Get(ctx, in, out)
}
