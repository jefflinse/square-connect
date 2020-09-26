all: clean generate build

clean:
	go clean -i -testcache ./...
	rm -rf ./client
	rm -rf ./models

generate: square.json
	swagger generate client -f square.json

build:
	go get -u -f ./...
	go build ./...
