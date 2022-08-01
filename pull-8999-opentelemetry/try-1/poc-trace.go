package main

// try-1-otel-collector-1  | 2022-04-22T14:08:01.583Z	INFO	loggingexporter/logging_exporter.go:41	TracesExporter	{"#spans": 3}
// try-1-otel-collector-1  | 2022-04-22T14:08:01.583Z	DEBUG	loggingexporter/logging_exporter.go:51	ResourceSpans #0
// try-1-otel-collector-1  | Resource labels:
// try-1-otel-collector-1  |      -> service.name: STRING(otlptrace-example)
// try-1-otel-collector-1  |      -> service.version: STRING(0.0.1)
// try-1-otel-collector-1  | InstrumentationLibrarySpans #0
// try-1-otel-collector-1  | InstrumentationLibrary github.com/instrumentron v0.1.0
// try-1-otel-collector-1  | Span #0
// try-1-otel-collector-1  |     Trace ID       : f1da42951cf1b2d6161f39b682ed2d14
// try-1-otel-collector-1  |     Parent ID      :
// try-1-otel-collector-1  |     ID             : ca84c1c4847866d6
// try-1-otel-collector-1  |     Name           : Multiplication
// try-1-otel-collector-1  |     Kind           : SPAN_KIND_INTERNAL
// try-1-otel-collector-1  |     Start time     : 2022-04-22 14:08:01.582277925 +0000 UTC
// try-1-otel-collector-1  |     End time       : 2022-04-22 14:08:01.582283482 +0000 UTC
// try-1-otel-collector-1  |     Status code    : STATUS_CODE_UNSET
// try-1-otel-collector-1  |     Status message :
// try-1-otel-collector-1  | Span #1
// try-1-otel-collector-1  |     Trace ID       : 566a42f6e90af6b654a118d903f6fdab
// try-1-otel-collector-1  |     Parent ID      :
// try-1-otel-collector-1  |     ID             : 998cc5580627e538
// try-1-otel-collector-1  |     Name           : Multiplication
// try-1-otel-collector-1  |     Kind           : SPAN_KIND_INTERNAL
// try-1-otel-collector-1  |     Start time     : 2022-04-22 14:08:01.582286412 +0000 UTC
// try-1-otel-collector-1  |     End time       : 2022-04-22 14:08:01.58228678 +0000 UTC
// try-1-otel-collector-1  |     Status code    : STATUS_CODE_UNSET
// try-1-otel-collector-1  |     Status message :
// try-1-otel-collector-1  | Span #2
// try-1-otel-collector-1  |     Trace ID       : aafe57ffc3f6b493ae0625e387c63577
// try-1-otel-collector-1  |     Parent ID      :
// try-1-otel-collector-1  |     ID             : f11875aef8e634c1
// try-1-otel-collector-1  |     Name           : Addition
// try-1-otel-collector-1  |     Kind           : SPAN_KIND_INTERNAL
// try-1-otel-collector-1  |     Start time     : 2022-04-22 14:08:01.582287429 +0000 UTC
// try-1-otel-collector-1  |     End time       : 2022-04-22 14:08:01.582287691 +0000 UTC
// try-1-otel-collector-1  |     Status code    : STATUS_CODE_UNSET
// try-1-otel-collector-1  |     Status message :
// try-1-otel-collector-1  |

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	instrumentationName    = "github.com/instrumentron"
	instrumentationVersion = "v0.1.0"
)

var (
	tracer = otel.GetTracerProvider().Tracer(
		instrumentationName,
		trace.WithInstrumentationVersion(instrumentationVersion),
		trace.WithSchemaURL(semconv.SchemaURL),
	)
)

func add(ctx context.Context, x, y int64) int64 {
	var span trace.Span
	_, span = tracer.Start(ctx, "Addition")
	defer span.End()

	return x + y
}

func multiply(ctx context.Context, x, y int64) int64 {
	var span trace.Span
	_, span = tracer.Start(ctx, "Multiplication")
	defer span.End()

	return x * y
}

func newResource() *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("otlptrace-example"),
		semconv.ServiceVersionKey.String("0.0.1"),
	)
}

func installExportPipeline(ctx context.Context) func() {
	client := otlptracehttp.NewClient(otlptracehttp.WithInsecure())
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatalf("creating OTLP trace exporter: %v", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(newResource()),
	)
	otel.SetTracerProvider(tracerProvider)

	return func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			log.Fatalf("stopping tracer provider: %v", err)
		}
	}
}

func main() {
	ctx := context.Background()
	// Registers a tracer Provider globally.
	cleanup := installExportPipeline(ctx)
	defer cleanup()

	log.Println("the answer is", add(ctx, multiply(ctx, multiply(ctx, 2, 2), 10), 2))
}
