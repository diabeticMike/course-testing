package entity

import "github.com/google/uuid"

type Father struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Children []Child   `json:"children,omitempty"`
}
