package postgres

import (
	"context"
	"database/sql"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/jmoiron/sqlx"
)

type execLogger struct {
	logger cslog.Logger
	execer sqlx.Execer
}

func (el *execLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	logQuery(context.Background(), el.logger, query, args)
	return el.execer.Exec(query, args...)
}
