FROM golang:alpine as builder

RUN mkdir /identity-service
WORKDIR /identity-service
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git protobuf
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/micro/protoc-gen-micro
RUN go get github.com/qianlnk/protobuf
RUN GO111MODULE=off go get github.com/golang/protobuf/ptypes
RUN ls $GOPATH/src/

RUN go mod download
COPY . .

# Generate go proto files using protoc
RUN protoc --proto_path=proto:$GOPATH/src --micro_out=proto/. \
    --go_out=plugins=grpc+retag:proto/. identity-service.proto

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /identity-service
WORKDIR /identity-service

COPY --from=builder /identity-service/identity-service .

# Pull proto files from container we labeled as builder
COPY --from=builder /identity-service/proto/identity-service.pb.go .
COPY --from=builder /identity-service/proto/identity-service.micro.go .

CMD ["./identity-service"]