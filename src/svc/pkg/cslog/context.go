package cslog

import (
	"context"
	"net/http"
)

// CtxKey is the type for all ctx keys in the cslog package
type CtxKey string

const (
	// SpanIDCtxKey keys the SpanID in the context
	SpanIDCtxKey CtxKey = "span_id"

	// TraceIDCtxKey keys the TraceID in the context
	TraceIDCtxKey CtxKey = "trace_id"
)

// StoreSpanID returns a context with a stored spanID
func StoreSpanID(ctx context.Context, spanID string) context.Context {
	return context.WithValue(ctx, SpanIDCtxKey, spanID)
}

// StoreTraceID returns a context with a stored traceID
func StoreTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDCtxKey, traceID)
}

// StoreSpanIDTraceID store both SpanID and TraceID
func StoreSpanIDTraceID(ctx context.Context, spanID, traceID string) context.Context {
	return StoreTraceID(StoreSpanID(ctx, spanID), traceID)
}

// GetSpanIDFromHeaders returns the span ID from the request headers
func GetSpanIDFromHeaders(r *http.Request) string {
	return r.Header.Get(string(SpanIDCtxKey))
}

// GetTraceIDFromHeaders returns the trace ID from the request headers
func GetTraceIDFromHeaders(r *http.Request) string {
	return r.Header.Get(string(TraceIDCtxKey))
}

// SetSpanIDAndTraceIDInHeaders sets headers in the request to pass along
// the span and trace IDs in the context
func SetSpanIDAndTraceIDInHeaders(ctx context.Context, r *http.Request) {
	r.Header.Set(string(SpanIDCtxKey), GetSpanID(ctx))
	r.Header.Set(string(TraceIDCtxKey), GetTraceID(ctx))
}

// GetSpanID returns the current spanID in the context
func GetSpanID(ctx context.Context) string {
	sID, _ := ctx.Value(SpanIDCtxKey).(string)
	return sID
}

// GetTraceID returns the current traceID in the context
func GetTraceID(ctx context.Context) string {
	tID, _ := ctx.Value(TraceIDCtxKey).(string)
	return tID
}
