package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"
)

// Middleware acts as a bridge between the request and the controllers
type Middleware struct {
	l       cslog.Logger
	uuidGen fmt.Stringer
}

// NewMiddleware returns a new Middleware
func NewMiddleware(l cslog.Logger, uuidGen fmt.Stringer) (*Middleware, error) {
	if l == nil {
		return nil, errors.New("cannot have a nil logger")
	}

	if uuidGen == nil {
		return nil, errors.New("cannot have a nil uuidGen")
	}

	return &Middleware{l: l, uuidGen: uuidGen}, nil
}

func (m *Middleware) spanAndTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := cslog.GetTraceIDFromHeaders(r)
		if traceID == "" {
			traceID = m.uuidGen.String()
		}

		spanID := m.uuidGen.String()
		ctx := cslog.StoreSpanIDTraceID(r.Context(), spanID, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) logRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.l.Info(r.Context(), "started handling HTTP request",
			"uri", r.RequestURI,
			"method", r.Method,
		)
		next.ServeHTTP(w, r)
	})
}
