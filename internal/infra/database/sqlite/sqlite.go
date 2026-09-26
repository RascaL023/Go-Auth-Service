package sqlite

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewConnection(ctx context.Context, dbPathName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPathName)
	if err != nil {
		return  nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}
