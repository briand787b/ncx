package postgres

import (
	"context"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"
)

func logQuery(ctx context.Context, l cslog.Logger, qry string, args ...interface{}) {
	l.Query(ctx, qry,
		"args", args,
	)
}
