BINARY_NAME=qp
build:
	go build -o bin/$(BINARY_NAME) -v

run:build
	./bin/$(BINARY_NAME)

phony:build run