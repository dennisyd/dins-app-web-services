FROM golang:alpine as builder

RUN mkdir /internal-recipes-service
WORKDIR /internal-recipes-service
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /internal-recipes-service
WORKDIR /internal-recipes-service

COPY --from=builder /internal-recipes-service/internal-recipes-service .

CMD ["./internal-recipes-service"]