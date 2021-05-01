
SERVICE ?= entxid-test
IMAGE ?= $(SERVICE)

GOPRIVATE ?= github.com/shanna/*
GOENV = GOPRIVATE=$(GOPRIVATE)

.PHONY: all
all: generate
	$(GOENV) go build -trimpath ./cmd/$(SERVICE)

.PHONY: generate
generate: clean
	$(GOENV) go generate ./...
	$(GOENV) go mod tidy
	$(GOENV) go mod verify

.PHONY: update
update:
	$(GOENV) go get github.com/oligot/go-mod-upgrade
	$(GOENV) go run github.com/oligot/go-mod-upgrade
	$(GOENV) go mod tidy
	$(GOENV) go mod verify

.PHONY: test
test: generate
	$(GOENV) go test -count=1 ./...

.PHONY: docker
docker: generate
	$(GOENV) go mod vendor
	$(GOENV) go mod verify
	docker build \
		--build-arg=GOPRIVATE=$(GOPRIVATE) \
		--build-arg=SERVICE=$(SERVICE) \
		-t $(IMAGE) \
		-f Dockerfile \
		.

.PHONY: clean
clean:
	rm -rf $(SERVICE)
