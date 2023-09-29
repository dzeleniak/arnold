package models

type MovementSet struct {
	ID          int     `json:"id" db:"id"`
	MovementID  int     `json:"movement_id" db:"movement_id"`
	Weight      float32 `json:"weight"`
	Repetitions int     `json:"repetitions"`
}