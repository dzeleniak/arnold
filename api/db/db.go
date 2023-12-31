package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func New(uri string) (*sql.DB, error) {

	db, err := sql.Open("postgres", uri);

	if err != nil {
		return nil, err;
	}

	err = db.Ping();
	
	if err != nil {
		return nil, err;
	}

	return db, nil;
}