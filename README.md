[![Build](https://github.com/kserve/modelmesh-runtime-adapter/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/kserve/modelmesh-runtime-adapter/actions/workflows/build.yml)

# ModelMesh Runtime Adapter

This repo contains the unified puller/runtime-adapter image of the sidecar containers
which run in the ModelMesh Serving model server Pods. Take a look at the main
[ModelMesh Serving](https://github.com/kserve/modelmesh-serving) repo for more details.

Logical subcomponents within the image:

- [model-serving-puller](model-serving-puller)
- [model-mesh-mlserver-adapter](model-mesh-mlserver-adapter)
- [model-mesh-triton-adapter](model-mesh-triton-adapter)
- [model-mesh-ovms-adapter](model-mesh-ovms-adapter)
- [model-mesh-torchserve-adapter](model-mesh-torchserve-adapter)

## Generate sources

The gRPC code stubs, interfaces and data access classes have to be generated by the
[`protoc` compiler](https://protobuf.dev/getting-started/gotutorial/#compiling-protocol-buffers)
from the `.proto` source files under `internal/proto/*`.

If any of the `.proto` files were modified, run the `protoc` compiler to regenerate
the respective Go source code. It's recommended to use the developer image, which
has all the required libraries pre-installed, by running `make run proto.compile`
instead of `make proto.compile`.

```shell
make run proto.compile
```

## Test the code changes

After making code changes, ensure all existing and new functionality still works
properly by running the unit tests.

```shell
make test
```

## Format code

Run the linter to make sure all code style rules are adhered to. The code will
automatically be formatted if any code style violations are found.

It's recommended to use the developer image, which has all the required libraries
pre-installed, by running `make run fmt` instead of `make fmt`.

```shell
make run fmt
```

## Build the Docker image

Once the code changes have been tested and linted, build a new `modelmesh-runtime-adapter`
Container image. If you need to use another builder than `docker`, you can specify it by using the `ENGINE` variable:

```shell
ENGINE=podman make build
```

## Push the image to a container registry

Push the newly built `modelmesh-runtime-adapter` image to a container registry.
Replace the value of the `DOCKER_USER` environment variable to your docker user ID
and change the `IMAGE_TAG` to something meaningful.

```bash
export DOCKER_USER="<your-docker-userid>"
export IMAGE_TAG="dev"

docker tag kserve/modelmesh-runtime-adapter:latest \
    ${DOCKER_USER}/modelmesh-runtime-adapter:${IMAGE_TAG}

docker push ${DOCKER_USER}/modelmesh-runtime-adapter:${IMAGE_TAG}
```

## Update the ModelMesh Serving deployment

In order to test the code changes in an existing [ModelMesh Serving](https://github.com/kserve/modelmesh-serving)
deployment, the newly built container image needs to be added to the
`model-serving-config` ConfigMap.

### Check existing model serving configuration

First, check if your ModelMesh Serving deployment already has an existing
`model-serving-config` ConfigMap:

```Shell
kubectl get configmap

NAME                            DATA   AGE
kube-root-ca.crt                1      12d
model-serving-config            1      12d
model-serving-config-defaults   1      12d
tc-config                       2      12d
```

### Create a new model serving config

If you did not already have a `model-serving-config` ConfigMap on your cluster,
you can create one. Replace the `<your-docker-userid>` placeholder with your
Docker username. Make sure the value of the `IMAGE_TAG` variable matches
the one that was pushed to the container registry.

```shell
export DOCKER_USER="<your-docker-userid>"
export IMAGE_NAME="${DOCKER_USER}/modelmesh-runtime-adapter"
export IMAGE_TAG="dev"

kubectl apply -f - <<EOF
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: model-serving-config
data:
  config.yaml: |
    storageHelperImage:
      name: ${IMAGE_NAME}
      tag: ${IMAGE_TAG}
EOF
```

### Update an existing model serving config

If the ConfigMap list contains `model-serving-config`, save the contents of your
existing configuration in a local temp file:

```Bash
mkdir -p temp
kubectl get configmap model-serving-config -o yaml > temp/model-serving-config.yaml
```

Add the `storageHelperImage` property to the `config.yaml` string property.

```YAML
      storageHelperImage:
        name: your-docker-userid/modelmesh-runtime-adapter
        tag: latest
```

Replace the `your-docker-userid` placeholder with _your_ Docker username and make
sure the `tag` matches the one that was pushed to the container registry earlier.

The complete ConfigMap YAML file _may_ look like this:

```YAML
apiVersion: v1
kind: ConfigMap
metadata:
  name: model-serving-config
  namespace: modelmesh-serving
data:
  config.yaml: |
    podsPerRuntime: 1
    restProxy:
      enabled: true
    scaleToZero:
      enabled: false
      gracePeriodSeconds: 5
    storageHelperImage:
      name: your-docker-userid/modelmesh-runtime-adapter
      tag: dev
```

Apply the ConfigMap to your cluster:

```Bash
kubectl apply -f temp/model-serving-config.yaml
```

If you are comfortable using `vi`, you can forgo creating a temp file and edit
the ConfigMap directly in the terminal:

```Shell
kubectl edit configmap model-serving-config
```

### Verify the container images used by the model serving runtime

The `modelmesh-controller` watches the ConfigMap and responds to updates by
automatically restarting the serving runtime pods using the newly built
runtime adapter image.

You can check which container images are used by running the following command:

```Shell
kubectl get pods -o jsonpath='{range .items[*]}{"\n"}{.metadata.name}{"\t"}{range .spec.containers[*]}{.image}{", "}{end}{end}' | sort | column -ts $'\t' | sed 's/, *$//g'

etcd-78ff7867d5-45svw                            quay.io/coreos/etcd:v3.5.4
minio-6ddbfc9665-gtf7x                           kserve/modelmesh-minio-examples:latest
modelmesh-controller-64f5c8d6d6-k6rzc            kserve/modelmesh-controller:latest
modelmesh-serving-mlserver-1.x-84884c6849-s8dw6  kserve/rest-proxy:latest, seldonio/mlserver:1.3.2, your-docker-userid/modelmesh-runtime-adapter:dev, kserve/modelmesh:latest
modelmesh-serving-mlserver-1.x-84884c6849-xpdw4  kserve/rest-proxy:latest, seldonio/mlserver:1.3.2, your-docker-userid/modelmesh-runtime-adapter:dev, kserve/modelmesh:latest
```
