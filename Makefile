ifndef ($(GOPATH))
	GOPATH = $(HOME)/go
endif
generate:
	# Generating internal-recipes-service
	@rm -f proto/internal-recipes-service/internal-recipes-service.pb.go
	@protoc --proto_path=proto/internal-recipes-service:$(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
 		--gogoslick_out=plugins=grpc:proto/internal-recipes-service/. proto/internal-recipes-service/internal-recipes-service.proto
	# Generating identity-service
	@rm -f proto/identity-service/identity-service.pb.go
	@protoc --proto_path=proto/identity-service:$(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
 		--gogoslick_out=plugins=grpc:proto/identity-service/. proto/identity-service/identity-service.proto