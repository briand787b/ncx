package cslog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
)

// CSLogger controls the format of logs written to the underylying LogWriter
type CSLogger struct {
	l       LogWriter
	uuidGen fmt.Stringer
}

// LogJSON is the structure of the JSON logs
type LogJSON struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	SpanID  string                 `json:"span_id"`
	TraceID string                 `json:"trace_id"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// NewCSLogger returns a new CSLogger holding the provided LogWriter
//
// NOTE: uuidGenerator uses a valid a default for the constructor
// if one is not provided
func NewCSLogger(lw LogWriter, uuidGenerator fmt.Stringer) *CSLogger {
	if uuidGenerator == nil {
		uuidGenerator = uuid.New()
	}

	return &CSLogger{l: lw, uuidGen: uuidGenerator}
}

// Error writes ERROR-lvl logs, but takes input as a string
func (l *CSLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	l.write(ctx, "ERROR", msg, args...)
}

// Info writes INFO-lvl logs
func (l *CSLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	l.write(ctx, "INFO", msg, args...)
}

// Invalid writes logs for failed validation events
func (l *CSLogger) Invalid(ctx context.Context, subj interface{}, reason string) {
	l.write(ctx, "INVALID", fmt.Sprintf("%T failed validation", subj),
		[]string{"reason", reason},
	)
}

// Close closes the io.Closer and logs the returned error if it is non-nil
func (l *CSLogger) Close(ctx context.Context, c io.Closer) {
	if err := c.Close(); err != nil {
		l.Error(ctx, "could not close io.Closer", "error", err.Error())
	}
}

// Query writes QUERY-lvl logs
func (l *CSLogger) Query(ctx context.Context, qry string, args ...interface{}) {
	qry = strings.NewReplacer("\t", " ", "\n", " ").Replace(qry)
	l.write(ctx, "QUERY", "", "query_string", qry, "args", args)
}

// Warn writes WARN-lvl logs
func (l *CSLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	l.write(ctx, "WARNING ", msg, args...)
}

func (l *CSLogger) write(ctx context.Context, lvl, msg string, kvs ...interface{}) {
	var spanID, traceID string
	if ctx != nil {
		var missingSpanID, missingTraceID bool
		spanID = GetSpanID(ctx)
		if spanID == "" {
			missingSpanID = true
			spanID = l.uuidGen.String()
		}

		traceID = GetTraceID(ctx)
		if traceID == "" {
			missingTraceID = true
			traceID = l.uuidGen.String()
		}

		ctx = StoreSpanIDTraceID(ctx, spanID, traceID)
		if missingSpanID {
			l.Error(ctx, "spanID is empty string", "new_span_id", spanID)
		}

		if missingTraceID {
			l.Error(ctx, "traceID is empty string", "new_trace_id", traceID)
		}
	}

	lj := LogJSON{
		Level:   lvl,
		Message: msg,
		SpanID:  spanID,
		TraceID: traceID,
	}

	if kvs != nil && len(kvs)%2 != 0 {
		l.Error(ctx, "uneven number of args provided to CSLogger", "number_of_args_provided", len(kvs))

		lj.Data = map[string]interface{}{
			"state": "invalid",
			"args":  kvs,
		}
	} else if kvs != nil {
		m := make(map[string]interface{})
		for i := 0; i < len(kvs); i += 2 {
			key, ok := (kvs[i]).(string)
			if !ok {
				l.Error(ctx, "CSLogger key not assertable to string", "key", kvs[i])
				key = fmt.Sprintf("%v", kvs[i])
			}

			m[key] = kvs[i+1]
		}

		lj.Data = m
	}

	bs, err := json.Marshal(lj)
	if err != nil {
		l.Error(ctx, "could not marshal key values to JSON", "kvs", kvs)
	}

	l.l.Println(string(bs))
}
