

# load('ext://restart_process', 'docker_build_with_restart')
docker_build("yurikrupnik/users-api", ".",
  #  only=["./http"],

 #  live_update=[
  #      sync('.', '/'),
    #  sync("./", "./")
   #],
)

# k8s_yaml('k8s/base/deployment.yml')
# k8s_yaml('k8s/base/namespace.yml')
k8s_yaml(kustomize('k8s/base'))


# ports to container port that runs as container env var - both ways
k8s_resource("users-api", port_forwards="5001:8080")
# k8s_resource(workload='users-api', port_forwards="5001:8080")
