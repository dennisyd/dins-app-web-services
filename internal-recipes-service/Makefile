build:
	docker build -t internal-recipes-service .
	rm -f proto/*.go
	rm -f proto/*.ts
	rm -f proto/*.js
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.twirp.go > proto/internal-recipes-service.twirp.go
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.pb.go > proto/internal-recipes-service.pb.go
	docker run --entrypoint cat internal-recipes-service twirp.ts > proto/twirp.ts
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.ts > proto/internal-recipes-service.ts
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.pb.js > proto/internal-recipes-service.pb.js
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.pb.d.ts > proto/internal-recipes-service.pb.d.ts
	docker run --entrypoint cat internal-recipes-service internal-recipes-service.twirp.ts > proto/internal-recipes-service.twirp.ts
run:
	docker run -p 8080:8080 internal-recipes-service