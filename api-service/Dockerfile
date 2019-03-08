FROM golang:alpine as builder

RUN mkdir -p /lasagna-msa/api-service
WORKDIR /lasagna-msa
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git

RUN go mod download
WORKDIR /lasagna-msa/api-service
COPY api-service/. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /api-service
WORKDIR /api-service

COPY --from=builder /lasagna-msa/api-service/api-service .

CMD ["./api-service"]