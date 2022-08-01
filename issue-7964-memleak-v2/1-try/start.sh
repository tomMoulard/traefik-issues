#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

CLUSTER_NAME="test"
CLUSTER_IMAGE_VERSION="v1.22.3-k3s1"

k3d cluster create $CLUSTER_NAME --agents 2 --image rancher/k3s:$CLUSTER_IMAGE_VERSION --k3s-arg '--disable=traefik@server:0' -p '80:80@loadbalancer' -p '8080:8080@loadbalancer'

k3d image import --cluster test traefik/traefik:experimental-v2.5

echo "Installing manifests..."
kubectl apply -f manifests/1-rbac.yml -f manifests/2-traefik.yaml

# k3d cluster delete $CLUSTER_NAME
# kubectl delete namespaces ingresses
# kubectl get pods -A
# make build-image
# PRE_TARGET="" make build-image
# kubectl -n ingresses logs whoami-76c64d4749-p65tl
# kubectl logs traefik-7fd8d8f6fd-4jzjb
# wget http://localhost:8080/debug/pprof/heap
# go tool pprof -http=:9191 heap.out
# k top pods traefik-66f5cb756c-5bd8t
# k describe pods traefik-66f5cb756c-5bd8t
# k logs -f -l app=traefik # --tail=-1
# k exec $(k get pods -A | grep ' traefik' | awk '{print $2}') -it -- /bin/bash
# lsof -iTCP -P -r 2

# tried without metrics/prometheus, and k top still showed the MEMORY at ~150MB after ~30 minutes -> still leaky.

# stop using safe.Go in aggregator. -> no visible difference in pproc.

# tried completely resetting the config everytime we get a new message. -> still leaking.

# managed to display traefik_entrypoint_open_connection, but display too many (GET, POST, "traefik", "http") of them, how do I filter?

# k3d image import --cluster test traefik/traefik && k delete pods -l app=traefik
