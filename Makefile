.DEFAULT_GOAL := help

# init:
# 	podman run \
#   --rm \
#   -v ./:/app \
#   docker.io/node:20.18.1-bookworm \
#   sh -c 'cd /app && npm ci --verbose'

tidy:
	podman run \
  --rm \
  -v ./:/app \
  docker.io/golang:1.24.2-bookworm \
  sh -c 'cd /app && go mod tidy && go mod vendor'

build:
	podman run \
  --rm \
  -v ./:/app \
  docker.io/golang:1.24.2-bookworm \
  sh -c 'cd /app && go build -mod=vendor -v -o ./bin/app ./cmd'

test:
	podman run \
  --rm \
  -v ./:/app \
  docker.io/golang:1.24.2-bookworm \
  sh -c 'cd /app && go test ./internal/tests/... -v'

# build-dev:
# 	podman run \
#   --rm \
#   -v ./:/app \
#   docker.io/node:20.18.1-bookworm \
#   sh -c 'cd /app && npm run build-dev --verbose'

start:
	podman-compose up -d

start-app:
	podman-compose up -d app

stop-app:
	podman-compose down app

restart-app: stop-app start-app

logs-app:
	podman logs -f todo-manager_app

enter-app:
	podman exec -it todo-manager_app bash

stop:
	podman-compose down

restart: stop start


GREEN='\033[1;32m'
WHITE='\033[1;37m'
help:
# @echo -e ${GREEN}init'             '${WHITE}— initialize the project
# @echo -e ${GREEN}start'            '${WHITE}— start the project
# @echo -e ${GREEN}start-app'        '${WHITE}— start the project without a database
# @echo -e ${GREEN}stop'             '${WHITE}— stop the project
# @echo -e ${GREEN}restart'          '${WHITE}— restart the project



# migrate create -ext sql -dir migrations -sec create_users_table