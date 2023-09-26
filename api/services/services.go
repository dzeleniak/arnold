package services

import "github.com/dzeleniak/arnold/stores"

type Services struct {
	Movement MovementService
}

func New(s *stores.Stores) *Services {
	return &Services{
		Movement: &movementService{stores: s},
	}
}