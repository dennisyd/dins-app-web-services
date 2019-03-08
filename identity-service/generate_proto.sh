protoc --proto_path=proto:$GOPATH/src:$GOPATH/src/github.com/gogo/protobuf/protobuf \
 --micro_out=proto/. --gogoslick_out=proto/. identity-service.proto