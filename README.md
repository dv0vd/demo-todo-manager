# TODO Manager

## Live demo: https://dv0vd.dev/demo/todo-manager/swagger/index.html
## Podman/Docker image: https://hub.docker.com/r/dv0vd/demo-todo-manager

REST API for TODO management app developed using Go and the service-repository pattern. 
## Getting started

### Podman Compose (includes a built-in database)
1) Configure the `.env` file.
2) Check tests `make test`.
2) Run the command `make tidy`.
3) Run the command `make build`.
4) Start the project with: `make start`.
5) To stop or restart the project, use `make stop` and `make restart`, respectively.
6) Access the Swagger UI at `/swagger/index.html`. To generate the Swagger documentation, run the following command: `make swagger`.

### Podman image (use your own Postgres)
Run the container with your Postgres:
```
podman run \
	-d \
	-e DB_HOST=<your-postgres-host> \
	-e DB_PORT=<your-postgres-port> \
	-e DB_USER=<your-postgres-user> \
	-e DB_PASSWORD=<your-postgres-password> \
	-e DB_NAME=<database-name> \
	-e HOST=<your-host> \
	--name demo-todo-manager \
	--restart unless-stopped \
	--memory=128M \
	--cpus=0.25  \
	docker.io/dv0vd/demo-todo-manager
```

