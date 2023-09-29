package stores

import (
	"database/sql"

	"github.com/dzeleniak/arnold/models"
)

type (
	MovementSetStore interface {
		Get(tx *sql.Tx) ([]models.MovementSet, error)
		Create(tx *sql.Tx, set models.MovementSet) (int64, error)
		Update(tx *sql.Tx, set models.MovementSet) (error)
		Delete(tx *sql.Tx, id int) (error)
	}

	movementSetStore struct {
		*sql.DB
	}
)

func (s *movementSetStore) Get(tx *sql.Tx) ([]models.MovementSet, error) {
	sets := make([]models.MovementSet, 0)
	rows, err := s.Query("SELECT * FROM movementsets")

	if err != nil {
		return nil, err;
	}

	for rows.Next() {
		var s models.MovementSet;

		err := rows.Scan(&s) 

		if err != nil {
			return nil, err;
		}

		sets = append(sets, s)
	}

	return sets, nil;
}

func (s *movementSetStore) Create(tx *sql.Tx, set models.MovementSet) (int64, error) {
	var err error;

	query := "INSERT INTO movementsets (movement_id, weight, repetitions) VALUES ($1, $2, $3) RETURNING id";

	var id int64;

	if tx != nil {
		err = tx.QueryRow(query, set.MovementID, set.Weight, set.Repetitions).Scan(&id);
	} else {
		err = s.QueryRow(query, set.MovementID, set.Weight, set.Repetitions).Scan(&id);
	}

	if err != nil {
		return 0, err;
	}

	return id, nil;
}