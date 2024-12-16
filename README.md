# web-services

## Getting Started

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
