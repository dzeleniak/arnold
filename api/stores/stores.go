package stores

import "database/sql"

type Stores struct {
	DB *sql.DB
	Movement *movementStore
	MovementSet *movementSetStore
}

func New(db *sql.DB) *Stores {
	return &Stores{
		DB: db,
		Movement: &movementStore{db},
		MovementSet: &movementSetStore{db},
	}
}

func (s *Stores) Begin() (*sql.Tx, error) {
	return s.DB.Begin();
}

func (s *Stores) Commit(tx *sql.Tx) error {
	return tx.Commit();
}

func (s *Stores) Rollback(tx *sql.Tx) error {
	return tx.Rollback();
}