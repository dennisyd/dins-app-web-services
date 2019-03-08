ifndef ($(GOPATH))
	GOPATH = $(HOME)/go
endif
generate:
	@rm -f proto/identity-service.pb.go
	@rm -f proto/identity-service.micro.go
	@protoc --proto_path=proto:$(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
 		--micro_out=proto/. --gogoslick_out=proto/. identity-service.proto
	@echo Identity Service Protobufs Generated!