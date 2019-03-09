package main

import (
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

func main() {

	// create new service
	service := micro.NewService(
		micro.Name("lasagna.internal.recipes.service"),
		micro.Version("v1"),
	)

	// parse command line flags
	service.Init()

	// register handler
	pb.RegisterRecipesHandler(service.Server(), &handler{})

	// run service
	if err := service.Run(); err != nil {
		log.Fatalln("Unable to start internal recipes service", err)
	}
}
