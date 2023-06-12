package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
	sdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func main() {
	ctx := context.Background()

	// Instantiate the OTLP resources
	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("local"),
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

	gauge, _ := meter.Float64ObservableGauge(
		"cpu_gauge",
		instrument.WithDescription("Returns the current cpu usage as a percentage"),
		instrument.WithUnit("percentage"),
	)
	attrs := []attribute.KeyValue{
		attribute.Key("A").String("B"),
		attribute.Key("C").String("D"),
	}
	_, err = meter.RegisterCallback(func(_ context.Context, o metric.Observer) error {
		n := -10. + rand.New(rand.NewSource(time.Now().UnixNano())).Float64()*(90.) // [-10, 100)
		o.ObserveFloat64(gauge, n, attrs...)
		return nil
	}, gauge)
	if err != nil {
		log.Fatal(err)
	}
}
