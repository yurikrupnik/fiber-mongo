
# docker_compose('./compose.yaml')
# dc_resource('db', labels=["database"])
# todo how to connect to mongo cluster running in local machine
# todo test with mongo helm chart
# load('ext://helm_resource', 'helm_resource', 'helm_repo')
# helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')
# helm_repo('cowboysysop', 'https://cowboysysop.github.io/charts/')
# helm_resource('mongodb', 'cowboysysop/mongo-express')

local_resource(
  'build-fiber-mongo',
  cmd='GOOS=linux GOARCH=amd64 go build -o ./fiber-mongo',
  deps=['.'],
  ignore=["fiber-mongo", "dist/"]
)

docker_build("yurikrupnik/fiber-mongo", ".", only=["fiber-mongo"])

k8s_yaml(kustomize('k8s/base'))


# ports to container port that runs as container env var - both ways
k8s_resource("fiber-mongo", port_forwards="5022:8080")
# k8s_resource(workload='users-api', port_forwards="5001:8080")
