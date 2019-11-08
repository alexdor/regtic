.PHONY: generate build clean deploy

generate:
	go generate ./...
	go test ./models

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

functions := $(shell find lambda-functions -name \*main.go | awk -F'/' '{print $$2}')

dev:
	reflex -d fancy -r '\.go$$' -s -- sh -c 'echo "\n\nRestarting server\n\n"; make build';

build: ## Build golang binaries
	export GO111MODULE=on
	@for function in $(functions) ; do \
		go mod tidy ;\
		env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$$function lambda-functions/$$function/main.go ; \
	done
