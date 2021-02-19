# this makefile is to
# build binary of httpAPIserver
# creating docker image using dockerfile
# pushing the docker image in dockerhub
# and also
# make install # for doing helm chart install
# make uninstall # for doing helm chart uninstall


.PHONY: all build_binary image_create image_push

all: build_binary image_create image_push

IMAGE ?= api-server
REPO ?= shahincsejnu
TAG ?= v2.0.1

RELEASE_NAME ?= test
REPO_NAME ?= oka
CHART_NAME ?= apiserver

build_binary:
	@echo "binary is building with the name server.."
	@go build -o server .

image_create: build_binary
	@echo "docker image is creating.."
	@docker build -t $(REPO)/$(IMAGE):$(TAG) .

image_push: image_create
	#@docker login --username=$(REPO)
	@docker push $(REPO)/$(IMAGE):$(TAG)

install:
	@echo "Adding remote repo in your machine.."
	@helm repo add $(REPO_NAME) https://shahincsejnu.github.io/helm-charts-stuffs/
	@echo "Updating your repos.."
	@helm repo update
	@echo "Installing helm chart from the repo $(REPO_NAME)"
	@helm install $(RELEASE_NAME) $(REPO_NAME)/$(CHART_NAME)

uninstall:
	@echo "Uninstalling $(RELEASE_NAME)..."
	@helm uninstall $(RELEASE_NAME)

