package main

import (
	"net/http"

	pb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

func main() {
	server := pb.NewInternalRecipesServiceServer(&handler{}, nil)

	http.ListenAndServe(":8080", server)
}
