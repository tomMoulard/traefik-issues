OPTIONS ?= --concurrency=5000 --total=400000 --call whoami.Whoami.AskWhoAmI -d '{"name":"tata"}'
test:
	ghz frontend.localhost:80 ${OPTIONS} \
		--insecure
	ghz secured.frontend.localhost:443 ${OPTIONS}
