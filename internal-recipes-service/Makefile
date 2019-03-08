ifndef ($(GOPATH))
	GOPATH = $(HOME)/go
endif
generate:
	@rm -f proto/internal-recipes-service.pb.go
	@rm -f proto/internal-recipes-service.micro.go
	@protoc --proto_path=proto:$(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
 		--micro_out=proto/. --gogoslick_out=proto/. internal-recipes-service.proto
	@echo Internal Recipes Service Protobufs Generated!