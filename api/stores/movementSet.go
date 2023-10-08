package stores

import (
	"database/sql"

	"github.com/dzeleniak/arnold/models"
)

type (
	MovementSetStore interface {
		Get(tx *sql.Tx) ([]models.MovementSet, error)
		GetByMovementID(tx *sql.Tx) ([]models.MovementSet, error)
		Create(tx *sql.Tx, set *models.MovementSet) (int64, error)
		Update(tx *sql.Tx, set *models.MovementSet) (error)
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

		err := rows.Scan(&s.ID, &s.Weight, &s.Repetitions, &s.MovementID) 

		if err != nil {
			return nil, err;
		}

		sets = append(sets, s)
	}

	return sets, nil;
}

func (s *movementSetStore) GetByMovementID(tx *sql.Tx, id int64) ([]models.MovementSet, error) {
	sets := make([]models.MovementSet, 0)

	rows, err := s.Query("SELECT * FROM movementsets WHERE movementsets.movement_id=$1", id);

	if err != nil {
		return nil, err;
	}

	for rows.Next() {
		var s models.MovementSet;

		err := rows.Scan(&s.ID, &s.Weight, &s.Repetitions, &s.MovementID)
		if err != nil {
			return nil, err;
		}
		sets = append(sets, s)
	}
	
	return sets, nil;
}

func (s *movementSetStore) Create(tx *sql.Tx, set *models.MovementSet) (int64, error) {
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

func (s *movementSetStore) Update(tx *sql.Tx, set *models.MovementSet) (int64, error) {
	var err error;

	query, err := s.Prepare("UPDATE movementsets SET movement_id=$1, weight=$2, repetitions=$3 WHERE movementsets.id=$4 RETURNING id")
	
	if err != nil {
		return 0, err 
	}

	var id int64;

	err = query.QueryRow(set.MovementID, set.Weight, set.Repetitions, set.ID).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, sql.ErrNoRows
	} else if err != nil {
		return 0, err
	}
	
	return id, nil;
}

func (s *movementSetStore) Delete(tx *sql.Tx, id int64) error {
	res, err := s.Exec("DELETE FROM movementsets WHERE movementsets.id=$1 RETURNING movementsets.id", id);

	if err != nil {
		return err;
	}

	if r, err := res.RowsAffected(); err != nil {
		return err;
	} else if r == 0 {
		return sql.ErrNoRows;
	}

	return nil;
}