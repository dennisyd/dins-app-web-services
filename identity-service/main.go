package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	micro "github.com/micro/go-micro"
	pb "github.com/team-morpheus/lasagna-msa/identity-service/proto"
)

func main() {
	// Create db connection and defer db to close when method ends
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("could not connect to DB: %v", err)
	}

	// AutoMigrate database to match pb User struct
	db.AutoMigrate(&pb.User{})

	// set UserRepository to use db conection
	repo := &UserRepository{db}

	// set TokenService to use repo
	tokenService := &TokenService{repo}

	// Create new micro service
	srv := micro.NewService(
		micro.Name("lasagna.identity.service"),
		micro.Version("v1"),
	)

	// Initialize service with command line args
	srv.Init()

	// Register service handler
	pb.RegisterIdentityServiceHandler(srv.Server(), &handler{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
