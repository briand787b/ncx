package postgres

import (
	"context"
	"database/sql"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/jmoiron/sqlx"
)

type queryContextLogger struct {
	logger         cslog.Logger
	queryerContext sqlx.QueryerContext
}

func (qcl *queryContextLogger) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	logQuery(ctx, qcl.logger, query, args)
	return qcl.queryerContext.QueryContext(ctx, query, args...)
}

func (qcl *queryContextLogger) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	logQuery(ctx, qcl.logger, query, args)
	return qcl.queryerContext.QueryxContext(ctx, query, args...)
}

func (qcl *queryContextLogger) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	logQuery(ctx, qcl.logger, query, args)
	return qcl.queryerContext.QueryRowxContext(ctx, query, args...)
}
