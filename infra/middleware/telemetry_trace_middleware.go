package middleware

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func TelemetryTraceMiddleware(next http.Handler, ctx context.Context) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tr := otel.Tracer("http-server")
		ctx, span := tr.Start(ctx, "HTTP "+r.Method+" "+r.URL.Path)
		defer span.End()

		span.SetAttributes(
			attribute.String("http.method", r.Method),
			attribute.String("http.url", r.URL.String()),
		)

		reqWithCtx := r.WithContext(ctx)
		next.ServeHTTP(w, reqWithCtx)
	})
}
