build:
	docker build -t identity-service .
	rm -f proto/*.go
	docker run --entrypoint cat identity-service identity-service.pb.go > proto/identity-service.pb.go
	docker run --entrypoint cat identity-service identity-service.micro.go > proto/identity-service.micro.go
	docker run --entrypoint cat identity-service identity-service.pb.gorm.go > proto/identity-service.pb.gorm.go
run:
	docker run -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns identity-service