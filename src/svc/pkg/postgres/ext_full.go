package postgres

import (
	"fmt"
	"time"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/jmoiron/sqlx"
)

// ExtFullParams are the params taken to instantiate a new ExtFull
type ExtFullParams struct {
	Host     string
	DB       string
	User     string
	Password string
	Port     string
}

// ExtFull is the interface that abstracts all the querying
// and executing functions of the sqlx package.  It is satisfied
// by the sqlx.DB type.
type ExtFull struct {
	binder
	sqlx.Execer
	sqlx.ExecerContext
	sqlx.Queryer
	sqlx.QueryerContext
	txFull
}

// NewExtFull returns an implementation of ExtFull that uses postgres
//
// WARNING: GetExtFull blocks until a successful connection to the database is achieved
// This may result in a permanent block if the application cannot connect regardless
// of number of attempts
func NewExtFull(l cslog.Logger, p *ExtFullParams) *ExtFull {
	connStr := fmt.Sprintf("sslmode=disable host=%s dbname=%s user=%s password=%s port=%s",
		p.Host,
		p.DB,
		p.User,
		p.Password,
		p.Port,
	)

	// DEBUG
	fmt.Println("DEBUG: db connection string: ", connStr)

	var err error
	if db, err = sqlx.Connect("postgres", connStr); err == nil {
		fmt.Println("connected to postgres!")
		return newExtFullFromDB(l, db)
	}

	for {
		if db, err = sqlx.Connect("postgres", connStr); err == nil {
			fmt.Println("connected to postgres!")
			return newExtFullFromDB(l, db)
		}

		fmt.Println("WARNING: database connection failed: ", err)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("retrying...")
	}
}

// NewExtFullFromTx returns an implementation of ExtFullTx backed by
// whatever database underlies the db variable provided to it.
func NewExtFullFromTx(l cslog.Logger, tx *sqlx.Tx) *ExtFull {
	return &ExtFull{
		binder: tx,
		Execer: &execLogger{
			logger: l,
			execer: db,
		},
		ExecerContext: &execContextLogger{
			logger:        l,
			execerContext: db,
		},
		Queryer: &queryLogger{
			logger:  l,
			queryer: db,
		},
		QueryerContext: &queryContextLogger{
			logger:         l,
			queryerContext: db,
		},
		txFull: &txCloser{
			l,
			*tx,
		},
	}
}

func newExtFullFromDB(l cslog.Logger, db *sqlx.DB) *ExtFull {
	return &ExtFull{
		binder: db,
		Execer: &execLogger{
			logger: l,
			execer: db,
		},
		ExecerContext: &execContextLogger{
			logger:        l,
			execerContext: db,
		},
		Queryer: &queryLogger{
			logger:  l,
			queryer: db,
		},
		QueryerContext: &queryContextLogger{
			logger:         l,
			queryerContext: db,
		},
		txFull: &txBeginner{
			l,
			*db,
		},
	}
}
