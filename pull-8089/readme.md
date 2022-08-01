# Traefik
First of all we will install a k3s cluster with the following command:
  ```bash
    k3d cluster delete mycluster
    k3d cluster create mycluster --api-port 6550 -p 80:80@loadbalancer -p 8080:8080@loadbalancer -p 443:443@loadbalancer --k3s-server-arg '--no-deploy=traefik' -i rancher/k3s:v1.18.6-k3s1
  ```

Once the k3s is up and running:

  ## Deploy traefik

  ```
  helm repo add traefik https://helm.traefik.io/traefik
  helm repo update

  kubectl create namespace traefik
  helm install --namespace traefik traefik traefik/traefik
  ```

  ## Deploy ingress-routes

  ```
  kubectl apply -f ingress-routes.yaml
  ```