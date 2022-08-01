# notes

ENV VARS with setDefault -> env var are applied before config set by Traefik
https://github.com/open-telemetry/opentelemetry-go/blob/exporters/otlp/otlptrace/otlptracehttp/v1.7.0/exporters/otlp/otlptrace/otlptracehttp/client.go#L80

## Status and Releases

The current status of the major functional components for OpenTelemetry Go is as follows:

| Tracing | Metrics | Logging |
| ------- | ------- | ------- |
| Stable  | Alpha   | Not Yet Implemented |

# SPAN OUTPUT

## poc-trace.go -> otel collector

```
try-1-otel-collector-1  | 2022-04-22T14:08:01.583Z	INFO	loggingexporter/logging_exporter.go:41	TracesExporter	{"#spans": 3}
try-1-otel-collector-1  | 2022-04-22T14:08:01.583Z	DEBUG	loggingexporter/logging_exporter.go:51	ResourceSpans #0
try-1-otel-collector-1  | Resource labels:
try-1-otel-collector-1  |      -> service.name: STRING(otlptrace-example)
try-1-otel-collector-1  |      -> service.version: STRING(0.0.1)
try-1-otel-collector-1  | InstrumentationLibrarySpans #0
try-1-otel-collector-1  | InstrumentationLibrary github.com/instrumentron v0.1.0
try-1-otel-collector-1  | Span #0
try-1-otel-collector-1  |     Trace ID       : f1da42951cf1b2d6161f39b682ed2d14
try-1-otel-collector-1  |     Parent ID      :
try-1-otel-collector-1  |     ID             : ca84c1c4847866d6
try-1-otel-collector-1  |     Name           : Multiplication
try-1-otel-collector-1  |     Kind           : SPAN_KIND_INTERNAL
try-1-otel-collector-1  |     Start time     : 2022-04-22 14:08:01.582277925 +0000 UTC
try-1-otel-collector-1  |     End time       : 2022-04-22 14:08:01.582283482 +0000 UTC
try-1-otel-collector-1  |     Status code    : STATUS_CODE_UNSET
try-1-otel-collector-1  |     Status message :
try-1-otel-collector-1  | Span #1
try-1-otel-collector-1  |     Trace ID       : 566a42f6e90af6b654a118d903f6fdab
try-1-otel-collector-1  |     Parent ID      :
try-1-otel-collector-1  |     ID             : 998cc5580627e538
try-1-otel-collector-1  |     Name           : Multiplication
try-1-otel-collector-1  |     Kind           : SPAN_KIND_INTERNAL
try-1-otel-collector-1  |     Start time     : 2022-04-22 14:08:01.582286412 +0000 UTC
try-1-otel-collector-1  |     End time       : 2022-04-22 14:08:01.58228678 +0000 UTC
try-1-otel-collector-1  |     Status code    : STATUS_CODE_UNSET
try-1-otel-collector-1  |     Status message :
try-1-otel-collector-1  | Span #2
try-1-otel-collector-1  |     Trace ID       : aafe57ffc3f6b493ae0625e387c63577
try-1-otel-collector-1  |     Parent ID      :
try-1-otel-collector-1  |     ID             : f11875aef8e634c1
try-1-otel-collector-1  |     Name           : Addition
try-1-otel-collector-1  |     Kind           : SPAN_KIND_INTERNAL
try-1-otel-collector-1  |     Start time     : 2022-04-22 14:08:01.582287429 +0000 UTC
try-1-otel-collector-1  |     End time       : 2022-04-22 14:08:01.582287691 +0000 UTC
try-1-otel-collector-1  |     Status code    : STATUS_CODE_UNSET
try-1-otel-collector-1  |     Status message :
try-1-otel-collector-1  |
```

## Traefik -> otel collector

```
try-1-otel-collector-1  | 2022-04-25T08:02:18.451Z	INFO	loggingexporter/logging_exporter.go:41	TracesExporter	{"#spans": 2}
try-1-otel-collector-1  | 2022-04-25T08:02:18.452Z	DEBUG	loggingexporter/logging_exporter.go:51	ResourceSpans #0
try-1-otel-collector-1  | Resource labels:
try-1-otel-collector-1  |      -> service.name: STRING(unknown_service:___traefik_OTEL_docker)
try-1-otel-collector-1  |      -> telemetry.sdk.language: STRING(go)
try-1-otel-collector-1  |      -> telemetry.sdk.name: STRING(opentelemetry)
try-1-otel-collector-1  |      -> telemetry.sdk.version: STRING(1.6.3)
try-1-otel-collector-1  | InstrumentationLibrarySpans #0
try-1-otel-collector-1  | InstrumentationLibrary traefik dev
try-1-otel-collector-1  | Span #0
try-1-otel-collector-1  |     Trace ID       : a8774f481b33daf4b51e97a9b2652456
try-1-otel-collector-1  |     Parent ID      : dce580a15efd4547
try-1-otel-collector-1  |     ID             : 85c465904ddff66d
try-1-otel-collector-1  |     Name           : forward whoami/whoami@docker
try-1-otel-collector-1  |     Kind           : SPAN_KIND_INTERNAL
try-1-otel-collector-1  |     Start time     : 2022-04-25 08:02:18.433654357 +0000 UTC
try-1-otel-collector-1  |     End time       : 2022-04-25 08:02:18.436612021 +0000 UTC
try-1-otel-collector-1  |     Status code    : STATUS_CODE_UNSET
try-1-otel-collector-1  |     Status message :
try-1-otel-collector-1  | Attributes:
try-1-otel-collector-1  |      -> traefik.service.name: STRING(whoami)
try-1-otel-collector-1  |      -> traefik.router.name: STRING(whoami@docker)
try-1-otel-collector-1  |      -> http.method: STRING(GET)
try-1-otel-collector-1  |      -> http.url: STRING(/)
try-1-otel-collector-1  |      -> http.host: STRING(whoami.localhost:8001)
try-1-otel-collector-1  |      -> http.status_code: STRING(200)
try-1-otel-collector-1  | Span #1
try-1-otel-collector-1  |     Trace ID       : a8774f481b33daf4b51e97a9b2652456
try-1-otel-collector-1  |     Parent ID      :
try-1-otel-collector-1  |     ID             : dce580a15efd4547
try-1-otel-collector-1  |     Name           : EntryPoint web whoami.localhost:8001
try-1-otel-collector-1  |     Kind           : SPAN_KIND_INTERNAL
try-1-otel-collector-1  |     Start time     : 2022-04-25 08:02:18.433634613 +0000 UTC
try-1-otel-collector-1  |     End time       : 2022-04-25 08:02:18.436618242 +0000 UTC
try-1-otel-collector-1  |     Status code    : STATUS_CODE_UNSET
try-1-otel-collector-1  |     Status message :
try-1-otel-collector-1  | Attributes:
try-1-otel-collector-1  |      -> component: STRING(traefik)
try-1-otel-collector-1  |      -> http.method: STRING(GET)
try-1-otel-collector-1  |      -> http.url: STRING(/)
try-1-otel-collector-1  |      -> http.host: STRING(whoami.localhost:8001)
try-1-otel-collector-1  |      -> http.status_code: STRING(200)
try-1-otel-collector-1  |
```


# temp
```
[method protocol entrypoint entrypoint web method GET protocol http]                   -> [{Key:method Value:{vtype:4 numeric:0 stringly:protocol slice:<nil>}} {Key:entrypoint Value:{vtype:4 numeric:0 stringly:entrypoint slice:<nil>}} {Key:web Value:{vtype:4 numeric:0 stringly:method slice:<nil>}} {Key:GET Value:{vtype:4 numeric:0 stringly:protocol slice:<nil>}}]
[method protocol service service whoami@docker method GET protocol http]               -> [{Key:method Value:{vtype:4 numeric:0 stringly:protocol slice:<nil>}} {Key:service Value:{vtype:4 numeric:0 stringly:service slice:<nil>}} {Key:whoami@docker Value:{vtype:4 numeric:0 stringly:method slice:<nil>}} {Key:GET Value:{vtype:4 numeric:0 stringly:protocol slice:<nil>}}]
[code method protocol service service whoami@docker method GET protocol http code 200] -> [{Key:code Value:{vtype:4 numeric:0 stringly:method slice:<nil>}} {Key:protocol Value:{vtype:4 numeric:0 stringly:service slice:<nil>}} {Key:service Value:{vtype:4 numeric:0 stringly:whoami@docker slice:<nil>}} {Key:method Value:{vtype:4 numeric:0 stringly:GET slice:<nil>}} {Key:protocol Value:{vtype:4 numeric:0 stringly:http slice:<nil>}} {Key:code Value:{vtype:4 numeric:0 stringly:200 slice:<nil>}}]
[code method protocol service service whoami@docker method GET protocol http code 200] -> [{Key:code Value:{vtype:4 numeric:0 stringly:method slice:<nil>}} {Key:protocol Value:{vtype:4 numeric:0 stringly:service slice:<nil>}} {Key:service Value:{vtype:4 numeric:0 stringly:whoami@docker slice:<nil>}} {Key:method Value:{vtype:4 numeric:0 stringly:GET slice:<nil>}} {Key:protocol Value:{vtype:4 numeric:0 stringly:http slice:<nil>}} {Key:code Value:{vtype:4 numeric:0 stringly:200 slice:<nil>}}]
[method protocol service service whoami@docker method GET protocol http]               -> [{Key:method Value:{vtype:4 numeric:0 stringly:protocol slice:<nil>}} {Key:service Value:{vtype:4 numeric:0 stringly:service slice:<nil>}} {Key:whoami@docker Value:{vtype:4 numeric:0 stringly:method slice:<nil>}} {Key:GET Value:{vtype:4 numeric:0 stringly:protocol slice:<nil>}}]
[code method protocol entrypoint entrypoint web method GET protocol http code 200]     -> [{Key:code Value:{vtype:4 numeric:0 stringly:method slice:<nil>}} {Key:protocol Value:{vtype:4 numeric:0 stringly:entrypoint slice:<nil>}} {Key:entrypoint Value:{vtype:4 numeric:0 stringly:web slice:<nil>}} {Key:method Value:{vtype:4 numeric:0 stringly:GET slice:<nil>}} {Key:protocol Value:{vtype:4 numeric:0 stringly:http slice:<nil>}} {Key:code Value:{vtype:4 numeric:0 stringly:200 slice:<nil>}}]
[code method protocol entrypoint entrypoint web method GET protocol http code 200]     -> [{Key:code Value:{vtype:4 numeric:0 stringly:method slice:<nil>}} {Key:protocol Value:{vtype:4 numeric:0 stringly:entrypoint slice:<nil>}} {Key:entrypoint Value:{vtype:4 numeric:0 stringly:web slice:<nil>}} {Key:method Value:{vtype:4 numeric:0 stringly:GET slice:<nil>}} {Key:protocol Value:{vtype:4 numeric:0 stringly:http slice:<nil>}} {Key:code Value:{vtype:4 numeric:0 stringly:200 slice:<nil>}}]
[method protocol entrypoint entrypoint web method GET protocol http]                   -> [{Key:method Value:{vtype:4 numeric:0 stringly:protocol slice:<nil>}} {Key:entrypoint Value:{vtype:4 numeric:0 stringly:entrypoint slice:<nil>}} {Key:web Value:{vtype:4 numeric:0 stringly:method slice:<nil>}} {Key:GET Value:{vtype:4 numeric:0 stringly:protocol slice:<nil>}}]
```
