build:
	docker build -t internal-recipes-service .
	rm -f proto/*.go
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.pb.go > proto/internal-recipes-service.pb.go
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.micro.go > proto/internal-recipes-service.micro.go
run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns internal-recipes-service