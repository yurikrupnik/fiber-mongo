GITHUB_TOKEN=""
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
	-ctlptl create registry ctlptl-registry
	-ctlptl create cluster kind --registry=ctlptl-registry
	-tilt up
down: ## Kind down cluster with 3 worker nodes
	tilt down
	kind delete cluster

.PHONY: release-dry
release-dry: ## Create dev release - able to run while changing code per snapshot.
	IMAGE=${IMAGE} GITHUB_TOKEN= goreleaser release --snapshot --rm-dist
.PHONY: release-dry
release-build:
	IMAGE=${IMAGE} goreleaser build
release-release:
	GITHUB_TOKEN=$GITHUB_TOKEN IMAGE=$IMAGE goreleaser release

build-image:
	echo IMAGE ${IMAGE}
	docker build . -t ${IMAGE}
build-push:
	docker push $IMAGE

