service {
  name = "whoami2"
  namespace = "namespace2"
  tags = [
    "traefik.enable=true"
  ]
  address = "whoami"
  port = 80
}
