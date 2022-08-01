URL=https://badrequest.localhost
http1:
	curl --http1.1 \
		--data '[]' --insecure \
		-X PUT '${URL}' \
		-H 'Content-Length: 5' -v

http2:
	curl --http2 \
		--data '[]' --insecure \
		-X PUT '${URL}' \
		-H 'Content-Length: 5' -v
