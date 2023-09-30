package services

import (
	"github.com/dzeleniak/arnold/models"
	"github.com/dzeleniak/arnold/stores"
)

type (
	MovementSetService interface {
		GetMovementSets() ([]models.MovementSet, error)
		CreateMovementSet(set *models.MovementSet) (int64, error)
		UpdateMovementSet(set *models.MovementSet) (int64, error)
		DeleteMovementSet(id int64) error
	}

	movementSetService struct {
		stores *stores.Stores
	}
)

func (s *movementSetService) GetMovementSets() ([]models.MovementSet, error) {
	sets, err := s.stores.MovementSet.Get(nil)

	return sets, err;
}	

func (s *movementSetService) CreateMovementSet(set *models.MovementSet) (int64, error) {
	id, err := s.stores.MovementSet.Create(nil, set)

	return id, err
}

func (s *movementSetService) UpdateMovementSet(set *models.MovementSet) (int64, error) {

	id, err := s.stores.MovementSet.Update(nil, set)

	return id, err;
}

func (s *movementSetService) DeleteMovementSet(id int64) error {
	err := s.stores.MovementSet.Delete(nil, id);
	return err;
}