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
	@rm -rf cover.out result.json ./bin ./deployments/charts/*.tgz
	@echo Successfuly removed artifacts

## go
.PHONY: lint
lint: ## run golangci-lint
	@golangci-lint run ./...

.PHONY: cover
cover: ## run coverage
	@go test -json -coverprofile=cover.out ./... > result.json

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
	@k6 run --vus=1 --iterations=1 --out=cloud ./test/k6/api.test.js

.PHONY: test-smoke
test-smoke: ## test smoke
	@k6 run --vus=5 --duration='30s' --out=cloud ./test/k6/api.test.js

.PHONY: test-load
test-load: ## test load
	@k6 run --out=cloud ./test/k6/api.test.js

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
