Hello @fabifont,

Thanks for your interest in Traefik!

You need to define two ratelimit middleware and attach them to the corresponding router.

Here is an example:
```yml
version: '3.6'

services:
  traefik:
    image: traefik:v2.6
    command:
      - --api.insecure=true
      - --providers.docker
    ports:
      - "80:80"
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
  whoami:
    image: traefik/whoami # https://github.com/traefik/whoami
    labels:
      traefik.http.middlewares.login-ratelimit.ratelimit.average: 100
      traefik.http.middlewares.login-ratelimit.ratelimit.burst: 10
      traefik.http.routers.login-whoami.rule: Host(`whoami.localhost`) && PathPrefix(`/api/login`)
      traefik.http.routers.login-whoami.middlewares: login-ratelimit

      traefik.http.middlewares.api-ratelimit.ratelimit.average: 500
      traefik.http.middlewares.api-ratelimit.ratelimit.burst: 50
      traefik.http.routers.api-whoami.rule: Host(`whoami.localhost`) && PathPrefix(`/api`)
      traefik.http.routers.api-whoami.middlewares: api-ratelimit

```

Here I define 2 routers, `api-whoami` and `login-whoami` to whom I attach respectively `api-ratelimit` and `login-ratelimit` that are two ratelimit middlewares.

I can see the configuration received by traefik as:
```json
{
  "routers": {
    "api-whoami@docker": {
      "entryPoints": [
        "http"
      ],
      "middlewares": [
        "api-ratelimit@docker"
      ],
      "service": "whoami-example",
      "rule": "Host(`whoami.localhost`) && PathPrefix(`/api`)",
      "status": "enabled",
      "using": [
        "http"
      ]
    },
    "login-whoami@docker": {
      "entryPoints": [
        "http"
      ],
      "middlewares": [
        "login-ratelimit@docker"
      ],
      "service": "whoami-example",
      "rule": "Host(`whoami.localhost`) && PathPrefix(`/api/login`)",
      "status": "enabled",
      "using": [
        "http"
      ]
    }
  }
}
```

As your issue might be a configuration problem, this issue will be closed.
To confirm this, please join our [Community Forum](https://community.traefik.io/) and reach out to us on the [Traefik section](https://community.traefik.io/c/traefik/5).
