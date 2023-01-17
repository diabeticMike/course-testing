package entity

import "github.com/google/uuid"

type Passport struct {
	ID     uuid.UUID `json:"id"`
	Code   int16     `json:"code"`
	UserID uuid.UUID
}
