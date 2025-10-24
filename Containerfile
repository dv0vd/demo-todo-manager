FROM docker.io/golang:1.25.1-alpine3.22
WORKDIR /app
COPY ./ /app

RUN apk add --no-cache gettext
RUN go mod tidy && go mod vendor
RUN go build -mod=vendor -v -o ./bin/app ./cmd

ENTRYPOINT ["sh", "entrypoint.sh"]