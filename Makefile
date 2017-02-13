all: build test

deps:
	go get -v -t ./...

test:
	go test -v ./...

build:
	cd cmd/chttp && make -e build
