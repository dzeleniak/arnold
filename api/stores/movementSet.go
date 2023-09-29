package stores

import (
	"database/sql"

	"github.com/dzeleniak/arnold/models"
)

type (
	MovementSetStore interface {
		Get(tx *sql.Tx) ([]models.MovementSet, error)
		Create(tx *sql.Tx, set models.MovementSet) (error)
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