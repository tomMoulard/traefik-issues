TAILSCALE_VERSION := 1.26.1
ARCH := $(shell dpkg --print-architecture)
TAILSCALE_DIR := tailscale_${TAILSCALE_VERSION}_${ARCH}
TAILSCALED := ${TAILSCALE_DIR}/tailscaled
TAILSCALE := ${TAILSCALE_DIR}/tailscale
TGZ := https://pkgs.tailscale.com/stable/${TAILSCALE_DIR}.tgz

${TAILSCALED} ${TAILSCALE}:
	aria2c ${TGZ} -o ${TAILSCALE_DIR}.tgz
	tar xvf ${TAILSCALE_DIR}.tgz
	$(RM) ${TAILSCALE_DIR}.tgz

tailscaled: ${TAILSCALED}
	sudo ${TAILSCALED} --state=tailscaled.state

tailscale: ${TAILSCALE}
	sudo ${TAILSCALE} up

ip:
	${TAILSCALE} ip -4

clean: ${TAILSCALED}
	sudo ${TAILSCALED} -cleanup
	$(RM) -r tailscaled.state ${TAILSCALED} ${TAILSCALE} ${TAILSCALE_DIR} ${TAILSCALE_DIR}.tgz
