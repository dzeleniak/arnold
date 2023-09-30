package services

import (
	"github.com/dzeleniak/arnold/models"
	"github.com/dzeleniak/arnold/stores"
)

type (
	MovementService interface {
		GetMovements() ([]models.Movement, error)
		CreateMovement(m *models.Movement) (int64, error)
		UpdateMovementById(m *models.Movement) (int64, error)
		DeleteMovement(id int) error
	}

	movementService struct {
		stores *stores.Stores
	}
)

func (s *movementService) GetMovements() ([]models.Movement, error) {
	r, err := s.stores.Movement.Get(nil);
	return r, err;
}

func (s *movementService) CreateMovement(m *models.Movement) (int64, error) {
	r, err := s.stores.Movement.Create(nil, m);
	return r, err;
}

func (s *movementService) UpdateMovementById(m *models.Movement) (int64, error) {
	r, err := s.stores.Movement.UpdateById(nil, m);
	return r, err;
}

func (s *movementService) DeleteMovement(id int) error {
	err := s.stores.Movement.DeleteById(nil, id);
	return err;
}