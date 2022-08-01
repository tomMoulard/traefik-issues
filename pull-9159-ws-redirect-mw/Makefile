test:
	curl -H "Connection: upgrade" -H "Upgrade: websocket" whoami.localhost -vvv 2>&1 | grep -e 'HTTP/1.1 302 Found' -e 'Location: https://whoami.localhost/' -c | grep 2
