NAME = rchicoli/docker-go-input-validation
VERSION = 0.0.2-dev
WORKDIR = /go/src/go-input-validation

.PHONY: all build tag release

all: build

build:
	docker run --rm -ti -v $(PWD):$(WORKDIR) -w $(WORKDIR) golang:1.7.1-alpine go build -v
	mv go-input-validation docker/
	docker build --rm -t $(NAME):$(VERSION) docker/

tag:
	docker tag $(NAME):$(VERSION) $(NAME):latest

release: tag
	docker push $(NAME):$(VERSION)
	docker push $(NAME):latest
