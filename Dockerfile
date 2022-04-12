FROM golang:1.18

WORKDIR /auth-service

COPY . .

RUN ["go", "mod", "tidy"]
RUN ["go", "mod", "download"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"
