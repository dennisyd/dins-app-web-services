package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	pb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

func main() {
	mux := http.NewServeMux()

	// register server with handler
	server := pb.NewInternalRecipesServiceServer(&handler{}, nil)

	mux.Handle(pb.InternalRecipesServicePathPrefix, server)

	handle := cors.AllowAll().Handler(mux)

	addr := fmt.Sprintf(":%d", 8080)
	log.Printf("server listening at: %s", addr)

	err := http.ListenAndServe(addr, handle)
	if err != nil {
		log.Fatalln(err)
	}
}
