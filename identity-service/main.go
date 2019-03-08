package main

import (
	"log"

	micro "github.com/micro/go-micro"
)

func main() {

	// create new service
	service := micro.NewService(
		micro.Name("lasagna.identity.service"),
		micro.Version("v1"),
	)

	// parse command line flags
	service.Init()

	// register handler
	// pb.RegisterInternalRecipesServiceHandler(service.Server(), &handler{})

	// run service
	if err := service.Run(); err != nil {
		log.Fatalln("Unable to start internal recipes service", err)
	}
}
