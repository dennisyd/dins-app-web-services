FROM golang:alpine as builder

RUN mkdir -p /web-services/api-service
WORKDIR /web-services
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git

WORKDIR /web-services/api-service
COPY api-service/. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /api-service
WORKDIR /api-service

COPY --from=builder /web-services/api-service/api-service .

CMD ["./api-service"]