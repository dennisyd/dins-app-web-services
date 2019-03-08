build:
	protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gogoslick_out=proto/. --micro_out=proto/. \
    --go_out=proto/. identity-service.proto
run:
	docker run -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns identity-service