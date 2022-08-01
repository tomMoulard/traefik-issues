# Response
Hello @mrbrown-09,

Thanks for your interest in Traefik!

As the RFC7230 states, you cannot have both `Content-Length` and `Transfer-Encoding` headers in the same message:

> A sender must not send a `Content-Length` header field in any message that
> contains a `Transfer-Encoding` header field.
> ~ https://greenbytes.de/tech/webdav/rfc7230.html#rfc.section.3.3.2.p.3

Also, as `Transfer-Encoding` [is](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Transfer-Encoding) a `hop-by-hop header`, the golang reverse proxy removes it:

https://github.com/golang/go/blob/master/src/net/http/httputil/reverseproxy.go#L174-L184

https://github.com/golang/go/blob/master/src/net/http/httputil/reverseproxy.go#L305-L307

Moreover, I could not see the `Transfer-Encoding` header on traefik 2.1.4.

Accordingly, the issue will be tagged as a question.
If you feel that I'm missing something here, please let me know.

# TODO

Check `content-length` header value

# Notes
## Custom RP
```bash
> curl -I 'localhost:8000/custom?hk=Content-Length&hv=424242&hk=Transfer-Encoding&hv=chunked'
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Date: Fri, 26 Feb 2021 10:32:49 GMT

> curl -I 'localhost:8080/custom?hk=Content-Length&hv=424242&hk=Transfer-Encoding&hv=chunked'
HTTP/1.1 200 OK
Date: Fri, 26 Feb 2021 10:32:56 GMT
X-Testing: Custom-RP
```

## no traefik
```bash
> curl -I localhost:8000
HTTP/1.1 200 OK
Date: Tue, 16 Feb 2021 15:17:06 GMT
Content-Length: 52
Content-Type: text/plain; charset=utf-8

> curl -I localhost:8000/both
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Date: Tue, 16 Feb 2021 15:17:09 GMT

> curl -I localhost:8000/cl
HTTP/1.1 200 OK
Content-Length: 4242424242
Date: Tue, 16 Feb 2021 15:17:14 GMT
Content-Type: text/plain; charset=utf-8

> curl -I localhost:8000/te
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Date: Tue, 16 Feb 2021 15:17:16 GMT

```
## Version 2.1.4
```bash
> curl -I app.localhost
HTTP/1.1 200 OK
Content-Length: 52
Content-Type: text/plain; charset=utf-8
Date: Tue, 16 Feb 2021 15:20:19 GMT

> curl -I app.localhost/both
HTTP/1.1 200 OK
Date: Tue, 16 Feb 2021 15:20:22 GMT

> curl -I app.localhost/cl
HTTP/1.1 200 OK
Content-Length: 4242424242
Content-Type: text/plain; charset=utf-8
Date: Tue, 16 Feb 2021 15:20:26 GMT

> curl -I app.localhost/te
HTTP/1.1 200 OK
Date: Tue, 16 Feb 2021 15:20:28 GMT

```

# Traefik 2.4.2
```bash
> curl -I app.localhost
HTTP/1.1 200 OK
Content-Length: 52
Content-Type: text/plain; charset=utf-8
Date: Tue, 16 Feb 2021 15:24:39 GMT

> curl -I app.localhost/both
HTTP/1.1 200 OK
Date: Tue, 16 Feb 2021 15:24:41 GMT

> curl -I app.localhost/cl
HTTP/1.1 200 OK
Content-Length: 4242424242
Content-Type: text/plain; charset=utf-8
Date: Tue, 16 Feb 2021 15:24:44 GMT

> curl -I app.localhost/te
HTTP/1.1 200 OK
Date: Tue, 16 Feb 2021 15:24:46 GMT

```
