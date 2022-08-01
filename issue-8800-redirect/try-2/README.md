I have tried using your ingresses, but I still cannot have a request being redirected to a private IP address.

<details>
  <summary markdown="span">My -not- reproducible use case</summary>

`ressources.yml`:
```yml
apiVersion: v1
items:
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      cert-manager.io/cluster-issuer: letsencrypt-production
      external-dns.alpha.kubernetes.io/hostname: mgr.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      kubernetes.io/tls-acme: "true"
      meta.helm.sh/release-name: dbmgr
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "true"
    creationTimestamp: "2021-07-15T16:52:31Z"
    generation: 1
    labels:
      app.kubernetes.io/instance: XXXXX
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: mgr
      environment: preproduction
      helm.sh/chart: mgr-1.4.0
      method: traefik
      visibility: internal
    name: dbmgr-mgr
    namespace: default
    resourceVersion: "91095703"
    uid: b251ed66-7508-4286-aa5e-28a1d3dfcfc4
  spec:
    rules:
    - host: mgr.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                name: http
          path: /
          pathType: Prefix
    tls:
    - hosts:
      - mgr.uat.mycompany.com
      secretName: chart-mgr-tls
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXe-bXXXXXXXXXXXXXXe.elb.us-east-1.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      cert-manager.io/cluster-issuer: letsencrypt-production
      external-dns.alpha.kubernetes.io/hostname: app-e.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      kubernetes.io/tls-acme: "true"
      meta.helm.sh/release-name: app-e
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2022-01-20T20:49:47Z"
    generation: 2
    labels:
      app.kubernetes.io/instance: app-e
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: app-e-component-d
      app.kubernetes.io/version: 2.3.0
      environment: uat
      helm.sh/chart: component-d-1.1.0
      method: traefik
      tls: aws
      visibility: public
    name: app-e-component-d
    namespace: default
    resourceVersion: "97207026"
    uid: 6e6b0c36-2182-4c11-b8d4-20528539a88e
  spec:
    rules:
    - host: app-e.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 80
          path: /
          pathType: Prefix
    tls:
    - hosts:
      - app-e.uat.mycompany.com
      secretName: component-d-tls
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXX17-1XXXXXXXX6.us-east-1.elb.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      external-dns.alpha.kubernetes.io/hostname: app-a-bot.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      meta.helm.sh/release-name: mycompany-app-a
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2021-10-28T20:51:53Z"
    generation: 3
    labels:
      app.kubernetes.io/instance: mycompany-app-a
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: mycompany-app-a
      app.kubernetes.io/version: 0.0.1
      environment: uat
      helm.sh/chart: teams-suite-0.2.4
      method: traefik
      tls: aws
      visibility: public
    name: mycompany-app-a-bot
    namespace: default
    resourceVersion: "53179125"
    uid: 1f0f268b-a6c2-4114-b09a-91d7210dc87f
  spec:
    rules:
    - host: app-a-bot.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 80
          path: /
          pathType: Prefix
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX7-1XXXXXXXX6.us-east-1.elb.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      external-dns.alpha.kubernetes.io/hostname: app-a-connector.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      meta.helm.sh/release-name: mycompany-app-a
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2021-10-27T14:04:01Z"
    generation: 3
    labels:
      app.kubernetes.io/instance: mycompany-app-a
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: mycompany-app-a
      app.kubernetes.io/version: 0.0.1
      environment: uat
      helm.sh/chart: teams-suite-0.2.4
      method: traefik
      tls: aws
      visibility: public
    name: mycompany-app-a-connector
    namespace: default
    resourceVersion: "53179129"
    uid: 0d811951-c83e-488d-9880-6ec2f46a098c
  spec:
    rules:
    - host: app-a-connector.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 80
          path: /
          pathType: Prefix
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX7-1XXXXXXXX6.us-east-1.elb.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      external-dns.alpha.kubernetes.io/hostname: app-a-core.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      meta.helm.sh/release-name: mycompany-app-a
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2021-11-11T21:09:31Z"
    generation: 3
    labels:
      app.kubernetes.io/instance: mycompany-app-a
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: mycompany-app-a
      app.kubernetes.io/version: 0.0.1
      environment: uat
      helm.sh/chart: teams-suite-0.2.4
      method: traefik
      tls: aws
      visibility: public
    name: mycompany-app-a-core
    namespace: default
    resourceVersion: "53179133"
    uid: d201069a-ede5-4437-aafe-2b5f1fb1b681
  spec:
    rules:
    - host: app-a-core.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 80
          path: /
          pathType: Prefix
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX7-1XXXXXXXX6.us-east-1.elb.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      external-dns.alpha.kubernetes.io/hostname: app-a-helpdesk.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      meta.helm.sh/release-name: mycompany-app-a
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2021-10-28T20:51:53Z"
    generation: 3
    labels:
      app.kubernetes.io/instance: mycompany-app-a
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: mycompany-app-a
      app.kubernetes.io/version: 0.0.1
      environment: uat
      helm.sh/chart: teams-suite-0.2.4
      method: traefik
      tls: aws
      visibility: public
    name: mycompany-app-a-helpdesk
    namespace: default
    resourceVersion: "53179136"
    uid: 4961510b-dfa1-4899-bc37-4c2429e8547b
  spec:
    rules:
    - host: app-a-helpdesk.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 80
          path: /
          pathType: Prefix
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX7-1XXXXXXXX6.us-east-1.elb.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      external-dns.alpha.kubernetes.io/hostname: app-a-jira.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      meta.helm.sh/release-name: mycompany-app-a
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2021-10-29T12:32:08Z"
    generation: 4
    labels:
      app.kubernetes.io/instance: mycompany-app-a
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: mycompany-app-a
      app.kubernetes.io/version: 0.0.1
      environment: uat
      helm.sh/chart: teams-suite-0.2.4
      method: traefik
      tls: aws
      visibility: public
    name: mycompany-app-a-component-d
    namespace: default
    resourceVersion: "53179140"
    uid: 5dd37d77-6be6-449f-8b83-d715410ff8eb
  spec:
    rules:
    - host: app-a-jira.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 80
          path: /
          pathType: Prefix
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX7-1XXXXXXXX6.us-east-1.elb.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      external-dns.alpha.kubernetes.io/hostname: app-a-tabs.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      meta.helm.sh/release-name: mycompany-app-a
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2021-10-27T14:04:01Z"
    generation: 3
    labels:
      app.kubernetes.io/instance: mycompany-app-a
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: mycompany-app-a
      app.kubernetes.io/version: 0.0.1
      environment: uat
      helm.sh/chart: teams-suite-0.2.4
      method: traefik
      tls: aws
      visibility: public
    name: mycompany-app-a-tabs
    namespace: default
    resourceVersion: "53179142"
    uid: 0b8e405c-fad6-49f0-b69d-228d262085f0
  spec:
    rules:
    - host: app-a-tabs.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 80
          path: /
          pathType: Prefix
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX7-1XXXXXXXX6.us-east-1.elb.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      cert-manager.io/cluster-issuer: letsencrypt-production
      external-dns.alpha.kubernetes.io/hostname: mstjold.mycompany.ai
      external-dns/domain: mycompany.ai
      kubernetes.io/tls-acme: "true"
      meta.helm.sh/release-name: mycompany-app-a-temporary-component-d
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2022-02-16T16:56:00Z"
    generation: 3
    labels:
      app.kubernetes.io/instance: mycompany-app-a-temporary-component-d
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: nginx
      environment: preproduction
      helm.sh/chart: nginx-9.7.6
      method: traefik
      tls: acme
      visibility: public
    name: mycompany-app-a-temporary-component-d-nginx
    namespace: default
    resourceVersion: "105028364"
    uid: dc08e92a-337c-4c63-84ec-b7c79141c16c
  spec:
    rules:
    - host: mstjold.mycompany.ai
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                name: http
          path: /
          pathType: ImplementationSpecific
    tls:
    - hosts:
      - mstjold.mycompany.ai
      secretName: mstjold.mycompany.ai-tls
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXc-8XXXXXXXXXXXXXXa.elb.us-east-1.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      external-dns.alpha.kubernetes.io/hostname: pd.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      meta.helm.sh/release-name: app-c
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2021-12-15T20:22:48Z"
    generation: 1
    labels:
      app.kubernetes.io/instance: app-c
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: app-c-component-c
      app.kubernetes.io/version: 0.2.0
      environment: uat
      helm.sh/chart: component-c-1.0.0
      method: traefik
      tls: aws
      visibility: public
    name: app-c-component-c
    namespace: default
    resourceVersion: "68037979"
    uid: e71b26d8-b61f-4d52-8d82-4145133dd33c
  spec:
    rules:
    - host: pd.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 3300
          path: /
          pathType: Prefix
    tls:
    - hosts:
      - pd.uat.mycompany.com
      secretName: myapp-tls
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX7-1XXXXXXXX6.us-east-1.elb.amazonaws.com
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      external-dns.alpha.kubernetes.io/hostname: account.uat.mycompany.com
      external-dns/domain: uat.mycompany.com
      meta.helm.sh/release-name: app-b
      meta.helm.sh/release-namespace: default
      traefik.ingress.kubernetes.io/preserve-host: "true"
      traefik.ingress.kubernetes.io/router.entrypoints: web, websecure
      traefik.ingress.kubernetes.io/router.tls: "false"
    creationTimestamp: "2021-11-13T15:33:50Z"
    generation: 1
    labels:
      app.kubernetes.io/instance: app-b
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: app-b
      app.kubernetes.io/version: "5.9"
      environment: uat
      helm.sh/chart: app-b-1.9.0
      method: traefik
      tls: aws
      visibility: public
    name: app-b
    namespace: default
    resourceVersion: "104009999"
    uid: 63856992-0aea-4c7e-a04f-c1e762f4cb22
  spec:
    rules:
    - host: account.uat.mycompany.com
      http:
        paths:
        - backend:
            service:
              name: whoami
              port:
                number: 8000
          path: /
          pathType: Prefix
  status:
    loadBalancer:
      ingress:
      - hostname: aXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX7-1XXXXXXXX6.us-east-1.elb.amazonaws.com
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
```

`whoami.yml`:
```yml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: default
  name: whoami
  labels:
    app: whoami

spec:
  replicas: 2
  selector:
    matchLabels:
      app: whoami
  template:
    metadata:
      labels:
        app: whoami
    spec:
      containers:
        - name: whoami
          image: traefik/whoami:latest
          ports:
            - name: web
              containerPort: 80

---
apiVersion: v1
kind: Service
metadata:
  name: whoami

spec:
  ports:
    - protocol: TCP
      name: web
      port: 80
  selector:
    app: whoami
```

```sh
$ k3d cluster create mycluster \
	-p 80:80@loadbalancer \
	-p 8080:8080@loadbalancer \
	-p 443:443@loadbalancer \
	--k3s-arg '--no-deploy=traefik@server:0' \
	-i rancher/k3s:v1.21.4-k3s1
...
INFO[2022-03-01T16:08:59+01:00] Cluster 'mycluster' created successfully!
$ helm install -f values.yml traefik traefik/traefik
...
$ kubectl apply -f whoami.yml
deployment.apps/whoami created
service/whoami created
$ kubectl apply -f resources.yml
ingress.networking.k8s.io/dbmgr-mgr created
ingress.networking.k8s.io/app-e-component-d created
ingress.networking.k8s.io/mycompany-app-a-bot created
ingress.networking.k8s.io/mycompany-app-a-connector created
ingress.networking.k8s.io/mycompany-app-a-core created
ingress.networking.k8s.io/mycompany-app-a-helpdesk created
ingress.networking.k8s.io/mycompany-app-a-component-d created
ingress.networking.k8s.io/mycompany-app-a-tabs created
ingress.networking.k8s.io/mycompany-app-a-temporary-component-d-nginx created
ingress.networking.k8s.io/app-c-component-c created
ingress.networking.k8s.io/app-b created
$ curl --http1.0 -H "Host: " app-a-tabs.uat.mycompany.com -v
* Couldn't find host app-a-tabs.uat.mycompany.com in the (nil) file; using defaults
*   Trying 127.0.0.1:80...
* Connected to app-a-tabs.uat.mycompany.com (127.0.0.1) port 80 (#0)
> GET / HTTP/1.0
> Host:
> User-Agent: curl/7.79.0-DEV
> Accept: application/json, application/xml, text/plain, */*
>
* Mark bundle as not supporting multiuse
* HTTP 1.0, assume close after body
< HTTP/1.0 404 Not Found
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Tue, 01 Mar 2022 15:10:03 GMT
< Content-Length: 19
<
404 page not found
* Closing connection 0
$ curl app-a-tabs.uat.mycompany.com -v
* Couldn't find host app-a-tabs.uat.mycompany.com in the (nil) file; using defaults
*   Trying 127.0.0.1:80...
* Connected to app-a-tabs.uat.mycompany.com (127.0.0.1) port 80 (#0)
> GET / HTTP/1.1
> Host: app-a-tabs.uat.mycompany.com
> User-Agent: curl/7.79.0-DEV
> Accept: application/json, application/xml, text/plain, */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 301 Moved Permanently
< Location: https://app-a-tabs.uat.mycompany.com/
< Date: Tue, 01 Mar 2022 15:10:07 GMT
< Content-Length: 17
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host app-a-tabs.uat.mycompany.com left intact
Moved Permanently
```
</details>

Since I am not able to have access to your backend, I would assume that this is the reason about redirection. Thus I will close this issue as a question, but let me know if you disagree.
