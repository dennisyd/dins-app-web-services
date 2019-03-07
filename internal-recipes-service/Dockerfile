FROM golang:alpine as builder

RUN mkdir /internal-recipes-service
WORKDIR /internal-recipes-service
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git protobuf nodejs nodejs-npm
RUN npm i -g protobufjs
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/twitchtv/twirp/protoc-gen-twirp
RUN go get go.larrymyers.com/protoc-gen-twirp_typescript

RUN go mod download
COPY . .

# Generate go proto files using protoc
RUN protoc --proto_path=proto:. \
    --twirp_out=proto/. --go_out=proto/. \
    --twirp_typescript_out=version=v6:proto/. \
    --twirp_typescript_out=library=pbjs:proto/. internal-recipes-service.proto

# Generate typescript proto files using protobuf.js
RUN pbjs -t static-module -w commonjs -o proto/internal-recipes-service.pb.js proto/internal-recipes-service.proto
RUN pbts --no-comments -o proto/internal-recipes-service.pb.d.ts proto/internal-recipes-service.pb.js

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /service
WORKDIR /service

COPY --from=builder /internal-recipes-service/internal-recipes-service .

# Pull proto files from container we labeled as builder
COPY --from=builder /internal-recipes-service/proto/internal-recipes-service.pb.go .
COPY --from=builder /internal-recipes-service/proto/internal-recipes-service.twirp.go .
COPY --from=builder /internal-recipes-service/proto/internal-recipes-service.ts .
COPY --from=builder /internal-recipes-service/proto/twirp.ts .
COPY --from=builder /internal-recipes-service/proto/internal-recipes-service.pb.d.ts .
COPY --from=builder /internal-recipes-service/proto/internal-recipes-service.pb.js .
COPY --from=builder /internal-recipes-service/proto/internal-recipes-service.twirp.ts .

CMD ["./internal-recipes-service"]