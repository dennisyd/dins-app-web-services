FROM golang:alpine as builder

RUN mkdir /identity-service
WORKDIR /identity-service
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /identity-service
WORKDIR /identity-service

COPY --from=builder /identity-service/identity-service .

CMD ["./identity-service"]