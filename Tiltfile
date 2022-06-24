

local_resource(
  'local-myserver',
  cmd='GOARCH=arm64 go build',
  serve_cmd='PORT=8080 ./fiber-mongo',
  deps=['cmd/fiber-mongo']
)
# docker_build("yurikrupnik/users-api", "./")

k8s_yaml('k8s/base/deployment.yml')
# k8s_yaml('k8s/base/namespace.yml')
# k8s_yaml(kustomize('k8s/base'))