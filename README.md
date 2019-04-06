# web-services

Backend microservices for dins.app

- Identity Service (Login/Signup/3rd party Oauth)
- Internal Recipe Service (CRUD recipes in our database)
- External Recipe Service (Queries external services for data and returns data to client; for example, add a recipe from a url — would scrape info from the url and return)
- Recommendation Service (Takes a user id and gets user info from identity service, then uses user preferences to query internal recipe service for weighted results)
- Api Service (Exposes a REST api of available enpoints to interact with gRPC services above)
  - GET `/v1/recipes` -> Get all recipes from the database
  - POST `/v1/recipes` -> Create a new recipe in the database

## Getting Started

- Docker and docker-compose is required for development. Get Docker [here](https://www.docker.com/get-started).

To develop locally simply clone the repo and run the following commands to start the services:

```bash
docker swarm init    # to make your current machine a swarm manager
docker-compose build # builds the services stack
docker-compose up    # starts the services, use -d flag to run in background
```

To bring down the services:

```bash
docker-compose down
```

To deploy using stack deployment (will pull images from docker repo instead of building):

```bash
docker stack deploy --compose-file=docker-compose.yml
```

If you have access to push images to the repo, you can push new builds like so:

```bash
docker-compose push
```
