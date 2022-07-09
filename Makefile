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

release:
	goreleaser build --snapshot --rm-dist

build-image:
	echo IMAGE ${IMAGE}
	docker build . -t ${IMAGE}
build-push:
	docker push $IMAGE

