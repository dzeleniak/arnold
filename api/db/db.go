package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New(uri string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", uri);

	if err != nil {
		return nil, err;
	}

	err = db.Ping();
	
	if err != nil {
		return nil, err;
	}

	return db, nil;
}