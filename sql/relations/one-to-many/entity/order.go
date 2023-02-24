package entity

import "github.com/google/uuid"

type Order struct {
	ID         uuid.UUID
	CustomerID uuid.UUID
}

func (Order) TableName() string {
	return "order"
}
