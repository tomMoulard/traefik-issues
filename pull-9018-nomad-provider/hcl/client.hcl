log_level  = "DEBUG"
data_dir   = "/tmp"
datacenter = "dc-1"

client {
    enabled = true
    servers = ["${NOMAD_SERVER_ADDRESS}"]
    options {
        "docker.privileged.enabled" = "true"
    }
}

ports {
    http = ${NOMAD_CLIENT_HTTP_PORT}
}

