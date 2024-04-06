PROJECT_NAME := $(shell basename $(CURDIR))
VERSION := $(shell git describe --tags --always)

## common
.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## show version
	@echo $(VERSION)

.PHONY: clean
clean:  ## remove artifacts
	@rm -rf coverage.txt profile.out ./bin ./deployments/charts/*.tgz
	@echo Successfuly removed artifacts

## go
.PHONY: lint
lint: ## run golangci-lint
	@golangci-lint run ./...

.PHONY: gazelle
gazelle: ## run gazelle with bazel
	@bazel run //:gazelle

.PHONY: build
build: ## build go binary
	@bazel build //...

.PHONY: test
test: ## test go binary
	@bazel test //...

.PHONY: test-e2e
test-e2e: ## test e2e
	@k6 run ./test/k6/api.v1.players.post.test.js
	@k6 run ./test/k6/api.v1.players.delete.test.js

.PHONY: gen-swagger
gen-swagger: ## generate swagger
	@swag init -q -d ./adapter,./pkg,./entity -o ./adapter/api/docs

## docker
.PHONY: docker-build
docker-build: ## build docker image
	@docker build -t $(PROJECT_NAME):$(VERSION) .

.PHONY: docker-run
docker-run: docker-build ## run docker container
	@docker run -it --rm -p 30000:30000 $(PROJECT_NAME):$(VERSION) start api

.PHONY: docker-push
docker-push: ## push docker image
	@bazel run //adapter:push -- --tag=$(VERSION)
