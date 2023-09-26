package models

type Movement struct {
	ID int `json:"id" db:"id"`
	Name string `json:"name"`
}