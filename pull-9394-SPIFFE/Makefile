FINGERPRINT := $(shell openssl x509 -in "./agent/agent.crt.pem" -outform DER | openssl sha1 -r | cut -d' ' -f 1)
create-workload-registration-entry:
	docker-compose exec -T spire-server \
		/opt/spire/bin/spire-server entry create \
		-parentID spiffe://example.org/spire/agent/x509pop/$(FINGERPRINT) \
		-spiffeID spiffe://example.org/workload \
		-selector unix:uid:1001 \
		-ttl 120
