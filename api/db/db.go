package db

import "database/sql"

func New(development bool) (*sql.DB, error) {
	var uri string;

	if development {
		uri = "arnoldDev.db"
	} else {
		uri = "arnoldDev.db"
	}

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