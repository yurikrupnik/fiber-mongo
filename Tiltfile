
local_resource(
  'local-myserver',
  cmd='GOOS=linux GOARCH=amd64 go build -o ./fiber-mongo',
  deps=['.'],
  ignore=["fiber-mongo"]
)

docker_build("yurikrupnik/users-api", ".",
   only=["fiber-mongo"],
)

k8s_yaml(kustomize('k8s/base'))


# ports to container port that runs as container env var - both ways
k8s_resource("users-api", port_forwards="5001:8080")
# k8s_resource(workload='users-api', port_forwards="5001:8080")
