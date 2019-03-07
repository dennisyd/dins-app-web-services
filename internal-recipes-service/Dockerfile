FROM golang:alpine as builder

RUN mkdir /internal-recipes-service
WORKDIR /internal-recipes-service
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git protobuf
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/micro/protoc-gen-micro

RUN go mod download
COPY . .

# Generate go proto files using protoc
RUN protoc --proto_path=proto:. --micro_out=proto/. \
    --go_out=proto/. internal-recipes-service.proto

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /internal-recipes-service
WORKDIR /internal-recipes-service

COPY --from=builder /internal-recipes-service/internal-recipes-service .

# Pull proto files from container we labeled as builder
COPY --from=builder /internal-recipes-service/proto/internal-recipes-service.pb.go .
COPY --from=builder /internal-recipes-service/proto/internal-recipes-service.micro.go .

CMD ["./internal-recipes-service"]