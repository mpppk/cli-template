SHELL = /bin/bash

.PHONY: lint
lint:
	gometalinter

.PHONY: test
test:
	go test ./...

.PHONY: integration-test
integration-test: deps
	go test -tags=integration ./...

.PHONY: coverage
coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: codecov
codecov:  coverage
	bash <(curl -s https://codecov.io/bash)

.PHONY: build
build:
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