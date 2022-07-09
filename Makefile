GITHUB_TOKEN="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"
DOCKER_REGISTRY=yurikrupnik
SERVICE=fiber-mongo
IMAGE=${DOCKER_REGISTRY}/${SERVICE}

##@ General
# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: release-dev
release-dev: ## goreleaser build snapshot
	 goreleaser build --rm-dist --snapshot

##@ Cluster
up: ## Kind up cluster with 3 worker nodes
	-kind create cluster --name test-env --image kindest/node:v1.21.1 --config cluster.yaml
	-nvm install node
	#tilt up
	make tilt
down: ## Kind down cluster with 3 worker nodes
	-tilt down
	kind delete cluster --name test-env --config cluster.yaml

release-dry:
	GITHUB_TOKEN= goreleaser build --snapshot --rm-dist
release-build:
	goreleaser build
release-release:
	goreleaser release

build-image:
	echo IMAGE ${IMAGE}
	docker build . -t ${IMAGE}
build-push:
	docker push $IMAGE

