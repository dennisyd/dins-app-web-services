build:
	docker build -t internal-recipes-service .
	rm -f proto/*.go
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.micro.go > proto/internal-recipes-service.micro.go
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.pb.go > proto/internal-recipes-service.pb.go
run:
	docker run -p 9201:9201 internal-recipes.service