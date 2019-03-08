protoc --proto_path=proto:$GOPATH/src --micro_out=proto/. \
    --go_out=proto/. \
    --gorm_out="engine=postgres:proto/." identity-service.proto