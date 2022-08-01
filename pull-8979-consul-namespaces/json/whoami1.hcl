service {
  name = "whoami1"
  namespace = "namespace1"
  tags = [
    "traefik.enable=true"
  ]
  address = "whoami"
  port = 80
}
