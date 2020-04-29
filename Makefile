SHELL = /bin/bash

.PHONY: setup
setup:
	mv go.mod go.mod.org
	mv go.sum go.sum.org
	go get github.com/google/wire/cmd/wire
	go get github.com/goreleaser/goreleaser
	mv go.mod.org go.mod
	mv go.sum.org go.sum

.PHONY: lint
lint: generate
	go vet ./...
	goreleaser check

.PHONY: test
test: generate
	go test ./...

.PHONY: integration-test
integration-test:
	go test -tags=integration ./...

.PHONY: coverage
coverage: generate
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: codecov
codecov:  coverage
	bash <(curl -s https://codecov.io/bash)

.PHONY: wire
wire:
	go generate -tags=wireinject ./...

.PHONY: generate
generate: wire
	go generate ./...

.PHONY: build
build: generate
	go build

.PHONY: cross-build-snapshot
cross-build:
	goreleaser --rm-dist --snapshot

.PHONY: install
install:
	go install

.PHONY: circleci
circleci:
	circleci build -e GITHUB_TOKEN=$GITHUB_TOKEN