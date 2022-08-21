package rest

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strconv"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// ServerParams structures the parameters for a new Server
type ServerParams struct {
	Port   string
	Logger cslog.Logger
	// StatsStore stats.Store
}

// Server holds the variables for running a server
type Server struct {
	port int
	l    cslog.Logger
	// statsStore stats.Store
}

// NewServer instantiates a new Server
func NewServer(p *ServerParams) (*Server, error) {
	port, err := strconv.Atoi(p.Port)
	if err != nil {

	}
	return &Server{
		port: port,
		l:    p.Logger,
		// statsStore: p.StatsStore,
	}, nil
}

// Run runs the server
func (s *Server) Run() error {
	mw, err := NewMiddleware(s.l, uuid.New())
	if err != nil {
		return errors.Wrap(err, "could not create new middleware")
	}

	// sc := NewStatsController(s.l, s.statsStore)

	// initialize router
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(mw.spanAndTrace)
	r.Use(mw.logRoute)

	r.Route("/v1.0", func(r chi.Router) {
		// r.Get("/locations/{location_id}/items/most-ordered", sc.handleGetMostOrderedItemByLocationID)
	})

	ctx := cslog.StoreSpanIDTraceID(context.Background(), "main", "main")
	walkFn := func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		s.l.Info(ctx, "API Route",
			"method", method,
			"route", route,
			"handler", funcName,
		)
		return nil
	}

	if err := chi.Walk(r, walkFn); err != nil {
		return errors.Wrap(err, "could not print API routes")
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%v", s.port), r); err != nil {
		return errors.Wrap(err, "could not listen and serve")
	}

	return nil
}
