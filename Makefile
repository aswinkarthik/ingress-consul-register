GOBIN=go
GLIDE=glide
GLIDE_NO_VENDOR=$(shell glide novendor)
PROJECT_NAME=ingress-consul-register
DOCKERBIN=docker

test:
	$(GOBIN) test -v $(GLIDE_NO_VENDOR)

build:
	$(GOBIN) build -o $(PROJECT_NAME)

docker-build:
	$(DOCKERBIN) build . -t aswinkarthik93/$(PROJECT_NAME)

