FROM golang:alpine as builder

RUN mkdir /identity-service
WORKDIR /identity-service
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git protobuf build-base
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/micro/protoc-gen-micro
RUN GO111MODULE=off go get -d github.com/infobloxopen/protoc-gen-gorm
RUN cd $GOPATH/src/github.com/infobloxopen/protoc-gen-gorm && make vendor && make install && cd /identity-service

RUN go mod download
COPY . .

RUN ls $GOPATH/pkg/mod/github.com
# Generate go proto files using protoc
RUN protoc -I. -I$GOPATH/src -I./proto --micro_out=proto/. \
    --go_out=proto/. \
    --gorm_out="engine=postgres:proto/." identity-service.proto

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /identity-service
WORKDIR /identity-service

COPY --from=builder /identity-service/identity-service .

# Pull proto files from container we labeled as builder
COPY --from=builder /identity-service/proto/identity-service.pb.go .
COPY --from=builder /identity-service/proto/identity-service.micro.go .
COPY --from=builder /identity-service/proto/identity-service.pb.gorm.go .

CMD ["./identity-service"]