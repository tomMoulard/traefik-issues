log_level  = "DEBUG"
data_dir   = "/tmp"
datacenter = "dc-1"

server {
    enabled          = true
    bootstrap_expect = 1
    retry_join = ["${NOMAD_SERVER_ADDRESS}"]
}

ports {
  http = ${NOMAD_SERVER_HTTP_PORT}
  rpc  = ${NOMAD_SERVER_RPC_PORT}
  serf = ${NOMAD_SERVER_SERF_PORT}
}

