# Configuration variables with defaults
IMAGE_NAME ?= cosmo-custom-router
IMAGE_TAG ?= latest
IMAGE ?= $(IMAGE_NAME):$(IMAGE_TAG)

TARGETOS ?= linux
TARGETARCH ?= amd64
PLATFORM ?= $(TARGETOS)/$(TARGETARCH)

# Dynamic metadata build args
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT ?= $(shell git rev-parse HEAD 2>/dev/null || echo unknown)
DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")/

.PHONY: all build docker-build help

all: docker-build

build: docker-build

docker-build:
	go mod tidy && \
	docker build \
		--platform $(PLATFORM) \
		--build-arg TARGETOS=$(TARGETOS) \
		--build-arg TARGETARCH=$(TARGETARCH) \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg DATE=$(DATE) \
		-t $(IMAGE) .

help:
	@echo "Usage: make [target] [VARIABLES...]"
	@echo ""
	@echo "Targets:"
	@echo "  docker-build  Build the Docker image (default)"
	@echo "  build         Alias for docker-build"
	@echo "  help          Show this help message"
	@echo ""
	@echo "Variables:"
	@echo "  IMAGE_NAME    Name of the Docker image (default: $(IMAGE_NAME))"
	@echo "  IMAGE_TAG     Tag for the Docker image (default: $(IMAGE_TAG))"
	@echo "  IMAGE         Full image reference (default: $(IMAGE))"
	@echo "  TARGETOS      Target operating system (default: $(TARGETOS))"
	@echo "  TARGETARCH    Target architecture (default: $(TARGETARCH))"
	@echo "  PLATFORM      Target platform string (default: $(PLATFORM))"
	@echo "  VERSION       Build version (default from git describe)"
	@echo "  COMMIT        Git commit SHA (default from git rev-parse)"
	@echo "  DATE          Build date in RFC3339 format"
