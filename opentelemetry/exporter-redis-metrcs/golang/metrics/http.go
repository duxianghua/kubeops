package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/unit"
	sdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func main() {
	ctx := context.Background()

	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("service"),
		semconv.ServiceVersionKey.String("v0.0.0"),
	)

	// Instantiate the OTLP HTTP exporter
	exporter, err := otlpmetrichttp.New(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Instantiate the OTLP HTTP exporter
	meterProvider := sdk.NewMeterProvider(
		sdk.WithResource(resources),
		sdk.WithReader(sdk.NewPeriodicReader(exporter)),
	)
	defer func() {
		err := meterProvider.Shutdown(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Create an instance on a meter for the given instrumentation scope
	meter := meterProvider.Meter(
		"github.com/.../example/manual-instrumentation",
		metric.WithInstrumentationVersion("v0.0.0"),
	)

	// Create two synchronous instruments: counter and histogram
	requestCount, err := meter.Int64Counter(
		"request_count",
		instrument.WithDescription("Incoming request count"),
		instrument.WithUnit("request"),
	)
	if err != nil {
		log.Fatalln(err)
	}
	requestDuration, err := meter.Float64Histogram(
		"duration",
		instrument.WithDescription("Incoming end to end duration"),
		instrument.WithUnit(string(unit.Milliseconds)),
	)
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		requestStartTime := time.Now()

		_, _ = w.Write([]byte("UP"))

		elapsedTime := float64(time.Since(requestStartTime)) / float64(time.Millisecond)

		// Record measurements
		attrs := semconv.HTTPServerMetricAttributesFromHTTPRequest("", req)
		requestCount.Add(ctx, 1, attrs...)
		requestDuration.Record(ctx, elapsedTime, attrs...)
	})

	http.ListenAndServe(":8081", nil)
}
