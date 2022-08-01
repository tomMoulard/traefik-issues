Hello @fomojola,

Thanks for your interest in Traefik!

Traefik tries to detect when 2 (or more) services have the same configuration in order to make a loadbalancer between those services (e.g. when scaling up deployments on docker-compose).
Also, when building a router with no user defined service, Traefik will generate the service name using values from the context(i.e. the service name and the docker compose project for [docker-compose](https://github.com/traefik/traefik/blob/ba822acb23db3d620f6b3773ac9eff3e98ce313f/pkg/provider/docker/config.go#L414)).
For example, given this service:

```yaml
services:
  whoami:
    image: traefik/whoami
    labels:
      traefik.http.routers.whoami.rule: Host(`whoami.localhost`)
```

Traefik will generate the service name `whoami-example` if running docker compose like: `docker-compose -p example up`.

In the example below, there are 2 definition of the same router but with different services.
In this case, I get the error described in the issue:

```yaml
version: '3.6'
services:
  traefik:
    image: traefik:v2.6
    command:
      - --providers.docker
      - --log.level=DEBUG
    ports:
      - "80:80"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
  whoami: # service name: whoami-example
    image: traefik/whoami
    labels:
      traefik.http.routers.whoami.rule: Host(`whoami.localhost`)
  not-whoami: # service name: not-whoami-example
    image: traefik/whoami
    labels:
      traefik.http.routers.whoami.rule: Host(`whoami.localhost`)
```
```
> docker-compose up
...
example-traefik-1     | time="2022-00-00T00:00:00Z" level=error msg="Router defined multiple times with different configurations in [not-whoami-example-5d0d871caa159531c32c84ad43d06b73996f133f595faf4b5e8c03e2c83dedfd whoami-example-991695a2af818413afe9eefbfbdcb96203adc9a669554ba9bfb6b8cb3d1afab8]" routerName=whoami providerName=docker
...
```

To fix this, each docker services should to specify the **same** service name.
This allows Traefik to match routers with one another and load balance between them.

Here is a fixed version of the example above:

```yaml
services:
  whoami:
    image: traefik/whoami
    labels:
      traefik.http.routers.whoami.rule: Host(`whoami.localhost`)
      traefik.http.services.whoami-service.loadbalancer.server.port: 80 # set the service name to whoami-service
  not-whoami:
    image: traefik/whoami
    labels:
      traefik.http.routers.whoami.rule: Host(`whoami.localhost`)
      traefik.http.services.whoami-service.loadbalancer.server.port: 80 # set the service name to whoami-service
```
