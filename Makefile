SHELL := /bin/bash

HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_COUNTERFEITER := $(shell command -v counterfeiter;)
PLATFORM := $(shell uname -s)

# #### DEPS ####
.PHONY: deps-counterfeiter deps-ginkgo deps-modules

deps-counterfeiter:
ifndef HAS_COUNTERFEITER
	go install github.com/maxbrunsfeld/counterfeiter/v6@latest
endif

deps-ginkgo:
ifndef HAS_GINKGO
	go install github.com/onsi/ginkgo/v2/ginkgo
endif

deps-modules:
	go mod download

# #### SRC ####
internal/internalfakes/fake_device_service.go: internal/device_service.go
	go generate internal/device_service.go

internal/internalfakes/fake_firmware_service.go: internal/firmware_service.go
	go generate internal/firmware_service.go

internal/internalfakes/fake_updater.go: internal/updater.go
	go generate internal/updater.go

# #### TEST ####
.PHONY: lint test-units test-features test

lint: deps-modules
ifndef HAS_GOLANGCI_LINT
ifeq ($(PLATFORM), Darwin)
	brew install golangci-lint
endif
ifeq ($(PLATFORM), Linux)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
endif
endif
	golangci-lint run

test-units: internal/internalfakes/fake_device_service.go internal/internalfakes/fake_firmware_service.go internal/internalfakes/fake_updater.go deps-modules deps-ginkgo
	ginkgo -r -skip-package test .

test-features: deps-modules deps-ginkgo
	ginkgo -r test

test: lint test-units test-features

# #### BUILD ####
.PHONY: build
SOURCES = $(shell find . -name "*.go" | grep -v "_test\." )

build/ota-service: $(SOURCES) deps-modules
	go build -o build/ota-service github.com/petewall/ota-service/v2

build: build/ota-service

build-image:
	docker build -t petewall/ota-service .

# #### RUN ####
.PHONY: run

run: build/ota-service
	./build/ota-service
