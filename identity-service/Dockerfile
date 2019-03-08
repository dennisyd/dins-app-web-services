FROM golang:alpine as builder

RUN mkdir -p /lasagna-msa/identity-service
WORKDIR /lasagna-msa
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git

RUN go mod download
WORKDIR /lasagna-msa/identity-service
COPY identity-service/. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /identity-service
WORKDIR /identity-service

COPY --from=builder /lasagna-msa/identity-service/identity-service .

CMD ["./identity-service"]