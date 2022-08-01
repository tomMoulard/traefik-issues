package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/unit"
)

// type instruments struct {
// reqCounter metric.Int64Counter
// reqGauge   metric.Int64UpDownCounter
// }

func main() {
	// In a library or program this would be provided by otel.GetMeterProvider().
	// meterProvider := otel.GetMeterProvider()
	// meterProvider := nonrecording.NewNoopMeterProvider()
	meterProvider := global.MeterProvider()
	meter := meterProvider.Meter("go.opentelemetry.io/otel/metric#AsyncExample")

	memoryUsage, err := meter.AsyncInt64().Gauge(
		"MemoryUsage",
		instrument.WithUnit(unit.Bytes),
	)
	if err != nil {
		fmt.Println("Failed to register instrument")
		panic(err)
	}

	err = meter.RegisterCallback([]instrument.Asynchronous{memoryUsage},
		func(ctx context.Context) {
			// instrument.WithCallbackFunc(func(ctx context.Context) {
			// Do Work to get the real memoryUsage
			// mem := GatherMemory(ctx)
			mem := 75000

			memoryUsage.Observe(ctx, int64(mem))
		})
	if err != nil {
		fmt.Println("Failed to register callback")
		panic(err)
	}

	memoryUsage.Observe(context.Background(), 64)
}

// func newInstruments(meter metric.Meter) *instruments {
// // meter = global.Meter("io.opentelemetry.metrics.hello")
// // instrument := newInstruments(metter)
// // instruments.reqGauge.Add(ctx, 1,
// // label.String("method", "GET"),
// // label.String("route", "/"),
// // )
// mm := metric.Must(meter)

// return &instruments{
// reqCounter: mm.NewInt64Counter(
// "http_requests_total",
// metric.WithDescription("The total number of http requests"),
// metric.WithUnit(unit.Dimensionless),
// ),
// reqGauge: mm.NewInt64UpDownCounter(
// "http_requests_active",
// metric.WithDescription("The number of in-flight http requests"),
// metric.WithUnit(unit.Dimensionless),
// ),
// }
// }

// type instruments struct {
// 	// queryCount is the number of queries executed.
// 	queryCount syncint64.Counter
//
// 	// queryRows is the number of rows returned by a query.
// 	queryRows syncint64.Histogram
//
// 	// batchCount is the number of batch queries executed.
// 	batchCount syncint64.Counter
//
// 	// connectionCount is the number of connections made
// 	// with the traced session.
// 	connectionCount syncint64.Counter
//
// 	// latency is the sum of attempt latencies.
// 	latency syncint64.Histogram
// }
//
// // newInstruments will create instruments using a meter
// // from the given provider p.
// func newInstruments(p metric.MeterProvider) *instruments {
// 	meter := p.Meter(
// 		"",
// 		metric.WithInstrumentationVersion(SemVersion()),
// 	)
// 	instruments := &instruments{}
// 	var err error
//
// 	if instruments.queryCount, err = meter.SyncInt64().Counter(
// 		"db.cassandra.queries",
// 		instrument.WithDescription("Number queries executed"),
// 	); err != nil {
// 		log.Printf("failed to create iQueryCount instrument, %v", err)
// 	}
//
// 	if instruments.queryRows, err = meter.SyncInt64().Histogram(
// 		"db.cassandra.rows",
// 		instrument.WithDescription("Number of rows returned from query"),
// 	); err != nil {
// 		log.Printf("failed to create iQueryRows instrument, %v", err)
// 	}
//
// 	if instruments.batchCount, err = meter.SyncInt64().Counter(
// 		"db.cassandra.batch.queries",
// 		instrument.WithDescription("Number of batch queries executed"),
// 	); err != nil {
// 		log.Printf("failed to create iBatchCount instrument, %v", err)
// 	}
//
// 	if instruments.connectionCount, err = meter.SyncInt64().Counter(
// 		"db.cassandra.connections",
// 		instrument.WithDescription("Number of connections created"),
// 	); err != nil {
// 		log.Printf("failed to create iConnectionCount instrument, %v", err)
// 	}
//
// 	if instruments.latency, err = meter.SyncInt64().Histogram(
// 		"db.cassandra.latency",
// 		instrument.WithDescription("Sum of latency to host in milliseconds"),
// 		instrument.WithUnit(unit.Milliseconds),
// 	); err != nil {
// 		log.Printf("failed to create iLatency instrument, %v", err)
// 	}
//
// 	return instruments
// }
