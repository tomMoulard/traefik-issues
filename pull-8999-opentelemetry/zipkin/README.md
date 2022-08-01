# Open Telemetry support

Data sources:

 - Traces https://opentelemetry.io/docs/concepts/data-sources/#traces
 - Metrics https://opentelemetry.io/docs/concepts/data-sources/#metrics
 - Logs https://opentelemetry.io/docs/concepts/data-sources/#logs
 - baggage : for indexing observability events in one service with attributes provided by a prior service in the same transaction. :  https://opentelemetry.io/docs/concepts/data-sources/#baggage

View traces with zipkin
View metrics with prometheus
View observability with New Relic

| Exporter   | Metrics | Traces |
| ---------- | :-----: | :----: |
| Jaeger     |         | ✓      |
| OTLP       | ✓       | ✓      |
| Prometheus | ✓       |        |
| stdout     | ✓       | ✓      |
| Zipkin     |         | ✓      |
