package stores

import (
	"database/sql"

	"github.com/dzeleniak/arnold/models"
)

type (
	MovementStore interface {
		Get(tx *sql.Tx) ([]models.Movement, error)
		Create(tx *sql.Tx, movement *models.Movement) (int64, error)
		UpdateById(tx *sql.Tx, movement *models.Movement)
		DeleteById(tx *sql.Tx, id int) error
	}

	movementStore struct {
		*sql.DB
	}
)

func (s *movementStore) Get(tx *sql.Tx) ([]models.Movement, error) {
	movements := make([]models.Movement, 0);
	rows, err := s.Query("SELECT * FROM movements")

	if err != nil {
		return nil, err;
	}

	for rows.Next() {
		var m models.Movement;
		err = rows.Scan(&m.ID, &m.Name);

		if  err != nil {
			return nil, err;
		}
		
		movements = append(movements, m);
	}
	return movements, nil;
}

func (s *movementStore) Create(tx *sql.Tx, movement *models.Movement) (int64, error) {
	var err error

	query := "INSERT INTO movements (name) VALUES ($1) RETURNING id"

	var id int64

	if tx != nil {
		err = tx.QueryRow(query, movement.Name).Scan(&id);
	} else {
		err = s.QueryRow(query, movement.Name).Scan(&id);
	}

	if err != nil {
		return 0, err;
	}

	return id, nil;
} 

func (s *movementStore) UpdateById(tx *sql.Tx, movement *models.Movement) (int64, error) {
	query, err := s.Prepare("UPDATE movements SET name = $1 WHERE movements.id = $2")
	if err != nil {
		return 0, err;
	}

	var id int64;

	err = query.QueryRow(movement.Name).Scan(&id);

	if id == 0 {
		return 0, sql.ErrNoRows;
	} else if err != nil {
		return 0, err;
	}

	return id, nil;
}

func (s *movementStore) DeleteById(tx *sql.Tx, id int) error {
	row, err := s.Exec("DELETE FROM movements WHERE movements.id = $1 RETURNING movements.id", id);
	if err != nil {
		return err;
	}

	if r, err := row.RowsAffected(); err != nil {
		return err;
	} else if r == 0 {
		return sql.ErrNoRows;
	}

	return nil;
}