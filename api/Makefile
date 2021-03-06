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

dev:
	reflex -d fancy -r '\.go$$' -s -- sh -c 'echo "\n\nRestarting server\n\n"; make build';

api:
	go run -tags api --race .

build: ## Build golang binaries
	export GO111MODULE=on
	go mod tidy
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/main ./main.go
