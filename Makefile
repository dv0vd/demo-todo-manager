.DEFAULT_GOAL := help

tidy:
	podman run \
  --rm \
  -v ./:/app \
  docker.io/golang:1.25.1-alpine3.22 \
  sh -c 'cd /app && go mod tidy && go mod vendor'

build:
	podman run \
  --rm \
  -v ./:/app \
  docker.io/golang:1.25.1-alpine3.22 \
  sh -c 'cd /app && go build -mod=vendor -v -o ./bin/app ./cmd'

test:
	podman run \
  --rm \
  -v ./:/app \
  docker.io/golang:1.25.1-alpine3.22 \
  sh -c 'cd /app && go test ./internal/tests/... -v'

start:
	podman-compose up -d

logs-app:
	podman logs -f todo-manager-app

enter-app:
	podman exec -it todo-manager-app sh

stop:
	podman-compose down

restart: stop start

restart-fresh: stop swagger build start logs-app

create-migration-postgres:
	podman run --rm -v ./migrations/postgres:/migrations docker.io/migrate/migrate:v4.19.0 create -ext sql -dir migrations -seq ${m}

swagger:
	podman run --rm -v ./:/app -w /app ghcr.io/swaggo/swag:v1.16.6 init -g ./cmd/main.go

GREEN='\033[1;32m'
WHITE='\033[1;37m'
help:
	@echo -e ${GREEN}tidy'                         '${WHITE}— tidy project's dependencies${RESET}
	@echo -e ${GREEN}build'                        '${WHITE}— build the project${RESET}
	@echo -e ${GREEN}test'                         '${WHITE}— run unit tests${RESET}
	@echo -e ${GREEN}start'                        '${WHITE}— start the project${RESET}
	@echo -e ${GREEN}stop'                         '${WHITE}— stop the project${RESET}
	@echo -e ${GREEN}restart'                      '${WHITE}— restart the project${RESET}
	@echo -e ${GREEN}restart-fresh'                '${WHITE}— restart the project for development${RESET}
	@echo -e ${GREEN}logs-app'                     '${WHITE}— get project's logs${RESET}
	@echo -e ${GREEN}enter-app'                    '${WHITE}— enter the database container${RESET}
	@echo -e ${GREEN}create-migration-postgres'    '${WHITE}— enter the database container${RESET}
	@echo -e ${GREEN}swagger'                      '${WHITE}— generate API docs${RESET}