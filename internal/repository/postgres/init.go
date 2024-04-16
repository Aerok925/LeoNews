package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type repos struct {
	db *sqlx.DB
}

func Init(db *sql.DB) *repos {
	dbSqlx := sqlx.NewDb(db, "postgres")
	return &repos{
		dbSqlx,
	}
}
