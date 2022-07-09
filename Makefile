GITHUB_TOKEN=""
DOCKER_REGISTRY=yurikrupnik
SERVICE=fiber-mongo
IMAGE=${DOCKER_REGISTRY}/${SERVICE}

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

##@ Development
up:
	-kind create cluster --name test-env --image kindest/node:v1.21.1 --config cluster.yaml
	-nvm install node
	#tilt up
	make tilt
down:
	-tilt down
	kind delete cluster --name test-env

release-dry:
	goreleaser build --snapshot --rm-dist
release-build:
	goreleaser build
release-release:
	goreleaser release

build-image:
	echo IMAGE ${IMAGE}
	docker build . -t ${IMAGE}
build-push:
	docker push $IMAGE

