package postgres

import (
	"context"
	"database/sql"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/jmoiron/sqlx"
)

type queryLogger struct {
	logger  cslog.Logger
	queryer sqlx.Queryer
}

func (ql *queryLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	logQuery(context.Background(), ql.logger, query, args)
	return ql.queryer.Query(query, args...)
}

func (ql *queryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	logQuery(context.Background(), ql.logger, query, args)
	return ql.queryer.Queryx(query, args...)
}

func (ql *queryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	logQuery(context.Background(), ql.logger, query, args)
	return ql.queryer.QueryRowx(query, args...)
}
