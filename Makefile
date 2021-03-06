all: clean generate build

clean:
	go clean -i -testcache ./...

generate: square.json
	rm -rf ./clients
	rm -rf ./models
	swagger generate client -f square.json

build:
	go get -u -d ./...
	go build ./...
