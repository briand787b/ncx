package postgres

import (
	"context"
	"database/sql"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/jmoiron/sqlx"
)

type execContextLogger struct {
	logger        cslog.Logger
	execerContext sqlx.ExecerContext
}

func (ecl *execContextLogger) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	logQuery(ctx, ecl.logger, query, args)
	return ecl.execerContext.ExecContext(ctx, query, args)
}
