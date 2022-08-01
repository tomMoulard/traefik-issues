```bash
$ curl whoami.localhost
Hostname: 15ced77f4edc
IP: 127.0.0.1
IP: 172.24.0.2
RemoteAddr: 172.24.0.3:43984
GET / HTTP/1.1
Host: whoami.localhost
User-Agent: curl/7.68.0
Accept: application/json, application/xml, text/plain, */*
Accept-Encoding: gzip
X-Forwarded-For: 172.24.0.1
X-Forwarded-Host: whoami.localhost
X-Forwarded-Port: 80
X-Forwarded-Proto: http
X-Forwarded-Server: 52c93c91659f
X-Real-Ip: 172.24.0.1

$ curl localhost:8000
Hostname: 15ced77f4edc
IP: 127.0.0.1
IP: 172.24.0.2
RemoteAddr: 172.24.0.1:53426
GET / HTTP/1.1
Host: localhost:8000
User-Agent: curl/7.68.0
Accept: application/json, application/xml, text/plain, */*
```

Accept-Encoding: gzip
Accept: application/json, application/xml, text/plain, */*
GET / HTTP/1.1
Host: whoami.localhost
Hostname: 15ced77f4edc
IP: 127.0.0.1
IP: 172.24.0.2
RemoteAddr: 172.24.0.3:43984
User-Agent: curl/7.68.0
X-Forwarded-For: 172.24.0.1
X-Forwarded-Host: whoami.localhost
X-Forwarded-Port: 80
X-Forwarded-Proto: http
X-Forwarded-Server: 52c93c91659f
X-Real-Ip: 172.24.0.1

Accept: application/json, application/xml, text/plain, */*
GET / HTTP/1.1
Host: localhost:8000
Hostname: 15ced77f4edc
IP: 127.0.0.1
IP: 172.24.0.2
RemoteAddr: 172.24.0.1:53426
User-Agent: curl/7.68.0
