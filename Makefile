GITHUB_TOKEN=""
DOCKER_REGISTRY=yurikrupnik
SERVICE=fiber-mongo
IMAGE=${DOCKER_REGISTRY}/${SERVICE}

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
release:
	goreleaser build
build-image:
	echo IMAGE ${IMAGE}
	docker build . -t ${IMAGE}
build-push:
	docker push $IMAGE

