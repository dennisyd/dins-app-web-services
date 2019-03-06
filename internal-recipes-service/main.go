package main

import (
	"log"

	micro "github.com/micro/go-micro"
	proto "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

func main() {

	// create service
	service := micro.NewService(
		micro.Name("lasagna.srv.internal.recipes.service"),
		micro.Version("latest"),
	)

	// parse command line flags
	service.Init()

	// register handler (generated from protobuf definition)
	proto.RegisterInternalRecipesServiceHandler(service.Server(), &handler{})

	// start the service
	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}
