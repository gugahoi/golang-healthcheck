
REPOSITORY := gugahoi/golang-healthcheck
TAG ?= latest
TAG := $(if $(TRAVIS_TAG),$(TRAVIS_TAG),$(TAG))

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: build
build: SHA=$(shell git log -1 --pretty=format:"%H")
build:
	CGO_ENABLED=0 go build -o build/app \
		-ldflags="-X main.Version=$(TAG) -X main.CommitSHA=$(SHA)" \
		.

.PHONY: docker-build
docker-build:
	docker build --build-arg TAG=$(TAG) -t $(REPOSITORY):$(TAG) .

.PHONY: docker-run
docker-run: docker-build
	docker run -it -p 8080:80 $(REPOSITORY):$(TAG)

.PHONY: docker-push
docker-push: docker-build
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker push $(REPOSITORY):$(TAG)
