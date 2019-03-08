protoc --proto_path=proto:$GOPATH/src:$GOPATH/src/github.com/gogo/protobuf/protobuf --gogoslick_out=proto/. --micro_out=proto/. \
 identity-service.proto