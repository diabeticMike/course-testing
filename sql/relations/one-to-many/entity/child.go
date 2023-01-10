package entity

import "github.com/google/uuid"

type Child struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	FatherID uuid.UUID
}
