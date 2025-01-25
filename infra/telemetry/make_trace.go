package telemetry

import (
	"context"
	"regexp"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func MakeTrace(ctx context.Context, structName string, actionName string) (context.Context, trace.Span) {

	tr := otel.Tracer(structName)
	return tr.Start(ctx, actionName)
}

func MakeTraceCall(ctx context.Context) (context.Context, trace.Span) {

	pc, file, line, _ := runtime.Caller(1)
	fullFunctionName := runtime.FuncForPC(pc).Name()

	re := regexp.MustCompile(`\(\*(\w+)\)\.(\w+)`)
	matches := re.FindStringSubmatch(fullFunctionName)

	structName := matches[1]
	actionName := matches[2]

	tr := otel.Tracer(actionName)
	ctx, span := tr.Start(ctx, structName+"/"+actionName)

	span.SetAttributes(
		attribute.KeyValue{
			Key: "source.file", Value: attribute.StringValue(file),
		},
		attribute.KeyValue{
			Key: "source.line", Value: attribute.IntValue(line),
		})

	return ctx, span
}
