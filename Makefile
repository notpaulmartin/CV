all: build compile-cv

build:
	go build -o bin/main main.go

compile-cv:
	./bin/main
