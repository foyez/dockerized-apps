# Node Application

## Build and run

```sh
# build and run the containers
docker compose -f docker-compose.yaml up -d --build
docker compose --env-file dev.env -f docker-compose.yaml -f docker-compose.dev.yaml up -d --build

# stop and remove the containers
docker compose -f docker-compose.yaml -f docker-compose.dev.yaml down
```
