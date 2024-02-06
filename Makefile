APP_NAME := "echo-boilerplate"

build:
	go build -o ./bin/$(APP_NAME) ./cmd/server

dev: build
	./bin/$(APP_NAME)