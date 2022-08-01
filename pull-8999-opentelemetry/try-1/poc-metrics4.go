package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/unit"
	otelHistogram "go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
)

func main() {
	ctx := context.Background()
	client := otlpmetricgrpc.NewClient(otlpmetricgrpc.WithEndpoint("localhost:4317"), otlpmetricgrpc.WithInsecure())
	exp, err := otlpmetric.New(ctx, client)
	if err != nil {
		log.Fatalf("Failed to create the collector exporter: %v", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err := exp.Shutdown(ctx); err != nil {
			otel.Handle(err)
		}
	}()

	optsResource := []resource.Option{
		// resource.WithSchemaURL(),
		resource.WithAttributes(),
		resource.WithContainer(),
		resource.WithContainerID(),
		resource.WithDetectors(),
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithOS(),
		resource.WithOSDescription(),
		resource.WithOSType(),
		resource.WithProcess(),
		resource.WithProcessCommandArgs(),
		resource.WithProcessExecutableName(),
		resource.WithProcessExecutablePath(),
		resource.WithProcessOwner(),
		resource.WithProcessPID(),
		resource.WithProcessRuntimeDescription(),
		resource.WithProcessRuntimeName(),
		resource.WithProcessRuntimeVersion(),
		resource.WithTelemetrySDK(),
	}

	r, err := resource.New(ctx, optsResource...)
	if err != nil {
		panic(err)
	}

	pusher := controller.New(
		processor.NewFactory(
			simple.NewWithHistogramDistribution(
				otelHistogram.WithExplicitBoundaries([]float64{0.001, 0.01, 0.1, 0.5, 1, 2, 5, 10}),
			),
			exp,
		),
		controller.WithExporter(exp),
		controller.WithCollectPeriod(10*time.Second),
		controller.WithResource(r),
	)

	global.SetMeterProvider(pusher)

	if err := pusher.Start(ctx); err != nil {
		log.Fatalf("could not start metric controller: %v", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		// pushes any last exports to the receiver
		if err := pusher.Stop(ctx); err != nil {
			otel.Handle(err)
		}
	}()

	meter := global.Meter("", metric.WithInstrumentationVersion("v0.42.42"))

	// Recorder metric example

	// counter, _ := meter.SyncFloat64().Counter("an_important_metric", instrument.WithDescription("Measures the cumulative epicness of the app"))
	histo, _ := meter.SyncFloat64().Histogram("some_histo", instrument.WithDescription("woaw"), instrument.WithUnit(unit.Milliseconds))
	// gauge, _ := meter.AsyncFloat64().Gauge("such_gauge", instrument.WithDescription("nice"), instrument.WithUnit(unit.Milliseconds))

	for i := 0; i < 10; i++ {
		log.Printf("Doing really hard work (%d / 10) (%q)\n", i+1, fmt.Sprintf("%d", i))
		// time.Sleep(time.Second)
		// counter.Add(ctx, 1.0, attribute.String("foo", "bar-"+fmt.Sprintf("%d", i)))
		histo.Record(context.Background(), float64(i), attribute.String("baz", "buz-"+fmt.Sprintf("%d", i)))
		// gauge.Observe(context.Background(), float64(i), attribute.String("toto", "tata-"+fmt.Sprintf("%d", i)))
	}
}
