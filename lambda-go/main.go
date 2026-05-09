package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

	"github.com/aws/aws-lambda-go/lambda"
)

func initExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("63b9-223-181-114-155.ngrok-free.app"),
		otlptracehttp.WithInsecure(),
	)
}

const jsLambdaUrl = "https://df6sdws2d56b3cybddvhmj42gu0iukhh.lambda-url.us-east-1.on.aws/"

func handler(ctx context.Context) (string, error) {
	tracer := otel.Tracer("go-lambda-tracer")

	ctx, span := tracer.Start(ctx, "parent")
	defer span.End()

	_, child := tracer.Start(ctx, "child-span")

	// simulate work
	time.Sleep(1 * time.Second)
	child.End()

	//Call the JS lambda
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, jsLambdaUrl, nil)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	span.AddEvent("called js lambda")

	return string(body), nil

}
func main() {
	ctx := context.Background()
	// Initialize exporter
	exp, err := initExporter(ctx)
	if err != nil {
		panic(err)
	}

	//Create Trace Provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp), // Where to send the created trace
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("go-lambda"),
		)),
	)

	//Register this trace provider globally
	otel.SetTracerProvider(tp)

	defer func() {
		_ = tp.ForceFlush(ctx)
		_ = tp.Shutdown(ctx)
	}()

	lambda.Start(handler)
}
