
REPOSITORY := gugahoi/golang-healthcheck
TAG ?= latest

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: build
build:
	go build -o build/golang-healthcheck ./...

.PHONY: docker-build
docker-build:
	docker build -t $(REPOSITORY):$(TAG) .

.PHONY: docker-run
docker-run: docker-build
	docker run -it -p 8080:80 $(REPOSITORY):$(TAG)
