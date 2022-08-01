job "traefik" {
  datacenters = ["dc1"]

  group "proxy" {
    network {
      mode = "host"
      port "ingress" {
        static = 80
      }
    }

    service {
      name     = "traefik"
      provider = "nomad"
      port     = "http"
      tags = [ ]
    }

    task "traefik" {
      driver = "docker"
      config {
        image = "traefik/traefik:nomad"
        args = [
          "--log.level=DEBUG",
          "--entryPoints.web.address=:80",
          "--providers=nomad",
          "--providers.nomad.refreshInterval=1s",
        ]
      }

      resources {
        cpu    = 10
        memory = 32
      }
    }
  }
}
